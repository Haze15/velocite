#!/bin/bash
# Construit Velocite.ipa (iPhone / iPad, non signée : LiveContainer, SideStore, AltStore… la signent eux-mêmes)
# et, avec « sim », une version pour le simulateur iPhone (test automatique). À lancer sur un Mac avec Xcode.
set -euo pipefail
cd "$(dirname "$0")/.."
VER=$(python3 -c "import json;print(json.load(open('version.json'))['v'])")
BUILD=${GITHUB_RUN_NUMBER:-1}
build() {   # $1 = sdk, $2 = cible, $3 = dossier, $4 = plateforme Info.plist
  local APP="$3/Payload/Velocite.app"
  rm -rf "$3"; mkdir -p "$APP/www"
  xcrun --sdk "$1" swiftc -parse-as-library -O -target "$2" -sdk "$(xcrun --sdk "$1" --show-sdk-path)" ios/main.swift -o "$APP/Velocite"
  sed -e "s/__VERSION__/$VER/" -e "s/__BUILD__/$BUILD/" -e "s/__PLATFORM__/$4/" -e "s/__PLATFORM_LC__/$(echo "$4" | tr 'A-Z' 'a-z')/" ios/Info.plist > "$APP/Info.plist"
  plutil -lint "$APP/Info.plist"
  plutil -convert binary1 "$APP/Info.plist"
  printf 'APPL????' > "$APP/PkgInfo"
  for s in "60x60@2x:120" "60x60@3x:180" "76x76@2x:152" "83.5x83.5@2x:167"; do
    sips -s format png -z "${s#*:}" "${s#*:}" assets/icon-512.png --out "$APP/AppIcon${s%%:*}.png" >/dev/null
  done
  cp index.html version.json manifest.webmanifest "$APP/www/"
  cp -R assets "$APP/www/"
  codesign --force --sign - --timestamp=none "$APP"
}
build iphoneos arm64-apple-ios15.0 build-ios iPhoneOS
(cd build-ios && rm -f ../Velocite.ipa && zip -qry ../Velocite.ipa Payload)
echo "OK : Velocite.ipa ($VER, $(du -h Velocite.ipa | cut -f1))"
if [ "${1:-}" = sim ]; then
  build iphonesimulator "$(uname -m)-apple-ios15.0-simulator" build-sim iPhoneSimulator
  echo "OK : build-sim/Payload/Velocite.app"
fi
