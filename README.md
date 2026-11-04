# Vélocité

FPS dans le navigateur au mouvement nerveux : glissade avec boost, bhop, air-strafe, tap-strafe, escalade, wall-jump, superglide et jump pads. Dix cartes (Canyon, Raffinerie, Forêt, Château, Volcan, Fusion, Ville fantôme, Océan, Arctique et Monde qui les réunit toutes), des buggys, tanks, hélicos, avions, chasseurs et bateaux à piloter, des armes spéciales à ramasser au sol, un mode spectateur, 5 armes, 6 couteaux, la hache double Morgane et un trident qui se lance et revient, quatre personnages jouables avec chacun sa classe et ses capacités (stim, dôme, grappin, portail, scan, tremplin), skins d'arme (dont des skins animés), emotes, robots d'entraînement animés, mini-carte et multijoueur en pair-à-pair jusqu'à 8 joueurs. Interface en français, anglais, afrikaans et isiZulu.

**Jouer :** https://haze15.github.io/velocite/

## Mouvement

- **Glissade :** sprinte puis accroupis-toi. Un boost toutes les 2 s.
- **Bhop :** saute dans les 0,12 s qui suivent l'atterrissage. La vitesse perdue au sol est rendue et chaque saut enchaîné ajoute environ 0,4 m/s, sans limite de vitesse.
- **Snap strafe facile :** activé pour tout le monde, au clavier comme à la manette (désactivable dans les réglages).
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
| Nova | Sentinelle | Scan : révèle robots et joueurs à travers les murs et sur la mini-carte 7 s (recharge 18 s) | Tremplin : pad violet posé au sol que tout le monde peut utiliser, 30 s (recharge 22 s) |
| Titan | Assaut lourd | Jetpack : bond propulsé puis vol stationnaire en maintenant Espace (recharge 7 s) | Frappe aérienne : marque le point visé, 8 obus tombent du ciel 2 s plus tard (recharge 40 s) |
| Spectre | Fantôme | Dash : propulsion éclair au sol ou en l'air dans la direction de tes touches (recharge 3,5 s) | Téléport : 50 m droit devant en gardant l'élan, s'arrête avant un mur (recharge 20 s) |
| Lynx | Coureuse de murs | Dash (recharge 3,5 s). Passif : **course sur les murs** — en l'air, en avançant avec un mur sur le côté, elle court dessus jusqu'à 1,7 s ; Espace la propulse loin du mur | Survoltage : 8 s à +30 % de vitesse, course sur les murs sans limite de temps (recharge 45 s) |
| Corbeau | Voltigeur | Double grappin : deux câbles tirés de part et d'autre de la visée (70 m), tu te balances entre eux ; **Avancer** = jet de gaz dans la direction regardée, **Espace** = lâcher, Q de nouveau = lâcher avec un petit saut. Le gaz se recharge (vite au sol) (recharge 1,2 s) | Spirale : ruée en vrille de 1,3 s qui tranche tout autour de toi (40 dégâts par passage, 3 passages max par cible) (recharge 40 s) |


Soldat, X-Bot, Nova, Spectre, Titan, Lynx ou Corbeau, avec dix-huit couleurs. Le Corbeau porte un harnais croisé, deux bonbonnes de gaz aux hanches et une cape courte. Lynx porte un bandeau lumineux et une écharpe aux couleurs choisies. Le Spectre porte une capuche, un masque blanc aux yeux lumineux, une cape déchirée, des épaulières et des brassards ; ses yeux, ses articulations et ses runes prennent la couleur choisie. Le choix se fait dans le menu avec un aperçu 3D, et c'est ce que les autres joueurs voient en ligne.

Les autres joueurs tiennent leur arme à deux mains devant eux (le couteau d'une main), et leur buste suit leur visée vers le haut ou le bas. À la première personne, les bras et les mains (doigts refermés sur l'arme) ont les couleurs de ton perso : gants et manches du Soldat, peau et bracelet de Nova, os du squelette, bracelets lumineux du Spectre et de Lynx… Les animations ont été refaites : balancement de marche, respiration, sprint avec l'arme en travers, atterrissage, rechargement en trois temps (la main gauche va chercher le chargeur), changement d'arme.

