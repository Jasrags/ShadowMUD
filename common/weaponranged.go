package common

// const (
// 	RangedWeaponsFilepath = "_data/items/weapons/ranged"

// 	WeaponRangedReloadBreakAction        WeaponRangedReload = "b"
// 	WeaponRangedReloadDetachableMagazine WeaponRangedReload = "c"
// 	WeaponRangedReloadDrum               WeaponRangedReload = "d"
// 	WeaponRangedReloadMuzzleLoader       WeaponRangedReload = "ml"
// 	WeaponRangedReloadInternalMagazine   WeaponRangedReload = "m"
// 	WeaponRangedReloadCylinder           WeaponRangedReload = "cy"
// 	WeaponRangedReloadBelt               WeaponRangedReload = "belt"

// 	WeaponFiringModeSingleShot    WeaponFiringMode = "Single-Shot"
// 	WeaponFiringModeSemiAutomatic WeaponFiringMode = "Semi-Automatic"
// 	WeaponFiringModeBurstFire     WeaponFiringMode = "Burst Fire"
// 	WeaponFiringModeLongBurst     WeaponFiringMode = "Long Burst"
// 	WeaponFiringModeFullAuto      WeaponFiringMode = "Full Auto"
// )

// type (
// 	WeaponRangedReload string
// 	WeaponFiringMode   string
// 	WeaponRangedSpec   struct {
// 		ID               string               `yaml:"id,omitempty"`
// 		Name             string               `yaml:"name,omitempty"`
// 		Description      string               `yaml:"description,omitempty"`
// 		Accuracy         int                  `yaml:"accuracy,omitempty"`
// 		DamageValue      int                  `yaml:"damage_value,omitempty"`
// 		DamageType       DamageType           `yaml:"damage_type,omitempty"`
// 		ArmorPenatration int                  `yaml:"armor_penatration,omitempty"`
// 		FiringModes      []WeaponFiringMode   `yaml:"firing_modes"`
// 		Recoil           int                  `yaml:"recoil,omitempty"`
// 		AmmoCapacity     int                  `yaml:"ammo_capacity,omitempty"`
// 		Reload           WeaponRangedReload   `yaml:"reload,omitempty"`
// 		Availability     int                  `yaml:"availability,omitempty"`
// 		LegalityType     LegalityType         `yaml:"legality_type,omitempty"`
// 		ItemTags         []ItemTag            `yaml:"tags"`
// 		Modifications    []WeaponModification `yaml:"modifications"`
// 		Modifiers        []Modifier           `yaml:"modifiers"`
// 		Cost             int                  `yaml:"cost,omitempty"`
// 		RuleSource       RuleSource           `yaml:"rule_source,omitempty"`
// 	}
// 	WeaponRanged struct {
// 		ID                 string               `yaml:"id,omitempty"`
// 		SelectedFiringMode WeaponFiringMode     `yaml:"selected_firing_mode,omitempty"`
// 		AmmoType           *WeaponAmunitionSpec `yaml:"ammo_type,omitempty"`
// 		AmmoRemaining      int                  `yaml:"ammo_remaining,omitempty"`
// 		Modifications      []WeaponModification `yaml:"modifications"`
// 		Modifiers          []Modifier           `yaml:"modifiers"`
// 		Spec               WeaponRangedSpec     `yaml:"-"`
// 	}
// )

// func (w *WeaponRanged) ToggleFiringMode() string {
// 	if len(w.Spec.FiringModes) == 0 {
// 		return fmt.Sprintf(MessageNoFiringModes, w.Spec.Name)
// 	}

// 	for i, v := range w.Spec.FiringModes {
// 		if v == w.SelectedFiringMode {
// 			if i+1 < len(w.Spec.FiringModes) {
// 				w.SelectedFiringMode = w.Spec.FiringModes[i+1]
// 			} else {
// 				w.SelectedFiringMode = w.Spec.FiringModes[0]
// 			}
// 			break
// 		}
// 	}

// 	return fmt.Sprintf(MessageFiringModeChanged, w.Spec.Name, w.SelectedFiringMode)
// }

