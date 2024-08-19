package common

const (
	BiowareFilepath = "_data/items/bioware"

	BiowareTypeBasic    BiowareType = "Basic"
	BiowareTypeCultured BiowareType = "Cultured"
)

type (
	BiowareType string
	BiowareSpec struct {
		ID           string         `yaml:"id,omitempty"`
		Name         string         `yaml:"name,omitempty"`
		Description  string         `yaml:"description,omitempty"`
		BiowareType  BiowareType    `yaml:"bioware_type,omitempty"`
		EssenceCost  float64        `yaml:"essence_cost,omitempty"`
		Capacity     Attribute[int] `yaml:"capacity,omitempty"`
		Rating       int            `yaml:"rating,omitempty,omitempty"`
		ToggleAction ActionType     `yaml:"toggle_action,omitempty,omitempty"`
		IsActive     bool           `yaml:"is_active,omitempty"`
		Modifiers    []Modifier     `yaml:"modifiers"`
		Cost         int            `yaml:"cost,omitempty"`
		Availability int            `yaml:"availability,omitempty"`
		Legality     LegalityType   `yaml:"legality,omitempty"`
		Notes        string         `yaml:"notes,omitempty"`
		RuleSource   RuleSource     `yaml:"rule_source,omitempty"`
	}
	Bioware struct {
		ID        string      `yaml:"id,omitempty"`
		Rating    int         `yaml:"rating,omitempty,omitempty"`
		Modifiers []Modifier  `yaml:"modifiers"`
		Spec      BiowareSpec `yaml:"-"`
	}
)

var CoreBioware = []Bioware{
	{
		ID:        "adrenaline_pump",
		Rating:    1,
		Modifiers: []Modifier{},
		Spec: BiowareSpec{
			ID:           "adrenaline_pump",
			Name:         "Adrenaline Pump",
			Description:  "Adrenaline Pump Description",
			BiowareType:  BiowareTypeBasic,
			EssenceCost:  0.75, // Rating * 0.75
			Capacity:     Attribute[int]{},
			ToggleAction: ActionFree,
			IsActive:     false,
			Modifiers:    []Modifier{},
			Cost:         55000, // Rating * 55,000¥
			Availability: 0,     // (Rating * 6)F
			Legality:     LegalityTypeLegal,
			Notes:        "",
			RuleSource:   RuleSourceSR5Core,
		},
	},
}

// BASIC
// Type 	Device 	Essence 	Avail 	Cost 	Source
// Basic 	Adrenaline Pump (R 1-3) 	Rating * 0.75 	(Rating * 6)F 	Rating * 55,000¥ 	Core
// Basic 	Bone Density Augmentation (R 1-4) 	Rating * 0.3 	Rating * 4 	Rating * 5,000¥ 	Core
// Basic 	Cat's Eye 	0.1 	4 	4,000¥ 	Core
// Basic 	Enhanced Articulation 	0.3 	12 	24,000¥ 	Core
// Basic 	Muscle Augmentation (R 1-4) 	Rating * 0.2 	(Rating * 5)R 	Rating * 31,000¥ 	Core
// Basic 	Muscle Toner (R 1-4) 	Rating * 0.2 	(Rating * 5)R 	Rating * 32,000¥ 	Core
// Basic 	Orthoskin (R 1-4) 	Rating * 0.25 	(Rating * 4)R 	Rating * 6,000¥ 	Core
// Basic 	Pathogenic Defense (R 1-6) 	Rating * 0.1 	Rating * 2 	Rating * 4,500¥ 	Core
// Basic 	Platelet Factories 	0.2 	12 	17,000¥ 	Core
// Basic 	Skin Pocket 	0.1 	4 	12,000¥ 	Core
// Basic 	Suprathyroid Gland 	0.7 	20R 	140,000¥ 	Core
// Basic 	Symbiotes (R 1-4) 	Rating * 0.2 	Rating * 5 	Rating * 3,500¥ 	Core
// Basic 	Synthacardium (R 1-3) 	Rating * 0.1 	Rating * 4 	Rating * 30,000¥ 	Core
// Basic 	Tailored Pheromones (R 1-3) 	Rating * 0.2 	(Rating * 4)R 	31,000¥ 	Core
// Basic 	Toxin Extractor (R 1-6) 	Rating * 0.2 	Rating * 3 	Rating * 4,800¥ 	Core
// Basic 	Tracheal Filter (R 1-6) 	Rating * 0.1 	Rating * 3 	Rating * 4,500¥ 	Core

// CULTURED
// Type 	Device 	Essence 	Avail 	Cost 	Source
// Cultured 	Cerebral Booster (R 1-3) 	Rating * 0.2 	Rating * 6 	Rating * 31,500¥ 	Core
// Cultured 	Damage Compensators (R 1-12) 	Rating * 0.1 	(Rating * 3)F 	Rating * 2,000¥ 	Core
// Cultured 	Mnemonic Enhancer (R 1-3) 	Rating * 0.1 	Rating * 5 	Rating * 9,000¥ 	Core
// Cultured 	Pain Editor 	0.3 	18F 	48,000¥ 	Core
// Cultured 	Reflex Recorder (Skill) 	0.1 	10 	14,000¥ 	Core
// Cultured 	Sleep Regulator 	0.1 	6 	12,000¥ 	Core
// Cultured 	Synaptic Booster (R 1-3) 	Rating * 0.5 	(Rating * 6)R 	Rating * 95,000¥ 	Core
