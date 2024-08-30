package drugcomponent

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	VectorInhalation Vector = "Inhalation"
	VectorInjection  Vector = "Injection"
)

type (
	Vectors  []Vector
	Vector   string
	Category string
	Grade    struct {
		ID         string            `yaml:"id"`
		Name       string            `yaml:"name"`
		Cost       int               `yaml:"cost"`
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
	Spec struct {
		ID           string              `yaml:"id"`
		Name         string              `yaml:"name"`
		Description  string              `yaml:"description"`
		Category     Category            `yaml:"category"`
		Rating       int                 `yaml:"rating"`
		Availability int                 `yaml:"availability"`
		Legality     shared.LegalityType `yaml:"legality"`
		Cost         int                 `yaml:"cost"`
		Speed        int                 `yaml:"speed"`
		Vectors      Vectors             `yaml:"vectors"`
		Duration     string              `yaml:"duration"`
		// bonus
		//   attribute
		//   limit
		//   quality
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
)
