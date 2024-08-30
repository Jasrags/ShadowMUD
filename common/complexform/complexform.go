package complexform

import "github.com/Jasrags/ShadowMUD/common/shared"

type (
	Spec struct {
		ID          string            `yaml:"id"`
		Name        string            `yaml:"name"`
		Description string            `yaml:"description"`
		Target      string            `yaml:"target"`
		Duration    string            `yaml:"duration"`
		ForceValue  int               `yaml:"force_value"`
		RuleSource  shared.RuleSource `yaml:"rule_source"`
	}
)
