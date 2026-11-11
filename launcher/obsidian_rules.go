package main

import (
	"path/filepath"
	"strings"
)

// Obsidian (anti-triche) : règles pour juger une DLL chargée dans le jeu ou un programme ouvert à côté.
// Les DLL de Windows, des programmes installés (Program Files) et du moteur WebView2 sont normales ;
// les overlays connus (Discord, Steam, OBS, NVIDIA, Medal, Overwolf…) aussi. Une DLL venue d'ailleurs
// (Bureau, Téléchargements, Temp…) a été injectée.

type obsEnv struct {
	WinDir   string   // C:\Windows
	Programs []string // C:\Program Files, C:\Program Files (x86)
	ExeDir   string   // dossier de Velocite.exe
}

var obsOkParts = []string{
	`\microsoft\edgewebview\`, `\microsoft\edge\`, `\appdata\local\microsoft\`, `\programdata\microsoft\`,
	`\discord`, `\steam\`, `\obs-studio`, `\nvidia`, `\medal\`, `\overwolf\`, `\rivatuner`, `\msi afterburner`,
	`\razer`, `\logitech`, `\amd\`, `\intel\`, `\nahimic`, `\sonic studio`, `\geforce`, `\epic games\`, `\ubisoft`,
}

func obsLower(p string) string { return strings.ToLower(filepath.Clean(strings.ReplaceAll(p, "/", `\`))) }

func obsUnder(p, dir string) bool {
	if dir == "" {
		return false
	}
	d := obsLower(dir)
	return strings.HasPrefix(p, strings.TrimRight(d, `\`)+`\`)
}

// true = DLL suspecte (injectée)
func obsSuspiciousModule(path string, env obsEnv) bool {
	p := obsLower(path)
	if p == "" || !strings.HasSuffix(p, ".dll") {
		return false
	}
	if obsUnder(p, env.WinDir) {
		return false
	}
	for _, d := range env.Programs {
		if obsUnder(p, d) {
			return false
		}
	}
	for _, s := range obsOkParts {
		if strings.Contains(p, s) {
			return false
		}
	}
	// une DLL posée à côté de Velocite.exe n'en fait pas partie (le jeu n'en livre aucune) : suspecte
	return true
}

// Programmes de triche connus ouverts en même temps que le jeu
var obsTools = []string{"cheatengine", "cheat engine", "extremeinjector", "extreme injector", "xenos", "ghinjector", "gh injector", "dllinjector", "dll injector", "injector.exe", "artmoney", "wemod", "speedhack", "ce-x86", "ce-x64"}

func obsSuspiciousProcess(name string) string {
	n := strings.ToLower(name)
	for _, t := range obsTools {
		if strings.Contains(n, t) {
			return name
		}
	}
	return ""
}
