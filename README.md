# Vélocité

FPS dans le navigateur au mouvement nerveux : glissade avec boost, bhop, air-strafe, tap-strafe, escalade, wall-jump, superglide et jump pads. Quatre cartes (Canyon, Raffinerie, Forêt, Château), 5 armes, 6 couteaux et la hache double Morgane, trois personnages jouables avec couleurs, capacités (stim, dôme-bouclier, portail), skins d'arme (dont des skins animés), emotes, robots d'entraînement animés, mini-carte et multijoueur en pair-à-pair jusqu'à 8 joueurs. Interface en français, anglais, afrikaans et isiZulu.

**Jouer :** https://haze15.github.io/velocite/

## Mouvement

- **Glissade :** sprinte puis accroupis-toi. Un boost toutes les 2 s.
- **Bhop :** saute dans les 0,12 s qui suivent l'atterrissage. La vitesse perdue au sol est rendue et chaque saut enchaîné ajoute environ 0,4 m/s, jusqu'à 17 m/s.
- **Wall-jump :** en l'air contre un mur, saute pour rebondir en gardant ton élan le long du mur. Trois rebonds avant de retoucher le sol.
- **Superglide :** en te hissant sur un rebord, saute puis accroupis-toi presque en même temps (fenêtre de 0,14 s). Sortie à 18 m/s, plus si tu arrives lancé.
- **Son de glissade :** gravier, métal, souffle, électrique, aucun, ou ton propre fichier son (mémorisé jusqu'à 2 Mo).
- **Touches :** toutes les actions se règlent dans le menu, deux touches par action, clavier, souris ou molette.

## Personnages

Soldat, X-Bot ou Robot, avec dix-huit couleurs. Le choix se fait dans le menu avec un aperçu 3D, et c'est ce que les autres joueurs voient en ligne.

Les avatars ont une pose de saut (jambes repliées) et une pose de superglide (jambes tendues vers l'avant). Pendant un superglide, tu vois aussi tes jambes tendues à la première personne.

**Emotes :** appuie sur **J** puis sur **1 à 6** : salut, danse, pouce levé, bravo, assis, muscles. La caméra passe à la troisième personne pendant l'emote (la souris tourne autour du perso), et bouger, sauter ou tirer l'arrête. Les autres joueurs voient l'emote en ligne.

## Cartes

- **Canyon** : désert au coucher du soleil, grande pente de glisse, tour centrale, arène de robots.
- **Raffinerie** : zone industrielle au crépuscule, half-pipe, passerelles en hauteur reliées par un pont, couloir étroit pour enchaîner les wall-jumps, réservoirs à escalader.

- **Forêt** : pins, ruisseau, cabane avec toit accessible, mirador sur pilotis, colline de glisse de 9 m.
- **Château** : forteresse originale avec douves, pont-levis, cour intérieure, chemin de ronde crénelé, tours d'angle, donjon avec grand escalier de glisse et pad vers la tour haute.

La carte se choisit à côté du bouton Jouer. En ligne, c'est l'hôte qui choisit et les invités suivent automatiquement.

## Multijoueur

1. Un joueur clique sur **Créer une partie** : il héberge la partie et reçoit un code de 5 caractères.
2. Les autres entrent ce code et cliquent sur **Rejoindre**, ou ouvrent directement le lien partagé (`…/velocite/#CODE`).
3. Tout le monde clique sur **Entrer dans l'arène**.

La liste **Parties en cours** du menu affiche les parties publiques avec leur hôte, leur carte et le nombre de joueurs : un clic sur **Rejoindre** suffit. Une partie y apparaît tant que la case « Afficher ma partie dans la liste » est cochée chez l'hôte. L'annonce passe par un broker MQTT public (EMQX et HiveMQ) ; si ton réseau le bloque, la liste reste vide mais les codes marchent toujours.

Fermer l'onglet ou la fenêtre fait quitter la partie tout de suite (les autres voient ton départ), et le bouton **Quitter la partie** fait pareil sans fermer le jeu.

La connexion passe par WebRTC avec [PeerJS](https://peerjs.com/) et son serveur de mise en relation public. Il n'y a pas de serveur de jeu : l'hôte relaie positions, tirs et éliminations. Si l'hôte quitte, la partie se ferme.

Sur certains réseaux (4G, réseaux d'entreprise), la connexion directe entre deux joueurs peut échouer faute de serveur TURN.

## Commandes

Toutes les touches se changent dans le menu (Échap) : clavier, boutons de souris ou molette, deux touches par action. Touches par défaut :

| Touche | Action |
| --- | --- |
| ZQSD / WASD | Se déplacer (touches physiques, marche en AZERTY comme en QWERTY) |
| Maj | Sprint (automatique par défaut) |
| Espace | Sauter |
| C | S'accroupir, glisser |
| Molette ↓ | Avancer (tap-strafe) |
| Clic gauche / droit | Tirer / viser |
| 1 à 5 | Carabine, mitraillette, pompe, sniper, pistolet |
| 6 ou V | Couteau |
| Q / A en AZERTY | Stim : +35 % de vitesse pendant 6 s (recharge 9 s) |
| X | Inspecter l'arme ou le couteau (chaque couteau a sa manipulation) |
| J | Emotes (puis 1 à 6) |
| G | Dôme-bouclier : bulle de 5 m qui arrête les balles pendant 12 s (recharge 25 s) |
| F | Portail : pose l'entrée, va à au moins 6 m, réappuie pour ouvrir la sortie (40 s) |
| M | Mini-carte |
| R | Recharger |
| Tab | Tableau des scores en ligne |
| T | Réapparaître |
| P | Couper le son |

## Armes

| Arme | Dégâts | Cadence | Chargeur |
| --- | --- | --- | --- |
| Carabine K-31 | 14 (tête ×1,75) | 780 coups/min | 28 |
| Mitraillette VX | 11 (tête ×1,5) | 1080 coups/min | 25 |
| Pompe KS-8 | 8 × 11 | 72 coups/min | 6 |
| Épervier .50 (sniper, lunette ×4) | 85 (tête ×2,1) | 46 coups/min | 5 |
| Pistolet HX | 18 (tête ×1,6) | semi-auto | 14 |
| Couteau de combat | 50, portée 2,7 m | 115 coups/min | — |
| Karambit | 42, portée 2,4 m | 145 coups/min | — |
| Machette | 65, portée 3,3 m | 85 coups/min | — |
| Dague | 55, portée 2,6 m | 110 coups/min | — |
| Papillon | 46, portée 2,5 m | 135 coups/min | — |
| Katana | 70, portée 3,6 m | 80 coups/min | — |
| Morgane (hache double) | 85, portée 3,4 m | 62 coups/min | — |

Avec le couteau en main, tu cours 10 % plus vite. Vingt-quatre skins d'arme au choix dans les réglages : standard, or, camo désert, néon, carbone, arctique, rouge sang, jungle, chrome, lave, galaxie, tigre, bonbon, sakura, océan, toxique, royal, camo urbain, coucher de soleil, et cinq skins animés (plasma, arc-en-ciel, pulsation, circuit, flammes).

**Katana :** combo de trois coups (diagonale, remontée, horizontal) avec une traînée lumineuse, moulinet en sortant la lame et inspection dédiée. Il a ses propres skins : tamahagane, Damas, lame de sang, dragon, énergie, sakura, obsidienne, éclair et néant (dragon, énergie, éclair et néant sont animés). La traînée prend la couleur du skin.

**Morgane :** hache double à lames argent-violet et runes roses. Coup vertical puis diagonal, lents mais lourds. Un joueur éliminé à la Morgane réapparaît en tenue Morgane (rose, oreilles de chat, nœud et jupe plissée) jusqu'à sa prochaine mort ; les robots d'entraînement la gardent 90 s.

Chaque joueur a 75 de bouclier et 100 de vie. Le bouclier se recharge après 4 s sans dégâts.

## Appli Windows

`Velocite.exe` ouvre le jeu dans sa propre fenêtre, sans barre d'adresse, avec Chrome, Edge ou Brave (sinon dans le navigateur par défaut). Le jeu et ses modèles sont dans l'exécutable : le solo marche hors ligne, le multijoueur demande internet. Les liens de partie partagés depuis l'appli pointent vers la version en ligne. L'appli se ferme toute seule quand on ferme la fenêtre du jeu.

Pour la reconstruire : `launcher/build-exe.sh` (Go 1.21+, et [rsrc](https://github.com/akavel/rsrc) pour l'icône). Le résultat arrive dans `dist/Velocite.exe`.

## Lancer en local

Les modèles 3D sont chargés depuis `assets/`, il faut donc un petit serveur :

```sh
python3 -m http.server 8000
# puis http://localhost:8000
```

## Structure

- `index.html` : tout le jeu (rendu three.js r128, physique de mouvement à 120 Hz, armes, HUD, réseau).
- `assets/models/` : modèles de personnages animés.

## Crédits

- Robot : « RobotExpressive » de Tomás Laulhé ([Quaternius](https://quaternius.com/)), modifié par Don McCurdy, licence CC0 1.0.
- Soldat et X-Bot : modèles et animations [Mixamo](https://www.mixamo.com/) (Adobe), tels que fournis dans les exemples de three.js.
- Moteur : [three.js](https://threejs.org/) r128 (MIT), copié dans `launcher/vendor/` pour l'appli Windows. Réseau : [PeerJS](https://peerjs.com/) (MIT).
