package weapon

import "github.com/Jasrags/ShadowMUD/common/shared"

var CoreWeapons = []Spec{
	//Melee
	{
		ID:           "club",
		Name:         "Club",
		Type:         WeaponTypeMelee,
		Group:        WeaponGroupClubs,
		Category:     WeaponCategoryClub,
		Description:  "The weapon they named the skill after.",
		Accuracy:     4,
		Reach:        1,
		DamageValue:  2,
		DamageType:   shared.DamageTypePhysical,
		Availability: 4,
		Legality:     shared.LegalityTypeLegal,
		Cost:         30,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagClub},
	},
	{
		ID:           "extendable_baton",
		Name:         "Extendable Baton",
		Type:         WeaponTypeMelee,
		Group:        WeaponGroupClubs,
		Category:     WeaponCategoryClub,
		Description:  "A baton that extends to a full length.",
		Accuracy:     5,
		Reach:        1,
		DamageValue:  2,
		DamageType:   shared.DamageTypePhysical,
		Availability: 4,
		Legality:     shared.LegalityTypeLegal,
		Cost:         100,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBaton},
		// Simple Action
		//  Collapse/Extend
		// Collapsed
		//  Concealability Modifier 0
		// Extended
		//  Concealability Modifier +2
		// Wireless
		//  Readying the extendable baton is a Free Action instead of a Simple Action.
	},
	{
		ID:             "sap",
		Name:           "Sap",
		Type:           WeaponTypeMelee,
		Group:          WeaponGroupClubs,
		Category:       WeaponCategoryClub,
		Description:    "A small, weighted club.",
		Concealability: -2,
		Accuracy:       5,
		DamageValue:    2,
		DamageType:     shared.DamageTypePhysical,
		Availability:   2,
		Legality:       shared.LegalityTypeLegal,
		Cost:           30,
		RuleSource:     shared.RuleSourceSR5Core,
		Tags:           []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagClub},
		// Concealability
		//  Modifier +2
	},
	{
		ID:           "staff",
		Name:         "Staff",
		Type:         WeaponTypeMelee,
		Group:        WeaponGroupClubs,
		Category:     WeaponCategoryClub,
		Description:  "A long stick.",
		Accuracy:     6,
		Reach:        2,
		DamageValue:  3,
		DamageType:   shared.DamageTypePhysical,
		Availability: 3,
		Legality:     shared.LegalityTypeLegal,
		Cost:         100,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagStaff},
	},
	{
		ID:               "stun_baton",
		Name:             "Stun Baton",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "A baton that delivers an electrical shock.",
		Accuracy:         4,
		Reach:            1,
		DamageValue:      9,
		DamageType:       shared.DamageTypeStun,
		ArmorPenatration: -5,
		Availability:     6,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             750,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBaton},
		// Modifiers: []Modifier{
		// 	{
		// 		Type:   ModifierTypeElectric,
		// 		Effect: ModifierEffectAttack,
		// 	},
		// },
		// Wireless
		//  The stun baton recharges by induction, regaining one charge per full hour of wireless-enabled time.
		// Modifier
		//  Electric attack
	},
	{
		ID:           "telescoping_staff",
		Name:         "Telescoping Staff",
		Type:         WeaponTypeMelee,
		Group:        WeaponGroupClubs,
		Category:     WeaponCategoryClub,
		Description:  "A staff that extends to a full length.",
		Accuracy:     4,
		Reach:        2,
		DamageValue:  4,
		DamageType:   shared.DamageTypePhysical,
		Availability: 4,
		Legality:     shared.LegalityTypeLegal,
		Cost:         350,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagStaff},
	},
	// Blades
	{
		ID:               "combat_axe",
		Name:             "Combat Axe",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "A large axe designed for combat.",
		Accuracy:         4,
		Reach:            2,
		DamageValue:      5,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -4,
		Availability:     12,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             4000,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade, shared.ItemTagAxe},
	},
	{
		ID:               "combat_knife",
		Name:             "Combat Knife",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "A knife designed for combat.",
		Accuracy:         6,
		DamageValue:      2,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -3,
		Availability:     4,
		Legality:         shared.LegalityTypeLegal,
		Cost:             300,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade},
	},
	{
		ID:               "forearm_snap_blades",
		Name:             "Forearm Snap-Blades",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "Blades that extend from the forearm.",
		Accuracy:         4,
		DamageValue:      2,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     7,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             200,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade},
		// Wireless
		//  Readying the forearm snap blades is a Free Action instead of a Simple Action.
	},
	{
		ID:               "katana",
		Name:             "Katana",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "A traditional Japanese sword.",
		Accuracy:         7,
		Reach:            1,
		DamageValue:      3,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -3,
		Availability:     9,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             1000,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade, shared.ItemTagKatana},
	},
	{
		ID:               "knife",
		Name:             "Knife",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "A small knife.",
		Accuracy:         5,
		DamageValue:      1,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -1,
		Legality:         shared.LegalityTypeLegal,
		Cost:             10,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade, shared.ItemTagKnife},
	},
	{
		ID:               "pole_arm",
		Name:             "Pole Arm",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "A pole arm is a large weapon with a blade on the end.",
		Accuracy:         5,
		Reach:            3,
		DamageValue:      3,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     6,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             1000,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade, shared.ItemTagPolearm},
	},
	{
		ID:               "survival_knife",
		Name:             "Survival Knife",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      `A survival knife is a large knife with a serrated edge. It is designed for use in the wilderness, and is often used by military personnel and survivalists.`,
		Accuracy:         5,
		DamageValue:      2,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -1,
		Legality:         shared.LegalityTypeLegal,
		Cost:             100,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade, shared.ItemTagKnife},
		// Wireless
		//  The knife displays an ARO of local maps, your GPS position, and can be used to make commcalls.
	},
	{
		ID:               "sword",
		Name:             "Sword",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupClubs,
		Category:         WeaponCategoryClub,
		Description:      "A sword is a long, sharp blade.",
		Accuracy:         6,
		Reach:            1,
		DamageValue:      3,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     5,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             500,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagBlade, shared.ItemTagSword},
	},
	// Exotic Melee Weapons
	{
		ID:               "monofilament_chainsaw",
		Name:             "Monofilament Chainsaw",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupExotic,
		Description:      "A chainsaw with a monofilament blade.",
		Accuracy:         3,
		Reach:            1,
		DamageValue:      8,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -6,
		Availability:     8,
		Legality:         shared.LegalityTypeLegal,
		Cost:             500,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagExotic, shared.ItemTagChainsaw, shared.ItemTagMonofiliament},
		// When used against barriers, double the monofilament chainsaw’s Damage Value of 8P
	},
	{
		ID:               "monofilament_whip",
		Name:             "Monofilament Whip",
		Type:             WeaponTypeMelee,
		Group:            WeaponGroupExotic,
		Description:      "A whip with a monofilament blade.",
		Accuracy:         5,
		Reach:            2,
		DamageValue:      12,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: -8,
		Availability:     12,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             10000,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagExotic, shared.ItemTagWhip, shared.ItemTagMonofiliament},
		// Skill
		//  Exotic Melee Weapon (Monofilament Whip)
		// Glitch
		//  You catch the weighted tip on something nearby and need to disentangle it before you make another proper attack with it.
		// Critical Glitch
		//  You hit yourself with the whip and take its base damage (resisted normally)
		// Wireless
		//  The whip can be readied with a Free Action, rather than a Simple Action.
		//  The whip's built-in safety system retracts automatically instead of getting you entangled on a glitch.
		//  Accuracy increased by 2.
	},
	// Misc Melee Weapons
	{
		ID:          "knucks",
		Name:        "Knucks",
		Type:        WeaponTypeMelee,
		Group:       WeaponGroupMisc,
		Description: "These may be traditional fist-load weapons like brass knuckles, or modern Hardliner Gloves with a thin layer of densiplast located the knuckles and the edge of the hand. Either way they substantially boost the impact of a punch, making it potentially deadly.",
		// Accuracy:    0, // Physical
		DamageValue:  1,
		DamageType:   shared.DamageTypePhysical,
		Availability: 2,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         100,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagMiscellaneous},
		// Skill
		//  Unarmed Combat Skill
	},
	{
		ID:          "plasteel_toe_boots",
		Name:        "Plasteel Toe Boots",
		Type:        WeaponTypeMelee,
		Group:       WeaponGroupMisc,
		Description: "These boots have a reinforced toe made of plasteel, making them a formidable weapon in a kick.",
		// Accuracy:    0, // Physical
		DamageValue:  1,
		DamageType:   shared.DamageTypePhysical,
		Availability: 2,
		Legality:     shared.LegalityTypeLegal,
		Cost:         200,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagMiscellaneous},
		// Skill
		//  Unarmed Combat Skill
	},
	{
		ID:          "shock_gloves",
		Name:        "Shock Gloves",
		Type:        WeaponTypeMelee,
		Group:       WeaponGroupMisc,
		Description: "These gloves are lined with a layer of conductive material that delivers an electric shock to the target.",
		// Accuracy:    0, // Physical
		DamageValue:      8,
		DamageType:       shared.DamageTypeStun,
		ArmorPenatration: -5,
		Availability:     6,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             550,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagMelee, shared.ItemTagMiscellaneous},
		// Modifiers:        []Modifier{{Type: "Electric attack", Effect: "Attack"}},
		// Skill
		//  Unarmed Combat Skill
		// Wireless
		//  The shock gloves recharge by induction, regaining one charge per full hour of wireless-enabled time.
		// Modifier
		//  Electric attack
	},
	// Ranged
	{
		ID:           "shuriken",
		Name:         "Shuriken",
		Type:         WeaponTypeRanged,
		Group:        WeaponGroupThrowingWeapons,
		Description:  "A shuriken is a small, star-shaped piece of metal with sharpened edges, designed for throwing. It is also known as a “throwing star” or “ninja star.”",
		DamageValue:  1,
		DamageType:   shared.DamageTypePhysical,
		FiringModes:  []FiringMode{WeaponFiringModeSingleShot},
		Recoil:       -1,
		Availability: 4,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         25,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagThrowing},
		// Wireless
		//  If all the throwing knives or shuriken you throw in a single Combat Turn are wireless and you have a smartlink system, each knife you throw receives a +1 dice pool bonus per knife thrown that Combat Turn at your current target, as the knives inform and adjust for wind and other atmospheric conditions. So you’d get no bonus on the first throw, a +1 bonus on the second throw, a +2 bonus on the third throw, etc. (assuming you aimed all three knives at the same target).
	},
	{
		ID:           "throwing_knife",
		Name:         "Throwing Knife",
		Type:         WeaponTypeRanged,
		Group:        WeaponGroupThrowingWeapons,
		Description:  "A throwing knife is a knife that is specially designed and weighted so that it can be thrown effectively.",
		DamageValue:  1,
		DamageType:   shared.DamageTypePhysical,
		FiringModes:  []FiringMode{WeaponFiringModeSingleShot},
		Recoil:       -1,
		Availability: 4,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         25,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagThrowing},
		// Wireless
		//  If all the throwing knives or shuriken you throw in a single Combat Turn are wireless and you have a smartlink system, each knife you throw receives a +1 dice pool bonus per knife thrown that Combat Turn at your current target, as the knives inform and adjust for wind and other atmospheric conditions. So you’d get no bonus on the first throw, a +1 bonus on the second throw, a +2 bonus on the third throw, etc. (assuming you aimed all three knives at the same target).
	},
	{
		ID:               "bow",
		Name:             "Bow",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupBallisticProjectiles,
		Description:      "A bow is a flexible arc that shoots aerodynamic projectiles called arrows. A string joins the two ends of the bow and when the string is drawn back, the ends of the bow are flexed. When the string is released, the potential energy of the flexed bow limbs is transformed into the velocity of the arrow.",
		Accuracy:         6,
		DamageValue:      2, // (Rating+2)
		DamageType:       shared.DamageTypePhysical,
		FiringModes:      []FiringMode{WeaponFiringModeSingleShot},
		ArmorPenatration: 0,   // -(Rating/4)
		Availability:     1,   // Rating
		Cost:             100, // Rating×100¥
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagBow, shared.ItemTagBallistic},
		// Max Rating
		//     10
		// Strength Minimum
		//     If Strength is less than Rating then -3 DP per point below Rating.
		// Damage Rating
		//     Lowest value of Strength, Bows Rating, or Arrow Rating.
		// Range Rating
		//     Lowest value of Strength, Bows Rating, or Arrow Rating.
		// Reload
		//     Simple Action
		// When attacking with a bow, a character whose Strength is less than the Rating suffers a –3 dice pool modifier per point below the minimum
		// Use the lowest value of your Strength, the bow’s rating, or the arrow Rating for range and damage when attacking a target, because your average Rating 10
	},
	{
		ID:               "light_crossbow",
		Name:             "Light Crossbow",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupBallisticProjectiles,
		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
		Accuracy:         7,
		DamageValue:      5,
		DamageType:       shared.DamageTypePhysical,
		FiringModes:      []FiringMode{WeaponFiringModeSemiAutomatic},
		ArmorPenatration: -1,
		AmmoCapacity:     4,
		Reload:           WeaponRangedReloadInternalMagazine,
		Availability:     2,
		Legality:         shared.LegalityTypeLegal,
		Cost:             300,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagCrossbow, shared.ItemTagBallistic},
	},
	{
		ID:               "medium_crossbow",
		Name:             "Medium Crossbow",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupBallisticProjectiles,
		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
		Accuracy:         6,
		DamageValue:      7,
		DamageType:       shared.DamageTypePhysical,
		FiringModes:      []FiringMode{WeaponFiringModeSemiAutomatic},
		ArmorPenatration: -2,
		AmmoCapacity:     4,
		Reload:           WeaponRangedReloadInternalMagazine,
		Availability:     4,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             500,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagCrossbow, shared.ItemTagBallistic},
	},
	{
		ID:               "heavy_crossbow",
		Name:             "Heavy Crossbow",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupBallisticProjectiles,
		Description:      "A crossbow is a type of ranged weapon based on the bow and consisting of a horizontal bow-like assembly mounted on a frame which is handheld in a similar fashion to the stock of a gun. It shoots arrow-like projectiles called bolts or quarrels.",
		Accuracy:         5,
		DamageValue:      10,
		DamageType:       shared.DamageTypePhysical,
		FiringModes:      []FiringMode{WeaponFiringModeSemiAutomatic},
		ArmorPenatration: -3,
		AmmoCapacity:     4,
		Reload:           WeaponRangedReloadInternalMagazine,
		Availability:     8,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             1000,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagCrossbow, shared.ItemTagBallistic},
	},
	{
		ID:               "grapple_gun",
		Name:             "Grapple Gun",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupBallisticProjectiles,
		Description:      "A grapple gun is a device that allows the user to fire a grappling hook or a similar object to a distant location.",
		Accuracy:         3,
		DamageValue:      7,
		DamageType:       shared.DamageTypeStun,
		ArmorPenatration: -2,
		FiringModes:      []FiringMode{WeaponFiringModeSingleShot},
		AmmoCapacity:     1,
		Reload:           WeaponRangedReloadMuzzleLoader,
		Availability:     8,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             500,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagExotic},
	},
	{
		ID:               "defiance_ex_shocker",
		Name:             "Defiance EX Shocker",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupTasers,
		Accuracy:         4,
		DamageValue:      9,
		DamageType:       shared.DamageTypeStun,
		ArmorPenatration: -5,
		FiringModes:      []FiringMode{WeaponFiringModeSingleShot},
		AmmoCapacity:     4,
		Reload:           WeaponRangedReloadInternalMagazine,
		Cost:             250,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagTaser},
		// Modifiers:        []Modifier{{Type: "add", Effect: "ElectricDamage"}},
		// Wireless
		//  A successful hit informs you of the status of the target’s basic health (and Condition Monitors).
	},
	{
		ID:               "yamaha_pulsar",
		Name:             "Yamaha Pulsar",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupFirearms,
		Category:         WeaponCategoryPistol,
		SubCategory:      WeaponSubCategoryHoldOutPistol,
		Accuracy:         5,
		DamageValue:      7,
		DamageType:       shared.DamageTypeStun,
		ArmorPenatration: -5,
		FiringModes:      []FiringMode{WeaponFiringModeSemiAutomatic},
		AmmoCapacity:     4,
		Reload:           WeaponRangedReloadInternalMagazine,
		Cost:             180,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagTaser},
		// Modifiers:        []Modifier{{Type: "add", Effect: "ElectricDamage"}},
		// Wireless
		//  A successful hit informs you of the status of the target’s basic health (and Condition Monitors).
	},
	{
		ID:               "ficetti_tiffani_needler",
		Name:             "Fichetti Tiffani Needler",
		Type:             WeaponTypeRanged,
		Group:            WeaponGroupTasers,
		Description:      "The Fichetti Tiffani Needler is a hold-out pistol that fires flechette ammunition. It is a small, easily concealed weapon that is popular with shadowrunners and criminals.",
		Accuracy:         5,
		DamageValue:      8,
		DamageType:       shared.DamageTypePhysical,
		ArmorPenatration: 5,
		FiringModes:      []FiringMode{WeaponFiringModeSemiAutomatic},
		AmmoCapacity:     4,
		Reload:           WeaponRangedReloadDetachableMagazine,
		Availability:     6,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             1000,
		RuleSource:       shared.RuleSourceSR5Core,
		Tags:             []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagPistol, shared.ItemTagHoldOutPistol},
		// 8P(f)
		// Wireless
		//  You can change the color of the Tiffani Needler with a Simple Action.
		// Can only fire flechette rounds
	},
	{
		ID:           "streetline_special",
		Name:         "Streetline Special",
		Type:         WeaponTypeRanged,
		Group:        WeaponGroupFirearms,
		Category:     WeaponCategoryPistol,
		SubCategory:  WeaponSubCategoryHoldOutPistol,
		Description:  "The Streetline Special is a hold-out pistol that is popular with shadowrunners and criminals. It is a small, easily concealed weapon.",
		Accuracy:     4,
		DamageValue:  6,
		DamageType:   shared.DamageTypePhysical,
		FiringModes:  []FiringMode{WeaponFiringModeSemiAutomatic},
		AmmoCapacity: 6,
		Reload:       WeaponRangedReloadDetachableMagazine,
		Availability: 4,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         120,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagPistol, shared.ItemTagHoldOutPistol},
		// MAD Scanner
		//  -2 DP to detect Streetline Special
	},
	{
		ID:           "walther_palm_pistol",
		Name:         "Walther Palm Pistol",
		Type:         WeaponTypeRanged,
		Group:        WeaponGroupFirearms,
		Category:     WeaponCategoryPistol,
		SubCategory:  WeaponSubCategoryHoldOutPistol,
		Description:  "The Walther Palm Pistol is a hold-out pistol that is popular with shadowrunners and criminals. It is a small, easily concealed weapon.",
		Accuracy:     4,
		DamageValue:  7,
		DamageType:   shared.DamageTypePhysical,
		FiringModes:  []FiringMode{WeaponFiringModeSingleShot, WeaponFiringModeBurstFire},
		AmmoCapacity: 2,
		Reload:       WeaponRangedReloadBreakAction,
		Availability: 4,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         180,
		RuleSource:   shared.RuleSourceSR5Core,
		Tags:         []shared.ItemTag{shared.ItemTagRanged, shared.ItemTagPistol, shared.ItemTagHoldOutPistol},
	},
}

