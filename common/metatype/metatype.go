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

var CoreMetatypes = []Metatype{
	{
		ID:          "human",
		Name:        "Human",
		PointCost:   0,
		Category:    CategoryMetahuman,
		Description: "Humans are the most common metatype in the world. They are known for their adaptability and versatility.",
		Attributes: Attributes{
			Body:       Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Agility:    Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Reaction:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Strength:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Willpower:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Logic:      Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Intuition:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Charisma:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Edge:       Attribute[int]{Min: 2, Max: 7, AugMax: 7},
			Initiative: Attribute[int]{Min: 2, Max: 12, AugMax: 20},
			Essence:    Attribute[float64]{Min: 0, Max: 6, AugMax: 6},
			Magic:      Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Resonance:  Attribute[int]{Min: 0, Max: 6, AugMax: 10},
		},
		QualityRestrictions: []string{},
		Qualities:           []string{},
		RuleSource:          shared.RuleSourceSR5Core,
	},
	{
		ID:          "elf",
		Name:        "Elf",
		PointCost:   40,
		Category:    CategoryMetahuman,
		Description: "Elves are known for their grace, beauty, and long lifespans.",
		Attributes: Attributes{
			Body:       Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Agility:    Attribute[int]{Min: 2, Max: 7, AugMax: 11},
			Reaction:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Strength:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Willpower:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Logic:      Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Intuition:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Charisma:   Attribute[int]{Min: 3, Max: 8, AugMax: 12},
			Edge:       Attribute[int]{Min: 2, Max: 7, AugMax: 7},
			Initiative: Attribute[int]{Min: 2, Max: 12, AugMax: 20},
			Essence:    Attribute[float64]{Min: 0, Max: 6, AugMax: 6},
			Magic:      Attribute[int]{Min: 1, Max: 6, AugMax: 6},
			Resonance:  Attribute[int]{Min: 1, Max: 6, AugMax: 6},
		},
		QualityRestrictions: []string{},
		Qualities:           []string{"low_light_vision"},
		RuleSource:          shared.RuleSourceSR5Core,
	},
	{
		ID:          "dwarf",
		Name:        "Dwarf",
		PointCost:   50,
		Category:    CategoryMetahuman,
		Description: "Dwarves are known for their toughness, resilience, and their ability to see in the dark.",
		Attributes: Attributes{
			Body:       Attribute[int]{Min: 3, Max: 8, AugMax: 12},
			Agility:    Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Reaction:   Attribute[int]{Min: 1, Max: 5, AugMax: 9},
			Strength:   Attribute[int]{Min: 3, Max: 8, AugMax: 12},
			Willpower:  Attribute[int]{Min: 2, Max: 7, AugMax: 11},
			Logic:      Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Intuition:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Charisma:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Edge:       Attribute[int]{Min: 1, Max: 6, AugMax: 6},
			Initiative: Attribute[int]{Min: 2, Max: 11, AugMax: 19},
			Essence:    Attribute[float64]{Min: 0, Max: 6, AugMax: 6},
			Magic:      Attribute[int]{Min: 1, Max: 6, AugMax: 6},
			Resonance:  Attribute[int]{Min: 1, Max: 6, AugMax: 6},
		},
		QualityRestrictions: []string{},
		Qualities:           []string{"thermographic_vision", "resistance_to_pathogens_toxins"},
		RuleSource:          shared.RuleSourceSR5Core,
		// <bonus>
		//     <lifestylecost>20</lifestylecost>
		// </bonus>
	},
	{
		ID:          "ork",
		Name:        "Ork",
		PointCost:   50,
		Category:    CategoryMetahuman,
		Description: "Orks are known for their strength, toughness, and tusks.",
		Attributes: Attributes{
			Body:       Attribute[int]{Min: 4, Max: 9, AugMax: 13},
			Agility:    Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Reaction:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Strength:   Attribute[int]{Min: 3, Max: 8, AugMax: 12},
			Willpower:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Logic:      Attribute[int]{Min: 1, Max: 5, AugMax: 9},
			Intuition:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Charisma:   Attribute[int]{Min: 1, Max: 5, AugMax: 9},
			Edge:       Attribute[int]{Min: 1, Max: 6, AugMax: 6},
			Initiative: Attribute[int]{Min: 2, Max: 12, AugMax: 20},
			Essence:    Attribute[float64]{Min: 0, Max: 6, AugMax: 6},
			Magic:      Attribute[int]{Min: 1, Max: 6, AugMax: 6},
			Resonance:  Attribute[int]{Min: 1, Max: 6, AugMax: 6},
		},
		QualityRestrictions: []string{},
		Qualities:           []string{"low_light_vision"},
		RuleSource:          shared.RuleSourceSR5Core,
	},
	{
		ID:          "troll",
		Name:        "Troll",
		PointCost:   90,
		Category:    CategoryMetahuman,
		Description: "Trolls are known for their size, strength, and toughness.",
		Attributes: Attributes{
			Body:       Attribute[int]{Min: 5, Max: 10, AugMax: 14},
			Agility:    Attribute[int]{Min: 1, Max: 5, AugMax: 9},
			Reaction:   Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Strength:   Attribute[int]{Min: 5, Max: 10, AugMax: 14},
			Willpower:  Attribute[int]{Min: 1, Max: 6, AugMax: 10},
			Logic:      Attribute[int]{Min: 1, Max: 5, AugMax: 9},
			Intuition:  Attribute[int]{Min: 1, Max: 5, AugMax: 9},
			Charisma:   Attribute[int]{Min: 1, Max: 4, AugMax: 8},
			Edge:       Attribute[int]{Min: 1, Max: 6, AugMax: 6},
			Initiative: Attribute[int]{Min: 2, Max: 11, AugMax: 19},
			Essence:    Attribute[float64]{Min: 0, Max: 6, AugMax: 6},
			Magic:      Attribute[int]{Min: 1, Max: 6, AugMax: 6},
			Resonance:  Attribute[int]{Min: 1, Max: 6, AugMax: 6},
		},
		QualityRestrictions: []string{},
		Qualities:           []string{"thermographic_vision"},
		RuleSource:          shared.RuleSourceSR5Core,
		// <bonus>
		//     <armor group="0">1</armor>
		//     <reach>1</reach>
		//     <lifestylecost>100</lifestylecost>
		// </bonus>
	},
}
