package common

// const (
// 	WeaponsFilePath = "_data/items/weapons"

// 	WeaponTypeMelee  WeaponType = "Melee"
// 	WeaponTypeRanged WeaponType = "Ranged"
// 	// Melee
// 	WeaponGroupImprovised WeaponGroup = "Improvised"
// 	WeaponGroupClubs      WeaponGroup = "Clubs"
// 	WeaponGroupBlades     WeaponGroup = "Blades"
// 	WeaponGroupExotic     WeaponGroup = "Exotic"
// 	WeaponGroupMisc       WeaponGroup = "Misc"
// 	// Ranged
// 	WeaponGroupThrowingWeapons         WeaponGroup = "Throwing Weapons"
// 	WeaponGroupBallisticProjectiles    WeaponGroup = "Ballistic Projectiles"
// 	WeaponGroupFlamethrowers           WeaponGroup = "Flamethrowers"
// 	WeaponGroupExoticRangedWeapons     WeaponGroup = "Exotic Ranged Weapons"
// 	WeaponGroupTasers                  WeaponGroup = "Tasers"
// 	WeaponGroupFirearms                WeaponGroup = "Firearms"
// 	WeaponGroupLasers                  WeaponGroup = "Lasers"
// 	WeaponGroupLargeCaliberProjectiles WeaponGroup = "Large-Caliber Projectiles"
// 	WeaponGroupImplantWeapons          WeaponGroup = "Implant Weapons"
// 	// Ranged
// 	WeaponCategoryPistol         WeaponCategory = "Pistol"
// 	WeaponCategorySubmachineGun  WeaponCategory = "Submachine Gun"
// 	WeaponCategoryRifle          WeaponCategory = "Rifle"
// 	WeaponCategoryShotgun        WeaponCategory = "Shotgun"
// 	WeaponCategoryMachineGun     WeaponCategory = "Machine Gun"
// 	WeaponCategoryExoticFirearms WeaponCategory = "Exotic Firearms"
// 	// Melee
// 	WeaponCategoryClub   WeaponCategory = "Club"
// 	WeaponCategoryStaff  WeaponCategory = "Staff"
// 	WeaponCategoryKnife  WeaponCategory = "Knife"
// 	WeaponCategoryAxe    WeaponCategory = "Axe"
// 	WeaponCategorySword  WeaponCategory = "Sword"
// 	WeaponCategoryKatana WeaponCategory = "Katana"
// 	WeaponCategoryDagger WeaponCategory = "Dagger"
// 	// Ranged
// 	WeaponSubCategoryHoldOutPistol      WeaponSubCategory = "Hold-Out Pistol"
// 	WeaponSubCategoryLightPistol        WeaponSubCategory = "Light Pistol"
// 	WeaponSubCategoryHeavyPistol        WeaponSubCategory = "Heavy Pistol"
// 	WeaponSubCategoryMachinePistol      WeaponSubCategory = "Machine Pistol"
// 	WeaponSubCategoryAssaultRifle       WeaponSubCategory = "Assault Rifle"
// 	WeaponSubCategorySniperRifle        WeaponSubCategory = "Sniper Rifle"
// 	WeaponSubCategorySportingRifle      WeaponSubCategory = "Sporting Rifle"
// 	WeaponSubCategoryLightMachineGun    WeaponSubCategory = "Light Machine Gun"
// 	WeaponSubCategoryMediumMachineGun   WeaponSubCategory = "Medium Machine Gun"
// 	WeaponSubCategoryHeavyMachineGun    WeaponSubCategory = "Heavy Machine Gun"
// 	WeaponSubCategoryAssaultCannon      WeaponSubCategory = "Assault Cannon"
// 	WeaponSubCategoryGrenadeLauncher    WeaponSubCategory = "Grenade Launcher"
// 	WeaponSubCategoryMissileLauncher    WeaponSubCategory = "Missile Launcher"
// 	WeaponSubCategoryImplantMeleeWeapon WeaponSubCategory = "Implant Melee Weapon"
// 	WeaponSubCategoryImplantFirearm     WeaponSubCategory = "Implant Firearm"
// )