var CoreAmunition = []AmmunitionSpec{
	{
		ID:                       "apds",
		Name:                     "APDS",
		Description:              "These are military-grade armor piercing rounds—their full name is armor piercing discarding sabot. They are designed to travel at high velocities and punch through personal body armor.",
		ArmorPenatrationModifier: -4,
		Availability:             12,
		Legality:                 shared.LegalityTypeForbidden,
		Cost:                     120,
		CostFor:                  10,
		Modifiers: shared.Modifiers{
			{
				Type:   shared.ModifierTypeArmorPenetration,
				Effect: shared.ModifierEffectAdd,
				Value:  -4,
			},
		},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "assault_cannon",
		Name:         "Assault Cannon",
		Description:  "These are for assault cannons only, and they’re the only thing assault cannons can load.",
		Availability: 12,
		Legality:     shared.LegalityTypeForbidden,
		Cost:         400,
		CostFor:      10,
		RuleSource:   shared.RuleSourceSR5Core,
	},
	{
		ID:                       "explosive_rounds",
		Name:                     "Explosive Rounds",
		Description:              "These slugs carry a shaped-charge explosive, designed to explode and fragment on impact. Explosive rounds misfire whenever you roll a critical glitch. When this happens, you must resist one “attack” with a Damage Value equal to the normal damage done by the weapon (and don’t forget the modifier for the explosive rounds). The attack misses its intended target, and the weapon firing the bullets is destroyed.",
		DamageValueModifier:      1,
		ArmorPenatrationModifier: -1,
		Availability:             9,
		Legality:                 shared.LegalityTypeForbidden,
		Cost:                     80,
		CostFor:                  10,
		RuleSource:               shared.RuleSourceSR5Core,
	},
	{
		ID:                       "flechette_rounds",
		Name:                     "Flechette Rounds",
		Description:              "The payload of a flechette round is made up of tiny, tightly packed metal slivers. The round breaks up and shatters on impact, becoming a tumbling hail of shrapnel. Flechette rounds are devastating against unprotected targets, but not as effective against hardened armor.",
		DamageValueModifier:      2,
		ArmorPenatrationModifier: 5,
		Availability:             6,
		Legality:                 shared.LegalityTypeForbidden,
		Cost:                     65,
		RuleSource:               shared.RuleSourceSR5Core,
	},
	{
		ID:                       "gel_rounds",
		Name:                     "Gel Rounds",
		Description:              "Gel rounds are designed to deliver a non-lethal payload. They are often used by security forces and police to subdue suspects without causing permanent harm. Gel rounds are less effective against armored targets.",
		DamageType:               shared.DamageTypeStun,
		ArmorPenatrationModifier: 1,
		Availability:             2,
		Legality:                 shared.LegalityTypeRestricted,
		Cost:                     25,
		RuleSource:               shared.RuleSourceSR5Core,
	},
	{
		ID:                       "hollow_points",
		Name:                     "Hollow Points",
		Description:              "Hollow points are designed to expand on impact, causing more damage to soft targets. They are less effective against armored targets.",
		DamageValueModifier:      1,
		ArmorPenatrationModifier: 2,
		Availability:             4,
		Legality:                 shared.LegalityTypeForbidden,
		Cost:                     70,
		RuleSource:               shared.RuleSourceSR5Core,
	},
	{
		ID:           "injection_darts",
		Name:         "Injection Darts",
		Description:  "Injection darts are used to deliver drugs or toxins to a target. They are fired from a dart pistol or a dart rifle. Injection darts are not designed to cause damage, but to deliver a payload to the target. Injection darts are not available for all weapons, and are not available for all drugs or toxins.",
		Availability: 4,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         75,
		RuleSource:   shared.RuleSourceSR5Core,
	},
	{
		ID:           "regular_ammo",
		Name:         "Regular Ammo",
		Description:  "Regular ammo is the standard ammunition for most firearms. It is a lead slug encased in a copper jacket. Regular ammo is effective against most targets, but is less effective against armored targets.",
		Availability: 2,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         20,
		RuleSource:   shared.RuleSourceSR5Core,
	},
	{
		ID:                       "stick_n_shock",
		Name:                     "Stick-n-Shock",
		Description:              "Stick-n-Shock rounds are designed to deliver an electrical shock to the target. They are often used by security forces and police to subdue suspects without causing permanent harm. Stick-n-Shock rounds are less effective against armored targets.",
		DamageValueModifier:      -2,
		DamageType:               shared.DamageTypeStun,
		ArmorPenatrationModifier: -5,
		Availability:             6,
		Legality:                 shared.LegalityTypeRestricted,
		Cost:                     80,
		// Modifiers: []Modifier{
		// 	{
		// 		Type:   "add",
		// 		Effect: "ElectricDamage",
		// 	},
		// },
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "taser_darts",
		Name:         "Taser Dart",
		Description:  "Taser darts are used to deliver an electrical shock to a target. They are fired from a dart pistol or a dart rifle. Taser darts are not designed to cause damage, but to deliver an electrical shock to the target. Taser darts are not available for all weapons.",
		Availability: 3,
		Legality:     shared.LegalityTypeLegal,
		Cost:         50,
		RuleSource:   shared.RuleSourceSR5Core,
	},
	{
		ID:               "tracer_rounds",
		Name:             "Tracer",
		Description:      "Tracer rounds are designed to leave a visible trail in the air, making it easier to track the path of the bullet. Tracer rounds are often used to help a shooter adjust their aim. Tracer rounds are less effective against armored targets.",
		AccuracyModifier: 1,
		Availability:     6,
		Legality:         shared.LegalityTypeRestricted,
		Cost:             60,
		RuleSource:       shared.RuleSourceSR5Core,
	},
}

