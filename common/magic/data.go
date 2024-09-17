package magic

import "github.com/Jasrags/ShadowMUD/common/shared"

var CoreMagicTypes MagicTypes = map[string]MagicType{
	"none": {
		ID:          "none",
		Name:        "None",
		Description: "No magical abilities.",
		PointCost:   0,
		Hidden:      false,
		RuleSource:  shared.RuleSourceSR5Core,
	},
	"adept": {
		ID:          "adept",
		Name:        "Adept",
		Description: "Adept magic users who enhance their physical abilities.",
		PointCost:   20,
		Hidden:      false,
		RuleSource:  shared.RuleSourceSR5Core,
	},
	"magician": {
		ID:          "magician",
		Name:        "Magician",
		Description: "Magicians who can cast spells and summon spirits.",
		PointCost:   15,
		Hidden:      false,
		RuleSource:  shared.RuleSourceSR5Core,
	},
	"aspected_magician": {
		ID:          "aspected_magician",
		Name:        "Aspected Magician",
		Description: "Magicians with a focus on a specific type of magic.",
		PointCost:   30,
		Hidden:      false,
		RuleSource:  shared.RuleSourceSR5Core,
	},
	"mystic_adept": {
		ID:          "mystic_adept",
		Name:        "Mystic Adept",
		Description: "A combination of adept and magician abilities.",
		PointCost:   35,
		Hidden:      false,
		RuleSource:  shared.RuleSourceSR5Core,
	},
	"technomancer": {
		ID:          "technomancer",
		Name:        "Technomancer",
		Description: "Users of magic that interact with technology.",
		PointCost:   15,
		Hidden:      false,
		RuleSource:  shared.RuleSourceSR5Core,
	},
}
