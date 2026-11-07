#!/usr/bin/env node
// Génère les clés publiques du mode propriétaire et du mot de passe admin de Vélocité.
//   node tools/keygen.js                 → tire deux mots de passe aléatoires (12 caractères) et affiche les clés
//   node tools/keygen.js "monMdp" "adm"  → calcule les clés pour des mots de passe choisis (6 à 16 caractères)
// Colle ensuite les deux valeurs dans index.html : const X_PUB = '…' (propriétaire) et const ADM_PUB = '…' (admin).
// Principe : clé privée Ed25519 = SHA-256(mot de passe) ; seule la clé publique est dans le jeu. Un joueur qui lit le code
// ne peut ni retrouver le mot de passe (sauf force brute : choisis-le long), ni se faire passer pour toi auprès d'un hôte.
const { webcrypto: crypto } = require('crypto');
const PKCS8 = Buffer.from('302e020100300506032b657004220420', 'hex');
const ALPHA = 'abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789';
const randPw = n => Array.from(crypto.getRandomValues(new Uint8Array(n))).map(b => ALPHA[b % ALPHA.length]).join('');
async function pubOf(pw) {
  const seed = Buffer.from(await crypto.subtle.digest('SHA-256', Buffer.from(pw, 'utf8')));
  const key = await crypto.subtle.importKey('pkcs8', Buffer.concat([PKCS8, seed]), { name: 'Ed25519' }, true, ['sign']);
  return (await crypto.subtle.exportKey('jwk', key)).x;
}
(async () => {
  const [, , o, a] = process.argv;
  const owner = o || randPw(12), admin = a || randPw(12);
  for (const [n, pw] of [['propriétaire', owner], ['admin', admin]]) if (pw.length < 6 || pw.length > 16) { console.error(`Mot de passe ${n} : 6 à 16 caractères.`); process.exit(1); }
  console.log(`Mot de passe propriétaire : ${owner}\n  const X_PUB = '${await pubOf(owner)}';`);
  console.log(`Mot de passe admin        : ${admin}\n  const ADM_PUB = '${await pubOf(admin)}';`);
})();