// {
//     "id": "4d96ee55-0929-480e-923d-178e69675541",
//     "name": "Throwing Syringe",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "183",
//     "avail": "6F",
//     "addweapon": "Throwing Syringe",
//     "cost": "40",
//     "costfor": "1"
//   },
//   {
//     "id": "60ab05a7-2850-410c-9261-c5df6b5cddf1",
//     "name": "Seeker Shaft",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "187",
//     "avail": "12F",
//     "cost": "45",
//     "costfor": "1",
//     "requireparent": null,
//     "required": {
//       "parentdetails": {
//         "ammoforweapontype": "bow"
//       }
//     },
//     "weaponbonus": {
//       "smartlinkpool": "1"
//     }
//   },
//   {
//     "id": "39042c0a-4b22-4762-8d56-623033860178",
//     "name": "Arrow: Monotip Head",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "HT",
//     "page": "187",
//     "avail": "8R",
//     "cost": "(Rating * 30)",
//     "costfor": "1",
//     "weaponbonus": {
//       "ap": "-2"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "0198f16e-8bb2-4618-ab2a-f010abe0a1c6",
//     "name": "Ammo: Depleted Uranium",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "189",
//     "avail": "28F",
//     "cost": "1000",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "-5",
//       "damage": "1"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "adef0fa4-67b7-41e1-b534-f15308c9c8c1",
//     "name": "Ammo: Silver",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "189",
//     "avail": "12R",
//     "cost": "250",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "2"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "60b7ad79-55b7-4803-8d1f-9145218eaf5b",
//     "name": "Ammo: Wood Pulp",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "189",
//     "avail": "6R",
//     "cost": "10",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "4",
//       "damage": "-4"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "b286a62d-c7cb-472a-bea8-0f41a300bb9c",
//     "name": "Ammo: Hi-De Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "189",
//     "avail": "10F",
//     "cost": "150",
//     "costfor": "10",
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "c429fc81-d11b-434b-b98b-5fbdda9db905",
//     "name": "Ammo: Subsonic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "189",
//     "avail": "8F",
//     "cost": "40",
//     "costfor": "10",
//     "weaponbonus": {
//       "damage": "-1"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "75ccb148-e774-429c-b854-a27816439626",
//     "name": "Spare Clip",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "433",
//     "avail": "4",
//     "bonus": {
//       "selectweapon": {
//         "+@weapondetails": "(contains(ammo, '(c)') or contains(ammo, '(d)')) and name != 'HK Urban Fighter'"
//       }
//     },
//     "cost": "5"
//   },
//   {
//     "id": "f87701a0-4ea2-47db-bcac-f5b8396c369e",
//     "name": "Speed Loader",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "433",
//     "avail": "2",
//     "bonus": {
//       "selectweapon": {
//         "+@weapondetails": "contains(ammo, '(cy)')"
//       }
//     },
//     "cost": "25"
//   },
//   {
//     "id": "ef9c8aae-26df-4fe6-88b3-79fbb5eb77c5",
//     "name": "Ammo: APDS",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "12F",
//     "cost": "120",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "-4"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "0bf4113c-dedb-411d-8981-7b6af169f056",
//     "name": "Ammo: Assault Cannon",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "12F",
//     "cost": "400",
//     "costfor": "10",
//     "ammoforweapontype": "cannon"
//   },
//   {
//     "id": "1315cecd-1c13-4d69-9828-a3ea535675da",
//     "name": "Ammo: Explosive Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "9F",
//     "cost": "80",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "-1",
//       "damage": "1"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "95bff6e2-d788-407b-9069-093250f89fcb",
//     "name": "Ammo: Flechette Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "6R",
//     "cost": "65",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "5",
//       "damage": "2",
//       "damagetype": "P(f)"
//     },
//     "isflechetteammo": "True",
//     "flechetteweaponbonus": {
//       "damagetype": "P(f)"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "0dbcca24-c152-47c5-bae2-8feadff22321",
//     "name": "Ammo: Stick-n-Shock Flechette Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "TSG",
//     "page": "29",
//     "avail": "6R",
//     "cost": "110",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "S(e)"
//     },
//     "isflechetteammo": "True",
//     "flechetteweaponbonus": {
//       "ap": "-5",
//       "damage": "-2",
//       "damagetype": "S(e)"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "0c8d16cb-6e96-4d95-8454-104a36091cf9",
//     "name": "Ammo: Gel Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "2R",
//     "cost": "25",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "1",
//       "damagetype": "S"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "f486f414-ee2c-46db-92ea-c682861d8fe0",
//     "name": "Ammo: Hollow Points",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "4F",
//     "cost": "70",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "2",
//       "damage": "1"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "638c81a2-328b-4e22-8fb0-ee37c5e2f6c9",
//     "name": "Ammo: Injection Darts",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "4R",
//     "addoncategory": [
//       "Drugs",
//       "Toxins",
//       "Custom"
//     ],
//     "cost": "75",
//     "costfor": "10",
//     "ammoforweapontype": {
//       "+content": "dartgun",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "63AA2CA3-E0B7-4193-8B15-290F6B5DD21E",
//     "name": "Ammo: Taurus Omni-6 Heavy",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "2R",
//     "cost": "20",
//     "costfor": "10",
//     "weaponbonus": {
//       "apreplace": "-1",
//       "damagereplace": "7P",
//       "modereplace": "SS"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "d27dfc89-095d-440c-ad1c-c7f888df2824",
//     "name": "Ammo: Peak-Discharge Energy Units",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "22",
//     "avail": "0",
//     "cost": "0",
//     "ammoforweapontype": "energy"
//   },
//   {
//     "id": "b2a0b340-c793-4322-8422-8b03d18a6fae",
//     "name": "Ammo: Regular Ammo",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "2R",
//     "cost": "20",
//     "costfor": "10",
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "d9f69780-93eb-41ff-9a9c-893f8c52794e",
//     "name": "Ammo: Stick-n-Shock",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "6R",
//     "cost": "80",
//     "costfor": "10",
//     "weaponbonus": {
//       "apreplace": "-5",
//       "damage": "-2",
//       "damagetype": "S(e)"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "e0d7aea7-52ac-4670-bac3-b13ce144257c",
//     "name": "Ammo: Tracer",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "6R",
//     "cost": "60",
//     "costfor": "10",
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "8afe5065-5815-4a98-b047-2a7fed85db55",
//     "name": "Ammo: Taser Dart",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "434",
//     "avail": "3",
//     "cost": "50",
//     "costfor": "10",
//     "ammoforweapontype": "taser"
//   },
//   {
//     "id": "96eb3f57-26e0-40c5-8dac-62cf8a0bc52f",
//     "name": "Ammo: EX-Explosive Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "54",
//     "avail": "14F",
//     "cost": "120",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "-1",
//       "damage": "2"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "ecc8aafc-f8a3-408b-9b47-6f85a473de13",
//     "name": "Ammo: Frangible Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "54",
//     "avail": "2R",
//     "cost": "10",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "4",
//       "damage": "-1"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "506dd9c6-ea57-4ae2-8887-81ce93e5a7b2",
//     "name": "Ammo: Flare Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "55",
//     "avail": "6R",
//     "cost": "20",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "2",
//       "damage": "-2"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "f29f0642-65a5-4461-bca1-396907dbf82e",
//     "name": "Ammo: Tracker Rounds, Security Tag",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "55",
//     "avail": "8R",
//     "cost": "150",
//     "costfor": "10",
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "4b844ac3-20e8-4982-8711-20579b5295a8",
//     "name": "Ammo: Tracker Rounds, Stealth Tag",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "55",
//     "avail": "8R",
//     "cost": "150",
//     "costfor": "10",
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "4257a39e-41eb-45ee-976e-41c3c0bf960c",
//     "name": "Ammo: Capsule Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "55",
//     "avail": "2",
//     "addoncategory": [
//       "Drugs",
//       "Toxins",
//       "Custom"
//     ],
//     "cost": "5",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "4",
//       "damage": "-4",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "9bb893c0-2cab-4f7f-a0cf-bdacc940254b",
//     "name": "Ammo: DMSO Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "436",
//     "avail": "2R",
//     "addoncategory": [
//       "Drugs",
//       "Toxins",
//       "Custom"
//     ],
//     "cost": "20",
//     "costfor": "10",
//     "ammoforweapontype": "squirtgun"
//   },
//   {
//     "id": "6fbc580a-50f8-481b-99b0-84243e53fc12",
//     "name": "Ammo: Narcoject Gas Gun Canisters",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "43",
//     "avail": "2R",
//     "addoncategory": [
//       "Drugs",
//       "Toxins",
//       "Custom"
//     ],
//     "cost": "20",
//     "costfor": "10",
//     "ammoforweapontype": "gasgun"
//   },
//   {
//     "id": "ff6aa1f8-d0a5-4d24-af56-2bc2d75bfdd5",
//     "name": "Ammo: Narcoject Trackstopper Foam",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "44",
//     "avail": "15R",
//     "cost": "500",
//     "costfor": "6",
//     "ammoforweapontype": "trackstopper"
//   },
//   {
//     "id": "9e7a685f-e612-4b6b-955c-a7bc42b23682",
//     "name": "Arrow: Standard",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "SR5",
//     "page": "423",
//     "avail": "Rating",
//     "cost": "Rating * 2",
//     "costfor": "1",
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "775d2fed-57ab-4e63-b2e4-638ccf4a21d0",
//     "name": "Arrow: Injection",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "SR5",
//     "page": "424",
//     "avail": "(Rating + 2)R",
//     "cost": "Rating * 20",
//     "costfor": "1",
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "c3c2fa27-36a7-4296-aadd-627078e6e052",
//     "name": "Bolt: Standard",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "424",
//     "avail": "2",
//     "cost": "5",
//     "costfor": "1",
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "7d697849-0a11-4a38-8f07-2e5becf8efe0",
//     "name": "Bolt: Injection",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "424",
//     "avail": "8R",
//     "cost": "50",
//     "costfor": "1",
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "d9bf2003-1911-4e65-b6a1-8babb761dd85",
//     "name": "Throwing Knife",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "424",
//     "avail": "4R",
//     "addweapon": "Throwing Knife",
//     "cost": "25",
//     "costfor": "1"
//   },
//   {
//     "id": "b4bbdbd3-1f65-44e5-b196-6ccbee1cc182",
//     "name": "Shuriken",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "424",
//     "avail": "4R",
//     "addweapon": "Shuriken",
//     "cost": "25",
//     "costfor": "1"
//   },
//   {
//     "id": "13b3f7fb-87c1-41d1-9c98-54484b0d557c",
//     "name": "Elektro-Netz",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HAMG",
//     "page": "172",
//     "avail": "6",
//     "cost": "600",
//     "addweapon": "Elektro-Netz",
//     "costfor": "1"
//   },
//   {
//     "id": "dd6e4f80-8200-47fa-8f78-9204e9c0f70d",
//     "name": "Arrow: Barbed Head",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "RG",
//     "page": "23",
//     "avail": "(Rating)R",
//     "cost": "(Rating * 2) + 10",
//     "costfor": "1",
//     "weaponbonus": {
//       "damage": "1"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "d09f3ebd-28fb-4375-ad0f-49fd1b521c12",
//     "name": "Arrow: Explosive Head",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "RG",
//     "page": "23",
//     "avail": "(Rating)F",
//     "cost": "(Rating * 2) + 15",
//     "costfor": "1",
//     "weaponbonus": {
//       "accuracy": "-1",
//       "ap": "-1",
//       "damage": "2"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "8fb09c1d-f890-444b-bb2b-cf1a5e6825b8",
//     "name": "Arrow: Hammerhead",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "RG",
//     "page": "24",
//     "avail": "Rating",
//     "cost": "(Rating * 2) + 5",
//     "costfor": "1",
//     "weaponbonus": {
//       "accuracy": "-1",
//       "ap": "2",
//       "damage": "1",
//       "damagetype": "S"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "f245f0d0-22ba-4367-97ee-6762df9137bc",
//     "name": "Arrow: Incendiary Head",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "RG",
//     "page": "24",
//     "avail": "12F",
//     "cost": "(Rating * 2) + 100",
//     "costfor": "1",
//     "weaponbonus": {
//       "accuracy": "-1",
//       "damagereplace": "8P",
//       "apreplace": "-6"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "1cdea5f7-df85-4ebe-bc32-f8dbfa27dad3",
//     "name": "Arrow: Screamer Head",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "RG",
//     "page": "24",
//     "avail": "Rating",
//     "cost": "(Rating * 2) + 5",
//     "costfor": "1",
//     "weaponbonus": {
//       "ap": "6",
//       "damage": "-2",
//       "damagetype": "S",
//       "accuracy": "-2"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "4ddbecdf-fc04-4b1d-851b-63d5e9e57dbb",
//     "name": "Arrow: Stick-n-Shock",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "RG",
//     "page": "24",
//     "avail": "(Rating)R",
//     "cost": "(Rating * 2) + 25",
//     "costfor": "1",
//     "weaponbonus": {
//       "accuracy": "-1",
//       "ap": "-5",
//       "damagereplace": "8S(e)"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "258049dc-6f94-4d41-be97-cadeba20ff36",
//     "name": "Arrow: Stick-n-Shock w/Static Shaft",
//     "category": "Ammunition",
//     "rating": "12",
//     "source": "RG",
//     "page": "24",
//     "avail": "6R",
//     "cost": "(Rating * 25) + 25",
//     "costfor": "1",
//     "weaponbonus": {
//       "accuracy": "-1",
//       "ap": "-5",
//       "damagereplace": "12S(e)"
//     },
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "9517a1df-5e00-46e5-92f7-4669b8b72380",
//     "name": "Bolt: Barbed Head",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "23",
//     "avail": "5R",
//     "cost": "15",
//     "costfor": "1",
//     "weaponbonus": {
//       "damage": "1"
//     },
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "945c0c98-9996-44ca-bc70-013389434d25",
//     "name": "Bolt: Explosive Head",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "23",
//     "avail": "9F",
//     "cost": "20",
//     "costfor": "1",
//     "weaponbonus": {
//       "accuracy": "-1",
//       "ap": "-1",
//       "damage": "2"
//     },
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "7057d386-6398-4e7f-8f14-1f74bddc7dba",
//     "name": "Bolt: Hammerhead",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "5",
//     "cost": "10",
//     "costfor": "1",
//     "weaponbonus": {
//       "ap": "2",
//       "damage": "1",
//       "damagetype": "S"
//     },
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "5acb7b20-ffe5-49fa-8ff7-372a59999fb6",
//     "name": "Bolt: Incendiary Head",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "12F",
//     "cost": "105",
//     "costfor": "1",
//     "weaponbonus": {
//       "accuracy": "-1"
//     },
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "8939b519-657a-4081-920f-3a4021eb6704",
//     "name": "Bolt: Screamer Head",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "2",
//     "cost": "10",
//     "costfor": "1",
//     "weaponbonus": {
//       "ap": "6",
//       "damage": "-2",
//       "damagetype": "S"
//     },
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "aa378f89-8f73-4f05-b920-fe7e35b7b3bd",
//     "name": "Bolt: Stick-n-Shock",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "6R",
//     "cost": "30",
//     "costfor": "1",
//     "weaponbonus": {
//       "ap": "-5",
//       "damagereplace": "8S(e)"
//     },
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "c4fecc98-0cf7-4e81-adde-bbed821850d6",
//     "name": "Bolt: Stick-n-Shock w/Static Shaft",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "6R",
//     "cost": "50",
//     "costfor": "1",
//     "weaponbonus": {
//       "ap": "-5",
//       "damagereplace": "12S(e)"
//     },
//     "ammoforweapontype": "crossbow"
//   },
//   {
//     "id": "c1df71d5-845b-46c9-ab8d-9166c5cbac70",
//     "name": "Boomerang",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "4",
//     "addweapon": "Boomerang",
//     "cost": "50",
//     "costfor": "1"
//   },
//   {
//     "id": "5b33b094-5e25-4f0e-9a3d-4f86a3621a3f",
//     "name": "Horizon BoomerEye",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "25",
//     "avail": "4",
//     "addweapon": "Horizon BoomerEye",
//     "cost": "50",
//     "costfor": "1"
//   },
//   {
//     "id": "d3d99f54-4176-4a11-a8cb-8447e285d977",
//     "name": "Harpoon",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "6",
//     "addweapon": "Harpoon",
//     "cost": "125",
//     "costfor": "1",
//     "ammoforweapontype": "harpoongun"
//   },
//   {
//     "id": "42404ae5-f203-4780-a575-a644d6443d20",
//     "name": "Javelin",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "6",
//     "addweapon": "Javelin",
//     "cost": "125",
//     "costfor": "1"
//   },
//   {
//     "id": "d37cd956-84de-40a7-a327-0185fe8bca3c",
//     "name": "Net",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "24",
//     "avail": "6",
//     "addweapon": "Net",
//     "cost": "350",
//     "costfor": "1"
//   },
//   {
//     "id": "7752a51c-b39f-4270-a5ed-e3de956cceb1",
//     "name": "Net Gun Reload (Shocknet)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "28",
//     "avail": "10R",
//     "cost": "600",
//     "costfor": "4",
//     "weaponbonus": {
//       "ap": "-5",
//       "damagereplace": "8S(e)"
//     },
//     "ammoforweapontype": "netgun"
//   },
//   {
//     "id": "5200ce56-c541-4a7a-8db2-ed4385f0d94f",
//     "name": "XL Net Gun Reload (Shocknet)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "28",
//     "avail": "10R",
//     "cost": "650",
//     "costfor": "2",
//     "weaponbonus": {
//       "ap": "-5",
//       "damagereplace": "8S(e)"
//     },
//     "ammoforweapontype": "netgunxl"
//   },
//   {
//     "id": "608324BB-E4CE-40B7-AC76-FF3F67B2C8E0",
//     "name": "Ammo: Fuel Canister",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "55",
//     "avail": "16F",
//     "cost": "40",
//     "costfor": "4",
//     "ammoforweapontype": "flame"
//   },
//   {
//     "id": "8C9AE7B4-EF1F-4983-9008-795D3592B417",
//     "name": "Ammo: AV Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "55",
//     "avail": "14R",
//     "cost": "175",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "-1"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "911BB20E-E325-4399-8F7A-82FDCE13AE51",
//     "name": "Ammo: Gyrojet",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "27",
//     "avail": "14R",
//     "cost": "160",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "-5",
//       "damage": "-2S(e)"
//     },
//     "ammoforweapontype": "gyrojet"
//   },
//   {
//     "id": "FE737992-B03B-484C-8F5C-B3BD554B9381",
//     "name": "Ammo: Gauss",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "55",
//     "avail": "18F",
//     "cost": "400",
//     "costfor": "10",
//     "ammoforweapontype": "cannon"
//   },
//   {
//     "id": "784dfff3-7c0e-4166-8687-e3012ef3b434",
//     "name": "Net Gun Reload (Standard)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "28",
//     "avail": "9",
//     "cost": "350",
//     "costfor": "4",
//     "ammoforweapontype": "netgun"
//   },
//   {
//     "id": "6ac1a45c-d8cc-42fd-9306-63ec3f4618e2",
//     "name": "XL Net Gun Reload (Standard)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "28",
//     "avail": "9",
//     "cost": "400",
//     "costfor": "2",
//     "ammoforweapontype": "netgunxl"
//   },
//   {
//     "id": "48c4b8ed-80cd-4744-83b8-06842fd917aa",
//     "name": "Urban Tribe Tomahawk",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "25",
//     "avail": "4",
//     "addweapon": "Urban Tribe Tomahawk",
//     "cost": "200",
//     "costfor": "1"
//   },
//   {
//     "id": "9d1f9b41-0a43-4835-9dc4-a2f04f8a20a8",
//     "name": "Bola: Standard",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "26",
//     "avail": "6",
//     "addweapon": "Bola",
//     "cost": "75",
//     "costfor": "1",
//     "ammoforweapontype": "bola"
//   },
//   {
//     "id": "f7b819d7-9f6a-446e-b916-82b5973883d6",
//     "name": "Bola: Monofilament",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "26",
//     "avail": "18F",
//     "addweapon": "Monofilament Bola",
//     "cost": "4000",
//     "costfor": "1",
//     "ammoforweapontype": "bola"
//   },
//   {
//     "id": "f4b92e14-fe1f-4be4-ad73-aed10e1f73b4",
//     "name": "Grenade: Flash-Bang",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "6R",
//     "addweapon": "Grenade: Flash-Bang",
//     "cost": "100",
//     "costfor": "1"
//   },
//   {
//     "id": "8e04c493-b04d-4511-ba68-7b2e0edbf3aa",
//     "name": "Grenade: Flash-Pak",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4",
//     "addweapon": "Grenade: Flash-Pak",
//     "cost": "125"
//   },
//   {
//     "id": "7c5e7573-d75e-43e3-949c-cb4c9d70329b",
//     "name": "Grenade: Fragmentation",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "11F",
//     "addweapon": "Grenade: Fragmentation",
//     "cost": "100"
//   },
//   {
//     "id": "e61c5487-2074-4fc8-8da0-9891b15482e7",
//     "name": "Grenade: High Explosive",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "11F",
//     "addweapon": "Grenade: High Explosive",
//     "cost": "100"
//   },
//   {
//     "id": "06798b15-40bb-4464-b6a8-a01466ee9b9e",
//     "name": "Grenade: Gas",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4R",
//     "addoncategory": [
//       "Drugs",
//       "Toxins",
//       "Custom"
//     ],
//     "addweapon": "Grenade: Gas",
//     "cost": "40"
//   },
//   {
//     "id": "84c6921e-dfa0-42ca-a01f-87c459ffc000",
//     "name": "Grenade: Smoke",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4R",
//     "addweapon": "Grenade: Smoke",
//     "cost": "40"
//   },
//   {
//     "id": "f3472dc1-4c7b-403a-8001-6441bb65a687",
//     "name": "Grenade: Thermal Smoke",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "6R",
//     "addweapon": "Grenade: Thermal Smoke",
//     "cost": "60"
//   },
//   {
//     "id": "f092fca8-46a9-4351-a06a-362846e6546a",
//     "name": "Minigrenade: Flash-Bang",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "6R",
//     "addweapon": "Minigrenade: Flash-Bang",
//     "cost": "100",
//     "ammoforweapontype": {
//       "+content": "glauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "7c00891f-6554-497d-9726-058d7c4a598e",
//     "name": "Minigrenade: Flash-Pak",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4",
//     "addweapon": "Minigrenade: Flash-Pak",
//     "cost": "125",
//     "ammoforweapontype": {
//       "+content": "glauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "8e1583bf-4cc5-4498-a800-36cb61d3fb27",
//     "name": "Minigrenade: Fragmentation",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "11F",
//     "addweapon": "Minigrenade: Fragmentation",
//     "cost": "100",
//     "ammoforweapontype": {
//       "+content": "glauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "daecdfc8-15d5-4864-9e20-13e4a0dca88e",
//     "name": "Minigrenade: High Explosive",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "11F",
//     "addweapon": "Minigrenade: High Explosive",
//     "cost": "100",
//     "ammoforweapontype": {
//       "+content": "glauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "74284124-57a0-479b-9ab1-3433bdb9e3b7",
//     "name": "Minigrenade: Gas",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4R",
//     "addoncategory": [
//       "Drugs",
//       "Toxins",
//       "Custom"
//     ],
//     "addweapon": "Minigrenade: Gas",
//     "cost": "40",
//     "ammoforweapontype": {
//       "+content": "glauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "b6381691-1a96-4299-b9df-129ffee64c45",
//     "name": "Minigrenade: Smoke",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4R",
//     "addweapon": "Minigrenade: Smoke",
//     "cost": "40",
//     "ammoforweapontype": {
//       "+content": "glauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "e2252b3c-1bae-496e-adbb-f1c3d6e0b32b",
//     "name": "Minigrenade: Thermal Smoke",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "6R",
//     "addweapon": "Minigrenade: Thermal Smoke",
//     "cost": "60",
//     "ammoforweapontype": {
//       "+content": "glauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "50e68547-da12-4d87-9e05-8fefb9ccf723",
//     "name": "Grenade: Paint",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "102",
//     "avail": "8R",
//     "addweapon": "Grenade: Paint",
//     "cost": "100"
//   },
//   {
//     "id": "ee1789ab-c000-4c69-bc11-7a99370f89b2",
//     "name": "Grenade: Paint (Radioactive Tracking Dye)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "102",
//     "avail": "8R",
//     "addweapon": "Grenade: Paint (Radioactive Tracking Dye)",
//     "cost": "150"
//   },
//   {
//     "id": "77a7504c-0eb8-4268-9744-66c4bbaae509",
//     "name": "Grenade: Smokebomb",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "187",
//     "avail": "6",
//     "addweapon": "Grenade: Smokebomb",
//     "cost": "25"
//   },
//   {
//     "id": "922c2a9b-88a8-47bf-a3ac-378a0e86be9c",
//     "name": "Grenade: Flash-Bang, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "6R",
//     "addweapon": "Grenade: Flash-Bang, Aerodynamic",
//     "cost": "100",
//     "costfor": "1"
//   },
//   {
//     "id": "4d27ebbf-02e8-45e7-9c0a-7151f1e4faae",
//     "name": "Grenade: Flash-Pak, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4",
//     "addweapon": "Grenade: Flash-Pak, Aerodynamic",
//     "cost": "125"
//   },
//   {
//     "id": "afde3ea3-2db2-4049-b54a-f471e8c11102",
//     "name": "Grenade: Fragmentation, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "11F",
//     "addweapon": "Grenade: Fragmentation, Aerodynamic",
//     "cost": "100"
//   },
//   {
//     "id": "e7a602bd-33f4-485b-8a18-a253b7e73a77",
//     "name": "Grenade: High Explosive, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "11F",
//     "addweapon": "Grenade: High Explosive, Aerodynamic",
//     "cost": "100"
//   },
//   {
//     "id": "1ceef439-32d7-4e1b-a13b-0223dfadf784",
//     "name": "Grenade: Gas, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4R",
//     "addoncategory": [
//       "Drugs",
//       "Toxins",
//       "Custom"
//     ],
//     "addweapon": "Grenade: Gas, Aerodynamic",
//     "cost": "40"
//   },
//   {
//     "id": "13208e69-2cc6-43c4-98ff-e0295cdf3b68",
//     "name": "Grenade: Smoke, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "4R",
//     "addweapon": "Grenade: Smoke, Aerodynamic",
//     "cost": "40"
//   },
//   {
//     "id": "2b58af45-2097-4607-a0d4-e251b64332af",
//     "name": "Grenade: Thermal Smoke, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "435",
//     "avail": "6R",
//     "addweapon": "Grenade: Thermal Smoke, Aerodynamic",
//     "cost": "60"
//   },
//   {
//     "id": "ff32c75d-9cb8-499e-8926-eb0728b2748b",
//     "name": "Grenade: Paint, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "102",
//     "avail": "8R",
//     "addweapon": "Grenade: Paint, Aerodynamic",
//     "cost": "100"
//   },
//   {
//     "id": "022bc6c7-1f91-4852-b1ee-0f50433a1b97",
//     "name": "Grenade: Paint (Radioactive Tracking Dye),  Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "102",
//     "avail": "8R",
//     "addweapon": "Grenade: Paint (Radioactive Tracking Dye), Aerodynamic",
//     "cost": "150"
//   },
//   {
//     "id": "32d933f1-638e-4022-bed7-7cf846cb4f20",
//     "name": "Grenade: Smokebomb, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "HT",
//     "page": "187",
//     "avail": "6",
//     "addweapon": "Grenade: Smokebomb, Aerodynamic",
//     "cost": "25"
//   },
//   {
//     "id": "60d28c15-42dd-4837-bc5f-8d8a15393655",
//     "name": "Rocket: Anti-Vehicle",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "436",
//     "avail": "18F",
//     "addweapon": "Rocket: Anti-Vehicle",
//     "cost": "2800",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "4c5be084-c975-42d2-962c-dc2add203b6b",
//     "name": "Rocket: Fragmentation",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "436",
//     "avail": "12F",
//     "addweapon": "Rocket: Fragmentation",
//     "cost": "2000",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "d5136e8b-bf34-401d-967d-c52f8cc175c1",
//     "name": "Rocket: High Explosive",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "436",
//     "avail": "18F",
//     "addweapon": "Rocket: High Explosive",
//     "cost": "2100",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "413366b8-0ef5-4998-bfce-9ba29db761a9",
//     "name": "Rocket: Incendiary",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "TSG",
//     "page": "29",
//     "avail": "12F",
//     "addweapon": "Rocket: Incendiary",
//     "cost": "1900",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "a42b7377-5a2e-456d-b729-8783873465f0",
//     "name": "Missile: Anti-Vehicle",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "436",
//     "avail": "22F",
//     "addweapon": "Missile: Anti-Vehicle",
//     "cost": "2800",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "bb40694a-59de-4fb8-9cee-697f37377a55",
//     "name": "Missile: Fragmentation",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "436",
//     "avail": "16F",
//     "addweapon": "Missile: Fragmentation",
//     "cost": "2000",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "838a308e-a638-433e-b5bc-54222fb645c9",
//     "name": "Missile: High Explosive",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SR5",
//     "page": "436",
//     "avail": "22F",
//     "addweapon": "Missile: High Explosive",
//     "cost": "2100",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "9ff4b57d-4aaf-482f-ab67-66e10ed58e13",
//     "name": "Ammo: Renraku/Ingram Supermach",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "132",
//     "avail": "16F",
//     "cost": "20",
//     "costfor": "10",
//     "ammoforweapontype": {
//       "+content": "supermach",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "78babb62-e0f1-4561-9e2d-05b0b5fbb2ad",
//     "name": "Ammo: Water",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SOTG",
//     "page": "16",
//     "avail": "0",
//     "cost": "0",
//     "ammoforweapontype": {
//       "+content": "firefighting cannons",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "767ecc1c-1f47-4684-b36f-0bbe0868ceac",
//     "name": "Ammo: Fire Extinguishant",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SOTG",
//     "page": "16",
//     "avail": "6R",
//     "cost": "500",
//     "costfor": "50",
//     "weaponbonus": {
//       "damage": "-1"
//     },
//     "ammoforweapontype": {
//       "+content": "firefighting cannons",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "d71ea9c1-9b5f-428e-b849-7d665eb31c5e",
//     "name": "Ammo: Slingshot Acid Capsules",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SOTG",
//     "page": "13",
//     "avail": "0",
//     "cost": "100",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagereplace": "6P",
//       "damagetype": "Acid"
//     },
//     "ammoforweapontype": {
//       "+content": "slingshot",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "e20f021b-1d00-4443-a481-d417114e067e",
//     "name": "Ammo: Slingshot Paint Capsules",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SOTG",
//     "page": "13",
//     "avail": "0",
//     "cost": "10",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagereplace": "0S"
//     },
//     "ammoforweapontype": {
//       "+content": "slingshot",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "5ece577b-7c15-4f2c-9f4b-cd325154cf9c",
//     "name": "Ammo: Slingshot Poison Capsules",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SOTG",
//     "page": "13",
//     "avail": "0",
//     "cost": "100",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagereplace": "As Narcoject"
//     },
//     "ammoforweapontype": {
//       "+content": "slingshot",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "d82e68e5-08d1-4f27-af45-de54ad1dcd72",
//     "name": "Ammo: Slingshot Capsule Round",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "RG",
//     "page": "23",
//     "avail": "0",
//     "cost": "5",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagereplace": "As Drug/Toxin",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": {
//       "+content": "slingshot",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "bdb59ce5-5283-44dd-9d9c-363e3ffced96",
//     "name": "Grenade: Maker",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "76",
//     "avail": "18F",
//     "addweapon": "Grenade: Maker",
//     "cost": "500"
//   },
//   {
//     "id": "aae8f02a-3459-4df8-ac25-9f74c1b523df",
//     "name": "Grenade: Maker, Aerodynamic",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "76",
//     "avail": "18F",
//     "addweapon": "Grenade: Maker, Aerodynamic",
//     "cost": "500"
//   },
//   {
//     "id": "d51c31c2-3f55-44b2-a52f-c175d0aaa03a",
//     "name": "Missile: Maker",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "76",
//     "avail": "24F",
//     "addweapon": "Missile: Maker",
//     "cost": "5000",
//     "ammoforweapontype": {
//       "+content": "mlauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "c7104585-4c86-4e11-9677-5d7f1912658c",
//     "name": "Ammo: Spinstorm Ferrous Slugs",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "80",
//     "avail": "12R",
//     "cost": "10",
//     "ammoforweapontype": {
//       "+content": "spinstorm",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "0c2be0b1-1b88-4b35-abe1-c2b887788ff3",
//     "name": "Ammo: Grey Goo Armor Eater",
//     "category": "Ammunition",
//     "rating": "1000",
//     "source": "SL",
//     "page": "88",
//     "avail": "24F",
//     "cost": "50 * Rating",
//     "costfor": "10",
//     "weaponbonus": {
//       "accuracyreplace": "3",
//       "apreplace": "0",
//       "damagereplace": "Special",
//       "userange": "Holdouts"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "3425fbe8-d842-4a7c-8f0d-3262b19c2ba9",
//     "name": "Ammo: Grey Goo Penetrator",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "88",
//     "avail": "24F",
//     "cost": "1000",
//     "costfor": "10",
//     "weaponbonus": {
//       "accuracyreplace": "3",
//       "apreplace": "-8",
//       "damagereplace": "9P",
//       "userange": "Holdouts"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "105acad3-5877-40c5-a472-290f953c6e5a",
//     "name": "Ammo: Man-Catcher Ammo Compound",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "133",
//     "avail": "18",
//     "cost": "200",
//     "costfor": "10",
//     "ammoforweapontype": {
//       "+content": "man-catcher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "8fe79706-ff0e-4701-a8b9-a20c27609761",
//     "name": "Torpedo Grenade: HEAP",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "133",
//     "avail": "14F",
//     "addweapon": "Torpedo Grenade: HEAP",
//     "cost": "300",
//     "ammoforweapontype": {
//       "+content": "torpglauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "63d1284d-9cf1-4214-ae1e-cfb3b74be548",
//     "name": "Torpedo Grenade: Depth Charge",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "SL",
//     "page": "133",
//     "avail": "12F",
//     "addweapon": "Torpedo Grenade: Depth Charge",
//     "cost": "175",
//     "ammoforweapontype": {
//       "+content": "torpglauncher",
//       "+@noextra": "True"
//     }
//   },
//   {
//     "id": "7e104b61-69b3-4a7e-8221-a6190314c411",
//     "name": "Ammo: Zapper",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "48",
//     "avail": "12R",
//     "cost": "140",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "(M)"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "b8ad4fef-4092-42ac-87ad-518dd8b60fb7",
//     "name": "Ammo: Looper",
//     "category": "Ammunition",
//     "rating": "6",
//     "source": "KC",
//     "page": "49",
//     "avail": "FixedValues(7R,8R,9R,10R,11R,12R)",
//     "cost": "FixedValues(20,30,40,50,100,200)",
//     "costfor": "10",
//     "weaponbonus": {
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "c40a91b0-7896-4ad1-b1e6-f39c9c93bf1b",
//     "name": "Ammo: Fuzzy",
//     "category": "Ammunition",
//     "rating": "2",
//     "source": "KC",
//     "page": "50",
//     "avail": "FixedValues(10R,12R)",
//     "cost": "FixedValues(30,50)",
//     "costfor": "10",
//     "weaponbonus": {
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "efa87c80-d712-41d8-813e-768a2f4ffcf7",
//     "name": "Ammo: E0-E0 (Public Grid)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "51",
//     "avail": "5R",
//     "cost": "50",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "(M)",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "f4be287a-2ed8-4e7b-8244-32243905fe84",
//     "name": "Ammo: E0-E0 (Local Grid)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "51",
//     "avail": "5R",
//     "cost": "60",
//     "bonus": {
//       "selecttext": null
//     },
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "(M)",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "8ddf2080-2444-4d9a-a601-5be4147b183d",
//     "name": "Ammo: E0-E0 (Megacorporate Grid)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "51",
//     "avail": "5R",
//     "cost": "100",
//     "bonus": {
//       "selecttext": {
//         "+@xml": "lifemodules.xml",
//         "+@xpath": "/chummer/storybuilder/macros/mega/persistent/*",
//         "+@allowedit": "True"
//       }
//     },
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "(M)",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "bf9001db-6484-47ab-9449-5f068f5fe04b",
//     "name": "Ammo: E0-E0 Rifle (Public Grid)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "51",
//     "avail": "5R",
//     "cost": "50",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "(M)",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "d1e0807b-6c6e-44bb-a230-c6f328ee26a5",
//     "name": "Ammo: E0-E0 Rifle (Local Grid)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "51",
//     "avail": "5R",
//     "cost": "60",
//     "bonus": {
//       "selecttext": null
//     },
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "(M)",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "b580d8fc-9987-403a-9be5-760b1d88f926",
//     "name": "Ammo: E0-E0 Rifle (Megacorporate Grid)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "51",
//     "avail": "5R",
//     "cost": "100",
//     "bonus": {
//       "selecttext": {
//         "+@xml": "lifemodules.xml",
//         "+@xpath": "/chummer/storybuilder/macros/mega/persistent/*",
//         "+@allowedit": "True"
//       }
//     },
//     "costfor": "10",
//     "weaponbonus": {
//       "damage": "-4",
//       "ap": "-4",
//       "damagetype": "(M)",
//       "userange": "Light Pistol"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "05a9f4e2-b7c0-4bf8-8614-33b1ee0f406d",
//     "name": "Arrow: Arrowlink (50m)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "52",
//     "avail": "6R",
//     "cost": "25",
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "2d3b26aa-a2c1-4514-915e-62aa08203456",
//     "name": "Arrow: Arrowlink (100m)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "52",
//     "avail": "8R",
//     "cost": "75",
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "74906d00-018e-4d1a-8a9c-4b77918de761",
//     "name": "Arrow: Arrowlink (200m)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "52",
//     "avail": "10R",
//     "cost": "200",
//     "ammoforweapontype": "bow"
//   },
//   {
//     "id": "65a853d4-52d6-44c7-a867-10ea1beaece3",
//     "name": "Arrow: Arrowlink (500m)",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "52",
//     "avail": "12R",
//     "cost": "400",
//     "ammoforweapontype": "bow"
//   },
//   {
//     "+content": "-->",
//     "id": "051ceea6-65c2-4058-95bd-f63d86a45984",
//     "name": "Grenade: Fuzzy Boom Boom Bunnies",
//     "category": "Ammunition",
//     "rating": "20",
//     "source": "KC",
//     "page": "53",
//     "avail": "10R",
//     "cost": "Rating*20",
//     "addweapon": {
//       "+content": "Grenade: Fuzzy Boom Boom Bunnies",
//       "+@rating": "{Rating}"
//     }
//   },
//   {
//     "id": "7470ed31-0d81-4992-bae6-38f9de198020",
//     "name": "Grenade: COS",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KC",
//     "page": "54",
//     "avail": "10R",
//     "cost": "500",
//     "addweapon": "Grenade: COS"
//   },
//   {
//     "id": "0c98750a-4286-4ec5-9035-b76f36c03ce4",
//     "name": "Grenade: Douser",
//     "category": "Ammunition",
//     "rating": "10",
//     "source": "KC",
//     "page": "55",
//     "avail": "Rating*2F",
//     "cost": "Rating*50",
//     "addweapon": {
//       "+content": "Grenade: Douser",
//       "+@rating": "{Rating}"
//     }
//   },
//   {
//     "id": "965ccb60-f9e2-4a44-8a11-05e6c5bd4457",
//     "name": "Grenade: DumDum",
//     "category": "Ammunition",
//     "rating": "10",
//     "source": "KC",
//     "page": "56",
//     "avail": "Rating*2R",
//     "cost": "Rating*50",
//     "addweapon": {
//       "+content": "Grenade: DumDum",
//       "+@rating": "{Rating}"
//     }
//   },
//   {
//     "id": "9d4cd46c-1fb9-4182-8a82-0479cf2adaae",
//     "name": "Ammo: Krime Power Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "24",
//     "avail": "3R",
//     "cost": "75",
//     "costfor": "10",
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "6ffa3fc7-e390-43b5-b27e-3d01baecec92",
//     "name": "Ammo: Krime Penetrator Buckshot Shells",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "24",
//     "avail": "16F",
//     "cost": "250",
//     "costfor": "10",
//     "weaponbonus": {
//       "ap": "-2"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "959b02ff-b3b3-4027-8ada-799f31a7a323",
//     "name": "Ammo: Krime Crackle Fin-Stabilized HEAT Slugs",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "24",
//     "avail": "20F",
//     "cost": "300",
//     "costfor": "10",
//     "weaponbonus": {
//       "damage": "2",
//       "ap": "-3"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "97febd7d-6385-472e-87af-29313f9224ae",
//     "name": "Ammo: Krime Laser Bullets",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "24",
//     "avail": "6R",
//     "cost": "125",
//     "costfor": "10",
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "c18eed96-2697-4b95-9be7-be9313504a2f",
//     "name": "Ammo: Krime Punisher Assault Cannon Rounds",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "24",
//     "avail": "8F",
//     "cost": "200",
//     "costfor": "10",
//     "weaponbonus": {
//       "damage": "-3",
//       "ap": "-3"
//     },
//     "ammoforweapontype": "cannon"
//   },
//   {
//     "id": "17136c4b-af3b-44a4-b5cf-2eedf7e9d217",
//     "name": "Ammo: Krime Splash Self-Defense Ammunition",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "24",
//     "avail": "4R",
//     "cost": "100",
//     "costfor": "10",
//     "weaponbonus": {
//       "damagetype": "(S)",
//       "ap": "1"
//     },
//     "ammoforweapontype": "gun"
//   },
//   {
//     "id": "f2300388-1d51-41bf-b333-fae00776a375",
//     "name": "Grenade: Krime Cleaner",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "27",
//     "avail": "10F",
//     "cost": "200",
//     "addweapon": "Grenade: Krime Cleaner"
//   },
//   {
//     "id": "16c588e4-9961-4ddd-b5f1-e95ed100159f",
//     "name": "Grenade: Krime Party",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "27",
//     "avail": "14F",
//     "cost": "190",
//     "addweapon": "Grenade: Krime Party"
//   },
//   {
//     "id": "b654b65a-b9e4-4be2-ac96-53044f5d9b40",
//     "name": "Grenade: Krime Cocktail",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "28",
//     "avail": "14F",
//     "cost": "100",
//     "addweapon": "Grenade: Krime Cocktail"
//   },
//   {
//     "id": "9964cf3d-c608-4ab5-ba8a-5ea7c96f29fb",
//     "name": "Grenade: Krime Stinger",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "KK",
//     "page": "29",
//     "avail": "8R",
//     "cost": "125",
//     "addweapon": "Grenade: Krime Stinger"
//   },
//   {
//     "id": "5768d347-cd1e-454c-b5f7-012266c185a4",
//     "name": "Micro-Torpedo: Chemical",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "R5",
//     "page": "136",
//     "avail": "18F",
//     "addweapon": "Micro-Torpedo: Chemical",
//     "cost": "3000",
//     "ammoforweapontype": "microtorpedo"
//   },
//   {
//     "id": "0e1ec849-e332-4683-aabb-a94e46512c22",
//     "name": "Micro-Torpedo: Explosive",
//     "category": "Ammunition",
//     "rating": "0",
//     "source": "R5",
//     "page": "136",
//     "avail": "18F",
//     "addweapon": "Micro-Torpedo: Explosive",
//     "cost": "2500",
//     "ammoforweapontype": "microtorpedo"
//   },

