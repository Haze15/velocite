//go:build windows

package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"golang.org/x/sys/windows"
)

// Seulement dans Velocite-TricheTest.exe : de fausses triches inoffensives pour vérifier qu'Obsidian les voit.
//   /__cheat/dll  : copie une DLL de Windows dans %TEMP% sous un autre nom et la charge (comme un injecteur)
//   /__cheat/tool : lance un programme inoffensif (ping) renommé « cheatengine-test.exe » pendant 2 minutes
func cheatRoutes(mux *http.ServeMux) {
	if cheatTest != "1" {
		return
	}
	mux.HandleFunc("/__cheat/dll", func(w http.ResponseWriter, r *http.Request) {
		src := filepath.Join(os.Getenv("SystemRoot"), "System32", "version.dll")
		dst := filepath.Join(os.TempDir(), "obsidian-test-inject.dll")
		if err := copyFile(src, dst); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if _, err := windows.LoadLibrary(dst); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/__cheat/tool", func(w http.ResponseWriter, r *http.Request) {
		src := filepath.Join(os.Getenv("SystemRoot"), "System32", "PING.EXE")
		dst := filepath.Join(os.TempDir(), "cheatengine-test.exe")
		if err := copyFile(src, dst); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		c := exec.Command(dst, "-n", "120", "127.0.0.1")
		c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000} // CREATE_NO_WINDOW
		if err := c.Start(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		go func() { _ = c.Wait() }()
		w.WriteHeader(http.StatusNoContent)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		return err
	}
	return out.Close()
}