// type (
// 	WeaponType        string
// 	WeaponGroup       string
// 	WeaponCategory    string
// 	WeaponSubCategory string
// 	WeaponSpec        struct {
// 		ID               string               `yaml:"id"`
// 		Name             string               `yaml:"name"`
// 		Description      string               `yaml:"description"`
// 		Type             WeaponType           `yaml:"type"`
// 		Group            WeaponGroup          `yaml:"group"`
// 		Category         WeaponCategory       `yaml:"category"`
// 		SubCategory      WeaponSubCategory    `yaml:"sub_category"`
// 		Concealability   int                  `yaml:"concealability"`
// 		Accuracy         int                  `yaml:"accuracy"`
// 		Reach            int                  `yaml:"reach"`
// 		DamageValue      int                  `yaml:"damage_value"`
// 		DamageType       DamageType           `yaml:"damage_type"`
// 		ArmorPenatration int                  `yaml:"armor_penatration"`
// 		Recoil           int                  `yaml:"recoil"`
// 		Reload           WeaponRangedReload   `yaml:"reload"`
// 		FiringModes      []WeaponFiringMode   `yaml:"firing_modes"`
// 		AmmoCapacity     int                  `yaml:"ammo_capacity"`
// 		Availability     int                  `yaml:"availability"`
// 		Legality         LegalityType         `yaml:"legality"`
// 		Tags             []ItemTag            `yaml:"tags"`
// 		Modifications    []WeaponModification `yaml:"modifications"`
// 		Modifiers        []Modifier           `yaml:"modifiers"`
// 		Cost             int                  `yaml:"cost"`
// 		RuleSource       RuleSource           `yaml:"rule_source"`
// 	}
// 	Weapons map[string]*Weapon
// 	Weapon  struct {
// 		sync.Mutex `yaml:"-"`
// 		log        *logrus.Entry `yaml:"-"`

// 		ID                 string               `yaml:"id"`
// 		SelectedFiringMode WeaponFiringMode     `yaml:"selected_firing_mode"`
// 		AmmoType           *WeaponAmunitionSpec `yaml:"ammo_type"`
// 		AmmoRemaining      int                  `yaml:"ammo_remaining"`
// 		Tags               []ItemTag            `yaml:"tags"`
// 		Modifications      []WeaponModification `yaml:"modifications"`
// 		Modifiers          []Modifier           `yaml:"modifiers"`
// 		Spec               *WeaponSpec          `yaml:"-"`
// 	}
// )

// func NewWeapon(spec *WeaponSpec) *Weapon {
// 	w := &Weapon{
// 		ID:   uuid.New().String(),
// 		Spec: spec,
// 	}
// 	w.log = logrus.WithFields(logrus.Fields{"package": "common", "type": "weapon", "weapon_id": w.ID, "weapon_name": w.Spec.Name})

// 	return w
// }

// // func (w *Weapon) GetDamageValue() error {
// // 	/*
// // 	   (STR)P
// // 	   (STR+3)P
// // 	   9S(e)
// // 	   6P(fire)
// // 	   6S(e)
// // 	   8P(f)
// // 	   (STR+3)S / 12P
// // 	   (STR+1 / 3)P
// // 	   12P
// // 	   (Rating+2)P
// // 	*/
// // 	return nil
// // }

