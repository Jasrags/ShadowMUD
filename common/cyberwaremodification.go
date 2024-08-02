package common

const (
	CyberwareModificationsFilepath = "_data/items/cyberware/modifications"
)

type (
	CyberwareModifications struct {
		ID           string     `yaml:"id,omitempty"`
		Name         string     `yaml:"name"`
		Description  string     `yaml:"description"`
		EssenceCost  float64    `yaml:"essence_cost"`
		CapacityCost int        `yaml:"capacity,omitempty"`
		RuleSource   RuleSource `yaml:"rule_source"`
	}
)

var (
	CoreCyberwareModifications = []CyberwareModifications{}
)
