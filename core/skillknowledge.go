package core

const (
	KnowledgeSkillDataPath = "data/skills/knowledge"
	KnowledgeSkillFilename = KnowledgeSkillDataPath + "/%s.yaml"
)

type KnowledgeSkillSpec struct {
	ID              string     `yaml:"id,omitempty"`
	Name            string     `yaml:"name"`
	Description     string     `yaml:"description"`
	IsCommon        bool       `yaml:"is_common"`
	Specializations []string   `yaml:"specializations"`
	RuleSource      RuleSource `yaml:"rule_source"`
}

type KnowledgeSkill struct {
	ID                     string             `yaml:"id,omitempty"`
	SelectedSpecialization string             `yaml:"selected_specialization,omitempty"`
	Rating                 int                `yaml:"rating,omitempty"`
	Spec                   KnowledgeSkillSpec `yaml:"-"`
}

var CoreKnowledgeSkills = []KnowledgeSkillSpec{
	{
		ID:              "astronomy",
		Name:            "Astronomy",
		Description:     "The study of celestial bodies and their movements.",
		IsCommon:        true,
		Specializations: []string{"astrophysics", "astrology"},
		RuleSource:      RuleSourceSR5Core,
	},
}
