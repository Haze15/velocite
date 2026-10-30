# Vélocité

FPS dans le navigateur au mouvement nerveux : glissade avec boost, bhop, air-strafe, tap-strafe, escalade, wall-jump, superglide et jump pads. Sept cartes (Canyon, Raffinerie, Forêt, Château, Volcan, Fusion, Ville fantôme), des buggys, tanks, hélicos et avions à piloter, 5 armes, 6 couteaux, la hache double Morgane et un trident qui se lance et revient, quatre personnages jouables avec chacun sa classe et ses capacités (stim, dôme, grappin, portail, scan, tremplin), skins d'arme (dont des skins animés), emotes, robots d'entraînement animés, mini-carte et multijoueur en pair-à-pair jusqu'à 8 joueurs. Interface en français, anglais, afrikaans et isiZulu.

**Jouer :** https://haze15.github.io/velocite/

## Mouvement

- **Glissade :** sprinte puis accroupis-toi. Un boost toutes les 2 s.
- **Bhop :** saute dans les 0,12 s qui suivent l'atterrissage. La vitesse perdue au sol est rendue et chaque saut enchaîné ajoute environ 0,4 m/s, jusqu'à 17 m/s.
- **Wall-jump :** en l'air contre un mur, saute pour rebondir en gardant ton élan le long du mur. Trois rebonds avant de retoucher le sol.
- **Mantle jump :** en te hissant sur un rebord, appuie seulement sur saut : tu es propulsé environ 7,5 m au-dessus du rebord.
- **Superglide :** en te hissant sur un rebord, saute puis accroupis-toi presque en même temps (fenêtre de 0,14 s). Sortie à 18 m/s, plus si tu arrives lancé.
- **Son de glissade :** gravier, métal, souffle, électrique, aucun, ou ton propre fichier son (mémorisé jusqu'à 2 Mo).
- **Touches :** toutes les actions se règlent dans le menu, deux touches par action, clavier, souris ou molette.

## Personnages et classes

Chaque personnage a une classe avec une capacité tactique (**Q**, A en AZERTY) et une ultime (**F**) :

| Personnage | Classe | Tactique (Q) | Ultime (F) |
| --- | --- | --- | --- |
| Soldat | Assaut | Stim : +35 % de vitesse 6 s (recharge 9 s) | Dôme-bouclier : bulle de 5 m qui arrête les balles 12 s (recharge 25 s) |
| X-Bot | Éclaireur | Grappin : accroche un point visé jusqu'à 45 m et te tire vers lui ; Espace pour lâcher avec un saut (recharge 6 s) | Portail : pose l'entrée, va à au moins 6 m, réappuie pour la sortie (40 s, recharge 30 s) |
| Robot | Sentinelle | Scan : révèle robots et joueurs à travers les murs et sur la mini-carte 7 s (recharge 18 s) | Tremplin : pad violet posé au sol que tout le monde peut utiliser, 30 s (recharge 22 s) |
| Spectre | Fantôme | Dash : propulsion éclair au sol ou en l'air dans la direction de tes touches (recharge 3,5 s) | Téléport : 50 m droit devant en gardant l'élan, s'arrête avant un mur (recharge 20 s) |


Soldat, X-Bot, Robot ou Spectre, avec dix-huit couleurs. Le Spectre porte une capuche, un masque blanc aux yeux lumineux, une cape déchirée, des épaulières et des brassards ; ses yeux, ses articulations et ses runes prennent la couleur choisie. Le choix se fait dans le menu avec un aperçu 3D, et c'est ce que les autres joueurs voient en ligne.

Les avatars ont une pose de saut (jambes repliées) et une pose de superglide (jambes tendues vers l'avant). Pendant un superglide, tu vois aussi tes jambes tendues à la première personne (case « Voir mes jambes » dans les réglages pour les enlever).

**Emotes :** appuie sur **J** puis sur **1 à 6** : salut, danse, pouce levé, bravo, assis, muscles. La caméra passe à la troisième personne pendant l'emote (la souris tourne autour du perso), et bouger, sauter ou tirer l'arrête. Les autres joueurs voient l'emote en ligne.

## Cartes

- **Canyon** : désert au coucher du soleil, grande pente de glisse, tour centrale, arène de robots.
- **Raffinerie** : zone industrielle au crépuscule, half-pipe, passerelles en hauteur reliées par un pont, couloir étroit pour enchaîner les wall-jumps, réservoirs à escalader.

- **Forêt** : pins, ruisseau, cabane avec toit accessible, mirador sur pilotis, colline de glisse de 9 m.
- **Château** : forteresse originale avec douves, pont-levis, cour intérieure, chemin de ronde crénelé, tours d'angle, donjon avec grand escalier de glisse et pad vers la tour haute.

- **Volcan** : carte immense (420 m de côté) avec un vrai relief : volcan de 74 m avec cratère de lave, trois coulées qui descendent ses flancs (la lave brûle), ponts, avant-postes dans les coins, camp de base au sud, plaine de cendres avec tremplins pour les buggys.
- **Fusion** : les quatre cartes collées en carré, sans murs entre elles, avec un carrefour et un pad au centre.
- **Ville fantôme** : grande ville abandonnée de 320 m de côté au crépuscule, dans un brouillard verdâtre plein de poussière. Rues en damier, immeubles à fenêtres cassées ou condamnées (jusqu'à une tour de 52 m avec un pad sur le toit), immeuble en chantier à six étages avec escaliers, autoroute surélevée avec bretelles, voitures abandonnées, place centrale, parc d'herbes hautes, tyroliennes entre les toits et aérodrome à l'est avec piste, hangar, tour de contrôle et héliport.

L'herbe haute de la Forêt, du Château, du Canyon et de la Ville fantôme ondule au vent.

Chaque carte a son décor : tonneaux, caisses, pneus, cactus et rochers sur le Canyon ; palettes, cônes, glissières en béton, tuyaux et lampadaires à la Raffinerie ; troncs couchés, souches, champignons, buissons et fleurs en Forêt ; bottes de foin, charrettes, torches, bannières et fleurs au Château ; rochers de basalte, cristaux brûlants et arbres calcinés sur le Volcan ; lampadaires, feux tricolores, bouches d'incendie, poubelles, bennes, gravats et barricades dans la Ville fantôme. Les gros objets bloquent le passage et arrêtent les balles : ils servent de couverture.

La carte se choisit à côté du bouton Jouer. En ligne, c'est l'hôte qui choisit et les invités suivent automatiquement.

## Tyroliennes

Des câbles ambre lumineux relient les toits et les tours sur toutes les cartes (lampes ambre aux deux bouts, traits ambre sur la mini-carte), dont une verticale sur le Canyon et cinq géantes sur le Volcan qui montent jusqu'au cratère. Approche-toi d'un câble et appuie sur **E** : tu pars dans le sens où tu regardes (vers le haut sur un câble vertical) et tu montes jusqu'à 17 m/s. **S** fait demi-tour, **Espace** saute en gardant l'élan (double saut disponible ensuite), **C** lâche. Tu peux tirer pendant le trajet.

## Véhicules

Sur le Canyon, le Volcan, la Fusion et la Ville fantôme, des buggys attendent. Approche-toi et appuie sur **E** pour monter : Z/W accélère, S freine puis recule, Q/D (A/D) tourne, **Espace** est le frein à main pour déraper, **Maj** donne un petit turbo. La caméra passe derrière le buggy et la souris tourne autour. Les tremplins et les crêtes font décoller. Foncer dans un robot ou un joueur à plus de 9 m/s lui fait des dégâts. **E** pour descendre. En ligne, les autres te voient assis au volant.

| Véhicule | Où | Commandes | Arme |
| --- | --- | --- | --- |
| Tank | Ville fantôme (2), Volcan, Fusion | Z/S avance et recule, Q/D pivote sur place ; la tourelle et le canon suivent la souris | Obus au clic gauche (130 dégâts en zone, rechargement 2,2 s). Le blindage divise par 2,5 les dégâts reçus |
| Hélico | Ville fantôme (2), Volcan (2), Fusion | Il se tourne vers où tu regardes ; Z/Q/S/D le déplacent, **Espace** monte, **C** descend, **Maj** accélère | Mitrailleuse au clic gauche |
| Avion | Ville fantôme (2, sur la piste), Volcan | Z/S règlent les gaz, la souris donne la direction, Q/D penchent ; il décolle en tirant vers le haut au-dessus de 30 m/s | Mitrailleuse au clic gauche |
| Chasseur | Ville fantôme (sur la piste), Volcan | Comme l'avion, mais bien plus rapide (jusqu'à 112 m/s avec Maj), décolle au-dessus de 40 m/s, postcombustion visible | Deux canons qui tirent en alternance |

**Réticule :** comme dans Battlefield, le petit rond au centre de l'écran montre où tu regardes et le grand réticule montre où partent vraiment tes tirs. Celui de l'hélico, de l'avion et du chasseur est devant le nez de l'appareil. Celui du tank est au bout du canon : il tient compte de la chute de l'obus et se remplit pendant le rechargement. Quand tu tournes la souris, le réticule rattrape la visée avec un temps de retard : la tourelle, le canon de l'hélico et le nez de l'avion pivotent moins vite que la caméra.

Les modèles sont détaillés : tank camouflé avec chenilles qui défilent, tourelle, coupole, mitrailleuse et antenne ; hélico d'attaque avec cockpit en tandem, nacelles de roquettes et feux de position ; avion à hélice avec bandes et cocardes ; chasseur à aile delta ; buggy avec roue de secours, rampe de phares et fanion.

**Véhicules destructibles :** chaque véhicule a son blindage, affiché en bas de l'écran quand tu le conduis.

| Véhicule | Blindage | Balles et lames | Explosions |
| --- | --- | --- | --- |
| Buggy | 260 | 80 % | ×2 |
| Tank | 1000 | 5 % (presque insensible) | ×1,6 |
| Hélico | 500 | 50 % | ×2 |
| Avion | 320 | 60 % | ×2 |
| Chasseur | 380 | 50 % | ×2 |

Grenades, obus du tank, trident, couteaux, armes et mitrailleuses des véhicules les abîment, tout comme les chocs contre les murs, les atterrissages trop durs et les collisions entre véhicules (le tank écrase les buggys). Abîmé, un véhicule fume ; sous 20 %, la fumée devient noire et il crache des étincelles. À zéro, il explose. Le conducteur est éjecté et prend 140 dégâts, et l'élimination revient à celui qui l'a détruit. L'explosion blesse aussi autour, ce qui peut faire sauter d'autres véhicules. L'épave noircie brûle sur place (un appareil touché en vol s'écrase d'abord), puis le véhicule réapparaît à son point de départ 14 s plus tard. En ligne, les dégâts sont partagés entre tous les joueurs.

Un avion qui pique dans le sol ou percute un immeuble explose (tu es éjecté et blessé) et réapparaît à sa place 8 s plus tard. Près du bord de la carte, l'avion fait demi-tour tout seul. Un hélico abandonné en vol redescend doucement en autorotation.

## Multijoueur

1. Un joueur clique sur **Créer une partie** : il héberge la partie et reçoit un code de 5 caractères.
2. Les autres entrent ce code et cliquent sur **Rejoindre**, ou ouvrent directement le lien partagé (`…/velocite/#CODE`).
3. Tout le monde clique sur **Entrer dans l'arène**.

La liste **Parties en cours** du menu affiche les parties publiques avec leur hôte, leur carte et le nombre de joueurs : un clic sur **Rejoindre** suffit. Une partie y apparaît tant que la case « Afficher ma partie dans la liste » est cochée chez l'hôte. L'annonce passe par un broker MQTT public (EMQX et HiveMQ) ; si ton réseau le bloque, la liste reste vide mais les codes marchent toujours.

Fermer l'onglet ou la fenêtre fait quitter la partie tout de suite (les autres voient ton départ), et le bouton **Quitter la partie** fait pareil sans fermer le jeu.

La connexion passe par WebRTC avec [PeerJS](https://peerjs.com/) et son serveur de mise en relation public. Il n'y a pas de serveur de jeu : l'hôte relaie positions, tirs et éliminations. Si l'hôte quitte, la partie se ferme.

Sur certains réseaux (4G, réseaux d'entreprise), la connexion directe entre deux joueurs peut échouer faute de serveur TURN.

## Grenades

**G** lance la grenade choisie, **H** change de type. Deux grenades de chaque type, une revient toutes les 12 s (toutes à la réapparition) :

| Grenade | Effet |
| --- | --- |
| Frag | rebondit et explose après 2,4 s : jusqu'à 115 dégâts dans 7 m, souffle |
| Thermite | s'enflamme au contact : zone de feu de 3,5 m, 25 dégâts/s pendant 7 s |
| Arc | se colle à la première surface et électrocute 1,2 s plus tard : 72 dégâts dans 4 m |
| Fumée | nuage épais pendant 13 s |
| Impulsion | explose au contact et propulse tout le monde sans dégâts : à tes pieds, c'est un saut géant |

Les grenades ne traversent pas les murs (pas de dégâts derrière un obstacle) et tes propres grenades te blessent un peu. En ligne, tout le monde voit les lancers et les explosions.

## Manette

Une manette Xbox, PlayStation ou compatible se branche à tout moment (appuie sur un bouton pour qu'elle soit détectée, Start ou A pour lancer depuis le menu). Stick gauche : bouger · stick droit : viser · A : sauter · B : s'accroupir, glisser · X : recharger ou interagir (buggy, tyrolienne) · Y : arme suivante · LB / RB : capacités tactique et ultime · LT / RT : viser / tirer · L3 : sprint · R3 : couteau · croix : emotes, inspecter, arme précédente, mini-carte · Start : pause. Vibrations au tir.

Réglages dédiés : sensibilité, **aide à la visée** (ralentissement près d'une cible et légère traction quand tu bouges ou vises, de 0 à 100 %), **snap strafe facile** en l'air et inversion verticale. Au clavier, le **snap tap** est activé par défaut : si deux directions opposées sont enfoncées, la dernière appuyée gagne.

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
| Q / A en AZERTY | Capacité tactique de ta classe |
| X | Inspecter l'arme ou le couteau (chaque couteau a sa manipulation) |
| J | Emotes (puis 1 à 6) |
| E | S'accrocher à une tyrolienne, monter dans un véhicule / en descendre |
| G / H | Lancer une grenade / changer de type |
| F | Capacité ultime de ta classe |
| M | Mini-carte |
| R | Recharger |
| Tab | Tableau des scores en ligne |
| T | Réapparaître |
| P | Couper le son |

## Armes

| Arme | Dégâts | Cadence | Chargeur |
| --- | --- | --- | --- |
| Carabine K-31 | 14 (tête ×1,75) | 780 coups/min | 28 |
| Mitraillette VX | 9 (tête ×1,35) | 950 coups/min | 24 |
| Pompe KS-8 | 9 × 13 | 82 coups/min | 6 |
| Épervier .50 (sniper, lunette ×4) | 85 (tête ×2,1) | 46 coups/min | 5 |
| Pistolet HX | 18 (tête ×1,6) | semi-auto | 14 |
| Couteau de combat | 50, portée 2,7 m | 115 coups/min | — |
| Karambit | 42, portée 2,4 m | 145 coups/min | — |
| Machette | 65, portée 3,3 m | 85 coups/min | — |
| Dague | 55, portée 2,6 m | 110 coups/min | — |
| Papillon | 46, portée 2,5 m | 135 coups/min | — |
| Katana | 70, portée 3,6 m | 80 coups/min | — |
| Morgane (hache double) | 85, portée 3,4 m | 62 coups/min | — |
| Trident | 70 au corps à corps, 95 lancé (tête ×1,4) | 80 coups/min | — |

Avec le couteau en main, tu cours 10 % plus vite. Vingt-quatre skins d'arme au choix dans les réglages : standard, or, camo désert, néon, carbone, arctique, rouge sang, jungle, chrome, lave, galaxie, tigre, bonbon, sakura, océan, toxique, royal, camo urbain, coucher de soleil, et cinq skins animés (plasma, arc-en-ciel, pulsation, circuit, flammes).

**Katana :** combo de trois coups (diagonale, remontée, horizontal) avec une traînée lumineuse, moulinet en sortant la lame et inspection dédiée. Il a ses propres skins : tamahagane, Damas, lame de sang, dragon, énergie, sakura, obsidienne, éclair et néant (dragon, énergie, éclair et néant sont animés). La traînée prend la couleur du skin.

**Morgane :** hache double à lames argent-violet et runes roses. Coup vertical puis diagonal, lents mais lourds. Un joueur éliminé à la Morgane réapparaît en tenue Morgane (rose, oreilles de chat, nœud et jupe plissée) jusqu'à sa prochaine mort ; les robots d'entraînement la gardent 90 s.

**Trident :** clic gauche pour frapper d'estoc, clic droit pour le lancer. Il file droit devant à 48 m/s et touche tout ce qu'il traverse ; s'il touche un mur, il s'y plante une fraction de seconde, puis il revient vers toi en accélérant et touche encore sur le chemin du retour. Tu le rattrapes automatiquement. En ligne, les autres le voient voler.

**Sons personnalisés :** dans les réglages, tu peux remplacer le son de chaque arme, du couteau, du saut, de l'atterrissage, des éliminations, des touches, du superglide et des capacités par ton propre fichier (300 Ko et 4 s maximum). En ligne, les autres joueurs entendent tes sons quand c'est toi qui tires ou qui utilises la capacité (chacun peut couper ça dans ses réglages).

**Bras et armes :** tu vois tes bras (gants et manches à ta couleur) tenir chaque arme à deux mains quand c'est possible. Les autres joueurs portent leur vraie arme en main, avec leur skin. Quand un joueur vise au sniper dans ta direction, le reflet de sa lunette brille, même de loin.

**Musique :** « Celestial Butterfly » tourne en fond dans le menu et un peu plus bas en jeu. Curseur « Musique » dans les réglages (0 pour la couper, P coupe tout le son). Le bouton « Importer ma musique » remplace la musique par ton propre fichier, gardé seulement dans ton navigateur.

**Effets de caméra :** légère secousse au tir et aux atterrissages lourds, coup de champ de vision au superglide et sur les pads, balancement de la marche. Un curseur dans les réglages les règle de 0 (aucun) à 150 %. L'arme en main bouge beaucoup moins qu'avant.

Les bruitages des armes ont été refaits pour être moins agressifs : un coup grave, un claquement court, une légère réverbération et une petite variation de hauteur à chaque tir.

Chaque joueur a 75 de bouclier et 100 de vie. Le bouclier se recharge après 4 s sans dégâts.

## Appli Windows

`Velocite.exe` ouvre le jeu dans sa propre fenêtre (agrandie), grâce à WebView2, le moteur d'Edge déjà présent sur Windows 10 et 11. Comme la fenêtre appartient à `Velocite.exe`, Discord peut la détecter : Paramètres Discord → Jeux enregistrés → « Ajoute-le ! » → Velocite. Si WebView2 manque, ou si on lance `Velocite.exe -navigateur`, le jeu s'ouvre dans une fenêtre d'application Chrome, Edge ou Brave (sinon dans le navigateur par défaut). Le jeu et ses modèles sont dans l'exécutable : le solo marche hors ligne, le multijoueur demande internet. Les liens de partie partagés depuis l'appli pointent vers la version en ligne. L'appli se ferme toute seule quand on ferme la fenêtre du jeu.

### Statut Discord

Dans l'appli, Discord peut afficher ce que tu fais : carte, solo ou en ligne avec le nombre de joueurs, arme ou buggy, éliminations, temps de jeu, et un bouton « Rejoindre la partie » qui ouvre le lien de ta partie. Un site web ne peut pas parler à Discord, donc ça ne marche que dans `Velocite.exe`, avec Discord ouvert sur le même PC.

Discord demande une « application » pour afficher le nom du jeu. À faire une seule fois :

1. Va sur https://discord.com/developers/applications et clique sur **New Application**. Le nom que tu choisis (par exemple « Vélocité ») est celui qui s'affichera : « Joue à Vélocité ».
2. Copie l'**Application ID** (page General Information).
3. Facultatif : dans **Rich Presence → Art Assets**, ajoute une image (512×512 minimum) nommée exactement `logo` pour avoir l'icône dans le statut. `launcher/discord-logo.png` est prévue pour ça.
4. Dans le jeu, menu Réglages → **Statut Discord** : coche la case et colle l'ID. Le message passe à « Connecté » quand Discord reçoit le statut.

Dans Discord, le partage d'activité doit être activé (Paramètres → Confidentialité de l'activité → « Partager mon activité »). On peut aussi mettre l'ID dans un fichier `discord-id.txt` à côté de `Velocite.exe`.

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
- Musique de fond : « Celestial Butterfly », téléchargée sur [Pixabay](https://pixabay.com/music/) (licence de contenu Pixabay, utilisation libre dans un jeu).
- Moteur : [three.js](https://threejs.org/) r128 (MIT), copié dans `launcher/vendor/` pour l'appli Windows. Réseau : [PeerJS](https://peerjs.com/) (MIT).
- Fenêtre de l'appli : [go-webview2](https://github.com/jchv/go-webview2) (MIT) et [go-winloader](https://github.com/jchv/go-winloader) (ISC) de John Chadwick, [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) (BSD), copiés dans `launcher/third_party/`.