Les avatars ont une pose de saut (jambes repliées) et une pose de superglide (jambes tendues vers l'avant). Pendant un superglide, tu vois aussi tes jambes tendues à la première personne (case « Voir mes jambes » dans les réglages pour les enlever).

**Emotes :** appuie sur **J** puis sur **1 à 6** : salut, danse, pouce levé, bravo, assis, muscles. La caméra passe à la troisième personne pendant l'emote (la souris tourne autour du perso), et bouger, sauter ou tirer l'arrête. Les autres joueurs voient l'emote en ligne.

## Cartes

- **Canyon** : désert au coucher du soleil, grande pente de glisse, tour centrale, arène de robots.
- **Raffinerie** : zone industrielle au crépuscule, half-pipe, passerelles en hauteur reliées par un pont, couloir étroit pour enchaîner les wall-jumps, réservoirs à escalader.

- **Forêt** : pins, ruisseau, cabane avec toit accessible, mirador sur pilotis, colline de glisse de 9 m.
- **Château** : forteresse originale avec douves, pont-levis, cour intérieure, chemin de ronde crénelé, tours d'angle, donjon avec grand escalier de glisse et pad vers la tour haute.

- **Volcan** : carte immense (420 m de côté) avec un vrai relief : volcan de 74 m avec cratère de lave, trois coulées qui descendent ses flancs (la lave brûle), ponts, avant-postes dans les coins, camp de base au sud, plaine de cendres avec tremplins pour les buggys.
- **Fusion** : les quatre cartes collées en carré, sans murs entre elles, avec un carrefour et un pad au centre.
- **Océan** : six îles tropicales (palmiers, plages, cabanes, phare), pontons, deux ponts, une épave et une plateforme pétrolière avec héliport et tyrolienne. Dans l'eau profonde, on nage (lentement, sans glissade). Bateaux, chasseur et avion sur la piste de l'île principale, hélicos.
- **Arctique** : collines enneigées sous la neige qui tombe, lac gelé qui sert de piste, base polaire avec tour et pad, chalets, tremplin de saut, forêt de sapins enneigés. Tanks, buggys, hélico, avion.
- **Monde** : toutes les cartes réunies en une seule (Fusion au centre, Ville fantôme à l'est, Volcan à l'ouest, Océan au nord, Arctique au sud), reliées par une route en boucle et des routes vers chaque zone, avec lampadaires, pads et véhicules en plus.
- **Chambre géante** : tu fais la taille d'un jouet dans une chambre d'enfant de 160 m. Lit (on y monte par un toboggan) avec oreillers et nounours, bureau avec ordinateur allumé, lampe et crayons (une règle sert de rampe), chaise, étagère à livres où l'on grimpe planche par planche, cubes alphabet empilés, dominos, petite voiture, ballon, briques à tenons, trampolines et deux tyroliennes (ficelles) depuis l'étagère et le bureau.
- **Mine** : grande caverne sous un plafond rocheux de 40 m (idéal pour le double grappin) : pilier central entouré d'un échafaudage avec deux rampes, passerelles en bois le long des parois, tyroliennes, rails et wagonnets, cristaux de minerai qui brillent, stalactites ; tout autour, quatre galeries basses étayées de poutres, reliées à la caverne par des portes.
- **Ville fantôme** : grande ville abandonnée de 320 m de côté au crépuscule, dans un brouillard verdâtre plein de poussière. Rues en damier, immeubles à fenêtres cassées ou condamnées (jusqu'à une tour de 52 m avec un pad sur le toit), immeuble en chantier à six étages avec escaliers, autoroute surélevée avec bretelles, voitures abandonnées, place centrale, parc d'herbes hautes, tyroliennes entre les toits et aérodrome à l'est avec piste, hangar, tour de contrôle et héliport.

L'herbe haute de la Forêt, du Château, du Canyon et de la Ville fantôme ondule au vent.

Chaque carte a son décor : tonneaux, caisses, pneus, cactus et rochers sur le Canyon ; palettes, cônes, glissières en béton, tuyaux et lampadaires à la Raffinerie ; troncs couchés, souches, champignons, buissons et fleurs en Forêt ; bottes de foin, charrettes, torches, bannières et fleurs au Château ; rochers de basalte, cristaux brûlants et arbres calcinés sur le Volcan ; lampadaires, feux tricolores, bouches d'incendie, poubelles, bennes, gravats et barricades dans la Ville fantôme. Les gros objets bloquent le passage et arrêtent les balles : ils servent de couverture.

La carte se choisit à côté du bouton Jouer. En ligne, c'est l'hôte qui choisit et les invités suivent automatiquement.

## Équipement

En cliquant sur **Jouer**, l'écran **Équipement** s'ouvre : choisis ton personnage et ton couteau, avec la description de leurs capacités. Tu peux le rouvrir à tout moment en partie avec **L** (bouton ⚙ sur téléphone) pour changer de perso ou de couteau sans attendre de mourir. La case « Demander avant chaque partie » le désactive.

## Tyroliennes

Des câbles ambre lumineux relient les toits et les tours sur toutes les cartes (lampes ambre aux deux bouts, traits ambre sur la mini-carte), dont une verticale sur le Canyon et cinq géantes sur le Volcan qui montent jusqu'au cratère. Approche-toi d'un câble et appuie sur **E** : tu pars dans le sens où tu regardes (vers le haut sur un câble vertical) et tu montes jusqu'à 17 m/s. **S** fait demi-tour, **Espace** saute en gardant l'élan (double saut disponible ensuite), **C** lâche. Tu peux tirer pendant le trajet.

## Coffres

Par défaut, tu commences chaque vie avec seulement le **pistolet**. Les autres armes, les couteaux et les armes spéciales se trouvent dans des **coffres** dorés qui brillent un peu partout sur la carte. Approche-toi, **E** pour l'ouvrir : une ou deux armes jaillissent devant, **E** encore pour les ramasser. Le coffre se referme et se remplit 75 s plus tard. Tu perds tout à la mort. En ligne, tout le monde voit les mêmes coffres s'ouvrir. La case « Départ au pistolet, armes dans les coffres » des réglages remet l'arsenal complet.

## Armes spéciales

Les armes spéciales sortent des coffres et se rangent dans l'emplacement **7**. Munitions limitées, pas de rechargement : vide, elle disparaît. Si tu meurs, elle tombe par terre et quelqu'un d'autre peut la prendre. Une arme ramassée réapparaît 30 s plus tard.

| Arme | Particularité |
| --- | --- |
| Lame-fusil | Clic gauche : coup de lame (80). Clic droit : décharge de chevrotine (8 cartouches) |
| Minigun | Doit d'abord tourner une demi-seconde, puis 1500 coups/min, 240 balles |
| Lance-roquettes | 4 roquettes, 125 dégâts en zone, rocket jump possible |
| Railgun | 125 dégâts, traverse tous les ennemis alignés jusqu'au premier mur |
| Revolver .44 | 48 dégâts, tête ×2 |
| Lance-grenades | Grenades de 40 mm qui explosent à l'impact (100 en zone) |
| Fusil rafale | Rafales de 3 balles, 24 dégâts |
| Arbalète | Carreaux de 110 dégâts qui retombent avec la distance |
| Lance-flammes | Courte portée (14 m), arrose devant toi |
| Fusil laser | Rayon automatique, 11 dégâts, très précis |
| Pompe automatique | 8 plombs, 240 coups/min |
| Sniper lourd | 140 dégâts, tête ×2, lunette ×4,5 |
| Mitrailleuse lourde | 100 balles, 700 coups/min |

L'onglet **Armes** du menu montre chaque arme en 3D avec ton skin (fais-la tourner à la souris).

## Mode spectateur

Le bouton **Spectateur** du menu te fait regarder la partie sans jouer : clic gauche / droit pour passer d'un joueur à l'autre (ou d'un robot à l'autre en solo), **E** pour une caméra libre (ZQSD, Espace, C). Les autres ne te voient pas. **Jouer** te fait revenir dans la partie.

## Véhicules

**À plusieurs :** quand quelqu'un conduit, approche-toi et appuie sur la touche d'action (E) pour monter en passager. Tu peux viser et tirer depuis ta place (sauf dans l'avion et le chasseur, où le passager est à l'intérieur). Places passagers : buggy 1, tank 2 (sur la plage arrière), hélico 3, bateau 2, avion 1, chasseur 1. Si le véhicule explose, les passagers sont éjectés et blessés comme le conducteur.

Sur le Canyon, le Volcan, la Fusion et la Ville fantôme, des buggys attendent. Approche-toi et appuie sur **E** pour monter : Z/W accélère, S freine puis recule, Q/D (A/D) tourne, **Espace** est le frein à main pour déraper, **Maj** donne un petit turbo. La caméra passe derrière le buggy et la souris tourne autour. Les tremplins et les crêtes font décoller. Foncer dans un robot ou un joueur à plus de 9 m/s lui fait des dégâts. **E** pour descendre. En ligne, les autres te voient assis au volant.

| Véhicule | Où | Commandes | Arme |
| --- | --- | --- | --- |
| Tank | Ville fantôme (2), Volcan, Fusion | Z/S avance et recule, Q/D pivote sur place ; la tourelle et le canon suivent la souris | Obus au clic gauche (130 dégâts en zone, rechargement 2,2 s). Le blindage divise par 2,5 les dégâts reçus |
| Hélico | Ville fantôme (2), Volcan (2), Fusion | Il se tourne vers où tu regardes ; Z/Q/S/D le déplacent, **Espace** monte, **C** descend, **Maj** accélère | Mitrailleuse au clic gauche |
| Avion | Ville fantôme (2, sur la piste), Volcan | Z/S règlent les gaz, la souris donne la direction, Q/D penchent ; il décolle en tirant vers le haut au-dessus de 30 m/s | Mitrailleuse au clic gauche |
| Chasseur | Ville fantôme (sur la piste), Volcan | Comme l'avion, mais bien plus rapide (jusqu'à 112 m/s avec Maj), décolle au-dessus de 40 m/s, postcombustion visible | Deux canons qui tirent en alternance |

**Vue cockpit :** en véhicule, **V** passe à la première personne avec un HUD épuré : vitesse et altitude encadrées, cap gradué, horizon et barre des gaz, sans carlingue devant les yeux. Dans le tank, c'est un viseur circulaire. L'hélico se tourne maintenant directement vers la souris.

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

## Modes de jeu

Le mode se choisit dans le menu (à côté de la carte) ou dans le panneau admin. En ligne, c'est l'hôte qui décide.

- **Chacun pour soi :** le mode classique.
- **2 équipes :** bleus contre rouges, répartis automatiquement. Pas de tir ami. Le score de chaque équipe s'affiche en haut de l'écran.
- **Zombies :** un joueur tiré au sort devient le premier zombie. Les zombies n'ont que le couteau (dégâts doublés, 150 PV, un peu plus rapides) et ne peuvent ni ouvrir les coffres ni ramasser d'armes. Chaque humain tué par un zombie revient en zombie. Les humains gagnent s'ils tiennent 3 minutes. Une nouvelle manche démarre 7 s après la fin. En solo, ce sont les robots qui jouent les zombies et te foncent dessus.

## Comptes et signalements

- **Mot de passe :** dans l'onglet Multijoueur, à côté du pseudo. Le premier mot de passe tapé avec un pseudo crée le compte chez l'hôte ; ensuite il faut le même pour rejoindre avec ce pseudo. Sans mot de passe, on peut toujours jouer avec un pseudo libre (sauf si l'hôte rend le mot de passe obligatoire).
- **Panneau admin, onglet Comptes :** liste des comptes, dernière connexion, réinitialiser ou définir un mot de passe, bannir ou débannir, supprimer. Les mots de passe sont stockés sous forme d'empreinte SHA-256 salée : personne ne peut les lire, on peut seulement les remplacer. Après une réinitialisation, le prochain mot de passe tapé par le joueur devient le sien.
- **Signaler :** en ligne, l'onglet Multijoueur propose de signaler un joueur (triche, insultes, spam, inactif, autre). Les signalements arrivent chez l'hôte et chez les admins, dans l'onglet **Signalements** du panneau admin, avec un bouton pour expulser ou bannir.

## Skins réservés

- **Squelette** : le Soldat du propriétaire.
- **Titan admin** : le Titan devient « Gasmask girl » (capuche, masque à gaz), animée avec les mouvements du X-Bot.
- **Spectre admin** : carrosserie pastel rose et lavande, oreilles de chat, nœud, yeux brillants et petites ailes.
- **X-Bot admin** : quand tu as l'accès admin (mot de passe ou accès donné par l'hôte), ton X-Bot devient noir laqué avec articulations en or, auréole, visière rouge et insignes dorés. Les autres le voient en ligne.

Sur mobile, le mot de passe admin débloque aussi l'onglet Triche du panneau.

## Mon propre modèle 3D

Onglet Personnage → **Importer mon modèle (.glb)**, 4 Mo maximum. Un modèle avec un squelette Mixamo reprend les animations du jeu ; sinon il utilise ses propres animations nommées idle, walk et run. Une case permet de le tourner de 180° s'il regarde dans le mauvais sens. Le modèle est mémorisé dans le navigateur et envoyé aux autres joueurs de la partie, qui le voient à ta place.

## Mises à jour

Quand une nouvelle version est en ligne, un bandeau apparaît en haut de l'écran (vérification toutes les 3 minutes) avec un bouton pour recharger (dans l'appli Windows, le bouton télécharge la nouvelle appli). Les nouveautés s'affichent une fois à chaque nouvelle version, et le bouton **Nouveautés** du menu les remontre.

