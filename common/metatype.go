package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/sirupsen/logrus"
)

const (
	MetatypesFilepath = "_data/metatypes"

	RacialTraitLowLightVision                    RacialTrait = "Low-Light Vision"
	RacialTraitThermographicVision               RacialTrait = "Thermographic Vision"
	RacialTrait2DicForPathogenAndToxinResistance RacialTrait = "+2 dice for pathogen and toxin resistance"
	RacialTrait20PercentIncreasedLifestyleCost   RacialTrait = "+20% increased Lifestyle cost"
	RacialTrait1Reach                            RacialTrait = "+1 Reach"
	RacialTrait1DermalArmor                      RacialTrait = "+1 dermal armor"
	RacialTraitDoubleLifestyleCosts              RacialTrait = "Double Lifestyle costs"

	MetatypeCategoryMetahuman    MetatypeCategory = "Metahuman"
	MetatypeCategoryMetavariant  MetatypeCategory = "Metavariant"
	MetatypeCategoryMetasapient  MetatypeCategory = "Metasapient"
	MetatypeCategoryShapeshifter MetatypeCategory = "Shapeshifter"
)

type MetatypeManager struct {
	Metatypes Metatypes
}

func NewMetatypeManager() *MetatypeManager {
	return &MetatypeManager{
		Metatypes: LoadMetatypes(),
	}
}

func (mm *MetatypeManager) LoadChummerData(name string, v *Metatype) error {
	// filepath := fmt.Sprintf("%s/%s.json", MetatypesFilepath, name)
	// if err := utils.LoadStructFromJSON(filepath, &v); err != nil {
	// 	return err
	// }
	return nil
}

type (
	MetatypeCategory string

	RacialTrait                        string
	RacialTraits                       []RacialTrait
	MetatypeAttribute[T int | float64] struct {
		Min    T `yaml:"min"`
		Max    T `yaml:"max"`
		AugMax T `yaml:"aug_max"`
	}
	Metatypes map[string]*Metatype
	Metatype  struct {
		ID                  string                     `yaml:"id"`
		Name                string                     `yaml:"name"`
		Category            MetatypeCategory           `yaml:"category"`
		Description         string                     `yaml:"description"`
		Body                MetatypeAttribute[int]     `yaml:"body"`
		Agility             MetatypeAttribute[int]     `yaml:"agility"`
		Reaction            MetatypeAttribute[int]     `yaml:"reaction"`
		Strength            MetatypeAttribute[int]     `yaml:"strength"`
		Willpower           MetatypeAttribute[int]     `yaml:"willpower"`
		Logic               MetatypeAttribute[int]     `yaml:"logic"`
		Intuition           MetatypeAttribute[int]     `yaml:"intuition"`
		Charisma            MetatypeAttribute[int]     `yaml:"charisma"`
		Edge                MetatypeAttribute[int]     `yaml:"edge"`
		Initiative          MetatypeAttribute[int]     `yaml:"initiative"`
		Essence             MetatypeAttribute[float64] `yaml:"essence"`
		Magic               MetatypeAttribute[int]     `yaml:"magic"`
		Resonance           MetatypeAttribute[int]     `yaml:"resonance"`
		QualityRestrictions []string                   `yaml:"quality_restrictions"`
		Qualities           []string                   `yaml:"qualities"`
		RuleSource          RuleSource                 `yaml:"rule_source"`
	}
)

func NewMetatype() *Metatype {
	m := &Metatype{
		QualityRestrictions: []string{},
		Qualities:           []string{},
	}

	return m
}

func LoadMetatype(name string, v *Metatype) error {
	name = strings.ToLower(name)
	filepath := fmt.Sprintf("%s/%s.yaml", MetatypesFilepath, name)

	if err := utils.LoadStructFromYAML(filepath, &v); err != nil {
		return err
	}

	return nil
}

func LoadMetatypes() Metatypes {
	logrus.Info("Started loading metatypes")
	list := make(Metatypes)

	files, errReadDir := os.ReadDir(MetatypesFilepath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read metatype directory")
	}

	for _, file := range files {
		var v Metatype
		if strings.HasSuffix(file.Name(), ".yaml") {

			name := strings.TrimSuffix(file.Name(), ".yaml")
			if err := LoadMetatype(name, &v); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load metatype")
			}

			list[v.ID] = &v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded metatype file")
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(list)}).Info("Done loading metatypes")

	return list
}