var CoreWeaponModifications = []ModificationSpec{
	{
		//Under-Barrel
		ID:           "bipod",
		Name:         "Bipod",
		Description:  "A bipod is a two-legged support that attaches to the barrel of a firearm, allowing the shooter to rest the weapon on the ground or another surface to improve stability.",
		MountPoints:  []MountPoint{MountPointUnderBarrel},
		Availability: 2,
		Legality:     shared.LegalityTypeLegal,
		Cost:         200,
		RuleSource:   shared.RuleSourceSR5Core,
		// Modifiers: []Modifier{
		// 	{
		// 		Type:   "RecoilCompensation",
		// 		Effect: "Add",
		// 		Value:  2,
		// 	},
		// },
		// Attach Bipod
		//     One Minute
		// Fold/Deploy Bipod
		//     Simple Action
		// Remove Bipod
		//     Complex Action
		// Wireless
		//     Folding up or deploying the bipod is a Free Action.
		// A bipod can be attached to the underbarrel mount of a weapon and provides 2 points of Recoil Compensation when properly deployed.
		// Attaching a bipod takes one minute. Folding up or deploying a bipod is a Simple Action. Removing it is a Complex Action.
	},
	{
		ID:           "concealable_holster",
		Name:         "Concealable Holster",
		Description:  "A holster that is designed to be easily concealed under clothing.",
		Availability: 2,
		Legality:     shared.LegalityTypeLegal,
		Cost:         150,
		RuleSource:   shared.RuleSourceSR5Core,
		// Modifiers: []Modifier{
		// 	{
		// 		Type:   "Concealability",
		// 		Effect: "Add",
		// 		Value:  -1,
		// 	},
		// },
		// Wireless
		//  Wireless sensors and a smart-fabric coated weave allow the holster to alter color and texture in real time adding an additional –1 to the item’s Concealability.
		// The concealable holster adds –1 to the item’s Concealability.
		// Only pistols and tasers fit in a Concealable
	},
	{
		ID:           "gas_vent_system",
		Name:         "Gas Vent System",
		Description:  "Gas-vent recoil compensation systems are barrel-mounted accessories that vent a weapon’s barrel gases at a specific vector to counter muzzle climb. Once installed, a gas-vent cannot be removed.",
		MountPoints:  []MountPoint{MountPointBarrel},
		Availability: 4, // (Rating×3)R
		Legality:     shared.LegalityTypeRestricted,
		Cost:         200, // Rating×200¥
		RuleSource:   shared.RuleSourceSR5Core,
		// Modifiers: []Modifier{
		// 	{
		// 		Type:   "RecoilCompensation",
		// 		Effect: "Add",
		// 		Value:  2, // Rating
		// 	},
		// },
		//  Gas-vent systems provide a number of points of Recoil Compensation equal to their rating.
	},
	{
		ID:           "gyro_mount",
		Name:         "Gyro Mount",
		Description:  "This heavy upper-body harness features an attached, articulated, motorized gyro-stabilized arm that mounts an assault rifle or a heavy weapon.",
		MountPoints:  []MountPoint{MountPointUnderBarrel},
		Availability: 7,
		Legality:     shared.LegalityTypeLegal,
		Cost:         1400,
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
		// Attach/Remove Weapon
		//     Simple Action
		// Put On Gyro-Mount
		//     5 Minutes
		// Remove Gyro-Mount
		//     Complex Action
		// Wireless
		//     Activating the harness’s quick-release with a wireless signal to exit the harness is a Free Action.
		//
		// The system neutralizes up to 6 points of recoil and movement modifiers.
		// Attaching or removing a weapon from the mount takes a Simple Action.
		// Putting on a gyro-mount harness takes about five minutes, while the quick-release allows you to get out of it with a Complex Action
	},
	{
		ID:           "hidden_arm_slide",
		Name:         "Hidden Arm Slide",
		Description:  "A hidden arm slide is a concealed holster that allows the wearer to draw a weapon from under a sleeve or pant leg.",
		Availability: 4,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         350,
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
		// Wireless
		//  You can ready the weapon in the slide as a Free Action.
		//If you quick draw the weapon in this slide, the threshold for the quick draw is 2.
		// It also gives the weapon a –1 Concealability modifier.
	},
	{
		ID:           "imaging_scope",
		Name:         "Imaging Scope",
		Description:  "An imaging scope is a telescopic sight that uses a digital camera to provide a magnified image of the target. The image is displayed on a screen in the eyepiece of the scope, allowing the shooter to see the target in low-light conditions or through smoke or fog.",
		Availability: 2,
		Legality:     shared.LegalityTypeLegal,
		Cost:         300,
		MountPoints:  []MountPoint{MountPointTop},
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "laser_sight",
		Name:         "Laser Sight",
		Description:  "A laser sight is a small, battery-powered laser that attaches to the barrel of a firearm. The laser projects a red or green dot on the target, allowing the shooter to aim the weapon more accurately.",
		Availability: 2,
		Legality:     shared.LegalityTypeLegal,
		Cost:         125,
		MountPoints:  []MountPoint{MountPointTop, MountPointUnderBarrel},
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "periscope",
		Name:         "Periscope",
		Description:  "A periscope is a device that allows the user to see around corners or over obstacles without exposing themselves to enemy fire. The periscope consists of a tube with mirrors at each end that reflect light from the target to the eyepiece of the periscope.",
		Availability: 3,
		Legality:     shared.LegalityTypeLegal,
		Cost:         70,
		MountPoints:  []MountPoint{MountPointTop},
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "quick_draw_holster",
		Name:         "Quick-Draw Holster",
		Description:  "A quick-draw holster is a holster that is designed to allow the wearer to draw a weapon quickly and easily. The holster is usually made of leather or nylon and has a snap or Velcro closure to keep the weapon secure.",
		Availability: 4,
		Legality:     shared.LegalityTypeLegal,
		Cost:         175,
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "shock_pad",
		Name:         "Shock Pad",
		Description:  "A shock pad is a device that attaches to the grip of a firearm and delivers an electric shock to the shooter if the weapon is taken from them. The shock pad is designed to prevent unauthorized use of the weapon and can be activated or deactivated with a switch on the grip.",
		Availability: 2,
		Legality:     shared.LegalityTypeLegal,
		Cost:         50,
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "suppressor",
		Name:         "Suppressor",
		Description:  "A suppressor is a device that attaches to the barrel of a firearm and reduces the noise and muzzle flash produced by the weapon. The suppressor is usually made of metal or plastic and contains baffles that slow and cool the escaping gases from the barrel.",
		Availability: 9,
		Legality:     shared.LegalityTypeForbidden,
		Cost:         500,
		MountPoints:  []MountPoint{MountPointBarrel},
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "smart_firing_platform",
		Name:         "Smart Firing Platform",
		Description:  "A smart firing platform is a computerized system that attaches to a firearm and assists the shooter in aiming the weapon. The platform uses sensors and cameras to track the target and adjust the aim of the weapon to compensate for movement and environmental conditions.",
		Availability: 12,
		Legality:     shared.LegalityTypeForbidden,
		Cost:         2500,
		MountPoints:  []MountPoint{MountPointUnderBarrel},
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "smartgun_system_external",
		Name:         "Smartgun System, External",
		Description:  "A smartgun system is a computerized aiming system that attaches to a firearm and assists the shooter in aiming the weapon. The system uses sensors and cameras to track the target and adjust the aim of the weapon to compensate for movement and environmental conditions.",
		Availability: 4,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         200,
		MountPoints:  []MountPoint{MountPointTop, MountPointUnderBarrel},
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "smartgun_system_internal",
		Name:         "Smartgun System, Internal",
		Description:  "A smartgun system is a computerized aiming system that is built into a firearm and assists the shooter in aiming the weapon. The system uses sensors and cameras to track the target and adjust the aim of the weapon to compensate for movement and environmental conditions.",
		Availability: 6,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         2, // (2×Weapon Cost)¥
		MountPoints:  []MountPoint{MountPointInternal},
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "spare_clip",
		Name:         "Spare Clip",
		Description:  "A spare clip comes unloaded but can hold the maximum rounds for the weapon.",
		Availability: 4,
		Legality:     shared.LegalityTypeLegal,
		Cost:         5,
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "speed_loader",
		Name:         "Speed Loader",
		Description:  "A speed loader is a device that attaches to the cylinder of a revolver and allows the shooter to reload the weapon quickly and easily. The speed loader is usually made of metal or plastic and contains a spring-loaded mechanism that pushes the rounds into the cylinder when the loader is twisted.",
		Availability: 2,
		Legality:     shared.LegalityTypeLegal,
		Cost:         25,
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:           "tripod",
		Name:         "Tripod",
		Description:  "A tripod is a three-legged support that attaches to the barrel of a firearm, allowing the shooter to rest the weapon on the ground or another surface to improve stability.",
		MountPoints:  []MountPoint{MountPointUnderBarrel},
		Availability: 4,
		Legality:     shared.LegalityTypeLegal,
		Cost:         500,
		// Modifiers:    []Modifier{},
		RuleSource: shared.RuleSourceSR5Core,
	},
}