## Sur mobile

Le **panneau admin** s'ouvre aussi au doigt : bouton « Admin » dans le menu (demande le mot de passe) et bouton ADMIN en jeu quand tu as l'accès.


Le jeu se joue aussi sur téléphone, en mode paysage. Les commandes tactiles s'activent toutes seules sur un écran tactile (réglage « Commandes tactiles »).

- **Se déplacer :** pose le pouce n'importe où sur la moitié gauche de l'écran, un joystick apparaît. Pousse-le à fond vers l'avant pour sprinter.
- **Viser :** glisse le doigt sur la moitié droite. Le gros bouton TIR vise aussi quand tu glisses dessus ; un second bouton TIR est à gauche.
- **Boutons :** saut, accroupi (glissade), viser (appui = on/off), recharger, prendre / monter dans un véhicule, capacités et grenade en haut à droite, armes en bas au milieu, pause en haut à gauche.
- **Sensibilité tactile** et **inversion** dans les réglages.
- **iPhone :** Safari ne permet pas le plein écran dans une page. Bouton **Partager** → **Sur l'écran d'accueil**, puis lance Vélocité depuis l'icône : il s'ouvre en plein écran, sans barre Safari. Sur Android, le jeu passe en plein écran au lancement.

- **Disposer les boutons** (réglages → « Disposer les boutons tactiles ») : glisse chaque bouton où tu veux, change sa taille et l'action qu'il fait (tir, saut, sprint, viser maintenu, couteau, grenade, chat, micro…). « Par défaut » remet tout en place.
- Le viseur reste visible au centre quand tu vises (le point reste affiché).

