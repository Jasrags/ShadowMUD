package skill

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	SkillGroupsFilepath = "_data/skills/groups"
)

type (
	Groups []Group
	Group  struct {
		ID          string            `yaml:"id,omitempty"`
		Name        string            `yaml:"name,omitempty"`
		Description string            `yaml:"description,omitempty"`
		Skills      []string          `yaml:"skills"`
		RuleSource  shared.RuleSource `yaml:"rule_source,omitempty"`
	}
)