// var CoreWeapons = []WeaponSpec{
// 	//Melee
// 	{
// 		ID:           "club",
// 		Name:         "Club",
// 		Type:         WeaponTypeMelee,
// 		Group:        WeaponGroupClubs,
// 		Category:     WeaponCategoryClub,
// 		Description:  "The weapon they named the skill after.",
// 		Accuracy:     4,
// 		Reach:        1,
// 		DamageValue:  2,
// 		DamageType:   DamageTypePhysical,
// 		Availability: 4,
// 		Legality:     LegalityTypeLegal,
// 		Cost:         30,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagMelee, ItemTagClub},
// 	},
// 	{
// 		ID:           "extendable_baton",
// 		Name:         "Extendable Baton",
// 		Type:         WeaponTypeMelee,
// 		Group:        WeaponGroupClubs,
// 		Category:     WeaponCategoryClub,
// 		Description:  "A baton that extends to a full length.",
// 		Accuracy:     5,
// 		Reach:        1,
// 		DamageValue:  2,
// 		DamageType:   DamageTypePhysical,
// 		Availability: 4,
// 		Legality:     LegalityTypeLegal,
// 		Cost:         100,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagMelee, ItemTagBaton},
// 		// Simple Action
// 		//  Collapse/Extend
// 		// Collapsed
// 		//  Concealability Modifier 0
// 		// Extended
// 		//  Concealability Modifier +2
// 		// Wireless
// 		//  Readying the extendable baton is a Free Action instead of a Simple Action.
// 	},
// 	{
// 		ID:             "sap",
// 		Name:           "Sap",
// 		Type:           WeaponTypeMelee,
// 		Group:          WeaponGroupClubs,
// 		Category:       WeaponCategoryClub,
// 		Description:    "A small, weighted club.",
// 		Concealability: -2,
// 		Accuracy:       5,
// 		DamageValue:    2,
// 		DamageType:     DamageTypePhysical,
// 		Availability:   2,
// 		Legality:       LegalityTypeLegal,
// 		Cost:           30,
// 		RuleSource:     RuleSourceSR5Core,
// 		Tags:           []ItemTag{ItemTagMelee, ItemTagClub},
// 		// Concealability
// 		//  Modifier +2
// 	},
// 	{
// 		ID:           "staff",
// 		Name:         "Staff",
// 		Type:         WeaponTypeMelee,
// 		Group:        WeaponGroupClubs,
// 		Category:     WeaponCategoryClub,
// 		Description:  "A long stick.",
// 		Accuracy:     6,
// 		Reach:        2,
// 		DamageValue:  3,
// 		DamageType:   DamageTypePhysical,
// 		Availability: 3,
// 		Legality:     LegalityTypeLegal,
// 		Cost:         100,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagMelee, ItemTagStaff},
// 	},
// 	{
// 		ID:               "stun_baton",
// 		Name:             "Stun Baton",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "A baton that delivers an electrical shock.",
// 		Accuracy:         4,
// 		Reach:            1,
// 		DamageValue:      9,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -5,
// 		Availability:     6,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             750,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBaton},
// 		Modifiers: []Modifier{
// 			{
// 				Type:   ModifierTypeElectric,
// 				Effect: ModifierEffectAttack,
// 			},
// 		},
// 		// Wireless
// 		//  The stun baton recharges by induction, regaining one charge per full hour of wireless-enabled time.
// 		// Modifier
// 		//  Electric attack
// 	},
// 	{
// 		ID:           "telescoping_staff",
// 		Name:         "Telescoping Staff",
// 		Type:         WeaponTypeMelee,
// 		Group:        WeaponGroupClubs,
// 		Category:     WeaponCategoryClub,
// 		Description:  "A staff that extends to a full length.",
// 		Accuracy:     4,
// 		Reach:        2,
// 		DamageValue:  4,
// 		DamageType:   DamageTypePhysical,
// 		Availability: 4,
// 		Legality:     LegalityTypeLegal,
// 		Cost:         350,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagMelee, ItemTagStaff},
// 	},
// 	// Blades
// 	{
// 		ID:               "combat_axe",
// 		Name:             "Combat Axe",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "A large axe designed for combat.",
// 		Accuracy:         4,
// 		Reach:            2,
// 		DamageValue:      5,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -4,
// 		Availability:     12,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             4000,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagAxe},
// 	},
// 	{
// 		ID:               "combat_knife",
// 		Name:             "Combat Knife",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "A knife designed for combat.",
// 		Accuracy:         6,
// 		DamageValue:      2,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -3,
// 		Availability:     4,
// 		Legality:         LegalityTypeLegal,
// 		Cost:             300,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade},
// 	},
// 	{
// 		ID:               "forearm_snap_blades",
// 		Name:             "Forearm Snap-Blades",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "Blades that extend from the forearm.",
// 		Accuracy:         4,
// 		DamageValue:      2,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -2,
// 		Availability:     7,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             200,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade},
// 		// Wireless
// 		//  Readying the forearm snap blades is a Free Action instead of a Simple Action.
// 	},
// 	{
// 		ID:               "katana",
// 		Name:             "Katana",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "A traditional Japanese sword.",
// 		Accuracy:         7,
// 		Reach:            1,
// 		DamageValue:      3,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -3,
// 		Availability:     9,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             1000,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagKatana},
// 	},
// 	{
// 		ID:               "knife",
// 		Name:             "Knife",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "A small knife.",
// 		Accuracy:         5,
// 		DamageValue:      1,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -1,
// 		Legality:         LegalityTypeLegal,
// 		Cost:             10,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagKnife},
// 	},
// 	{
// 		ID:               "pole_arm",
// 		Name:             "Pole Arm",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "A pole arm is a large weapon with a blade on the end.",
// 		Accuracy:         5,
// 		Reach:            3,
// 		DamageValue:      3,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -2,
// 		Availability:     6,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             1000,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagPolearm},
// 	},
// 	{
// 		ID:               "survival_knife",
// 		Name:             "Survival Knife",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      `A survival knife is a large knife with a serrated edge. It is designed for use in the wilderness, and is often used by military personnel and survivalists.`,
// 		Accuracy:         5,
// 		DamageValue:      2,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -1,
// 		Legality:         LegalityTypeLegal,
// 		Cost:             100,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagKnife},
// 		// Wireless
// 		//  The knife displays an ARO of local maps, your GPS position, and can be used to make commcalls.
// 	},
// 	{
// 		ID:               "sword",
// 		Name:             "Sword",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupClubs,
// 		Category:         WeaponCategoryClub,
// 		Description:      "A sword is a long, sharp blade.",
// 		Accuracy:         6,
// 		Reach:            1,
// 		DamageValue:      3,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -2,
// 		Availability:     5,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             500,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagSword},
// 	},
// 	// Exotic Melee Weapons
// 	{
// 		ID:               "monofilament_chainsaw",
// 		Name:             "Monofilament Chainsaw",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupExotic,
// 		Description:      "A chainsaw with a monofilament blade.",
// 		Accuracy:         3,
// 		Reach:            1,
// 		DamageValue:      8,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -6,
// 		Availability:     8,
// 		Legality:         LegalityTypeLegal,
// 		Cost:             500,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagExotic, ItemTagChainsaw, ItemTagMonofiliament},
// 		// When used against barriers, double the monofilament chainsaw’s Damage Value of 8P
// 	},
// 	{
// 		ID:               "monofilament_whip",
// 		Name:             "Monofilament Whip",
// 		Type:             WeaponTypeMelee,
// 		Group:            WeaponGroupExotic,
// 		Description:      "A whip with a monofilament blade.",
// 		Accuracy:         5,
// 		Reach:            2,
// 		DamageValue:      12,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -8,
// 		Availability:     12,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             10000,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagExotic, ItemTagWhip, ItemTagMonofiliament},
// 		// Skill
// 		//  Exotic Melee Weapon (Monofilament Whip)
// 		// Glitch
// 		//  You catch the weighted tip on something nearby and need to disentangle it before you make another proper attack with it.
// 		// Critical Glitch
// 		//  You hit yourself with the whip and take its base damage (resisted normally)
// 		// Wireless
// 		//  The whip can be readied with a Free Action, rather than a Simple Action.
// 		//  The whip's built-in safety system retracts automatically instead of getting you entangled on a glitch.
// 		//  Accuracy increased by 2.
// 	},
// 	// Misc Melee Weapons
// 	{
// 		ID:          "knucks",
// 		Name:        "Knucks",
// 		Type:        WeaponTypeMelee,
// 		Group:       WeaponGroupMisc,
// 		Description: "These may be traditional fist-load weapons like brass knuckles, or modern Hardliner Gloves with a thin layer of densiplast located the knuckles and the edge of the hand. Either way they substantially boost the impact of a punch, making it potentially deadly.",
// 		// Accuracy:    0, // Physical
// 		DamageValue:  1,
// 		DamageType:   DamageTypePhysical,
// 		Availability: 2,
// 		Legality:     LegalityTypeRestricted,
// 		Cost:         100,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagMelee, ItemTagMiscellaneous},
// 		// Skill
// 		//  Unarmed Combat Skill
// 	},
// 	{
// 		ID:          "plasteel_toe_boots",
// 		Name:        "Plasteel Toe Boots",
// 		Type:        WeaponTypeMelee,
// 		Group:       WeaponGroupMisc,
// 		Description: "These boots have a reinforced toe made of plasteel, making them a formidable weapon in a kick.",
// 		// Accuracy:    0, // Physical
// 		DamageValue:  1,
// 		DamageType:   DamageTypePhysical,
// 		Availability: 2,
// 		Legality:     LegalityTypeLegal,
// 		Cost:         200,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagMelee, ItemTagMiscellaneous},
// 		// Skill
// 		//  Unarmed Combat Skill
// 	},
// 	{
// 		ID:          "shock_gloves",
// 		Name:        "Shock Gloves",
// 		Type:        WeaponTypeMelee,
// 		Group:       WeaponGroupMisc,
// 		Description: "These gloves are lined with a layer of conductive material that delivers an electric shock to the target.",
// 		// Accuracy:    0, // Physical
// 		DamageValue:      8,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -5,
// 		Availability:     6,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             550,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagMelee, ItemTagMiscellaneous},
// 		Modifiers:        []Modifier{{Type: "Electric attack", Effect: "Attack"}},
// 		// Skill
// 		//  Unarmed Combat Skill
// 		// Wireless
// 		//  The shock gloves recharge by induction, regaining one charge per full hour of wireless-enabled time.
// 		// Modifier
// 		//  Electric attack
// 	},
// 	// Ranged
// 	{
// 		ID:           "shuriken",
// 		Name:         "Shuriken",
// 		Type:         WeaponTypeRanged,
// 		Group:        WeaponGroupThrowingWeapons,
// 		Description:  "A shuriken is a small, star-shaped piece of metal with sharpened edges, designed for throwing. It is also known as a “throwing star” or “ninja star.”",
// 		DamageValue:  1,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		Recoil:       -1,
// 		Availability: 4,
// 		Legality:     LegalityTypeRestricted,
// 		Cost:         25,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagRanged, ItemTagThrowing},
// 		// Wireless
// 		//  If all the throwing knives or shuriken you throw in a single Combat Turn are wireless and you have a smartlink system, each knife you throw receives a +1 dice pool bonus per knife thrown that Combat Turn at your current target, as the knives inform and adjust for wind and other atmospheric conditions. So you’d get no bonus on the first throw, a +1 bonus on the second throw, a +2 bonus on the third throw, etc. (assuming you aimed all three knives at the same target).
// 	},
// 	{
// 		ID:           "throwing_knife",
// 		Name:         "Throwing Knife",
// 		Type:         WeaponTypeRanged,
// 		Group:        WeaponGroupThrowingWeapons,
// 		Description:  "A throwing knife is a knife that is specially designed and weighted so that it can be thrown effectively.",
// 		DamageValue:  1,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		Recoil:       -1,
// 		Availability: 4,
// 		Legality:     LegalityTypeRestricted,
// 		Cost:         25,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagRanged, ItemTagThrowing},
// 		// Wireless
// 		//  If all the throwing knives or shuriken you throw in a single Combat Turn are wireless and you have a smartlink system, each knife you throw receives a +1 dice pool bonus per knife thrown that Combat Turn at your current target, as the knives inform and adjust for wind and other atmospheric conditions. So you’d get no bonus on the first throw, a +1 bonus on the second throw, a +2 bonus on the third throw, etc. (assuming you aimed all three knives at the same target).
// 	},
// 	{
// 		ID:               "bow",
// 		Name:             "Bow",
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupBallisticProjectiles,
// 		Description:      "A bow is a flexible arc that shoots aerodynamic projectiles called arrows. A string joins the two ends of the bow and when the string is drawn back, the ends of the bow are flexed. When the string is released, the potential energy of the flexed bow limbs is transformed into the velocity of the arrow.",
// 		Accuracy:         6,
// 		DamageValue:      2, // (Rating+2)
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		ArmorPenatration: 0,   // -(Rating/4)
// 		Availability:     1,   // Rating
// 		Cost:             100, // Rating×100¥
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagBow, ItemTagBallistic},
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
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupBallisticProjectiles,
// 		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
// 		Accuracy:         7,
// 		DamageValue:      5,
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		ArmorPenatration: -1,
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Availability:     2,
// 		Legality:         LegalityTypeLegal,
// 		Cost:             300,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagCrossbow, ItemTagBallistic},
// 	},
// 	{
// 		ID:               "medium_crossbow",
// 		Name:             "Medium Crossbow",
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupBallisticProjectiles,
// 		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
// 		Accuracy:         6,
// 		DamageValue:      7,
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		ArmorPenatration: -2,
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Availability:     4,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             500,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagCrossbow, ItemTagBallistic},
// 	},
// 	{
// 		ID:               "heavy_crossbow",
// 		Name:             "Heavy Crossbow",
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupBallisticProjectiles,
// 		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
// 		Accuracy:         5,
// 		DamageValue:      10,
// 		DamageType:       DamageTypePhysical,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		ArmorPenatration: -3,
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Availability:     8,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             1000,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagCrossbow, ItemTagBallistic},
// 	},
// 	{
// 		ID:               "grapple_gun",
// 		Name:             "Grapple Gun",
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupBallisticProjectiles,
// 		Description:      "A grapple gun is a device that allows the user to fire a grappling hook or a similar object to a distant location.",
// 		Accuracy:         3,
// 		DamageValue:      7,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -2,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		AmmoCapacity:     1,
// 		Reload:           WeaponRangedReloadMuzzleLoader,
// 		Availability:     8,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             500,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagExotic},
// 	},
// 	{
// 		ID:               "defiance_ex_shocker",
// 		Name:             "Defiance EX Shocker",
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupTasers,
// 		Accuracy:         4,
// 		DamageValue:      9,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -5,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSingleShot},
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Cost:             250,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagTaser},
// 		Modifiers:        []Modifier{{Type: "add", Effect: "ElectricDamage"}},
// 		// Wireless
// 		//  A successful hit informs you of the status of the target’s basic health (and Condition Monitors).
// 	},
// 	{
// 		ID:               "yamaha_pulsar",
// 		Name:             "Yamaha Pulsar",
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupFirearms,
// 		Category:         WeaponCategoryPistol,
// 		SubCategory:      WeaponSubCategoryHoldOutPistol,
// 		Accuracy:         5,
// 		DamageValue:      7,
// 		DamageType:       DamageTypeStun,
// 		ArmorPenatration: -5,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadInternalMagazine,
// 		Cost:             180,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagTaser},
// 		Modifiers:        []Modifier{{Type: "add", Effect: "ElectricDamage"}},
// 		// Wireless
// 		//  A successful hit informs you of the status of the target’s basic health (and Condition Monitors).
// 	},
// 	{
// 		ID:               "ficetti_tiffani_needler",
// 		Name:             "Fichetti Tiffani Needler",
// 		Type:             WeaponTypeRanged,
// 		Group:            WeaponGroupTasers,
// 		Description:      "The Fichetti Tiffani Needler is a hold-out pistol that fires flechette ammunition. It is a small, easily concealed weapon that is popular with shadowrunners and criminals.",
// 		Accuracy:         5,
// 		DamageValue:      8,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: 5,
// 		FiringModes:      []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		AmmoCapacity:     4,
// 		Reload:           WeaponRangedReloadDetachableMagazine,
// 		Availability:     6,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             1000,
// 		RuleSource:       RuleSourceSR5Core,
// 		Tags:             []ItemTag{ItemTagRanged, ItemTagPistol, ItemTagHoldOutPistol},
// 		// 8P(f)
// 		// Wireless
// 		//  You can change the color of the Tiffani Needler with a Simple Action.
// 		// Can only fire flechette rounds
// 	},
// 	{
// 		ID:           "streetline_special",
// 		Name:         "Streetline Special",
// 		Type:         WeaponTypeRanged,
// 		Group:        WeaponGroupFirearms,
// 		Category:     WeaponCategoryPistol,
// 		SubCategory:  WeaponSubCategoryHoldOutPistol,
// 		Description:  "The Streetline Special is a hold-out pistol that is popular with shadowrunners and criminals. It is a small, easily concealed weapon.",
// 		Accuracy:     4,
// 		DamageValue:  6,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		AmmoCapacity: 6,
// 		Reload:       WeaponRangedReloadDetachableMagazine,
// 		Availability: 4,
// 		Legality:     LegalityTypeRestricted,
// 		Cost:         120,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagRanged, ItemTagPistol, ItemTagHoldOutPistol},
// 		// MAD Scanner
// 		//  -2 DP to detect Streetline Special
// 	},
// 	{
// 		ID:           "walther_palm_pistol",
// 		Name:         "Walther Palm Pistol",
// 		Type:         WeaponTypeRanged,
// 		Group:        WeaponGroupFirearms,
// 		Category:     WeaponCategoryPistol,
// 		SubCategory:  WeaponSubCategoryHoldOutPistol,
// 		Description:  "The Walther Palm Pistol is a hold-out pistol that is popular with shadowrunners and criminals. It is a small, easily concealed weapon.",
// 		Accuracy:     4,
// 		DamageValue:  7,
// 		DamageType:   DamageTypePhysical,
// 		FiringModes:  []WeaponFiringMode{WeaponFiringModeSingleShot, WeaponFiringModeBurstFire},
// 		AmmoCapacity: 2,
// 		Reload:       WeaponRangedReloadBreakAction,
// 		Availability: 4,
// 		Legality:     LegalityTypeRestricted,
// 		Cost:         180,
// 		RuleSource:   RuleSourceSR5Core,
// 		Tags:         []ItemTag{ItemTagRanged, ItemTagPistol, ItemTagHoldOutPistol},
// 	},
// }
