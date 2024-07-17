package core

type Bioware struct {
	ID            string                   `yaml:"id,omitempty"`
	Name          string                   `yaml:"name"`
	Description   string                   `yaml:"description"`
	EssenceCost   float64                  `yaml:"essence_cost"`
	Capacity      AttributesInfo           `yaml:"capacity"`
	Rating        int                      `yaml:"rating,omitempty"`
	ToggleAction  ActionType               `yaml:"toggle_action,omitempty"`
	IsActive      bool                     `yaml:"is_active,omitempty"`
	Modifications []CyberwareModifications `yaml:"modifications"`
	Modifiers     []CyberwareModifier      `yaml:"modifiers"`
	Cost          int                      `yaml:"cost"`
	Availability  int                      `yaml:"availability"`
	Legality      LegalityType             `yaml:"legality"`
	Notes         string                   `yaml:"notes"`
	RuleSource    RuleSource               `yaml:"rule_source"`
	FileVersion   string                   `yaml:"file_version"`
}
