//go:build !windows

package main

import (
	"errors"
	"io"
	"net"
	"os"
	"strconv"
)

// Linux et macOS : socket Unix dans le dossier temporaire (utilisé pour les tests).
func dialDiscord() (io.ReadWriteCloser, error) {
	dirs := []string{os.Getenv("XDG_RUNTIME_DIR"), os.Getenv("TMPDIR"), os.TempDir(), "/tmp"}
	for _, d := range dirs {
		if d == "" {
			continue
		}
		for i := 0; i < 10; i++ {
			c, err := net.Dial("unix", d+"/discord-ipc-"+strconv.Itoa(i))
			if err == nil {
				return c, nil
			}
		}
	}
	return nil, errors.New("discord introuvable")
}
