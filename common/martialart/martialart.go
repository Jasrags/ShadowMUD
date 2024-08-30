package martialart

import "github.com/Jasrags/ShadowMUD/common/shared"

const ()

type (
	Technique struct {
		ID         string            `yaml:"id"`
		Name       string            `yaml:"name"`
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
	Spec struct {
		ID          string `yaml:"id"`
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		// Bonus
		// Required
		Techniques []string          `yaml:"techniques"`
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
)
