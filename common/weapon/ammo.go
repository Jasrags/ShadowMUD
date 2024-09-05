package weapon

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	WeaponAmunitionFilepath = "_data/items/weapons/amunition"
)

type (
	AmmunitionSpec struct {
		ID                       string              `yaml:"id"`
		Name                     string              `yaml:"name"`
		Description              string              `yaml:"description"`
		DamageValue              int                 `yaml:"damage_value"`
		DamageType               shared.DamageType   `yaml:"damage_type"`
		ArmorPenatration         int                 `yaml:"armor_penatration"`
		AccuracyModifier         int                 `yaml:"accuracy_modifier"`
		DamageValueModifier      int                 `yaml:"damage_value_modifier"`
		ArmorPenatrationModifier int                 `yaml:"armor_penatration_modifier"`
		Availability             int                 `yaml:"availability"`
		Legality                 shared.LegalityType `yaml:"legality"`
		Cost                     int                 `yaml:"cost"`
		CostFor                  int                 `yaml:"cost_for"`
		ItemTags                 []shared.ItemTag    `yaml:"tags"`
		Modifiers                shared.Modifiers    `yaml:"modifiers"`
		RuleSource               shared.RuleSource   `yaml:"rule_source"`
	}
)