var CoreMetatypes = []Metatype{
	{
		ID:                  "human",
		Name:                "Human",
		Category:            MetatypeCategoryMetahuman,
		Description:         "Humans are the most common metatype in the world. They are known for their adaptability and versatility.",
		Body:                MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Agility:             MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Reaction:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Strength:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Willpower:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Logic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Intuition:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Charisma:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Edge:                MetatypeAttribute[int]{Min: 2, Max: 7, AugMax: 7},
		Initiative:          MetatypeAttribute[int]{Min: 2, Max: 12, AugMax: 20},
		Essence:             MetatypeAttribute[float64]{Min: 0, Max: 6, AugMax: 6},
		Magic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Resonance:           MetatypeAttribute[int]{Min: 0, Max: 6, AugMax: 10},
		QualityRestrictions: []string{},
		Qualities:           []string{},
		RuleSource:          RuleSourceSR5Core,
	},
	{
		ID:                  "elf",
		Name:                "Elf",
		Category:            MetatypeCategoryMetahuman,
		Description:         "Elves are known for their grace, beauty, and long lifespans.",
		Body:                MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Agility:             MetatypeAttribute[int]{Min: 2, Max: 7, AugMax: 11},
		Reaction:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Strength:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Willpower:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Logic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Intuition:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Charisma:            MetatypeAttribute[int]{Min: 3, Max: 8, AugMax: 12},
		Edge:                MetatypeAttribute[int]{Min: 2, Max: 7, AugMax: 7},
		Initiative:          MetatypeAttribute[int]{Min: 2, Max: 12, AugMax: 20},
		Essence:             MetatypeAttribute[float64]{Min: 0, Max: 6, AugMax: 6},
		Magic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		Resonance:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		QualityRestrictions: []string{},
		Qualities:           []string{"low_light_vision"},
		RuleSource:          RuleSourceSR5Core,
	},
	{
		ID:                  "dwarf",
		Name:                "Dwarf",
		Category:            MetatypeCategoryMetahuman,
		Description:         "Dwarves are known for their toughness, resilience, and their ability to see in the dark.",
		Body:                MetatypeAttribute[int]{Min: 3, Max: 8, AugMax: 12},
		Agility:             MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Reaction:            MetatypeAttribute[int]{Min: 1, Max: 5, AugMax: 9},
		Strength:            MetatypeAttribute[int]{Min: 3, Max: 8, AugMax: 12},
		Willpower:           MetatypeAttribute[int]{Min: 2, Max: 7, AugMax: 11},
		Logic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Intuition:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Charisma:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Edge:                MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		Initiative:          MetatypeAttribute[int]{Min: 2, Max: 11, AugMax: 19},
		Essence:             MetatypeAttribute[float64]{Min: 0, Max: 6, AugMax: 6},
		Magic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		Resonance:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		QualityRestrictions: []string{},
		Qualities:           []string{"thermographic_vision", "resistance_to_pathogens_toxins"},
		RuleSource:          RuleSourceSR5Core,
		// <bonus>
		//     <lifestylecost>20</lifestylecost>
		// </bonus>
	},
	{
		ID:                  "ork",
		Name:                "Ork",
		Category:            MetatypeCategoryMetahuman,
		Description:         "Orks are known for their strength, toughness, and tusks.",
		Body:                MetatypeAttribute[int]{Min: 4, Max: 9, AugMax: 13},
		Agility:             MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Reaction:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Strength:            MetatypeAttribute[int]{Min: 3, Max: 8, AugMax: 12},
		Willpower:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Logic:               MetatypeAttribute[int]{Min: 1, Max: 5, AugMax: 9},
		Intuition:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Charisma:            MetatypeAttribute[int]{Min: 1, Max: 5, AugMax: 9},
		Edge:                MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		Initiative:          MetatypeAttribute[int]{Min: 2, Max: 12, AugMax: 20},
		Essence:             MetatypeAttribute[float64]{Min: 0, Max: 6, AugMax: 6},
		Magic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		Resonance:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		QualityRestrictions: []string{},
		Qualities:           []string{"low_light_vision"},
		RuleSource:          RuleSourceSR5Core,
	},
	{
		ID:                  "troll",
		Name:                "Troll",
		Category:            MetatypeCategoryMetahuman,
		Description:         "Trolls are known for their size, strength, and toughness.",
		Body:                MetatypeAttribute[int]{Min: 5, Max: 10, AugMax: 14},
		Agility:             MetatypeAttribute[int]{Min: 1, Max: 5, AugMax: 9},
		Reaction:            MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Strength:            MetatypeAttribute[int]{Min: 5, Max: 10, AugMax: 14},
		Willpower:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 10},
		Logic:               MetatypeAttribute[int]{Min: 1, Max: 5, AugMax: 9},
		Intuition:           MetatypeAttribute[int]{Min: 1, Max: 5, AugMax: 9},
		Charisma:            MetatypeAttribute[int]{Min: 1, Max: 4, AugMax: 8},
		Edge:                MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		Initiative:          MetatypeAttribute[int]{Min: 2, Max: 11, AugMax: 19},
		Essence:             MetatypeAttribute[float64]{Min: 0, Max: 6, AugMax: 6},
		Magic:               MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		Resonance:           MetatypeAttribute[int]{Min: 1, Max: 6, AugMax: 6},
		QualityRestrictions: []string{},
		Qualities:           []string{"thermographic_vision"},
		RuleSource:          RuleSourceSR5Core,
		// <bonus>
		//     <armor group="0">1</armor>
		//     <reach>1</reach>
		//     <lifestylecost>100</lifestylecost>
		// </bonus>
	},
}
