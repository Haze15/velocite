# Vélocité — lots A à E (7 novembre 2026)

Contenu de l'archive `Velocite-lots-ABCDE.zip` : **uniquement les fichiers modifiés ou ajoutés**, à déposer par-dessus l'arborescence d'origine (`Velocite-jeu-source.zip`). Rien n'a été supprimé.

```
index.html                 remplace l'original (toutes les modifications sont dedans)
README.md                  sections « Comptes », « Sécurité en ligne », « Structure »
assets/vendor/             three.min.js r128, GLTFLoader.js, SkeletonUtils.js, peerjs.min.js 1.5.4 (+ licences)
tools/keygen.js            génère les clés publiques propriétaire/admin à partir de mots de passe
patches/                   un diff par lot (lotA-securite, lotB-robustesse, lotC-perf, lotD-rendu, lotE-modeles, readme) et tout.diff
captures/                  avant/après rendus par le jeu (Canyon, Forêt, Nécropolis) et planche récapitulative
```

## À faire de ton côté (2 minutes)

1. **Choisir tes mots de passe.** Les anciennes empreintes SHA-256 ne pouvaient pas être conservées (il fallait un secret que personne ne peut rejouer). Les deux mots de passe actuellement en vigueur dans `index.html` sont ceux qui t'ont été donnés dans la conversation. Pour en mettre d'autres :
   `node tools/keygen.js "monMdpProprietaire" "monMdpAdmin"` → coller `X_PUB` et `ADM_PUB` dans `index.html` (lignes 1419-1420).
2. **Rien d'autre.** Les comptes déjà créés chez un hôte restent valables (même format `salt` + `h`). Les joueurs doivent avoir la nouvelle version pour rejoindre un hôte à jour (le bandeau de mise à jour s'en charge).

## Ce qui change pour les joueurs

- Les CDN ne sont plus nécessaires : le jeu démarre sur les réseaux qui les bloquent, et dit clairement quand quelque chose manque (bandeau rouge).
- Un hôte ne peut plus être pris en main par n'importe qui, un joueur ne peut plus tuer à distance ni se rendre invisible, un message malformé ne gèle plus la partie.
- Les mots de passe (comptes, admin, propriétaire) ne transitent plus jamais sur le réseau.
- En visée lunette : jusqu'à −19 % de draw calls. En qualité « Auto » sur écran 60 Hz, la résolution remonte quand les fps le permettent.

## Lot D — rendu (prototype validé sur Canyon, Forêt, Nécropolis, puis appliqué à toutes les cartes)

- **Toutes qualités** : pipeline couleur correct (hex et textures lus en sRGB, sortie sRGB, tone mapping ACES, exposition 0,8–0,92), ambiant/soleil rééquilibrés, brouillard de hauteur réglé par carte, particules d'ambiance rondes et douces.
- **Élevé / Ultra** : matériaux PBR partout (rugosité par matière, relief dérivé des textures procédurales), environnement calculé depuis le ciel de la carte (PMREM), SSAO demi-résolution, rayons de lumière. Le changement de qualité reconstruit la carte.
- Réglages faciles à retoucher dans `index.html` : `RENDU` (exposition, hemiK, sunK, envK, aoK, raysK), `SURF_BUMP` / `SURF_ROUGH` (relief et rugosité par matière), `FOGH_OF` (brouillard de hauteur par carte), `th.exposure` / `th.hemiK` / `th.sunK` dans le thème d'une carte pour une exception.
- À savoir : en Élevé/Ultra, le SSAO remplace le MSAA par un FXAA (la profondeur doit être lisible) ; les PC moyens en « Élevé » peuvent perdre quelques fps (matériaux plus lourds), la résolution automatique compense en « Auto ».

## Lot E — modèles (ce qui était « cassé »)

- **Autres joueurs** : la pose « arme tenue » était faite d'angles fixes réglés sur le squelette X-Bot ; sur le Soldat et le Titan, ces angles levaient les bras et l'arme flottait. Remplacée par une IK à deux os : arme devant la poitrine (suit la visée), mains sur les prises, pour n'importe quel squelette (y compris un modèle importé). Vérifié sur les 7 personnages avec carabine, SMG, fusil à pompe, sniper, pistolet et couteau.
- **Couteaux blancs** après un changement de couteau en partie : un compteur de version était rangé dans la table des matériaux de l'arme, ce qui faisait planter l'application des skins (bug d'origine).
- **Aperçu du menu** : le Soldat posait « fusil prêt » sans fusil. L'aperçu tient maintenant la carabine avec le skin choisi, même tenue que les autres joueurs ; idem pour l'avatar à la troisième personne.
- Fausse alerte vérifiée : les bras et armes en vue subjective, les robots, les véhicules et les 11 couteaux sont corrects.

## Vérifications effectuées

- Syntaxe (`node --check`), chargement des 18 cartes sans erreur, partie solo.
- Multijoueur réel (hôte + 2 clients WebRTC, signalisation PeerJS locale) : 7 attaques bloquées (toutes passaient sur l'original), 10 scénarios légitimes OK (comptes, kill crédité, accès propriétaire/admin par signature, refus des faux).
- Lot B : libs locales (0 requête CDN), repli CDN fonctionnel, bandeau en cas d'exception ou de promesse rejetée, `store.set` qui prévient quand le stockage est plein, message clair si le moteur 3D manque.
- Lot C : draw calls mesurés en visée lunette (1 992 → 1 616 par image) ; `autoRes` 0,6 → 1,0 sur 60 Hz simulé.
- Lot D : 18 cartes construites sans erreur après les changements de rendu, captures avant/après dans `captures/`.
- Lot E : 18 cartes sans erreur, 7 attaques toujours bloquées en multijoueur, captures `avatars-avant-apres.png` et `apercu-menu-avant-apres.png`.

## Limites connues (volontaires)

- Pas de contrôle de plausibilité des déplacements : téléport, portails, grappins, jump pads et réapparitions rendent toute borne de vitesse fausse. Un « noclip » reste donc possible côté client ; c'est inhérent au pair-à-pair sans serveur.
- Les dégâts restent calculés par le tireur (bornés par l'hôte), et appliqués par la victime : un client modifié peut toujours s'auto-soigner. Seul un serveur autoritaire l'empêcherait.
- Ed25519 WebCrypto demande Chrome/Edge 113+, Safari 17+, Firefox 130+ : sur un navigateur plus ancien, le mode propriétaire et le mot de passe admin ne fonctionnent pas (message « navigateur à mettre à jour »). Le jeu, lui, marche.
- `privacy.html` n'a pas été modifiée (point 5 de la liste, non retenu) : elle devrait mentionner PeerJS cloud, le relais TURN Metered, `version.json` sur GitHub et les comptes stockés chez l'hôte.
