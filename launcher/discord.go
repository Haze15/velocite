package main

// Statut Discord (Rich Presence) : la page du jeu envoie son état à /__presence,
// le lanceur le transmet au client Discord du PC par son canal IPC local.

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

// Etat envoyé par la page.
type presenceReq struct {
	On       bool   `json:"on"`
	ClientID string `json:"cid"`
	Details  string `json:"details"`
	State    string `json:"state"`
	Start    int64  `json:"start"` // ms depuis l'époque Unix
	Party    string `json:"party"`
	Size     int    `json:"size"`
	Max      int    `json:"max"`
	Large    string `json:"large"`
	Buttons  []struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	} `json:"buttons"`
}

var idRe = regexp.MustCompile(`^[0-9]{15,22}$`)

type presence struct {
	mu      sync.Mutex
	want    *presenceReq
	dirty   bool
	wake    chan struct{}
	conn    io.ReadWriteCloser
	connID  string
	nonce   int
	lastErr string
}

var rp = &presence{wake: make(chan struct{}, 1)}

func (p *presence) set(r *presenceReq) {
	p.mu.Lock()
	p.want = r
	p.dirty = true
	p.mu.Unlock()
	select {
	case p.wake <- struct{}{}:
	default:
	}
}

func (p *presence) status() (bool, string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.conn != nil, p.lastErr
}

// Boucle : (re)connexion, envoi limité à une mise à jour toutes les 4 s (Discord en accepte 5 par 20 s).
func (p *presence) run() {
	var lastSend time.Time
	retry := time.NewTicker(15 * time.Second)
	defer retry.Stop()
	for {
		select {
		case <-p.wake:
		case <-retry.C:
		}
		if wait := 4*time.Second - time.Since(lastSend); wait > 0 {
			time.Sleep(wait)
		}
		p.mu.Lock()
		want, dirty := p.want, p.dirty
		p.mu.Unlock()
		if want == nil || !dirty {
			continue
		}
		if err := p.apply(want); err != nil {
			p.mu.Lock()
			p.lastErr = err.Error()
			p.mu.Unlock()
			p.close()
			continue
		}
		lastSend = time.Now()
		p.mu.Lock()
		if p.want == want {
			p.dirty = false
		}
		p.lastErr = ""
		p.mu.Unlock()
	}
}

func (p *presence) close() {
	p.mu.Lock()
	c := p.conn
	p.conn = nil
	p.mu.Unlock()
	if c != nil {
		_ = c.Close()
	}
}

func (p *presence) apply(r *presenceReq) error {
	cid := r.ClientID
	if !idRe.MatchString(cid) {
		cid = fileClientID()
	}
	if !r.On || cid == "" {
		if p.conn != nil {
			_ = p.send(1, p.activity(nil))
		}
		if !r.On {
			p.close()
		}
		return nil
	}
	if p.conn != nil && p.connID != cid {
		p.close()
	}
	if p.conn == nil {
		c, err := dialDiscord()
		if err != nil {
			return errors.New("discord fermé")
		}
		p.mu.Lock()
		p.conn = c
		p.connID = cid
		p.mu.Unlock()
		if err := writeFrame(c, 0, map[string]any{"v": 1, "client_id": cid}); err != nil {
			return err
		}
		op, body, err := readFrame(c)
		if err != nil {
			return err
		}
		if op != 1 {
			return errors.New("discord a refusé l'identifiant : " + string(body))
		}
	}
	return p.send(1, p.activity(r))
}

