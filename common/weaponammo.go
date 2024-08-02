package common

const (
	WeaponAmunitionFilepath = "_data/items/weapons/amunition"
)

type (
	WeaponAmunitionSpec struct {
		ID                       string       `yaml:"id"`
		Name                     string       `yaml:"name"`
		Description              string       `yaml:"description"`
		DamageValue              int          `yaml:"damage_value,omitempty"`
		DamageType               DamageType   `yaml:"damage_type,omitempty"`
		ArmorPenatration         int          `yaml:"armor_penatration,omitempty"`
		AccuracyModifier         int          `yaml:"accuracy_modifier,omitempty"`
		DamageValueModifier      int          `yaml:"damage_value_modifier,omitempty"`
		ArmorPenatrationModifier int          `yaml:"armor_penatration_modifier,omitempty"`
		Availability             int          `yaml:"availability"`
		Legality                 LegalityType `yaml:"legality"`
		Cost                     int          `yaml:"cost"`
		ItemTags                 []ItemTag    `yaml:"tags"`
		Modifiers                []Modifier   `yaml:"modifiers"`
		RuleSource               RuleSource   `yaml:"rule_source"`
	}
	WeaponAmunition struct {
		ID        string              `yaml:"id"`
		Quantity  int                 `yaml:"quantity"`
		Modifiers []Modifier          `yaml:"modifiers"`
		Spec      WeaponAmunitionSpec `yaml:"-"`
	}
)

var CoreWeaponAmunition = []WeaponAmunitionSpec{
	{
		ID:                       "apds",
		Name:                     "APDS",
		Description:              "These are military-grade armor piercing rounds—their full name is armor piercing discarding sabot. They are designed to travel at high velocities and punch through personal body armor.",
		ArmorPenatrationModifier: -4,
		Availability:             12,
		Legality:                 LegalityTypeForbidden,
		Cost:                     120,
		RuleSource:               RuleSourceSR5Core,
	},
	{
		ID:           "assault_cannon",
		Name:         "Assault Cannon",
		Description:  "These are for assault cannons only, and they’re the only thing assault cannons can load.",
		Availability: 12,
		Legality:     LegalityTypeForbidden,
		Cost:         400,
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:                       "explosive_rounds",
		Name:                     "Explosive Rounds",
		Description:              "These slugs carry a shaped-charge explosive, designed to explode and fragment on impact. Explosive rounds misfire whenever you roll a critical glitch. When this happens, you must resist one “attack” with a Damage Value equal to the normal damage done by the weapon (and don’t forget the modifier for the explosive rounds). The attack misses its intended target, and the weapon firing the bullets is destroyed.",
		DamageValueModifier:      1,
		ArmorPenatrationModifier: -1,
		Availability:             9,
		Legality:                 LegalityTypeForbidden,
		Cost:                     80,
		RuleSource:               RuleSourceSR5Core,
	},
	{
		ID:                       "flechette_rounds",
		Name:                     "Flechette Rounds",
		Description:              "The payload of a flechette round is made up of tiny, tightly packed metal slivers. The round breaks up and shatters on impact, becoming a tumbling hail of shrapnel. Flechette rounds are devastating against unprotected targets, but not as effective against hardened armor.",
		DamageValueModifier:      2,
		ArmorPenatrationModifier: 5,
		Availability:             6,
		Legality:                 LegalityTypeForbidden,
		Cost:                     65,
		RuleSource:               RuleSourceSR5Core,
	},
	{
		ID:                       "gel_rounds",
		Name:                     "Gel Rounds",
		Description:              "Gel rounds are designed to deliver a non-lethal payload. They are often used by security forces and police to subdue suspects without causing permanent harm. Gel rounds are less effective against armored targets.",
		DamageType:               DamageTypeStun,
		ArmorPenatrationModifier: 1,
		Availability:             2,
		Legality:                 LegalityTypeRestricted,
		Cost:                     25,
		RuleSource:               RuleSourceSR5Core,
	},
	{
		ID:                       "hollow_points",
		Name:                     "Hollow Points",
		Description:              "Hollow points are designed to expand on impact, causing more damage to soft targets. They are less effective against armored targets.",
		DamageValueModifier:      1,
		ArmorPenatrationModifier: 2,
		Availability:             4,
		Legality:                 LegalityTypeForbidden,
		Cost:                     70,
		RuleSource:               RuleSourceSR5Core,
	},
	{
		ID:           "injection_darts",
		Name:         "Injection Darts",
		Description:  "Injection darts are used to deliver drugs or toxins to a target. They are fired from a dart pistol or a dart rifle. Injection darts are not designed to cause damage, but to deliver a payload to the target. Injection darts are not available for all weapons, and are not available for all drugs or toxins.",
		Availability: 4,
		Legality:     LegalityTypeRestricted,
		Cost:         75,
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "regular_ammo",
		Name:         "Regular Ammo",
		Description:  "Regular ammo is the standard ammunition for most firearms. It is a lead slug encased in a copper jacket. Regular ammo is effective against most targets, but is less effective against armored targets.",
		Availability: 2,
		Legality:     LegalityTypeRestricted,
		Cost:         20,
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:                       "stick_n_shock",
		Name:                     "Stick-n-Shock",
		Description:              "Stick-n-Shock rounds are designed to deliver an electrical shock to the target. They are often used by security forces and police to subdue suspects without causing permanent harm. Stick-n-Shock rounds are less effective against armored targets.",
		DamageValueModifier:      -2,
		DamageType:               DamageTypeStun,
		ArmorPenatrationModifier: -5,
		Availability:             6,
		Legality:                 LegalityTypeRestricted,
		Cost:                     80,
		Modifiers: []Modifier{
			{
				Type:   "add",
				Effect: "ElectricDamage",
			},
		},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:           "taser_darts",
		Name:         "Taser Dart",
		Description:  "Taser darts are used to deliver an electrical shock to a target. They are fired from a dart pistol or a dart rifle. Taser darts are not designed to cause damage, but to deliver an electrical shock to the target. Taser darts are not available for all weapons.",
		Availability: 3,
		Legality:     LegalityTypeLegal,
		Cost:         50,
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:               "tracer_rounds",
		Name:             "Tracer",
		Description:      "Tracer rounds are designed to leave a visible trail in the air, making it easier to track the path of the bullet. Tracer rounds are often used to help a shooter adjust their aim. Tracer rounds are less effective against armored targets.",
		AccuracyModifier: 1,
		Availability:     6,
		Legality:         LegalityTypeRestricted,
		Cost:             60,
		RuleSource:       RuleSourceSR5Core,
	},
}
