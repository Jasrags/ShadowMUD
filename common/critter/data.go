package critter

import "github.com/Jasrags/ShadowMUD/common/shared"

var (
	CoreCritters = []Spec{
		{
			ID:          "dog",
			Name:        "Dog",
			Description: "A dog.",
			Category:    CategoryMundaneCritters,
			Attributes: Attributes{
				Body:       Attribute[int]{Min: 4, Max: 7, AugMax: 11},
				Agility:    Attribute[int]{Min: 3, Max: 6, AugMax: 10},
				Reaction:   Attribute[int]{Min: 4, Max: 7, AugMax: 11},
				Strength:   Attribute[int]{Min: 4, Max: 7, AugMax: 11},
				Charisma:   Attribute[int]{Min: 3, Max: 6, AugMax: 10},
				Intuition:  Attribute[int]{Min: 4, Max: 7, AugMax: 11},
				Logic:      Attribute[int]{Min: 2, Max: 5, AugMax: 9},
				Willpower:  Attribute[int]{Min: 3, Max: 6, AugMax: 10},
				Initiative: Attribute[int]{Min: 8, Max: 14, AugMax: 22},
				Edge:       Attribute[int]{Min: 3, Max: 6, AugMax: 6},
				Magic:      Attribute[int]{Min: 0, Max: 3, AugMax: 3},
				Resonance:  Attribute[int]{Min: 0, Max: 3, AugMax: 3},
				Essence:    Attribute[float64]{Min: 0, Max: 6, AugMax: 6},
			},
			// Powers
			//     "Domesticated",
			//     {
			//       "+content": "Enhanced Senses",
			//       "+@select": "Smell"
			//     },
			//     {
			//       "+content": "Enhanced Senses",
			//       "+@select": "Hearing"
			//     },
			//     {
			//       "+content": "Natural Weapon",
			//       "+@select": "Claws/Bite: DV (STR+1)P, AP 0"
			//     }
			//   ]
			// },
			// Skills
			Skills: Skills{
				"intimidation": &Skill{
					ID:     "intimidation",
					Rating: 4,
				},
				"perception": &Skill{
					ID:             "perception",
					Rating:         5,
					Specialization: "Smell",
				},
				"running": &Skill{
					ID:     "running",
					Rating: 5,
				},
				"tracking": &Skill{
					ID:     "tracking",
					Rating: 6,
				},
				"unarmed_combat": &Skill{
					ID:     "unarmed_combat",
					Rating: 5,
				},
			},

			RuleSource: shared.RuleSourceSR5Core,
		},
	}
)
