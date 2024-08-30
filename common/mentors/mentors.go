package mentors

import "github.com/Jasrags/ShadowMUD/common/shared"

const ()

type (
	Choice struct {
		ID   string `yaml:"id"`
		Name string `yaml:"name"`
		// Bonus
	}
	Spec struct {
		ID           string `yaml:"id"`
		Name         string `yaml:"name"`
		Description  string `yaml:"description"`
		Advantage    string `yaml:"advantage"`
		Disadvantage string `yaml:"disadvantage"`
		// Bonus
		Choices    []Choice          `yaml:"choices"`
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
)
