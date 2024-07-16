package core

const (
	ArmorDataPath        = "data/items/armor"
	ArmorFilename        = ArmorDataPath + "/%s.yaml"
	ArmorFFileMinVersion = "0.0.1"
)

var (
	ArmorItems = map[string]Armor{}

	// tam1 = ArmorMod{
	// 	Name:         "Chemical Seal",
	// 	Rating:       2,
	// 	Capacity:     6,
	// 	Cost:         3000,
	// 	Availability: 12,
	// 	Legality:     LegalityTypeRestricted,
	// 	RuleSource:   "SR5:Core",
	// 	FileVersion:  "0.0.1",
	// }
	// Activate Seal
	//  Complex Action
	// Air Supply
	//  1 Hour
	// Wireless
	//  Activating the chemical seal is a Free Action.

	// tam2 = ArmorMod{
	// 	Name:         "Radiation Shielding",
	// 	Rating:       2,
	// 	Capacity:     6,    // Rating
	// 	Cost:         3000, // Rating x 200
	// 	Availability: 12,   // Rating x 2
	// 	RuleSource:   "SR5:R&G",
	// 	FileVersion:  "0.0.1",
	// }
	// Special Rules
	//  Each point of Radiation shielding provides 1 extra die for resisting Radiation damage (see p. 152).

	// ta = Armor{
	// 	Name:         "Armante Dress",
	// 	ArmorRating:  8,
	// 	Capacity:     4,
	// 	Availability: 10,
	// 	Cost:         2500,
	// 	Mods:         []ArmorMod{},
	// 	RuleSource:   "SR5:R&G ",
	// 	FileVersion:  "0.0.1",
	// }
	// Features
	//  Increase Social Limit by 2
	// Wireless Bonus
	//  +1 dice pool bonus to Social Tests
)

type Armor struct {
	ID           string
	Name         string
	Description  string
	Environment  EnvironmentType
	ArmorRating  int
	Capacity     int
	Cost         int
	Availability int
	Legality     LegalityType
	Mods         []ArmorMod
	RuleSource   string
	FileVersion  string
}

func (a *Armor) AddMod(mod *ArmorMod) error {
	a.Mods = append(a.Mods, *mod)

	return nil
}

type ArmorMod struct {
	ID           string
	Name         string
	Description  string
	ArmorRating  int
	Rating       int
	Cost         int
	Capacity     int
	Availability int
	Legality     LegalityType
	RuleSource   string
	FileVersion  string
}

func (am *ArmorMod) GetCapacity(f func(am *ArmorMod) int) int {
	return f(am)
}
