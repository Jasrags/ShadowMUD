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

type (
	MetatypeName               string
	Category                   string
	RacialTrait                string
	RacialTraits               []RacialTrait
	Attribute[T int | float64] struct {
		Min    T `yaml:"min"`
		Max    T `yaml:"max"`
		AugMax T `yaml:"aug_max"`
	}
	Attributes struct {
		Body       Attribute[int]     `yaml:"body"`
		Agility    Attribute[int]     `yaml:"agility"`
		Reaction   Attribute[int]     `yaml:"reaction"`
		Strength   Attribute[int]     `yaml:"strength"`
		Willpower  Attribute[int]     `yaml:"willpower"`
		Logic      Attribute[int]     `yaml:"logic"`
		Intuition  Attribute[int]     `yaml:"intuition"`
		Charisma   Attribute[int]     `yaml:"charisma"`
		Edge       Attribute[int]     `yaml:"edge"`
		Initiative Attribute[int]     `yaml:"initiative"`
		Essence    Attribute[float64] `yaml:"essence"`
		Magic      Attribute[int]     `yaml:"magic"`
		Resonance  Attribute[int]     `yaml:"resonance"`
	}
	Metatypes map[string]Metatype
	Metatype  struct {
		ID                  string            `yaml:"id"`
		PointCost           int               `yaml:"point_cost"`
		Name                string            `yaml:"name"`
		Category            Category          `yaml:"category"`
		Description         string            `yaml:"description"`
		Attributes          Attributes        `yaml:"attributes"`
		QualityRestrictions []string          `yaml:"quality_restrictions"`
		Qualities           []string          `yaml:"qualities"`
		RuleSource          shared.RuleSource `yaml:"rule_source"`
	}
)

func NewMetatype() *Metatype {
	m := &Metatype{
		QualityRestrictions: []string{},
		Qualities:           []string{},
	}

	return m
}
