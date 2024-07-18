package core

const (
	SkillGroupsDataPath = "data/skills/groups"
	SkillGroupFilename  = SkillGroupsDataPath + "/%s.yaml"
)

type SkillGroup struct {
	ID          string     `yaml:"id,omitempty"`
	Name        string     `yaml:"name,omitempty"`
	Description string     `yaml:"description,omitempty"`
	Skills      []string   `yaml:"skills"`
	RuleSource  RuleSource `yaml:"rule_source,omitempty"`
}

var CoreSkillGroups = []SkillGroup{
	{
		ID:         "acting",
		Name:       "Acting",
		Skills:     []string{"impersonation", "performance"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "athletics",
		Name:       "Athletics",
		Skills:     []string{"gymnastics", "running", "swimming"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "biotech",
		Name:       "Biotech",
		Skills:     []string{"biotechnology", "cybertechnology", "first_aid", "medicine"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "close_combat",
		Name:       "Close Combat",
		Skills:     []string{"blades", "clubs", "unarmed_combat"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "conjuring",
		Name:       "Conjuring",
		Skills:     []string{"banishing", "binding", "summoning"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "cracking",
		Name:       "Cracking",
		Skills:     []string{"cybercombat", "electronic", "hacking"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "electronics",
		Name:       "Electronics",
		Skills:     []string{"computer", "hardware", "software"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "engineering",
		Name:       "Engineering",
		Skills:     []string{"aeronautics_mechanic", "automotive_mechanic", "industrial_mechanic", "nautical_mechanic"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "firearms",
		Name:       "Firearms",
		Skills:     []string{"automatics", "longarms", "pistols"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "influence",
		Name:       "Influence",
		Skills:     []string{"etiquette", "leadership", "negotiation"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "outdoors",
		Name:       "Outdoors",
		Skills:     []string{"navigation", "survival", "tracking"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "sorcery",
		Name:       "Sorcery",
		Skills:     []string{"counterspelling", "ritual_spellcasting", "spellcasting"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "stealth",
		Name:       "Stealth",
		Skills:     []string{"disguise", "palming", "sneaking"},
		RuleSource: RuleSourceSR5Core,
	},
	{
		ID:         "tasking",
		Name:       "Tasking",
		Skills:     []string{"compiling", "decompiling", "registering"},
		RuleSource: RuleSourceSR5Core,
	},
}

// TODO: LoadSkillGroups loads the skill groups from the data files.
func LoadSkillGroups() map[string]SkillGroup {
	data := make(map[string]SkillGroup)
	return data
}
