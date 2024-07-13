package metatype

type Metatype struct {
	Name string

	Body      []int
	Agility   []int
	Reaction  []int
	Strength  []int
	Willpower []int
	Logic     []int
	Intuition []int
	Charisma  []int
	Edge      []int
	Essence   float64
	// Initiative   int // REA+INT
	RacialTraits []string
}

var (
	Metatypes = map[string]Metatype{
		"Human": {
			Name:      "Human",
			Body:      []int{1, 6},
			Agility:   []int{1, 6},
			Reaction:  []int{1, 6},
			Strength:  []int{1, 6},
			Willpower: []int{1, 6},
			Logic:     []int{1, 6},
			Intuition: []int{1, 6},
			Charisma:  []int{1, 6},
			Edge:      []int{2, 7},
			Essence:   6,
		},
		"Elf": {
			Name:         "Elf",
			Body:         []int{1, 6},
			Agility:      []int{2, 7},
			Reaction:     []int{1, 6},
			Strength:     []int{1, 6},
			Willpower:    []int{1, 6},
			Logic:        []int{1, 6},
			Intuition:    []int{1, 6},
			Charisma:     []int{3, 8},
			Edge:         []int{1, 6},
			Essence:      6,
			RacialTraits: []string{RacialTraitLowLightVision},
		},
		"Dwarf": {
			Name:         "Dwarf",
			Body:         []int{3, 8},
			Agility:      []int{1, 6},
			Reaction:     []int{1, 5},
			Strength:     []int{3, 8},
			Willpower:    []int{2, 7},
			Logic:        []int{1, 6},
			Intuition:    []int{1, 6},
			Charisma:     []int{1, 6},
			Edge:         []int{1, 6},
			Essence:      6,
			RacialTraits: []string{RacialTraitThermographicVision},
		},
		"Ork": {
			Name:         "Ork",
			Body:         []int{4, 9},
			Agility:      []int{1, 6},
			Reaction:     []int{1, 6},
			Strength:     []int{3, 8},
			Willpower:    []int{1, 6},
			Logic:        []int{1, 5},
			Intuition:    []int{1, 6},
			Charisma:     []int{1, 6},
			Edge:         []int{1, 5},
			Essence:      6,
			RacialTraits: []string{RacialTrait2DicForPathogenAndToxinResistance, RacialTrait20PercentIncreasedLifestyleCost},
		},
		"Troll": {
			Name:         "Troll",
			Body:         []int{5, 10},
			Agility:      []int{1, 5},
			Reaction:     []int{1, 6},
			Strength:     []int{5, 10},
			Willpower:    []int{1, 6},
			Logic:        []int{1, 5},
			Intuition:    []int{1, 5},
			Charisma:     []int{1, 4},
			Edge:         []int{1, 6},
			Essence:      6,
			RacialTraits: []string{RacialTraitThermographicVision, RacialTrait1Reach, RacialTrait1DermalArmor, RacialTraitDoubleLifestyleCosts},
		},
	}
)

const (
	RacialTraitLowLightVision                    = "Low-Light Vision"
	RacialTraitThermographicVision               = "Thermographic Vision"
	RacialTrait2DicForPathogenAndToxinResistance = "+2 dice for pathogen and toxin resistance"
	RacialTrait20PercentIncreasedLifestyleCost   = "+20% increased Lifestyle cost"
	RacialTrait1Reach                            = "+1 Reach"
	RacialTrait1DermalArmor                      = "+1 dermal armor"
	RacialTraitDoubleLifestyleCosts              = "Double Lifestyle costs"
)
