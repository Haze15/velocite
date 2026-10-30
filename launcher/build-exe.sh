#!/bin/sh
# Construit dist/Velocite.exe (Windows 64 bits) avec le jeu embarqué.
# Prérequis : Go 1.21+ et rsrc (github.com/akavel/rsrc) pour l'icône.
set -e
cd "$(dirname "$0")"
rm -rf game && mkdir -p game/vendor game/assets
cp ../index.html game/
cp -r ../assets/models game/assets/
cp -r ../assets/music game/assets/
cp vendor/*.js game/vendor/
if command -v rsrc >/dev/null 2>&1; then rsrc -ico icon.ico -arch amd64 -o rsrc_windows_amd64.syso; fi
mkdir -p ../dist
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -trimpath -ldflags "-H windowsgui -s -w" -o ../dist/Velocite.exe .
echo "OK : dist/Velocite.exe"
