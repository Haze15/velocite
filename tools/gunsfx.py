# Génère des coups de feu réalistes (synthèse physique simplifiée) -> WAV puis MP3
# couches : onde de bouche (Friedlander + bruit filtré), corps grave, claquement mécanique, queue (réverbération extérieure + écho de falaise)
import numpy as np, sys, os, subprocess
from scipy.signal import butter, sosfilt, fftconvolve
FS = 44100
def bp(x, lo, hi, o=2): return sosfilt(butter(o, [lo, hi], 'band', fs=FS, output='sos'), x)
def lp(x, f, o=2): return sosfilt(butter(o, f, 'low', fs=FS, output='sos'), x)
def hp(x, f, o=2): return sosfilt(butter(o, f, 'high', fs=FS, output='sos'), x)
def env(n, a, d): t = np.arange(n) / FS; return np.minimum(1, t / max(a, 1e-5)) * np.exp(-t / d)

def friedlander(n, td, A=1.0):
    t = np.arange(n) / FS; x = A * (1 - t / td) * np.exp(-t / td); return x

def ir(dur, rt, rng, early=()):
    n = int(dur * FS); t = np.arange(n) / FS
    r = rng.standard_normal(n) * np.exp(-6.9 * t / rt)
    r = lp(r, 5000) * 0.5 + lp(r, 1800) * 0.5
    r[0] = 1.0
    for dt, g in early:  # échos discrets (bâtiments, falaise)
        k = int(dt * FS)
        if k < n: r[k:k + 400] += g * np.hanning(400) * rng.standard_normal(400)
    return r

def shot(kind, seed):
    rng = np.random.default_rng(seed)
    P = {  # crack: fréquences du claquement ; body: grave ; td: durée de l'onde ; tail: réverbération
        'rifle':   dict(dur=1.1, td=0.0022, crack=(900, 7000), cg=1.0, body=105, bg=0.9, bd=0.06, mech=0.25, rt=0.9, wet=0.32, echo=((0.19, .25), (0.41, .12))),
        'smg':     dict(dur=0.75, td=0.0016, crack=(1200, 8000), cg=0.85, body=140, bg=0.6, bd=0.04, mech=0.35, rt=0.6, wet=0.26, echo=((0.16, .18),)),
        'shotgun': dict(dur=1.5, td=0.004, crack=(500, 5000), cg=1.0, body=62, bg=1.25, bd=0.12, mech=0.15, rt=1.25, wet=0.38, echo=((0.22, .3), (0.5, .15))),
        'sniper':  dict(dur=2.0, td=0.0032, crack=(700, 9000), cg=1.15, body=70, bg=1.3, bd=0.1, mech=0.12, rt=1.6, wet=0.42, echo=((0.28, .32), (0.63, .18), (1.05, .08))),
        'pistol':  dict(dur=0.8, td=0.0018, crack=(1000, 7000), cg=0.9, body=150, bg=0.65, bd=0.045, mech=0.4, rt=0.65, wet=0.28, echo=((0.17, .16),)),
        'revolver':dict(dur=1.1, td=0.0026, crack=(800, 6500), cg=1.0, body=95, bg=1.0, bd=0.075, mech=0.3, rt=0.95, wet=0.34, echo=((0.2, .22), (0.45, .1))),
        'lmg':     dict(dur=0.9, td=0.0024, crack=(850, 6500), cg=1.0, body=95, bg=1.0, bd=0.065, mech=0.3, rt=0.8, wet=0.3, echo=((0.18, .2),)),
    }[kind]
    j = lambda a, s=0.06: a * (1 + rng.uniform(-s, s))
    n = int(P['dur'] * FS)
    x = np.zeros(n)
    # 1) onde de bouche : impulsion de Friedlander (très courte) + bruit large bande à attaque instantanée
    td = j(P['td'], 0.12)
    fr = friedlander(int(0.03 * FS), td)
    x[:len(fr)] += fr * 1.6
    nb = rng.standard_normal(int(0.25 * FS))
    cr = bp(nb, *P['crack'], o=2) * env(len(nb), 0.0003, j(0.022, 0.2)) * P['cg'] * 1.25
    x[:len(cr)] += cr * 1.3
    # 2) corps : coup grave qui descend (pression des gaz)
    m = int(0.5 * FS); t = np.arange(m) / FS
    f0 = j(P['body']); ph = 2 * np.pi * np.cumsum(f0 * (0.55 + 0.45 * np.exp(-t / 0.03))) / FS
    body = np.sin(ph) * env(m, 0.0008, j(P['bd'] * 0.6, 0.15)) * P['bg'] * 0.65
    body += lp(rng.standard_normal(m), 260) * env(m, 0.001, P['bd'] * 1.6) * P['bg'] * 1.4
    x[:m] += body
    # 3) mécanique : culasse / percuteur (petits clics métalliques)
    for dt, g in ((0.0, 1.0), (j(0.012, 0.3), 0.6), (j(0.045, 0.2), 0.45)):
        k = int(dt * FS); c = bp(rng.standard_normal(int(0.012 * FS)), 2500, 9000) * env(int(0.012 * FS), 0.0001, 0.0025)
        x[k:k + len(c)] += c * P['mech'] * g
    # 4) saturation douce (micro et compression du coup de feu)
    x = np.tanh(x * 1.8) / np.tanh(1.8)
    # 5) espace : réverbération extérieure + échos
    r = ir(P['dur'], P['rt'], rng, P['echo'])
    wet = fftconvolve(x, r)[:n]; wet = hp(wet, 120) / (np.max(np.abs(wet)) + 1e-9)
    y = x * (1 - P['wet']) + wet * P['wet'] * 1.6
    y = hp(y, 35)
    # fondu de fin
    fade = int(0.08 * FS); y[-fade:] *= np.linspace(1, 0, fade)
    y /= np.max(np.abs(y)) + 1e-9
    return (y * 0.92).astype(np.float32)

if __name__ == '__main__':
    out = sys.argv[1]; os.makedirs(out, exist_ok=True)
    import wave
    for kind in ['rifle', 'smg', 'shotgun', 'sniper', 'pistol', 'revolver', 'lmg']:
        for v in range(3):
            y = shot(kind, 1000 + v * 17 + sum(map(ord, kind)))
            w = os.path.join(out, f'{kind}{v}.wav')
            with wave.open(w, 'wb') as f:
                f.setnchannels(1); f.setsampwidth(2); f.setframerate(FS); f.writeframes((y * 32767).astype('<i2').tobytes())
            mp = w[:-4] + '.mp3'
            subprocess.run(['ffmpeg', '-y', '-loglevel', 'error', '-i', w, '-codec:a', 'libmp3lame', '-b:a', '96k', mp], check=True)
            print(kind, v, os.path.getsize(mp))
