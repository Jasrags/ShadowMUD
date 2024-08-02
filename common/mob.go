package common

const (
	MobsDataPath = "_data/mobs"
)

type (
	MobSpec struct {
		ID         string     `yaml:"id,omitempty"`
		Name       string     `yaml:"name,omitempty"`
		RuleSource RuleSource `yaml:"rule_source,omitempty"`
	}
	Mob struct {
		ID   string   `yaml:"id,omitempty"`
		Spec *MobSpec `yaml:"-"`
	}
)
