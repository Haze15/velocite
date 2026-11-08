# Vélocité — lots A à F (8 novembre 2026)

Contenu de l'archive `Velocite-lots-ABCDEF.zip` : **uniquement les fichiers modifiés ou ajoutés**, à déposer par-dessus l'arborescence d'origine (`Velocite-jeu-source.zip`). Rien n'a été supprimé.

```
index.html                 remplace l'original (toutes les modifications sont dedans)
README.md                  sections « Comptes », « Sécurité en ligne », « Structure », « Rendu », « Les Docks », tableau des personnages et des armes spéciales
assets/vendor/             three.min.js r128, GLTFLoader.js, SkeletonUtils.js, peerjs.min.js 1.5.4 (+ licences)
tools/keygen.js            génère les clés publiques propriétaire/admin à partir de mots de passe
patches/                   un diff par lot (lotA-securite, lotB-robustesse, lotC-perf, lotD-rendu, lotE-modeles, lotF-contenu, readme) et tout.diff
captures/                  avant/après des lots D et E, puis docks-*.png (la carte), armes-planche.png, ancre-planche.png (lot F)
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

## Lot F — nouveau contenu

- **Carte « Les Docks »** (`docks`, après Néo-Dune dans la liste) : port de commerce à l'heure bleue, bruine, quais surélevés au-dessus d'un bassin où l'on nage. Jetée nord avec deux portiques à conteneurs enjambant le cargo ARGONAUTE (pont chargé, château arrière à trois niveaux, passerelle d'embarquement, grue de pont, amarres), tour d'accès à tremplins jusqu'à la flèche du portique (26 m), capitainerie-phare au faisceau tournant ; entrepôt 7 ouvert (rayonnages, passerelle intérieure, chariot élévateur, établi, bobines), entrepôt 3 (camion porte-conteneur, frigos), entrepôt frigorifique au toit accessible, bureau du port, cour à conteneurs (5 files, piles de 1 à 3), parc à cuves (3 réservoirs avec passerelles, tuyauterie, vannes), voie ferrée (locotracteur, wagons, wagon-citerne) ; marina (pontons flottants, station carburant, bateaux et bouées qui tanguent), appontement isolé relié par une tyrolienne de 135 m ; digue et feux d'entrée, cargos au large, skyline éclairée. 12 points d'apparition, 10 robots, 5 tyroliennes, 11 tremplins, musique générative dédiée, entrée dans la progression.
  Performances : ~1 300 objets de décor sont **fusionnés par matériau** à la construction (`mergeDecoList`, réutilisable pour d'autres cartes) → 370 maillages, 61 k triangles, 135 ms de construction, dans la moyenne des cartes existantes. Trois lumières ponctuelles éclairent l'intérieur des hangars (pas en qualité Basse).
- **Personnage « Ancre »** (classe Ingénieure, base X-Bot, équipement procédural fixé aux os : casque de chantier à lampe, visière de soudure relevée, épaulières, sac à dos à antenne, ceinture à outils, genouillères ; articulations orange sécurité + couleur choisie ; bras à la première personne assortis).
  - Tactique **Tourelle** (18 s) : posée sur la surface visée à moins de 8 m (ou devant toi), 12 s, cible la plus proche à moins de 25 m en vue, 2 tirs/s de 8 dégâts, une seule à la fois. Chez les autres joueurs elle est visible et vise, mais seule la copie du propriétaire blesse (messages « hit » w = 23 relayés et bornés par l'hôte comme les grenades).
  - Ultime **Rempart** (45 s) : mur d'énergie 7 m × 3 m, 3 m devant toi, 10 s, aligné sur l'axe le plus proche ; c'est un vrai bloc de collision + un obstacle aux tirs et à la ligne de vue des explosions, chez tout le monde.
- **Cinq armes spéciales** (dans les coffres comme les autres, et posées à demeure sur les Docks) : **Fusil à plasma** (boules rapides, 64 à 3,4 m, 14 charges), **Mortier M-7** (tir en cloche calculé pour retomber sur le point visé, sifflement en descente, 150 à 8 m, 6 obus), **Arc Tesla** (éclair instantané 38 m qui saute sur 3 cibles de plus à moins de 8 m : 24/16/10/7), **Essaim** (salves de 3 micro-missiles guidés sur la cible la plus alignée, 48 chacun à 2,8 m, 6 salves), **Canon à impulsion** (cône de 10 m : 44 max, projette les ennemis, renvoie les projectiles, recul qui propulse le tireur — saut d'impulsion en visant le sol). Chacune a son modèle 3D procédural, ses prises de mains, son son, sa couleur de traçante et son étiquette dans le fil des kills (fr/en).
- **Réseau** : un seul nouveau type de message, `fx` (tourelle, rempart), relayé par l'hôte uniquement pour les types connus, avec position vérifiée et débit limité par type (0,4/s chacun : une rafale de 8 remparts n'en fait passer que 2). Les nouveaux projectiles passent par `gren` (déjà limité à 5/s) ; les dégâts par `hit` (budget inchangé, MAX_HIT 300 : aucune nouvelle arme ne dépasse 150 par coup).
- **Points d'accroche ajoutés dans le code** pour la suite (du contenu peut maintenant être ajouté sans toucher aux fonctions centrales) : `VM_EXT` / `VM_HANDS` (modèle et prises d'une arme spéciale), `FIRE_EXT` (tir instantané particulier), `GREN_EXT` / `GREN_MESH` / `GREN_BOOM` (vol, apparence, explosion d'un projectile), `SP_SOUND`, `SHOT_FX` (effet vu par les autres), `ABIL_EXT` + `abStepExt` + `abActiveExt` (capacités), `FX_OK` / `FX_RECV` (messages fx).

## Vérifications effectuées

- Syntaxe (`node --check`), chargement des 18 cartes (20 avec le lot F) sans erreur, partie solo.
- Multijoueur réel (hôte + 2 clients WebRTC, signalisation PeerJS locale) : 7 attaques bloquées (toutes passaient sur l'original), 10 scénarios légitimes OK (comptes, kill crédité, accès propriétaire/admin par signature, refus des faux).
- Lot B : libs locales (0 requête CDN), repli CDN fonctionnel, bandeau en cas d'exception ou de promesse rejetée, `store.set` qui prévient quand le stockage est plein, message clair si le moteur 3D manque.
- Lot C : draw calls mesurés en visée lunette (1 992 → 1 616 par image) ; `autoRes` 0,6 → 1,0 sur 60 Hz simulé.
- Lot D : 18 cartes construites sans erreur après les changements de rendu, captures avant/après dans `captures/`.
- Lot E : 18 cartes sans erreur, 7 attaques toujours bloquées en multijoueur, captures `avatars-avant-apres.png` et `apercu-menu-avant-apres.png`.
- Lot F : syntaxe, 20 cartes construites sans erreur (bench), partie solo ; les 5 armes tirent et consomment leurs munitions, la tourelle et le rempart se posent et expirent, Ancre se charge et s'affiche (aperçu, première et troisième personne) ; multijoueur réel hôte + client : Ancre vue avec son équipement, tourelle et rempart relayés (le rempart ajoute bien un bloc de collision chez l'hôte), tourelle de l'hôte qui blesse le client, 3 missiles / 1 obus / 1 boule de plasma vus par l'hôte, arc Tesla et impulsion qui blessent (et poussent) l'hôte, fx inconnu ignoré, rafale de remparts limitée ; les 7 attaques du lot A toujours bloquées. Captures : `docks-planche.png` (+ 8 vues), `armes-planche.png`, `ancre-planche.png`.

## Limites connues (volontaires)

- Pas de contrôle de plausibilité des déplacements : téléport, portails, grappins, jump pads et réapparitions rendent toute borne de vitesse fausse. Un « noclip » reste donc possible côté client ; c'est inhérent au pair-à-pair sans serveur.
- Les dégâts restent calculés par le tireur (bornés par l'hôte), et appliqués par la victime : un client modifié peut toujours s'auto-soigner. Seul un serveur autoritaire l'empêcherait.
- Ed25519 WebCrypto demande Chrome/Edge 113+, Safari 17+, Firefox 130+ : sur un navigateur plus ancien, le mode propriétaire et le mot de passe admin ne fonctionnent pas (message « navigateur à mettre à jour »). Le jeu, lui, marche.
- `privacy.html` n'a pas été modifiée (point 5 de la liste, non retenu) : elle devrait mentionner PeerJS cloud, le relais TURN Metered, `version.json` sur GitHub et les comptes stockés chez l'hôte.
