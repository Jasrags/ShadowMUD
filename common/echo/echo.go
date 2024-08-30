package echo

import "github.com/Jasrags/ShadowMUD/common/shared"

const ()

type (
	Spec struct {
		ID          string `yaml:"id"`
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Limit       int    `yaml:"limit"`
		// Bonus
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
)
