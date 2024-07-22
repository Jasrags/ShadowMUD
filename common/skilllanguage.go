package common

const (
	LanguageSkillDataPath = "data/skills/languages"
	LanguageSkillFilename = LanguageSkillDataPath + "/%s.yaml"
)

type LanguageSkillSpec struct {
	ID         string     `yaml:"id,omitempty"`
	Name       string     `yaml:"name"`
	IsCommon   bool       `yaml:"is_common"`
	Rating     int        `yaml:"rating,omitempty"`
	RuleSource RuleSource `yaml:"rule_source"`
}

type LanguageSkill struct {
	ID     string            `yaml:"id,omitempty"`
	Rating int               `yaml:"rating,omitempty"`
	Spec   LanguageSkillSpec `yaml:"-"`
}

var (
	CoreLanguageSkills = []LanguageSkillSpec{
		{Name: "English", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Spanish", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Lakota", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Dakota", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Diné (Navajo)", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Russian", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "French", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Italian", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "German", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Aztlaner Spanish", IsCommon: true, RuleSource: "SR5:Core"},
		{Name: "Sperethiel", IsCommon: false, RuleSource: "SR5:Core"},
		{Name: "Or’zet", IsCommon: false, RuleSource: "SR5:Core"},
		{Name: "Japanese", IsCommon: false, RuleSource: "SR5:Core"},
		{Name: "Mandarin", IsCommon: false, RuleSource: "SR5:Core"},
	}
)
