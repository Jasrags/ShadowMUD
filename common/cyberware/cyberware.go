package cyberware

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common/shared"
	"github.com/sirupsen/logrus"
)

const (
	CyberwareFilepath = "_data/items/cyberware"

	CyberwareGradeStandard  Grade = "Standard"
	CyberwareGradeAlphaware Grade = "Alphaware"
	CyberwareGradeBetaware  Grade = "Betaware"
	CyberwareGradeDeltaware Grade = "Deltaware"
	CyberwareGradeUsed      Grade = "Used"

	CyberwarePartBody            Part = "Body"
	CyberwarePartHead            Part = "Head"
	CyberwarePartLimb            Part = "Limb"
	CyberwarePartLimbAccessories Part = "Limb Accessories"
	CyberwarePartWeapon          Part = "Weapon"
	CyberwarePartEye             Part = "Eye"
	CyberwarePartEar             Part = "Ear"
)

type (
	Grade string
	Part  string
	Specs map[string]*Spec
	Spec  struct {
		ID            string                    `yaml:"id,omitempty"`
		Name          string                    `yaml:"name"`
		Description   string                    `yaml:"description"`
		EssenceCost   shared.Attribute[float64] `yaml:"essence_cost"`
		Capacity      shared.Attribute[int]     `yaml:"capacity"`
		Rating        int                       `yaml:"rating,omitempty"`
		CyberwarePart Part                      `yaml:"cyberware_part"`
		Grade         Grade                     `yaml:"grade,omitempty"`
		ToggleAction  shared.ActionType         `yaml:"toggle_action,omitempty"`
		IsActive      bool                      `yaml:"is_active,omitempty"`
		Modifications []Modifications           `yaml:"modifications"`
		Modifiers     shared.Modifiers          `yaml:"modifiers"`
		Cost          shared.Attribute[int]     `yaml:"cost"`
		Availability  int                       `yaml:"availability"`
		Legality      shared.LegalityType       `yaml:"legality"`
		Notes         string                    `yaml:"notes"`
		RuleSource    shared.RuleSource         `yaml:"rule_source"`
	}
	Cyberwares map[string]*Cyberware
	Cyberware  struct {
		ID            string           `yaml:"id,omitempty"`
		Rating        int              `yaml:"rating,omitempty"`
		Modifications []Modifications  `yaml:"modifications"`
		Modifiers     shared.Modifiers `yaml:"modifiers"`
		Spec          Spec             `yaml:"-"`
	}
)

var (
	ErrUnknownCyberwareGrade = fmt.Errorf("unknown cyberware grade")
)

func GetCyberwareGradeModifiers(grade Grade) (float64, int, float64, error) {
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

// func (c *Cyberware) Recalculate() {
// 	c.Spec.Capacity.Reset()

// 	essenceCostMultiplier, _, _, _ := GetCyberwareGradeModifiers(c.Spec.Grade)
// 	for _, mod := range c.Modifications {
// 		c.Spec.Capacity.Mods += mod.CapacityCost
// 		c.Spec.EssenceCost.Mods += mod.EssenceCost * essenceCostMultiplier
// 	}
// 	c.Spec.EssenceCost.Reset()
// 	for _, mod := range c.Modifications {
// 		c.Spec.EssenceCost.Mods += mod.EssenceCost
// 	}
// }

var CoreCyberware = []Cyberware{
	{
		ID:     "wired_reflexes_r1",
		Rating: 1,
		Spec: Spec{
			Name:        "Wired Reflexes R1",
			Description: "Wired reflexes increase the user's reaction time, allowing them to react more quickly to threats.",
			EssenceCost: shared.Attribute[float64]{
				Base: 2.0,
			},
			Capacity: shared.Attribute[int]{
				Base: 0,
			},
			Rating:        1,
			CyberwarePart: CyberwarePartBody,
			Grade:         CyberwareGradeStandard,
			ToggleAction:  shared.ActionFree,
			IsActive:      false,
			Modifications: []Modifications{},
			// Modifiers:     []Modifier{},
			Cost: shared.Attribute[int]{
				Base: 39000,
			},
			Availability: 8,
			Legality:     shared.LegalityTypeLegal,
			Notes:        "",
			RuleSource:   shared.RuleSourceSR5Core,
		},
	},
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
