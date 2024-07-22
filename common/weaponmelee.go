package common

const (
	WeaponMeleeDataPath = "data/items/weapons/melee"
	WeaponMeleeFilename = WeaponMeleeDataPath + "/%s.yaml"
)

type WeaponMeleeSpec struct {
	ID               string       `yaml:"id"`
	Name             string       `yaml:"name"`
	Type             WeaponType   `yaml:"type"`
	Description      string       `yaml:"description,omitempty"`
	Accuracy         int          `yaml:"accuracy,omitempty"`
	Reach            int          `yaml:"reach,omitempty"`
	Concealability   int          `yaml:"concealability,omitempty"`
	DamageValue      int          `yaml:"damage_value,omitempty"`
	DamageType       DamageType   `yaml:"damage_type,omitempty"`
	ArmorPenatration int          `yaml:"armor_penatration,omitempty"`
	Availability     int          `yaml:"availability,omitempty"`
	LegalityType     LegalityType `yaml:"legality_type,omitempty"`
	ItemTags         []ItemTag    `yaml:"tags"`
	Modifiers        []Modifier   `yaml:"modifiers"`
	Cost             int          `yaml:"cost,omitempty"`
	RuleSource       RuleSource   `yaml:"rule_source,omitempty"`
}

type WeaponMelee struct {
	ID        string              `yaml:"id"`
	ItemTags  []ItemTag           `yaml:"tags"`
	Modifiers []Modifier          `yaml:"modifiers"`
	Spec      WeaponAmunitionSpec `yaml:"-"`
}

