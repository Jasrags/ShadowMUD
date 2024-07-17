package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jasrags/ShadowMUD/core/util"

	"github.com/sirupsen/logrus"
)

const (
	CyberwareDataPath       = "data/items/cyberware"
	CyberwareFilename       = CyberwareDataPath + "/%s.yaml"
	CyberwareFileMinVersion = "0.0.1"
)

var (
	ErrUnknownCyberwareGrade = fmt.Errorf("unknown cyberware grade")
)

type CyberwareGrade string

const (
	CyberwareGradeStandard  CyberwareGrade = "Standard"
	CyberwareGradeAlphaware CyberwareGrade = "Alphaware"
	CyberwareGradeBetaware  CyberwareGrade = "Betaware"
	CyberwareGradeDeltaware CyberwareGrade = "Deltaware"
	CyberwareGradeUsed      CyberwareGrade = "Used"
)

func GetCyberwareGradeModifiers(grade CyberwareGrade) (float64, int, float64, error) {
	var essenceCostMultiplier float64
	var availMod int
	var costMultiplier float64

	switch grade {
	case CyberwareGradeStandard:
		essenceCostMultiplier = 1.0
		availMod = 0
		costMultiplier = 1.0
	case CyberwareGradeAlphaware:
		essenceCostMultiplier = 0.8
		availMod = 2
		costMultiplier = 1.2
	case CyberwareGradeBetaware:
		essenceCostMultiplier = 0.7
		availMod = 4
		costMultiplier = 1.5
	case CyberwareGradeDeltaware:
		essenceCostMultiplier = 0.5
		availMod = 8
		costMultiplier = 2.5
	case CyberwareGradeUsed:
		essenceCostMultiplier = 1.25
		availMod = -4
		costMultiplier = 0.75
	default:
		logrus.WithFields(logrus.Fields{"grade": grade, "function": "GetCyberwareGradeModifiers"}).Error("Unknown CyberwareGrade grade")

		return 0, 0, 0, ErrUnknownCyberwareGrade
	}

	return essenceCostMultiplier, availMod, costMultiplier, nil
}

type CyberwareModifier struct {
	Type   string `yaml:"type"`
	Effect string `yaml:"effect"`
	Value  int    `yaml:"value"`
}

type Cyberware struct {
	ID           string              `yaml:"id,omitempty"`
	Name         string              `yaml:"name"`
	Description  string              `yaml:"description"`
	EssenceCost  float64             `yaml:"essence_cost"`
	Capacity     int                 `yaml:"capacity,omitempty"`
	Rating       int                 `yaml:"rating,omitempty"`
	Grade        CyberwareGrade      `yaml:"grade,omitempty"`
	ToggleAction ActionType          `yaml:"toggle_action,omitempty"`
	IsActive     bool                `yaml:"is_active,omitempty"`
	Modifiers    []CyberwareModifier `yaml:"modifiers"`
	Cost         int                 `yaml:"cost"`
	Availability int                 `yaml:"availability"`
	Legality     LegalityType        `yaml:"legality"`
	Notes        string              `yaml:"notes"`
	RuleSource   RuleSource          `yaml:"rule_source"`
}

func NewCyberware() *Cyberware {
	return &Cyberware{
		Modifiers: make([]CyberwareModifier, 0),
	}
}

// Part 	Device 	            Essence 	Capacity 	Avail 	Cost 	Source
// Body 	Wired Reflexes R1 	2 	        - 	        8R 	    39,000¥ 	Core
// Body 	Wired Reflexes R2 	3 	        - 	        12R 	149,000¥ 	Core
// Body 	Wired Reflexes R3 	5 	        - 	        20R 	217,000¥ 	Core

// Part 	Device 						Essence 		Capacity 		Avail 			Cost 				Source
// Body 	Bone Lacing (Plastic) 		0.5 			- 				8R 				8,000¥ 				Core
// Body 	Bone Lacing (Aluminum) 		1 				- 				12R 			18,000¥ 			Core
// Body 	Bond Lacing (Titanium) 		1.5 			- 				16R 			30,000¥ 			Core
// Body 	Dermal Plating (R 1-6) 		Rating * 0.5 	- 				(Rating * 4)R 	Rating * 3,000¥ 	Core
// Body 	Fingertip Compartment 		0.1 			[1] 			4 				3,000¥ 				Core
// Body 	Grapple Gun 				0.5 			[4] 			8 				5,000¥ 				Core
// Body 	Internal Air Tank (R 1-3) 	0.25 			[3] 			Rating 			Rating * 4,500¥ 	Core
// Body 	Muscle Replacement (R 1-4) 	Rating 	- 		(Rating * 5)R 	Rating * 25,000¥ 	Core
// Body 	Reaction Enhancers (R 1-3) 	Rating * 0.3 	- 	(Rating * 5)R 	Rating * 13,000¥ 	Core
// Body 	Skillwires (R 1-6) 	Rating * 0.1 	- 	Rating * 4 	Rating * 20,000¥ 	Core
// Body 	Smuggling Compartment 	0.2 	[2] 	6 	7,500¥ 	Core
// Body 	Wired Reflexes R1 	2 	- 	8R 	39,000¥ 	Core
// Body 	Wired Reflexes R2 	3 	- 	12R 	149,000¥ 	Core
// Body 	Wired Reflexes R3 	5 	- 	20R 	217,000¥ 	Core

func LoadCyberware() map[string]Cyberware {
	logrus.Info("Started loading cyberware")

	files, errReadDir := os.ReadDir(CyberwareDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read cyberware directory")
	}

	list := make(map[string]Cyberware, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", CyberwareDataPath, file.Name())

			var v Cyberware
			if err := util.LoadStructFromYAML(filepath, &v); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load metatype")
			}

			list[v.Name] = v
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded metatype file")
	}

	logrus.WithFields(logrus.Fields{"count": len(list)}).Info("Done loading metatypes")

	return list
}

func LoadCyberwareFile(name string) (*Cyberware, error) {
	var v Cyberware
	if err := util.LoadStructFromYAML(fmt.Sprintf(CyberwareFilename, name), &v); err != nil {
		return nil, err
	}

	return &v, nil
}
