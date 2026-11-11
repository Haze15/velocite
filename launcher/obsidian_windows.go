//go:build windows

package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Obsidian côté Windows : toutes les 5 s, liste les DLL chargées dans Velocite.exe et dans ses fenêtres
// WebView2 (processus enfants), et les programmes de triche ouverts. La page lit le résultat sur /__obs.
var obsState struct {
	sync.Mutex
	Mods  []string `json:"mods"`
	Tools []string `json:"tools"`
	Done  bool     `json:"done"`
}

func obsEnvNow() obsEnv {
	exe, _ := os.Executable()
	win := os.Getenv("SystemRoot")
	if win == "" {
		win = `C:\Windows`
	}
	var progs []string
	for _, k := range []string{"ProgramFiles", "ProgramFiles(x86)", "ProgramW6432", "CommonProgramFiles", "CommonProgramFiles(x86)"} {
		if v := os.Getenv(k); v != "" {
			progs = append(progs, v)
		}
	}
	return obsEnv{WinDir: win, Programs: progs, ExeDir: filepath.Dir(exe)}
}

// processus : PID → (parent, nom)
func obsProcs() (map[uint32]uint32, map[uint32]string) {
	par, name := map[uint32]uint32{}, map[uint32]string{}
	snap, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return par, name
	}
	defer windows.CloseHandle(snap)
	var e windows.ProcessEntry32
	e.Size = uint32(unsafe.Sizeof(e))
	for err = windows.Process32First(snap, &e); err == nil; err = windows.Process32Next(snap, &e) {
		par[e.ProcessID] = e.ParentProcessID
		name[e.ProcessID] = windows.UTF16ToString(e.ExeFile[:])
	}
	return par, name
}

func obsModules(pid uint32) []string {
	var out []string
	snap, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPMODULE|windows.TH32CS_SNAPMODULE32, pid)
	if err != nil {
		return out
	}
	defer windows.CloseHandle(snap)
	var m windows.ModuleEntry32
	m.Size = uint32(unsafe.Sizeof(m))
	for err = windows.Module32First(snap, &m); err == nil; err = windows.Module32Next(snap, &m) {
		out = append(out, windows.UTF16ToString(m.ExePath[:]))
	}
	return out
}

func obsScan() {
	env := obsEnvNow()
	self := uint32(os.Getpid())
	par, names := obsProcs()
	// Velocite.exe et tous ses descendants (WebView2 : navigateur, rendu, GPU…)
	mine := map[uint32]bool{self: true}
	for changed := true; changed; {
		changed = false
		for pid, pp := range par {
			if !mine[pid] && mine[pp] && pid != pp {
				mine[pid] = true
				changed = true
			}
		}
	}
	seen := map[string]bool{}
	var mods []string
	for pid := range mine {
		for _, p := range obsModules(pid) {
			if obsSuspiciousModule(p, env) && !seen[p] {
				seen[p] = true
				mods = append(mods, filepath.Base(p))
			}
		}
	}
	var tools []string
	for _, n := range names {
		if t := obsSuspiciousProcess(n); t != "" && !seen[t] {
			seen[t] = true
			tools = append(tools, t)
		}
	}
	sort.Strings(mods)
	sort.Strings(tools)
	obsState.Lock()
	obsState.Mods, obsState.Tools, obsState.Done = mods, tools, true
	obsState.Unlock()
}

func obsRun() {
	for {
		func() { defer func() { _ = recover() }(); obsScan() }()
		time.Sleep(5 * time.Second)
	}
}

func obsHandler(w http.ResponseWriter, r *http.Request) {
	obsState.Lock()
	b, _ := json.Marshal(&obsState)
	obsState.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	_, _ = w.Write(b)
}
