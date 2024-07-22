package common

const (
	MobsDataPath = "data/mobs"
	MobFilename  = MobsDataPath + "/%s/%s.yaml"
)

type MobSpec struct {
	ID         string     `yaml:"id,omitempty"`
	Name       string     `yaml:"name,omitempty"`
	RuleSource RuleSource `yaml:"rule_source,omitempty"`
}

type Mob struct {
	ID   string   `yaml:"id,omitempty"`
	Spec *MobSpec `yaml:"-"`
}
