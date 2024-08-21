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

var CoreSkillGroups = Groups{
	{
		ID:         "acting",
		Name:       "Acting",
		Skills:     []string{"impersonation", "performance"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "athletics",
		Name:       "Athletics",
		Skills:     []string{"gymnastics", "running", "swimming"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "biotech",
		Name:       "Biotech",
		Skills:     []string{"biotechnology", "cybertechnology", "first_aid", "medicine"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "close_combat",
		Name:       "Close Combat",
		Skills:     []string{"blades", "clubs", "unarmed_combat"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "conjuring",
		Name:       "Conjuring",
		Skills:     []string{"banishing", "binding", "summoning"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "cracking",
		Name:       "Cracking",
		Skills:     []string{"cybercombat", "electronic", "hacking"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "electronics",
		Name:       "Electronics",
		Skills:     []string{"computer", "hardware", "software"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "engineering",
		Name:       "Engineering",
		Skills:     []string{"aeronautics_mechanic", "automotive_mechanic", "industrial_mechanic", "nautical_mechanic"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "firearms",
		Name:       "Firearms",
		Skills:     []string{"automatics", "longarms", "pistols"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "influence",
		Name:       "Influence",
		Skills:     []string{"etiquette", "leadership", "negotiation"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "outdoors",
		Name:       "Outdoors",
		Skills:     []string{"navigation", "survival", "tracking"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "sorcery",
		Name:       "Sorcery",
		Skills:     []string{"counterspelling", "ritual_spellcasting", "spellcasting"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "stealth",
		Name:       "Stealth",
		Skills:     []string{"disguise", "palming", "sneaking"},
		RuleSource: shared.RuleSourceSR5Core,
	},
	{
		ID:         "tasking",
		Name:       "Tasking",
		Skills:     []string{"compiling", "decompiling", "registering"},
		RuleSource: shared.RuleSourceSR5Core,
	},
}

// // TODO: LoadSkillGroups loads the skill groups from the data files.
// func LoadSkillGroups() map[string]SkillGroup {
// 	data := make(map[string]SkillGroup)
// 	return data
// }