// var CoreWeaponRanged = []WeaponRangedSpec{
// 	{
// 		ID:           "shuriken",
// 		Name:         "Shuriken",
// 		Description:  "A shuriken is a small, star-shaped piece of metal with sharpened edges, designed for throwing. It is also known as a “throwing star” or “ninja star.”",
// 		DamageValue:  1,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		Recoil:       -1,
// 		Availability: 4,
// 		LegalityType: LegalityTypeRestricted,
// 		Cost:         25,
// 		RuleSource:   RuleSourceSR5Core,
// 		ItemTags:     []ItemTag{ItemTagRanged, ItemTagThrowing},
// 		// Wireless
// 		//  If all the throwing knives or shuriken you throw in a single Combat Turn are wireless and you have a smartlink system, each knife you throw receives a +1 dice pool bonus per knife thrown that Combat Turn at your current target, as the knives inform and adjust for wind and other atmospheric conditions. So you’d get no bonus on the first throw, a +1 bonus on the second throw, a +2 bonus on the third throw, etc. (assuming you aimed all three knives at the same target).
// 	},
// 	{
// 		ID:           "throwing_knife",
// 		Name:         "Throwing Knife",
// 		Description:  "A throwing knife is a knife that is specially designed and weighted so that it can be thrown effectively.",
// 		DamageValue:  1,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		Recoil:       -1,
// 		Availability: 4,
// 		LegalityType: LegalityTypeRestricted,
// 		Cost:         25,
// 		RuleSource:   RuleSourceSR5Core,
// 		ItemTags:     []ItemTag{ItemTagRanged, ItemTagThrowing},
// 		// Wireless
// 		//  If all the throwing knives or shuriken you throw in a single Combat Turn are wireless and you have a smartlink system, each knife you throw receives a +1 dice pool bonus per knife thrown that Combat Turn at your current target, as the knives inform and adjust for wind and other atmospheric conditions. So you’d get no bonus on the first throw, a +1 bonus on the second throw, a +2 bonus on the third throw, etc. (assuming you aimed all three knives at the same target).
// 	},
// 	{
// 		ID:               "bow",
// 		Name:             "Bow",
// 		Description:      "A bow is a flexible arc that shoots aerodynamic projectiles called arrows. A string joins the two ends of the bow and when the string is drawn back, the ends of the bow are flexed. When the string is released, the potential energy of the flexed bow limbs is transformed into the velocity of the arrow.",
// 		Accuracy:         6,
// 		DamageValue:      2, // (Rating+2)
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		ArmorPenatration: 0,   // -(Rating/4)
// 		Availability:     1,   // Rating
// 		Cost:             100, // Rating×100¥
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagBow, ItemTagBallistic},
// 		// Max Rating
// 		//     10
// 		// Strength Minimum
// 		//     If Strength is less than Rating then -3 DP per point below Rating.
// 		// Damage Rating
// 		//     Lowest value of Strength, Bows Rating, or Arrow Rating.
// 		// Range Rating
// 		//     Lowest value of Strength, Bows Rating, or Arrow Rating.
// 		// Reload
// 		//     Simple Action
// 		// When attacking with a bow, a character whose Strength is less than the Rating suffers a –3 dice pool modifier per point below the minimum
// 		// Use the lowest value of your Strength, the bow’s rating, or the arrow Rating for range and damage when attacking a target, because your average Rating 10
// 	},
// 	{
// 		ID:               "light_crossbow",
// 		Name:             "Light Crossbow",
// 		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
// 		Accuracy:         7,
// 		DamageValue:      5,
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		ArmorPenatration: -1,
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Availability:     2,
// 		LegalityType:     LegalityTypeLegal,
// 		Cost:             300,
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagCrossbow, ItemTagBallistic},
// 	},
// 	{
// 		ID:               "medium_crossbow",
// 		Name:             "Medium Crossbow",
// 		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
// 		Accuracy:         6,
// 		DamageValue:      7,
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		ArmorPenatration: -2,
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Availability:     4,
// 		LegalityType:     LegalityTypeRestricted,
// 		Cost:             500,
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagCrossbow, ItemTagBallistic},
// 	},
// 	{
// 		ID:               "heavy_crossbow",
// 		Name:             "Heavy Crossbow",
// 		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
// 		Accuracy:         5,
// 		DamageValue:      10,
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		ArmorPenatration: -3,
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Availability:     8,
// 		LegalityType:     LegalityTypeRestricted,
// 		Cost:             1000,
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagCrossbow, ItemTagBallistic},
// 	},
// 	{
// 		ID:               "grapple_gun",
// 		Name:             "Grapple Gun",
// 		Description:      "A grapple gun is a device that allows the user to fire a grappling hook or a similar object to a distant location.",
// 		Accuracy:         3,
// 		DamageValue:      7,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -2,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		AmmoCapacity:     1,
// 		Reload:           WeaponRangedReloadMuzzleLoader,
// 		Availability:     8,
// 		LegalityType:     LegalityTypeRestricted,
// 		Cost:             500,
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagExotic},
// 	},
// 	{
// 		ID:               "defiance_ex_shocker",
// 		Name:             "Defiance EX Shocker",
// 		Accuracy:         4,
// 		DamageValue:      9,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -5,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Cost:             250,
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagTaser},
// 		Modifiers:        []Modifier{{Type: "add", Effect: "ElectricDamage"}},
// 		// Wireless
// 		//  A successful hit informs you of the status of the target’s basic health (and Condition Monitors).
// 	},
// 	{
// 		ID:               "yamaha_pulsar",
// 		Name:             "Yamaha Pulsar",
// 		Accuracy:         5,
// 		DamageValue:      7,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -5,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Cost:             180,
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagTaser},
// 		Modifiers:        []Modifier{{Type: "add", Effect: "ElectricDamage"}},
// 		// Wireless
// 		//  A successful hit informs you of the status of the target’s basic health (and Condition Monitors).
// 	},
// 	{
// 		ID:               "ficetti_tiffani_needler",
// 		Name:             "Fichetti Tiffani Needler",
// 		Description:      "The Fichetti Tiffani Needler is a hold-out pistol that fires flechette ammunition. It is a small, easily concealed weapon that is popular with shadowrunners and criminals.",
// 		Accuracy:         5,
// 		DamageValue:      8,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: 5,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadDetachableMagazine,
// 		Availability:     6,
// 		LegalityType:     LegalityTypeRestricted,
// 		Cost:             1000,
// 		RuleSource:       RuleSourceSR5Core,
// 		ItemTags:         []ItemTag{ItemTagRanged, ItemTagPistol, ItemTagHoldOutPistol},
// 		// 8P(f)
// 		// Wireless
// 		//  You can change the color of the Tiffani Needler with a Simple Action.
// 		// Can only fire flechette rounds
// 	},
// 	{
// 		ID:           "streetline_special",
// 		Name:         "Streetline Special",
// 		Description:  "The Streetline Special is a hold-out pistol that is popular with shadowrunners and criminals. It is a small, easily concealed weapon.",
// 		Accuracy:     4,
// 		DamageValue:  6,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		AmmoCapacity: 6,
// 		Reload:       WeaponRangedReloadDetachableMagazine,
// 		Availability: 4,
// 		LegalityType: LegalityTypeRestricted,
// 		Cost:         120,
// 		RuleSource:   RuleSourceSR5Core,
// 		ItemTags:     []ItemTag{ItemTagRanged, ItemTagPistol, ItemTagHoldOutPistol},
// 		// MAD Scanner
// 		//  -2 DP to detect Streetline Special
// 	},
// 	{
// 		ID:           "walther_palm_pistol",
// 		Name:         "Walther Palm Pistol",
// 		Description:  "The Walther Palm Pistol is a hold-out pistol that is popular with shadowrunners and criminals. It is a small, easily concealed weapon.",
// 		Accuracy:     4,
// 		DamageValue:  7,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSingleShot, WeaponFiringModeBurstFire},
// 		AmmoCapacity: 2,
// 		Reload:       WeaponRangedReloadBreakAction,
// 		Availability: 4,
// 		LegalityType: LegalityTypeRestricted,
// 		Cost:         180,
// 		RuleSource:   RuleSourceSR5Core,
// 		ItemTags:     []ItemTag{ItemTagRanged, ItemTagPistol, ItemTagHoldOutPistol},
// 	},
// }

