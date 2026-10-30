//go:build windows

// Vélocité pour Windows : sert le jeu embarqué sur 127.0.0.1 et l'ouvre dans une fenêtre
// d'application Chrome, Edge ou Brave (sinon dans le navigateur par défaut).
// L'appli se ferme toute seule quand la fenêtre du jeu est fermée.
package main

import (
	"embed"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

//go:embed game
var gameFS embed.FS

// dernier signe de vie envoyé par la page (secondes Unix)
var lastPing atomic.Int64

func main() {
	sub, err := fs.Sub(gameFS, "game")
	if err != nil {
		fatal("Fichiers du jeu introuvables dans l'application.")
	}
	raw, err := fs.ReadFile(sub, "index.html")
	if err != nil {
		fatal("index.html manquant dans l'application.")
	}
	page := patchPage(string(raw))

	files := http.FileServer(http.FS(sub))
	mux := http.NewServeMux()
	mux.HandleFunc("/__alive", func(w http.ResponseWriter, r *http.Request) {
		lastPing.Store(time.Now().Unix())
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/__presence", presenceHandler)
	go rp.run()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Cache-Control", "no-store")
			_, _ = w.Write([]byte(page))
			return
		}
		files.ServeHTTP(w, r)
	})

	// port fixe pour garder les réglages (le stockage du navigateur dépend du port)
	ln, err := net.Listen("tcp", "127.0.0.1:47321")
	if err != nil {
		ln, err = net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			fatal("Impossible de démarrer le serveur local du jeu.")
		}
	}
	url := "http://" + ln.Addr().String() + "/"
	go func() { _ = http.Serve(ln, mux) }()

	// 90 s pour que la fenêtre s'ouvre et que la page envoie son premier signe de vie
	lastPing.Store(time.Now().Unix() + 90)
	if !openAppWindow(url) {
		openDefaultBrowser(url)
	}
	for {
		time.Sleep(3 * time.Second)
		if time.Now().Unix()-lastPing.Load() > 20 {
			return
		}
	}
}

// Bibliothèques servies en local (jeu jouable hors ligne en solo) et battement de cœur toutes les 5 s.
func patchPage(html string) string {
	r := strings.NewReplacer(
		"https://cdnjs.cloudflare.com/ajax/libs/three.js/r128/three.min.js", "vendor/three.min.js",
		"https://cdn.jsdelivr.net/npm/three@0.128.0/examples/js/loaders/GLTFLoader.js", "vendor/GLTFLoader.js",
		"https://cdn.jsdelivr.net/npm/three@0.128.0/examples/js/utils/SkeletonUtils.js", "vendor/SkeletonUtils.js",
	)
	html = r.Replace(html)
	beat := "<script>(function(){function b(){fetch('/__alive',{cache:'no-store'}).catch(function(){})}b();setInterval(b,5000)})();</script>\n</body>"
	// signale à la page qu'elle tourne dans l'appli (statut Discord disponible)
	html = strings.Replace(html, "<script>\n(() => {", "<script>window.__VELOCITE_APP=1;</script>\n<script>\n(() => {", 1)
	return strings.Replace(html, "</body>", beat, 1)
}

func openAppWindow(url string) bool {
	local := os.Getenv("LOCALAPPDATA")
	pf, pf86 := os.Getenv("ProgramFiles"), os.Getenv("ProgramFiles(x86)")
	candidates := []string{
		filepath.Join(pf, `Google\Chrome\Application\chrome.exe`),
		filepath.Join(pf86, `Google\Chrome\Application\chrome.exe`),
		filepath.Join(local, `Google\Chrome\Application\chrome.exe`),
		filepath.Join(pf86, `Microsoft\Edge\Application\msedge.exe`),
		filepath.Join(pf, `Microsoft\Edge\Application\msedge.exe`),
		filepath.Join(pf, `BraveSoftware\Brave-Browser\Application\brave.exe`),
		filepath.Join(local, `BraveSoftware\Brave-Browser\Application\brave.exe`),
	}
	profile := filepath.Join(local, "Velocite", "profil")
	_ = os.MkdirAll(profile, 0o755)
	for _, exe := range candidates {
		if exe == "" || !fileExists(exe) {
			continue
		}
		cmd := exec.Command(exe,
			"--app="+url,
			"--user-data-dir="+profile,
			"--no-first-run",
			"--no-default-browser-check",
			"--window-size=1600,900",
			"--autoplay-policy=no-user-gesture-required",
		)
		if err := cmd.Start(); err == nil {
			go func() { _ = cmd.Wait() }()
			return true
		}
	}
	return false
}

func openDefaultBrowser(url string) {
	_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
}

func fileExists(p string) bool {
	st, err := os.Stat(p)
	return err == nil && !st.IsDir()
}

func fatal(msg string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	box := user32.NewProc("MessageBoxW")
	t, _ := syscall.UTF16PtrFromString("Vélocité")
	m, _ := syscall.UTF16PtrFromString(msg)
	_, _, _ = box.Call(0, uintptr(unsafe.Pointer(m)), uintptr(unsafe.Pointer(t)), 0x10)
	os.Exit(1)
}