La souris peut aussi être inversée (haut / bas et gauche / droite) dans les réglages.

## Viseur

Dans les réglages : forme (croix, croix en T, croix en X, cercle, point), couleur, longueur, épaisseur, écart, opacité, point au centre, contour noir, et s'il s'écarte avec la dispersion de l'arme. L'aperçu se met à jour en direct.

## Chat écrit et vocal de proximité

- **Écrit :** **Entrée** (ou **Y**) ouvre la barre de chat. Par défaut le message n'arrive qu'aux joueurs à moins de 40 m ; **Tab** (ou le bouton à gauche de la barre) bascule entre « Proche » et « Tous ».
- **Vocal :** réglages → Micro : « Appuyer pour parler » (touche **N**) ou « Toujours ouvert ». On entend les joueurs à moins de 45 m, de plus en plus fort en s'approchant, à gauche ou à droite selon leur position. Une icône 🔊 s'affiche à côté du nom de celui qui parle. La voix passe directement entre les joueurs (WebRTC), le navigateur demande l'autorisation du micro.
- Sur mobile : boutons CHAT et MICRO.

## Multijoueur

1. Un joueur clique sur **Créer une partie** : il héberge la partie et reçoit un code de 5 caractères.
2. Les autres entrent ce code et cliquent sur **Rejoindre**, ou ouvrent directement le lien partagé (`…/velocite/#CODE`).
3. Tout le monde clique sur **Entrer dans l'arène**.

