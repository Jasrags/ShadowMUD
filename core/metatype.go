package core

import (
	"fmt"
	"os"
	"strings"

	"shadowrunmud/core/util"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	MetatypeDataPath       = "data/metatypes"
	MetatypeFilename       = MetatypeDataPath + "/%s.yaml"
	MetatypeFileMinVersion = "0.0.1"

	RacialTraitLowLightVision                    = "Low-Light Vision"
	RacialTraitThermographicVision               = "Thermographic Vision"
	RacialTrait2DicForPathogenAndToxinResistance = "+2 dice for pathogen and toxin resistance"
	RacialTrait20PercentIncreasedLifestyleCost   = "+20% increased Lifestyle cost"
	RacialTrait1Reach                            = "+1 Reach"
	RacialTrait1DermalArmor                      = "+1 dermal armor"
	RacialTraitDoubleLifestyleCosts              = "Double Lifestyle costs"
)

var (
	Metatypes = map[string]Metatype{}
)

type MetatypeAttribute struct {
	Min int
	Max int
}

func NewMetatype() *Metatype {
	uuid := uuid.New().String()
	return &Metatype{
		ID: uuid,
	}
}

type Metatype struct {
	ID           string            `yaml:"id"`
	Name         string            `yaml:"name"`
	Body         MetatypeAttribute `yaml:"body"`
	Agility      MetatypeAttribute `yaml:"agility"`
	Reaction     MetatypeAttribute `yaml:"reaction"`
	Strength     MetatypeAttribute `yaml:"strength"`
	Willpower    MetatypeAttribute `yaml:"willpower"`
	Logic        MetatypeAttribute `yaml:"logic"`
	Intuition    MetatypeAttribute `yaml:"intuition"`
	Charisma     MetatypeAttribute `yaml:"charisma"`
	Edge         MetatypeAttribute `yaml:"edge"`
	RacialTraits []string          `yaml:"racial_traits"`
	RuleSource   string            `yaml:"rule_source"`
	FileVersion  string            `yaml:"file_version"`
}

func LoadMetatypes() {
	logrus.Info("Started loading metatypes")

	files, errReadDir := os.ReadDir(MetatypeDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read metatype directory")
	}

	// Create a map to store the metatypes
	metatypes := make(map[string]Metatype, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", MetatypeDataPath, file.Name())

			var metatype Metatype
			if err := util.LoadStructFromYAML(filepath, &metatype); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load metatype")
			}

			metatypes[metatype.Name] = metatype
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded metatype file")
	}

	logrus.WithFields(logrus.Fields{"count": len(metatypes)}).Info("Done loading metatypes")

	Metatypes = metatypes
}

func LoadMetatype(name string) (*Metatype, error) {
	var v Metatype
	if err := util.LoadStructFromYAML(fmt.Sprintf(MetatypeFilename, name), &v); err != nil {
		return nil, err
	}

	return &v, nil
}