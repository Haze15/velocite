package main

import "testing"

func TestObsModules(t *testing.T) {
	env := obsEnv{WinDir: `C:\Windows`, Programs: []string{`C:\Program Files`, `C:\Program Files (x86)`}, ExeDir: `C:\Users\H\Desktop`}
	ok := []string{
		`C:\Windows\System32\kernel32.dll`, `C:\WINDOWS\SYSTEM32\ntdll.dll`,
		`C:\Program Files (x86)\Microsoft\EdgeWebView\Application\130.0\msedge.dll`,
		`C:\Users\H\AppData\Local\Discord\app-1.0\modules\discord_hook\DiscordHook64.dll`,
		`C:\Program Files (x86)\Steam\GameOverlayRenderer64.dll`,
		`C:\ProgramData\obs-studio-hook\graphics-hook64.dll`,
		`C:\Users\H\AppData\Local\Medal\recorder\hook64.dll`,
		`C:\Users\H\Desktop\Velocite.exe`,
	}
	bad := []string{`C:\Users\H\Desktop\aimbot.dll`, `C:\Users\H\Downloads\cheat\esp.dll`, `C:\Users\H\AppData\Local\Temp\x.dll`, `D:\hacks\wall.dll`}
	for _, p := range ok {
		if obsSuspiciousModule(p, env) {
			t.Errorf("faux positif : %s", p)
		}
	}
	for _, p := range bad {
		if !obsSuspiciousModule(p, env) {
			t.Errorf("non détecté : %s", p)
		}
	}
	if obsSuspiciousProcess("cheatengine-x86_64-SSE4-AVX2.exe") == "" || obsSuspiciousProcess("Extreme Injector v3.exe") == "" {
		t.Error("outil non détecté")
	}
	if obsSuspiciousProcess("Discord.exe") != "" || obsSuspiciousProcess("chrome.exe") != "" || obsSuspiciousProcess("SystemInformer.exe") != "" {
		t.Error("faux positif programme")
	}
}
