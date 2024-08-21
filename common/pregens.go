package common

// const (
// 	PregenFilepath = "_data/pregens"
// )

// type (
// 	Pregens map[string]*Pregen
// 	Pregen  struct {
// 		ID          string     `yaml:"id"`
// 		Title       string     `yaml:"title"`
// 		Description string     `yaml:"description"`
// 		Character   *Character `yaml:"character"`
// 		// TODO: Add character fields (metatype, attributes, skills, etc...)
// 		// TODO: Add items (weapons, armor, etc...)
// 	}
// )

// func LoadPregen(title string, u *Pregen) error {
// 	title = strings.ToLower(title)
// 	filepath := fmt.Sprintf("%s/%s.yaml", PregenFilepath, title)

// 	// Check if the user file exists
// 	if _, err := os.Stat(filepath); os.IsNotExist(err) {
// 		return err
// 	}

// 	if err := utils.LoadStructFromYAML(filepath, &u); err != nil {
// 		return err
// 	}

// 	logrus.WithFields(logrus.Fields{"id": u.ID}).Debug("Loaded pregen")

// 	return nil
// }

// func LoadPregens() Pregens {
// 	logrus.Info("Started loading pregens")
// 	list := make(Pregens)

// 	files, errReadDir := os.ReadDir(PregenFilepath)
// 	if errReadDir != nil {
// 		logrus.WithError(errReadDir).Fatal("Could not read pregen directory")
// 	}

// 	for _, file := range files {
// 		var v Pregen
// 		if strings.HasSuffix(file.Name(), ".yaml") {
// 			name := strings.TrimSuffix(file.Name(), ".yaml")
// 			if err := LoadPregen(name, &v); err != nil {
// 				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load pregen")
// 			}

// 			list[v.ID] = &v
// 			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded pregen file")
// 		}
// 	}

// 	logrus.WithFields(logrus.Fields{"count": len(list)}).Info("Done loading pregens")

// 	return list
// }

// func NewPregen() *Pregen {
// 	p := &Pregen{}
// 	return p
// }

// var CorePregens = Pregens{
// 	"street_samurai": {
// 		ID:          "street_samurai",
// 		Title:       "Street Samurai",
// 		Description: "A street samurai is a combat-oriented character that specializes in physical combat and is often augmented with cyberware or bioware.",
// 		Character: &Character{
// 			MetatypeID: "ork",
// 			Attributes: Attributes{
// 				Body:      Attribute[int]{Base: 7},
// 				Agility:   Attribute[int]{Base: 6},
// 				Reaction:  Attribute[int]{Base: 5}, // Should be 7 modified
// 				Strength:  Attribute[int]{Base: 5},
// 				Willpower: Attribute[int]{Base: 3},
// 				Logic:     Attribute[int]{Base: 2},
// 				Intuition: Attribute[int]{Base: 3},
// 				Charisma:  Attribute[int]{Base: 2},
// 				Edge:      Attribute[int]{Base: 1},
// 				Essence:   Attribute[float64]{Base: 6}, // Should be 0.88 modified
// 				Magic:     Attribute[int]{Base: 0},
// 				Resonance: Attribute[int]{Base: 0},
// 			},
// ActiveSkills: ActiveSkills{
// 	"automatics": &ActiveSkill{
// 		ID:     "automatics",
// 		Rating: 5,
// 	},
// 	"blades": &ActiveSkill{
// 		ID:     "blades",
// 		Rating: 5, // Should be 6 modified
// 	},
// 	"longarms": &ActiveSkill{
// 		ID:     "longarms",
// 		Rating: 3, // Should be 4 modified
// 	},
// 	"pilot_ground_craft": &ActiveSkill{
// 		ID:     "pilot_ground_craft",
// 		Rating: 1,
// 	},
// 	"pistols": &ActiveSkill{
// 		ID:     "pistols",
// 		Rating: 4,
// 	},
// 	"sneaking": &ActiveSkill{
// 		ID:     "sneaking",
// 		Rating: 2, // Should be 3 modified
// 	},
// 	"unarmed_combat": &ActiveSkill{
// 		ID:     "unarmed_combat",
// 		Rating: 2, // Should be 3 modified
// 	},
// },
// CreatedAt: time.Now(),
// },
// },
// // {ID: "1", Title: "Street Samurai", Description: "A street samurai is a combat-oriented character that specializes in physical combat and is often augmented with cyberware or bioware."},
// // {ID: "2", Title: "Decker", Description: "A decker is a computer hacker that specializes in cybercombat and data manipulation."},
// // {ID: "3", Title: "Rigger", Description: "A rigger is a character that specializes in controlling drones and vehicles."},
// // {ID: "4", Title: "Mage", Description: "A mage is a character that specializes in spellcasting and astral combat."},
// // {ID: "5", Title: "Shaman", Description: "A shaman is a character that specializes in summoning spirits and casting spells."},
// // {ID: "6", Title: "Adept", Description: "An adept is a character that specializes in physical combat and has magical abilities."},
// // {ID: "7", Title: "Technomancer", Description: "A technomancer is a character that has the ability to interact with the Matrix without the use of a cyberdeck."},
// }