// // Light Pistols
// // -------------
// // | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// // |--------|-----|----|----|-------|----|------|-------|------|--------|
// // | Ares Light Fire 70 | 7 | 8P | – | SA | – | 16 (c) | 3R | 200¥ | Core |
// // | Ares Light Fire 75 | 6 (8) | 6P | – | SA | – | 16 (c) | 6F | 1,250¥ | Core |
// // | Beretta 201T | 6 | 6P | – | SA/BF | – (1) | 21 (c) | 7R | 210¥ | Core |
// // | Colt America L36 | 7 | 7P | – | SA | – | 11 (c) | 4R | 320¥ | Core |
// // | Fichetti Security 600 | 6 (7) | 7P | – | SA | – (1) | 30 (c) | 6R | 350¥ | Core |
// // | Taurus Omni-6, light pistol rounds | 5 (6) | 6P | – | SA | – | 6 (cy) | 3R | 300¥ | Core |
// // | Taurus Omni-6, heavy pistol rounds | 5 (6) | 7P | -1 | SS | – | 6 (cy) | 3R | 300¥ | Core |
// // Heavy Pistols
// // -------------
// // | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost | Source    |
// // |--------|-----|-----|----|-------|----|------|-------|------|-----------|
// // | Ares Predator V | 5 (7) | 8P  | -1 | SA | – | 15 (c) | 5R | 725¥ | Core |
// // | Ares Viper Silvergun | 4 | 9P(f) | +4 | SA/BF | – | 30 (c) | 8F | 380¥ | Core |
// // | Browning Ultra-Power | 5 (6) | 8P | -1 | SA | – | 10 (c) | 4R | 640¥ | Core |
// // | Colt Government 2066 | 6 | 7P | -1 | SA | – | 14 (c) | 7R | 425¥ | Core |
// // | Remington Roomsweeper | 4 | 7P | -1 | SA | – | 8 (m) | 6R | 250¥ | Core |
// // | Ruger Super Warhawk | 5 | 9P | -2 | SS | – | 6 (cy) | 4R | 400¥ | Core |
// // Machine Pistols
// // ---------------
// // | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// // |--------|-----|----|----|-------|----|------|-------|------|--------|
// // | Ares Crusader II | 5 (7) | 7P | – | SA/BF | 2 | 40 (c) | 9R | 830¥ | Core |
// // | Ceska Black Scorpion | 5 | 6P | – | SA/BF | – (1) | 35 (c) | 6R | 270¥ | Core |
// // | Steyr TMP | 4 | 7P | – | SA/BF/FA | – | 30 (c) | 8R | 350¥ | Core |
// // Submachine Guns
// // ===============
// // | Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source |
// // |--------|-----|----|----|-------|----|------|-------|------|--------|
// // | Colt Cobra TZ-120 | 4 (5) | 7P | – | SA/BF/FA | 2 (3) | 32 (c) | 5R | 660¥ | Core |
// // | FN P93 Praetor | 6 | 8P | – | SA/BF/FA | 1 (2) | 50 (c) | 11F | 900¥ | Core |
// // | HK-227 | 5 (7) | 7P | – | SA/BF/FA | – (1) | 28 (c) | 8R | 730¥ | Core |
// // | Ingram Smartgun X | 4 (6) | 8P | – | BF/FA | 2 | 32 (c) | 6R | 800¥ | Core |
// // | SCK Model 100 | 5 (7) | 8P | – | SA/BF | – (1) | 30 (c) | 6R | 875¥ | Core |
// // | Uzi IV | 4 (5) | 7P | – | BF | – (1) | 24 (c) | 4R | 450¥ | Core |
// // Assault Rifles
// // --------------
// // | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost | Source    |
// // |--------|-----|-----|----|-------|----|------|-------|------|-----------|
// // | AK-97  | 4   | 10P | -2 | SA/BF/FA | –  | 38 (c) | 4R    | 950¥ | Core      |
// // | Ares Alpha  | 5 (7) | 11P | -2 | SA/BF/FA | 2  | 42 (c) | 11F   | 2,650¥ | Core |
// // | Colt M23  | 4   | 9P  | -2 | SA/BF/FA | –  | 40 (c) | 4R    | 550¥ | Core      |
// // | FN HAR  | 5 (6) | 10P | -2 | SA/BF/FA | 2  | 35 (c) | 8R    | 1,500¥ | Core |
// // | Yamaha Raiden  | 6 (8) | 11P | -2 | BF/FA | 2  | 60 (c) | 14F   | 2,600¥ | Core |
// // Sniper Rifles
// // -------------
// // Weapon | Acc | DV | AP | Modes | RC | Ammo | Avail | Cost | Source
// // -------|-----|----|----|-------|----|------|-------|------|--------
// // Ares Desert Strike | 7 | 13P | -4 | SA | – (1) | 14 (c) | 10F | 17,500¥ | Core
// // Cavalier Arms Crockett EBR | 6 | 12P | -3 | SA/BF | – (1) | 20 (c) | 12F | 10,300¥ | Core
// // Ranger Arms SM-5 | 8 | 14P | -5 | SA | – (1) | 15 (c) | 16F | 28,000¥ | Core
// // Remington 950 | 7 | 12P | -4 | SS | – | 5 (m) | 4R | 2,100¥ | Core
// // Ruger 100 | 6 | 11P | -3 | SS | – (1) | 8 (m) | 4R | 1,300¥ | Core
// // Shotguns
// // ========
// // | Weapon | Acc | DV  | AP | Modes    | RC | Ammo | Avail | Cost  | Source    |
// // |--------|-----|-----|----|----------|----|------|-------|-------|-----------|
// // | Defiance T-250 | 4   | 10P | -1 | SS/SA    | –  | 5 (m)            | 4R  | 450¥   | Core     |
// // | Defiance T-250, short-barreled | 4 | 9P | -1 | SS/SA    | –  | 5 (m)            | 4R  | 450¥   | Core     |
// // | Enfield AS-7 | 4 (5) | 13P | -1 | SA/BF    | –  | 10 (c) or 24 (d) | 12F | 1,100¥ | Core     |
// // | PJSS Model 55 | 6   | 11P | -1 | SS       | – (1) | 2 (b)            | 9R  | 1,000¥ | Core     |
// // Light Machine Guns
// // ------------------
// // | Weapon | Acc | DV  | AP | Modes | RC | Ammo | Avail | Cost  | Source    |
// // |--------|-----|-----|----|-------|----|------|-------|-------|-----------|
// // | Ingram Valiant         | 5 (6) | 9P  | -2 | BF/FA | 2 (3) | 50 (c) or 100 (belt) | 12F   | 5,800¥ | Core    |
// // Medium Machine Guns
// // -------------------
// // | Weapon              | Acc | DV  | AP | Modes | RC    | Ammo              | Avail | Cost   | Source    |
// // |---------------------|-----|-----|----|-------|-------|-------------------|-------|--------|-----------|
// // | Stoner-Ares M202    | 5   | 10P | -3 | FA    | –     | 50 (c) or 100 (belt) | 12F   | 7,000¥ | Core      |
// // Heavy Machine Guns
// // ------------------
// // | Weapon         | Acc | DV  | AP | Modes | RC   | Ammo              | Avail | Cost    | Source    |
// // |----------------|-----|-----|----|-------|------|-------------------|-------|---------|-----------|
// // | RPK HMG        | 5   | 12P | -4 | FA    | – (6) | 50 (c) or 100 (belt) | 16F   | 16,300¥ | Core      |
// // Exotic Firearms
// // ===============
// // | Weapon                        | Acc | DV   | AP  | Modes | RC | Ammo | Avail | Cost    | Source    |
// // |-------------------------------|-----|------|-----|-------|----|------|-------|---------|-----------|
// // | Ares S-III Super Squirt       | 3   | Chem | SA  | –     | –  | 20   | 7R    | 950¥    | Core      |
// // | Fichetti Pain Inducer         | 3   | Spec | –   | SS    | –  | Spec | 11R   | 5,000¥  | Core      |
// // | Parashield Dart Pistol        | 5   | Drug | SA  | –     | –  | 5    | 4R    | 600¥    | Core      |
// // | Parashield Dart Rifle         | 6   | Drug | SA  | –     | –  | 6    | 6R    | 1,200¥  | Core      |
// // Assault Cannons
// // ----------------
// // | Weapon                           | Acc | DV  | AP | Modes | RC   | Ammo              | Avail | Cost    | Source    |
// // |----------------------------------|-----|-----|----|-------|------|-------------------|-------|---------|-----------|
// // | Krime Cannon                     | 4   | 16P | -6 | SA    | – (1) | 6 (m)             | 20F   | 21,000¥ | Core      |
// // | Panther XXL                      | 5 (7) | 17P | -6 | SA    | –    | 15 (c)            | 20F   | 43,000¥ | Core      |
// // Grenade Launchers
// // -----------------
// // | Weapon                      | Acc | DV      | AP | Modes | RC | Ammo | Avail | Cost    | Source    |
// // |-----------------------------|-----|---------|----|-------|----|------|-------|---------|-----------|
// // | Ares Alpha, Grenade Launcher| 4 (6) | Grenade | SS | –     | 6  | –    | –     | Core    |           |
// // | Ares Antioch-2              | 4 (6) | Grenade | SS | –     | 8  | 8F   | 3,200¥ | Core    |           |
// // | ArmTech MGL-12              | 4   | Grenade | SA | –     | 12 | 10F  | 5,000¥ | Core    |           |
// // Missile Launchers
// // -----------------
// // | Weapon                      | Acc | DV      | AP | Modes | RC   | Ammo | Avail | Cost    | Source    |
// // |-----------------------------|-----|---------|----|-------|------|------|-------|---------|-----------|
// // | Aztechnology Striker        | 5   | Missile | SS | –     | 2 (ml) | 10F  | 1,200¥ | Core    |
// // | Onotari Interceptor         | 4 (6) | Missile | SS | –     | 2 (ml) | 18F  | 14,000¥ | Core    |

// // TODO: Load the data from the yaml files
// func LoadRangedWeapons() map[string]WeaponRanged {
// 	data := make(map[string]WeaponRanged)
// 	return data
// }
