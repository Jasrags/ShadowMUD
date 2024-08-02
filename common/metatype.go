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
)

type (
	RacialTrait       string
	RacialTraits      []RacialTrait
	MetatypeAttribute struct {
		Min int `yaml:"min"`
		Max int `yaml:"max"`
	}
	Metatypes map[string]*Metatype
	Metatype  struct {
		ID           string            `yaml:"id"`
		Name         string            `yaml:"name"`
		Description  string            `yaml:"description"`
		Body         MetatypeAttribute `yaml:"body"`
		Agility      MetatypeAttribute `yaml:"agility"`
		Reaction     MetatypeAttribute `yaml:"reaction"`
		Strength     MetatypeAttribute `yaml:"strength"`
		Willpower    MetatypeAttribute `yaml:"willpower"`
		Logic        MetatypeAttribute `yaml:"logic"`
		Intuition    MetatypeAttribute `yaml:"intuition"`
		Charisma     MetatypeAttribute `yaml:"charisma"`
		Edge         MetatypeAttribute `yaml:"edge"`
		RacialTraits RacialTraits      `yaml:"racial_traits"`
		RuleSource   RuleSource        `yaml:"rule_source"`
	}
)

func NewMetatype() *Metatype {
	m := &Metatype{
		RacialTraits: make(RacialTraits, 0),
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