var CoreWeaponMelee = []WeaponMeleeSpec{
	{
		ID:           "club",
		Name:         "Club",
		Type:         WeaponTypeRanged,
		Description:  "The weapon they named the skill after.",
		Accuracy:     4,
		Reach:        1,
		DamageValue:  2,
		DamageType:   DamageTypePhysical,
		Availability: 4,
		LegalityType: LegalityTypeLegal,
		Cost:         30,
		RuleSource:   RuleSourceSR5Core,
		ItemTags:     []ItemTag{ItemTagMelee, ItemTagClub},
	},
	{
		ID:           "extendable_baton",
		Name:         "Extendable Baton",
		Type:         WeaponTypeRanged,
		Description:  "A baton that extends to a full length.",
		Accuracy:     5,
		Reach:        1,
		DamageValue:  2,
		DamageType:   DamageTypePhysical,
		Availability: 4,
		LegalityType: LegalityTypeLegal,
		Cost:         100,
		RuleSource:   RuleSourceSR5Core,
		ItemTags:     []ItemTag{ItemTagMelee, ItemTagBaton},
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
		ID:           "sap",
		Name:         "Sap",
		Type:         WeaponTypeRanged,
		Description:  "A small, weighted club.",
		Accuracy:     5,
		DamageValue:  2,
		DamageType:   DamageTypePhysical,
		Availability: 2,
		LegalityType: LegalityTypeLegal,
		Cost:         30,
		RuleSource:   RuleSourceSR5Core,
		ItemTags:     []ItemTag{ItemTagMelee, ItemTagClub},
		// Concealability
		//  Modifier +2
	},
	{
		ID:           "staff",
		Name:         "Staff",
		Type:         WeaponTypeRanged,
		Description:  "A long stick.",
		Accuracy:     6,
		Reach:        2,
		DamageValue:  3,
		DamageType:   DamageTypePhysical,
		Availability: 3,
		LegalityType: LegalityTypeLegal,
		Cost:         100,
		RuleSource:   RuleSourceSR5Core,
		ItemTags:     []ItemTag{ItemTagMelee, ItemTagStaff},
	},
	{
		ID:               "stun_baton",
		Name:             "Stun Baton",
		Type:             WeaponTypeRanged,
		Description:      "A baton that delivers an electrical shock.",
		Accuracy:         4,
		Reach:            1,
		DamageValue:      9,
		DamageType:       DamageTypeStun,
		ArmorPenatration: -5,
		Availability:     6,
		LegalityType:     LegalityTypeRestricted,
		Cost:             750,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBaton},
		// Wireless
		//  The stun baton recharges by induction, regaining one charge per full hour of wireless-enabled time.
		// Modifier
		//  Electric attack
	},
	{
		ID:           "telescoping_staff",
		Name:         "Telescoping Staff",
		Type:         WeaponTypeRanged,
		Description:  "A staff that extends to a full length.",
		Accuracy:     4,
		Reach:        2,
		DamageValue:  4,
		DamageType:   DamageTypePhysical,
		Availability: 4,
		LegalityType: LegalityTypeLegal,
		Cost:         350,
		RuleSource:   RuleSourceSR5Core,
		ItemTags:     []ItemTag{ItemTagMelee, ItemTagStaff},
	},
	// Blades
	{
		ID:               "combat_axe",
		Name:             "Combat Axe",
		Type:             WeaponTypeRanged,
		Description:      "A large axe designed for combat.",
		Accuracy:         4,
		Reach:            2,
		DamageValue:      5,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -4,
		Availability:     12,
		LegalityType:     LegalityTypeRestricted,
		Cost:             4000,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagAxe},
	},
	{
		ID:               "combat_knife",
		Name:             "Combat Knife",
		Type:             WeaponTypeRanged,
		Description:      "A knife designed for combat.",
		Accuracy:         6,
		DamageValue:      2,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -3,
		Availability:     4,
		LegalityType:     LegalityTypeLegal,
		Cost:             300,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade},
	},
	{
		ID:               "forearm_snap_blades",
		Name:             "Forearm Snap-Blades",
		Type:             WeaponTypeRanged,
		Description:      "Blades that extend from the forearm.",
		Accuracy:         4,
		DamageValue:      2,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     7,
		LegalityType:     LegalityTypeRestricted,
		Cost:             200,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade},
		// Wireless
		//  Readying the forearm snap blades is a Free Action instead of a Simple Action.
	},
	{
		ID:               "katana",
		Name:             "Katana",
		Type:             WeaponTypeRanged,
		Description:      "A traditional Japanese sword.",
		Accuracy:         7,
		Reach:            1,
		DamageValue:      3,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -3,
		Availability:     9,
		LegalityType:     LegalityTypeRestricted,
		Cost:             1000,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagKatana},
	},
	{
		ID:               "knife",
		Name:             "Knife",
		Type:             WeaponTypeRanged,
		Description:      "A small knife.",
		Accuracy:         5,
		DamageValue:      1,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -1,
		LegalityType:     LegalityTypeLegal,
		Cost:             10,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagKnife},
	},
	{
		ID:               "pole_arm",
		Name:             "Pole Arm",
		Type:             WeaponTypeRanged,
		Description:      "A pole arm is a large weapon with a blade on the end.",
		Accuracy:         5,
		Reach:            3,
		DamageValue:      3,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     6,
		LegalityType:     LegalityTypeRestricted,
		Cost:             1000,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagPolearm},
	},
	{
		ID:               "survival_knife",
		Name:             "Survival Knife",
		Type:             WeaponTypeRanged,
		Description:      `A survival knife is a large knife with a serrated edge. It is designed for use in the wilderness, and is often used by military personnel and survivalists.`,
		Accuracy:         5,
		DamageValue:      2,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -1,
		LegalityType:     LegalityTypeLegal,
		Cost:             100,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagKnife},
		// Wireless
		//  The knife displays an ARO of local maps, your GPS position, and can be used to make commcalls.
	},
	{
		ID:               "sword",
		Name:             "Sword",
		Type:             WeaponTypeRanged,
		Description:      "A sword is a long, sharp blade.",
		Accuracy:         6,
		Reach:            1,
		DamageValue:      3,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     5,
		LegalityType:     LegalityTypeRestricted,
		Cost:             500,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagBlade, ItemTagSword},
	},
	// Exotic Melee Weapons
	{
		ID:               "monofilament_chainsaw",
		Name:             "Monofilament Chainsaw",
		Type:             WeaponTypeRanged,
		Description:      "A chainsaw with a monofilament blade.",
		Accuracy:         3,
		Reach:            1,
		DamageValue:      8,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -6,
		Availability:     8,
		LegalityType:     LegalityTypeLegal,
		Cost:             500,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagExotic, ItemTagChainsaw, ItemTagMonofiliament},
		// When used against barriers, double the monofilament chainsaw’s Damage Value of 8P
	},
	{
		ID:               "monofilament_whip",
		Name:             "Monofilament Whip",
		Type:             WeaponTypeRanged,
		Description:      "A whip with a monofilament blade.",
		Accuracy:         5,
		Reach:            2,
		DamageValue:      12,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -8,
		Availability:     12,
		LegalityType:     LegalityTypeRestricted,
		Cost:             10000,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagExotic, ItemTagWhip, ItemTagMonofiliament},
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
		Type:        WeaponTypeRanged,
		Description: "These may be traditional fist-load weapons like brass knuckles, or modern Hardliner Gloves with a thin layer of densiplast located the knuckles and the edge of the hand. Either way they substantially boost the impact of a punch, making it potentially deadly.",
		// Accuracy:    0, // Physical
		DamageValue:  1,
		DamageType:   DamageTypePhysical,
		Availability: 2,
		LegalityType: LegalityTypeRestricted,
		Cost:         100,
		RuleSource:   RuleSourceSR5Core,
		ItemTags:     []ItemTag{ItemTagMelee, ItemTagMiscellaneous},
		// Skill
		//  Unarmed Combat Skill
	},
	{
		ID:          "plasteel_toe_boots",
		Name:        "Plasteel Toe Boots",
		Type:        WeaponTypeRanged,
		Description: "These boots have a reinforced toe made of plasteel, making them a formidable weapon in a kick.",
		// Accuracy:    0, // Physical
		DamageValue:  1,
		DamageType:   DamageTypePhysical,
		Availability: 2,
		LegalityType: LegalityTypeLegal,
		Cost:         200,
		RuleSource:   RuleSourceSR5Core,
		ItemTags:     []ItemTag{ItemTagMelee, ItemTagMiscellaneous},
		// Skill
		//  Unarmed Combat Skill
	},
	{
		ID:          "shock_gloves",
		Name:        "Shock Gloves",
		Type:        WeaponTypeRanged,
		Description: "These gloves are lined with a layer of conductive material that delivers an electric shock to the target.",
		// Accuracy:    0, // Physical
		DamageValue:      8,
		DamageType:       DamageTypeStun,
		ArmorPenatration: -5,
		Availability:     6,
		LegalityType:     LegalityTypeRestricted,
		Cost:             550,
		RuleSource:       RuleSourceSR5Core,
		ItemTags:         []ItemTag{ItemTagMelee, ItemTagMiscellaneous},
		Modifiers:        []Modifier{{Type: "Electric attack", Effect: "Attack"}},
		// Skill
		//  Unarmed Combat Skill
		// Wireless
		//  The shock gloves recharge by induction, regaining one charge per full hour of wireless-enabled time.
		// Modifier
		//  Electric attack
	},
}
