package metatype

import (
	"github.com/Jasrags/ShadowMUD/common/shared"
)

const (
	RacialTraitLowLightVision                    RacialTrait = "Low-Light Vision"
	RacialTraitThermographicVision               RacialTrait = "Thermographic Vision"
	RacialTrait2DicForPathogenAndToxinResistance RacialTrait = "+2 dice for pathogen and toxin resistance"
	RacialTrait20PercentIncreasedLifestyleCost   RacialTrait = "+20% increased Lifestyle cost"
	RacialTrait1Reach                            RacialTrait = "+1 Reach"
	RacialTrait1DermalArmor                      RacialTrait = "+1 dermal armor"
	RacialTraitDoubleLifestyleCosts              RacialTrait = "Double Lifestyle costs"

	CategoryMetahuman    Category = "Metahuman"
	CategoryMetavariant  Category = "Metavariant"
	CategoryMetasapient  Category = "Metasapient"
	CategoryShapeshifter Category = "Shapeshifter"

	MetatypeNameHuman MetatypeName = "Human"
	MetatypeNameElf   MetatypeName = "Elf"
	MetatypeNameDwarf MetatypeName = "Dwarf"
	MetatypeNameOrk   MetatypeName = "Ork"
	MetatypeNameTroll MetatypeName = "Troll"
)

func NewAttributes() Attributes {
	return Attributes{
		"agility":   {Name: "Agility"},
		"body":      {Name: "Body"},
		"charisma":  {Name: "Charisma"},
		"edge":      {Name: "Edge"},
		"intuition": {Name: "Intuition"},
		"logic":     {Name: "Logic"},
		"magic":     {Name: "Magic"},
		"reaction":  {Name: "Reaction"},
		"resonance": {Name: "Resonance"},
		"strength":  {Name: "Strength"},
		"willpower": {Name: "Willpower"},
	}
}

type (
	MetatypeName               string
	Category                   string
	RacialTrait                string
	RacialTraits               []RacialTrait
	Attributes                 map[string]*Attribute[int]
	Attribute[T int | float64] struct {
		Name   string `yaml:"name"`
		Min    T      `yaml:"min"`
		Max    T      `yaml:"max"`
		AugMax T      `yaml:"aug_max"`
	}
	// Attributes struct {
	// 	Body       Attribute[int]     `yaml:"body"`
	// 	Agility    Attribute[int]     `yaml:"agility"`
	// 	Reaction   Attribute[int]     `yaml:"reaction"`
	// 	Strength   Attribute[int]     `yaml:"strength"`
	// 	Willpower  Attribute[int]     `yaml:"willpower"`
	// 	Logic      Attribute[int]     `yaml:"logic"`
	// 	Intuition  Attribute[int]     `yaml:"intuition"`
	// 	Charisma   Attribute[int]     `yaml:"charisma"`
	// 	Edge       Attribute[int]     `yaml:"edge"`
	// 	Initiative Attribute[int]     `yaml:"initiative"`
	// 	Essence    Attribute[float64] `yaml:"essence"`
	// 	Magic      Attribute[int]     `yaml:"magic"`
	// 	Resonance  Attribute[int]     `yaml:"resonance"`
	// }
	Metatypes map[string]Metatype
	Metatype  struct {
		ID                  string             `yaml:"id"`
		PointCost           int                `yaml:"point_cost"`
		Name                string             `yaml:"name"`
		Category            Category           `yaml:"category"`
		Description         string             `yaml:"description"`
		Attributes          Attributes         `yaml:"attributes"`
		Essence             Attribute[float64] `yaml:"essence"`
		QualityRestrictions []string           `yaml:"quality_restrictions"`
		Hidden              bool               `yaml:"hidden"`
		Qualities           []string           `yaml:"qualities"`
		RuleSource          shared.RuleSource  `yaml:"rule_source"`
	}
)

func NewMetatype() *Metatype {
	m := &Metatype{
		Attributes:          NewAttributes(),
		QualityRestrictions: []string{},
		Qualities:           []string{},
	}

	return m
}