func (p *presence) activity(r *presenceReq) map[string]any {
	p.nonce++
	args := map[string]any{"pid": os.Getpid()}
	if r != nil {
		a := map[string]any{"instance": false}
		// Discord refuse les textes de moins de 2 caractères : champs vides omis
		if r.Details != "" {
			a["details"] = clip(r.Details, 120)
		}
		if r.State != "" {
			a["state"] = clip(r.State, 120)
		}
		if r.Start > 0 {
			a["timestamps"] = map[string]any{"start": r.Start / 1000}
		}
		large := r.Large
		if large == "" {
			large = "logo"
		}
		a["assets"] = map[string]any{"large_image": large, "large_text": "Vélocité"}
		if r.Party != "" && r.Max > 0 {
			size := r.Size
			if size < 1 {
				size = 1
			}
			a["party"] = map[string]any{"id": clip(r.Party, 64), "size": []int{size, r.Max}}
		}
		var btns []map[string]string
		for _, b := range r.Buttons {
			if len(btns) == 2 || b.Label == "" || len(b.URL) < 8 || b.URL[:8] != "https://" {
				continue
			}
			btns = append(btns, map[string]string{"label": clip(b.Label, 31), "url": clip(b.URL, 512)})
		}
		if len(btns) > 0 {
			a["buttons"] = btns
		}
		args["activity"] = a
	}
	return map[string]any{"cmd": "SET_ACTIVITY", "args": args, "nonce": strconv.Itoa(p.nonce)}
}

func (p *presence) send(op uint32, v any) error {
	c := p.conn
	if c == nil {
		return errors.New("pas connecté")
	}
	if err := writeFrame(c, op, v); err != nil {
		return err
	}
	// réponse de Discord (ou erreur) ; on lit pour vider le canal
	rop, body, err := readFrame(c)
	if err != nil {
		return err
	}
	if rop == 2 {
		return errors.New("discord a fermé la connexion : " + string(body))
	}
	var resp struct {
		Evt  string `json:"evt"`
		Data struct {
			Message string `json:"message"`
		} `json:"data"`
	}
	if json.Unmarshal(body, &resp) == nil && resp.Evt == "ERROR" {
		return errors.New(resp.Data.Message)
	}
	return nil
}

func clip(s string, n int) string {
	r := []rune(s)
	if len(r) > n {
		r = r[:n]
	}
	out := string(r)
	if len([]rune(out)) < 2 && out != "" {
		out += "  "
	}
	return out
}

// Trame IPC Discord : opcode (uint32 LE), longueur (uint32 LE), JSON.
func writeFrame(w io.Writer, op uint32, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	h := make([]byte, 8)
	binary.LittleEndian.PutUint32(h[0:], op)
	binary.LittleEndian.PutUint32(h[4:], uint32(len(b)))
	_, err = w.Write(append(h, b...))
	return err
}

func readFrame(r io.Reader) (uint32, []byte, error) {
	h := make([]byte, 8)
	if _, err := io.ReadFull(r, h); err != nil {
		return 0, nil, err
	}
	op, n := binary.LittleEndian.Uint32(h[0:]), binary.LittleEndian.Uint32(h[4:])
	if n > 1<<20 {
		return 0, nil, errors.New("trame trop grande")
	}
	b := make([]byte, n)
	if _, err := io.ReadFull(r, b); err != nil {
		return 0, nil, err
	}
	return op, b, nil
}

// Identifiant d'application de secours : fichier discord-id.txt à côté de l'exécutable.
func fileClientID() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	dir := exe[:len(exe)-len(baseName(exe))]
	b, err := os.ReadFile(dir + "discord-id.txt")
	if err != nil {
		return ""
	}
	s := string(regexp.MustCompile(`[0-9]+`).Find(b))
	if idRe.MatchString(s) {
		return s
	}
	return ""
}

func baseName(p string) string {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == '/' || p[i] == '\\' {
			return p[i+1:]
		}
	}
	return p
}

// Point d'entrée HTTP : POST /__presence avec l'état, GET pour connaître l'état de la connexion.
func presenceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req presenceReq
		if err := json.NewDecoder(io.LimitReader(r.Body, 8192)).Decode(&req); err != nil {
			http.Error(w, "bad", http.StatusBadRequest)
			return
		}
		rp.set(&req)
	}
	ok, msg := rp.status()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	_ = json.NewEncoder(w).Encode(map[string]any{"connected": ok, "error": msg})
}
