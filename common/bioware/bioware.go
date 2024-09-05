package bioware

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	BiowareTypeBasic    Type = "Basic"
	BiowareTypeCultured Type = "Cultured"
)

type (
	Type     string
	Capacity struct {
		Base       int `yaml:"base"`
		Delta      int `yaml:"delta"`
		TotalValue int `yaml:"total_value"`
	}
	Specs map[string]Spec
	Spec  struct {
		ID           string              `yaml:"id"`
		Name         string              `yaml:"name"`
		Description  string              `yaml:"description"`
		BiowareType  Type                `yaml:"bioware_type"`
		EssenceCost  float64             `yaml:"essence_cost"`
		Capacity     Capacity            `yaml:"capacity"`
		Rating       int                 `yaml:"rating"`
		ToggleAction shared.ActionType   `yaml:"toggle_action"`
		IsActive     bool                `yaml:"is_active"`
		Modifiers    shared.Modifiers    `yaml:"modifiers"`
		Cost         int                 `yaml:"cost"`
		Availability int                 `yaml:"availability"`
		Legality     shared.LegalityType `yaml:"legality"`
		Notes        string              `yaml:"notes"`
		RuleSource   shared.RuleSource   `yaml:"rule_source"`
	}
)

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
