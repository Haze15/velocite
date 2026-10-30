# Vélocité

FPS dans le navigateur au mouvement nerveux : glissade avec boost, bhop, air-strafe, tap-strafe, escalade, wall-jump, superglide et jump pads. Arène de canyon, 5 armes et un couteau, skins d'arme, robots d'entraînement animés, mini-carte et multijoueur en pair-à-pair jusqu'à 8 joueurs. Interface en français, anglais, afrikaans et isiZulu.

**Jouer :** https://haze15.github.io/velocite/

## Mouvement

- **Glissade :** sprinte puis accroupis-toi. Un boost toutes les 2 s.
- **Bhop :** saute dans les 0,12 s qui suivent l'atterrissage. La vitesse perdue au sol est rendue et chaque saut enchaîné ajoute environ 0,4 m/s, jusqu'à 17 m/s.
- **Wall-jump :** en l'air contre un mur, saute pour rebondir en gardant ton élan le long du mur. Trois rebonds avant de retoucher le sol.
- **Superglide :** en te hissant sur un rebord, saute puis accroupis-toi presque en même temps (fenêtre de 0,14 s). Sortie à 14,5 m/s, plus si tu arrives lancé.
- **Son de glissade :** gravier, métal, souffle, électrique, aucun, ou ton propre fichier son (mémorisé jusqu'à 2 Mo).
- **Touches :** toutes les actions se règlent dans le menu, deux touches par action, clavier, souris ou molette.

## Multijoueur

1. Un joueur clique sur **Créer une partie** : il héberge la partie et reçoit un code de 5 caractères.
2. Les autres entrent ce code et cliquent sur **Rejoindre**, ou ouvrent directement le lien partagé (`…/velocite/#CODE`).
3. Tout le monde clique sur **Entrer dans l'arène**.

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
| Couteau | 50 (tête ×1,5), portée 2,7 m | 115 coups/min | — |

Avec le couteau en main, tu cours 10 % plus vite. Six skins d'arme au choix dans les réglages : standard, or, camo désert, néon, carbone, arctique.

Chaque joueur a 75 de bouclier et 100 de vie. Le bouclier se recharge après 4 s sans dégâts.

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
- Soldat : modèle et animations [Mixamo](https://www.mixamo.com/) (Adobe), tels que fournis dans les exemples de three.js.
- Moteur : [three.js](https://threejs.org/) r128 (MIT). Réseau : [PeerJS](https://peerjs.com/) (MIT).
