/* Triche de test pour Obsidian — à coller dans la console (F12 → Console) sur le site de Vélocité.
   Un petit panneau apparaît en haut à droite : chaque bouton déclenche une triche différente.
   Recharge la page (F5) pour tout annuler (sauf « Crédits », qui est justement remis à zéro par Obsidian). */
(() => {
  if (document.getElementById('cheatTest')) return;
  const box = document.createElement('div');
  box.id = 'cheatTest';
  box.style.cssText = 'position:fixed;top:70px;right:12px;z-index:200000;background:#140a0d;border:1px solid #ff3a55;border-radius:10px;padding:10px;font:13px system-ui;color:#fff;display:flex;flex-direction:column;gap:6px;width:220px';
  box.innerHTML = '<b style="color:#ff6a7e">TRICHE DE TEST</b>';
  const btn = (label, fn) => {
    const b = document.createElement('button');
    b.textContent = label;
    b.style.cssText = 'padding:7px;border-radius:6px;border:1px solid #552;background:#2a1418;color:#fff;cursor:pointer;text-align:left';
    b.onclick = () => { try { fn(); b.textContent = '✓ ' + label; } catch (e) { b.textContent = '✕ ' + e.message; } };
    box.append(b);
  };

  // 1. Speedhack : le temps du jeu va 2× plus vite (Obsidian le voit en ~8 s)
  btn('Speedhack ×2', () => {
    const real = Performance.prototype.now; let t0 = null;
    Performance.prototype.now = function () { const r = real.call(this); if (t0 === null) t0 = r; return t0 + (r - t0) * 2; };
  });

  // 2. Code modifié : on « accroche » requestAnimationFrame comme le ferait un aimbot
  btn('Modifier le code du jeu', () => {
    const raf = window.requestAnimationFrame;
    window.requestAnimationFrame = f => raf.call(window, f);
  });

  // 3. Script injecté depuis un site extérieur
  btn('Injecter un script', () => {
    const s = document.createElement('script');
    s.src = 'https://triche.example.com/aimbot.js';
    document.head.append(s);
  });

  // 4. 99 999 crédits écrits à la main dans la sauvegarde, puis rechargement
  btn('Me donner 99 999 crédits', () => {
    localStorage.setItem('velocite_credits', '99999');
    localStorage.setItem('velocite_goldToken', 'true');
    setTimeout(() => location.reload(), 300);
  });

  document.body.append(box);
})();
