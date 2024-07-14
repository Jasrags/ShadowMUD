package metatype

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"shadowrunmud/util"

	"github.com/sirupsen/logrus"
)

const (
	MetatypeDataPath       = "data/metatypes"
	MetatypeFilename       = MetatypeDataPath + "/%s.yaml"
	MetatypeFileMinVersion = "0.0.1"
)

type MetatypeAttribute struct {
	Min int
	Max int
}

type Metatype struct {
	Name      string            `yaml:"name"`
	Body      MetatypeAttribute `yaml:"body"`
	Agility   MetatypeAttribute `yaml:"agility"`
	Reaction  MetatypeAttribute `yaml:"reaction"`
	Strength  MetatypeAttribute `yaml:"strength"`
	Willpower MetatypeAttribute `yaml:"willpower"`
	Logic     MetatypeAttribute `yaml:"logic"`
	Intuition MetatypeAttribute `yaml:"intuition"`
	Charisma  MetatypeAttribute `yaml:"charisma"`
	Edge      MetatypeAttribute `yaml:"edge"`
	Essence   float64           `yaml:"essence"`
	// Initiative   int // REA+INT
	RacialTraits []string `yaml:"racial_traits"`
	RuleSource   string   `yaml:"rule_source"`
	FileVersion  string   `yaml:"file_version"`
}

var (
	Metatypes = map[string]Metatype{}
)

const (
	RacialTraitLowLightVision                    = "Low-Light Vision"
	RacialTraitThermographicVision               = "Thermographic Vision"
	RacialTrait2DicForPathogenAndToxinResistance = "+2 dice for pathogen and toxin resistance"
	RacialTrait20PercentIncreasedLifestyleCost   = "+20% increased Lifestyle cost"
	RacialTrait1Reach                            = "+1 Reach"
	RacialTrait1DermalArmor                      = "+1 dermal armor"
	RacialTraitDoubleLifestyleCosts              = "Double Lifestyle costs"
)

// LoadMetatypes loads metatypes from YAML files in a specified directory.
// It populates the global `Metatypes` map with the loaded metatypes.
// The function takes a `sync.WaitGroup` pointer as a parameter to indicate completion.
// It is expected to be called as a goroutine.
func LoadMetatypes(wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Debug("Started loading metatypes")

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
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded metatype file")
	}

	logrus.WithFields(logrus.Fields{"count": len(metatypes)}).Info("Done loading metatypes")

	Metatypes = metatypes
}
