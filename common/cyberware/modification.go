package cyberware

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	CyberwareModificationsFilepath = "_data/items/cyberware/modifications"
)

type (
	Modifications struct {
		ID           string            `yaml:"id,omitempty"`
		Name         string            `yaml:"name"`
		Description  string            `yaml:"description"`
		EssenceCost  float64           `yaml:"essence_cost"`
		CapacityCost int               `yaml:"capacity,omitempty"`
		Modifiers    shared.Modifiers  `yaml:"modifiers"`
		RuleSource   shared.RuleSource `yaml:"rule_source"`
	}
)

var (
	CoreCyberwareModifications = []Modifications{}
)
