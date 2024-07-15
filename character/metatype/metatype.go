package metatype

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"shadowrunmud/util"

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

type Metatype interface {
	GetID() string
	SetID(string)
	GetName() string
	GetBody() MetatypeAttribute
	GetAgility() MetatypeAttribute
	GetReaction() MetatypeAttribute
	GetStrength() MetatypeAttribute
	GetWillpower() MetatypeAttribute
	GetLogic() MetatypeAttribute
	GetIntuition() MetatypeAttribute
	GetCharisma() MetatypeAttribute
	GetEdge() MetatypeAttribute
	GetEssence() float64
	GetRacialTraits() []string
	GetRuleSource() string
	GetFileVersion() string
}

func NewMetatype() Metatype {
	uuid := uuid.New().String()
	return &metatype{
		ID: uuid,
	}
}

type metatype struct {
	ID           string            `yaml:"id,omitempty"`
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
	Essence      float64           `yaml:"essence"`
	RacialTraits []string          `yaml:"racial_traits"`
	RuleSource   string            `yaml:"rule_source"`
	FileVersion  string            `yaml:"file_version"`
}

func (m *metatype) GetID() string {
	return m.ID
}

func (m *metatype) SetID(id string) {
	m.ID = id
}

func (m *metatype) GetName() string {
	return m.Name
}

func (m *metatype) GetBody() MetatypeAttribute {
	return m.Body
}

func (m *metatype) GetAgility() MetatypeAttribute {
	return m.Agility
}

func (m *metatype) GetReaction() MetatypeAttribute {
	return m.Reaction
}

func (m *metatype) GetStrength() MetatypeAttribute {
	return m.Strength
}

func (m *metatype) GetWillpower() MetatypeAttribute {
	return m.Willpower
}

func (m *metatype) GetLogic() MetatypeAttribute {
	return m.Logic
}

func (m *metatype) GetIntuition() MetatypeAttribute {
	return m.Intuition
}

func (m *metatype) GetCharisma() MetatypeAttribute {
	return m.Charisma
}

func (m *metatype) GetEdge() MetatypeAttribute {
	return m.Edge
}

func (m *metatype) GetEssence() float64 {
	return m.Essence
}

func (m *metatype) GetRacialTraits() []string {
	return m.RacialTraits
}

func (m *metatype) GetRuleSource() string {
	return m.RuleSource
}

func (m *metatype) GetFileVersion() string {
	return m.FileVersion
}

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

			metatypes[metatype.GetName()] = metatype
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded metatype file")
	}

	logrus.WithFields(logrus.Fields{"count": len(metatypes)}).Info("Done loading metatypes")

	Metatypes = metatypes
}

func LoadMetatype(name string) (Metatype, error) {
	var m Metatype
	if err := util.LoadStructFromYAML(fmt.Sprintf(MetatypeFilename, name), &m); err != nil {
		return nil, err
	}

	return m, nil
}
