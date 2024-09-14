package jason

type MagicType struct {
	ID          string
	Name        string
	Description string
	PointCost   int
}

var CoreMagicTypes = map[string]MagicType{
	"none": {
		ID:          "none",
		Name:        "None",
		Description: "No magical abilities.",
		PointCost:   0,
	},
	"adept": {
		ID:          "adept",
		Name:        "Adept",
		Description: "Adept magic users who enhance their physical abilities.",
		PointCost:   20,
	},
	"magician": {
		ID:          "magician",
		Name:        "Magician",
		Description: "Magicians who can cast spells and summon spirits.",
		PointCost:   15,
	},
	"aspected_magician": {
		ID:          "aspected_magician",
		Name:        "Aspected Magician",
		Description: "Magicians with a focus on a specific type of magic.",
		PointCost:   30,
	},
	"mystic_adept": {
		ID:          "mystic_adept",
		Name:        "Mystic Adept",
		Description: "A combination of adept and magician abilities.",
		PointCost:   35,
	},
	"technomancer": {
		ID:          "technomancer",
		Name:        "Technomancer",
		Description: "Users of magic that interact with technology.",
		PointCost:   15,
	},
}
