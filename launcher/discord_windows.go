//go:build windows

package main

import (
	"errors"
	"io"
	"os"
	"strconv"
)

// Discord ouvre un tube nommé \\.\pipe\discord-ipc-0 (jusqu'à 9 s'il y a plusieurs clients).
func dialDiscord() (io.ReadWriteCloser, error) {
	for i := 0; i < 10; i++ {
		f, err := os.OpenFile(`\\.\pipe\discord-ipc-`+strconv.Itoa(i), os.O_RDWR, 0)
		if err == nil {
			return f, nil
		}
	}
	return nil, errors.New("discord introuvable")
}