La liste **Parties en cours** du menu affiche les parties publiques avec leur hôte, leur carte et le nombre de joueurs : un clic sur **Rejoindre** suffit. Une partie y apparaît tant que la case « Afficher ma partie dans la liste » est cochée chez l'hôte. L'annonce passe par un broker MQTT public (EMQX et HiveMQ) ; si ton réseau le bloque, la liste reste vide mais les codes marchent toujours.

Fermer l'onglet ou la fenêtre fait quitter la partie tout de suite (les autres voient ton départ), et le bouton **Quitter la partie** fait pareil sans fermer le jeu.

La connexion passe par WebRTC avec [PeerJS](https://peerjs.com/) et son serveur de mise en relation public. Il n'y a pas de serveur de jeu : l'hôte relaie positions, tirs et éliminations. Si l'hôte quitte, la partie se ferme.

Sur certains réseaux (4G, réseaux d'entreprise), la connexion directe entre deux joueurs échoue : le jeu passe alors par le relais TURN gratuit « Open Relay » de Metered (identifiants temporaires calculés dans le navigateur, 20 Go par mois partagés).

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

**Boutons réassignables :** dans les réglages (section Manette), clique sur une action puis appuie sur le bouton de la manette à utiliser ; Start annule, « Boutons par défaut » remet tout.

Réglages dédiés : sensibilité, **aide à la visée** (ralentissement près d'une cible et légère traction quand tu bouges ou vises, de 0 à 100 %), **snap strafe facile** en l'air et inversion verticale. Au clavier, le **snap tap** est activé par défaut : si deux directions opposées sont enfoncées, la dernière appuyée gagne.

## Optimisation

Les collisions utilisent une grille spatiale (seuls les blocs proches sont testés), le décor est fusionné par morceaux de 96 m pour que la carte graphique et les tirs ignorent ce qui est hors champ, l'herbe disparaît au-delà de 130 m et la distance d'affichage suit le brouillard. Sur la carte Monde, la physique est environ 4 fois plus rapide et les tirs 2,5 fois. En qualité « auto », la résolution baisse toute seule si les images par seconde chutent, et sur les petites machines (4 Go de mémoire ou moins, ou 3 cœurs ou moins) le jeu passe directement en qualité basse (pas d'anticrénelage, pas d'ombres, moins d'herbe, résolution réduite). Les gros modèles (Nova, Squelette) ne se chargent que quand quelqu'un les utilise.

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
| L | Équipement (changer de perso ou de couteau) |
| Clic droit (couteau en main) | Capacité du couteau |
| R | Recharger |
| Tab | Tableau des scores en ligne |
| T | Réapparaître |
| P | Couper le son |

## Armes

**Couteaux :** combat, karambit, machette, dague, papillon, katana, hachette, épée longue, épée géante, Morgane, trident. Le karambit est un vrai modèle 3D (Astral, CC BY 4.0) ; les autres lames sont modélisées dans le jeu (lames à facettes, gorge, garde, fusée, pommeau).

**Capacités des couteaux (clic droit / viser) :** combat — Fente (bond de 4 m qui poignarde, 65) ; karambit — Frénésie (4 s de frappes deux fois plus rapides) ; machette — Moulinet (coup circulaire, 70) ; dague — Lancer (85, tête ×1,5) ; hachette — Lancer (100) ; papillon — Esquive (pas de côté, invulnérable 0,5 s) ; katana — Iai (ruée de 8 m qui tranche, 95) ; épée longue — Parade (1,2 s à 20 % des dégâts) ; épée géante — Fracas (saut puis onde à l'atterrissage, 110 sur 5,5 m) ; Morgane — Tourbillon (4 coups de 34 autour de toi) ; trident — se lance et revient. Le katana a été refait : lame courbe à section en losange (tranchant, arête, dos), ligne de trempe d'une autre teinte, pointe kissaki, gorge le long du dos, habaki doré, tsuba ajourée à quatre lobes, poignée tressée en losanges avec ornements dorés. **Inspection (X)** : le katana monte à la verticale devant les yeux avec un reflet qui glisse sur l'acier, se balance pour montrer les deux faces, double moulinet et retour en arc ; les couteaux en modèle tournent dans le poignet, sont lancés en l'air avec double vrille, rattrapés et font un moulinet entre les doigts. Le katana a un skin spécial **★ Lame céleste** : galaxie qui défile le long de la lame, tranchant arc-en-ciel, garde dorée et étincelles qui remontent la lame.

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
| Épée géante | 120, portée 4,4 m | 42 coups/min | — |
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

**Télécharger :** [Velocite.exe](https://github.com/Haze15/velocite/raw/main/dist/Velocite.exe) (Windows 64 bits).

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

- **Gasmask girl** (`assets/models/GasmaskGirl.glb`) : créée par **Xinus22**, publiée sur Sketchfab par photon (that one larry), licence [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) — [sketchfab.com/3d-models/gasmask-girl-d5b54d5358fb42368faf47b1cfbab18f](https://sketchfab.com/3d-models/gasmask-girl-d5b54d5358fb42368faf47b1cfbab18f). Squelette renommé pour les animations du jeu, une seule expression du visage gardée.
- **Karambit** (`assets/models/knives/karambit.glb`) : modèle « Karambit » par **Astral** sur Sketchfab, licence [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) ; simplifié (moins de triangles) pour le jeu.

- **Squelette** (`assets/models/Skeleton.glb`) : modèle Sketchfab riggé Mixamo, fourni par le propriétaire du jeu ; c'est l'apparence de son Soldat. Il est animé avec les clips du X-Bot.

- Robots d’entraînement : « RobotExpressive » de Tomás Laulhé ([Quaternius](https://quaternius.com/)), modifié par Don McCurdy, licence CC0 1.0.
- Soldat, X-Bot et Nova (modèle « Michelle ») : modèles et animations [Mixamo](https://www.mixamo.com/) (Adobe), tels que fournis dans les exemples de three.js. Nova est animée avec les mouvements du X-Bot.
- Musique de fond : « Celestial Butterfly », téléchargée sur [Pixabay](https://pixabay.com/music/) (licence de contenu Pixabay, utilisation libre dans un jeu).
- Moteur : [three.js](https://threejs.org/) r128 (MIT), copié dans `launcher/vendor/` pour l'appli Windows. Réseau : [PeerJS](https://peerjs.com/) (MIT).
- Fenêtre de l'appli : [go-webview2](https://github.com/jchv/go-webview2) (MIT) et [go-winloader](https://github.com/jchv/go-winloader) (ISC) de John Chadwick, [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) (BSD), copiés dans `launcher/third_party/`.
