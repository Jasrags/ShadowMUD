package common

// const (
// 	SkillsFilepath = "_data/skills"

// 	// Skill Categories
// 	SkillCategoryCombat        SkillCategory = "Combat Active"
// 	SkillCategoryPhysical      SkillCategory = "Physical Active"
// 	SkillCategorySocial        SkillCategory = "Social Active"
// 	SkillCategoryMagical       SkillCategory = "Magical Active"
// 	SkillCategoryPseudoMagical SkillCategory = "Pseudo-Magical Active"
// 	SkillCategoryResonance     SkillCategory = "Resonance Active"
// 	SkillCategoryTechnical     SkillCategory = "Technical Active"
// 	SkillCategoryVehicle       SkillCategory = "Vehicle Active"
// 	// Knowledge Skills
// 	SkillCategoryAcademic     SkillCategory = "Academic"
// 	SkillCategoryInterest     SkillCategory = "Interest"
// 	SkillCategoryLanguage     SkillCategory = "Language"
// 	SkillCategoryProfessional SkillCategory = "Professional"
// 	SkillCategoryStreet       SkillCategory = "Street"
// )

// type (
// 	SkillID       string
// 	SkillCategory string
// 	SkillSpecs    map[string]*SkillSpec
// 	SkillSpec     struct {
// 		ID                     string        `yaml:"id"`
// 		Name                   string        `yaml:"name"`
// 		Hidden                 bool          `yaml:"hidden"`
// 		Category               SkillCategory `yaml:"category"`
// 		Description            string        `yaml:"description"`
// 		Defaultable            bool          `yaml:"defaultable"`
// 		Exotic                 bool          `yaml:"exotic"`
// 		RequiresGroundMovement bool          `yaml:"requires_ground_movement"`
// 		RequiresSwimMovement   bool          `yaml:"requires_swim_movement"`
// 		RequiresFlyMovement    bool          `yaml:"requires_fly_movement"`
// 		LinkedAttribute        AttributeType `yaml:"linked_attribute"`
// 		Group                  string        `yaml:"group"`
// 		Specializations        []string      `yaml:"specializations"`
// 		RuleSource             RuleSource    `yaml:"rule_source"`
// 	}
// 	Skills map[string]*Skill
// 	Skill  struct {
// 		ID             string `yaml:"id"`
// 		Specialization string `yaml:"specialization"`
// 		Rating         int    `yaml:"rating"`
// 		// Modifiers              Modifiers  `yaml:"modifiers"`
// 		Spec *SkillSpec `yaml:"-"`
// 	}
// )

// type SkillManager struct {
// 	Active    SkillSpecs
// 	Knowledge SkillSpecs
// }

// func NewSkillManager() *SkillManager {
// 	return &SkillManager{
// 		Active:    make(SkillSpecs),
// 		Knowledge: make(SkillSpecs),
// 	}
// }

// func (sm *SkillManager) LoadSkill(id string) (*SkillSpec, error) {
// 	logrus.WithFields(logrus.Fields{"id": id}).Debug("Loading skill")
// 	return nil, nil
// }

// func (sm *SkillManager) LoadSkills() error {
// 	logrus.Info("Loading skills")

// 	for _, skill := range CoreSkills {
// 		switch skill.Category {
// 		case SkillCategoryCombat, SkillCategoryPhysical, SkillCategorySocial, SkillCategoryMagical, SkillCategoryPseudoMagical, SkillCategoryResonance, SkillCategoryTechnical, SkillCategoryVehicle:
// 			sm.Active[skill.ID] = &skill
// 		case SkillCategoryAcademic, SkillCategoryInterest, SkillCategoryLanguage, SkillCategoryProfessional, SkillCategoryStreet:
// 			sm.Knowledge[skill.ID] = &skill
// 		}
// 	}

// 	logrus.Infof("Loaded %d active skills", len(sm.Active))
// 	logrus.Infof("Loaded %d knowledge skills", len(sm.Knowledge))

// 	// sm.Active = Core
// 	// var list []Skill
// 	// if err := utils.LoadStructFromYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "active"), &list); err != nil {
// 	// 	logrus.WithError(err).Error("Error loading active skills")
// 	// 	return err
// 	// }
// 	// for _, skill := range sl.Skill {
// 	// 	logrus.WithFields(logrus.Fields{"id": skill.ID, "name": skill.Name}).Debug("Loaded active skill")
// 	// }
// 	// if err := utils.LoadStructFromYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "knowledge"), &sm.Knowledge); err != nil {
// 	// 	logrus.WithError(err).Error("Error loading knowledge skills")
// 	// 	return err
// 	// }
// 	// for _, skill := range sm.Knowledge {
// 	// 	logrus.WithFields(logrus.Fields{"id": skill.ID, "name": skill.Name}).Debug("Loaded knowledge skill")
// 	// }

// 	return nil
// }

// func (sm *SkillManager) SaveSkill(skill *Skill) error {
// 	logrus.WithFields(logrus.Fields{"id": skill.ID}).Debug("Saving skill")
// 	return nil
// }

// func (sm *SkillManager) SaveSkills() error {
// 	logrus.Debug("Saving skills")
// 	// // Save Active Skills
// 	// if err := utils.SaveStructToYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "active"), sm.Active); err != nil {
// 	// 	logrus.WithError(err).Error("Error saving active skills")
// 	// 	return err
// 	// }
// 	// // Save Knowledge Skills
// 	// if err := utils.SaveStructToYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "knowledge"), sm.Knowledge); err != nil {
// 	// 	logrus.WithError(err).Error("Error saving knowledge skills")
// 	// 	return err
// 	// }

// 	return nil
// }

// func (sm *SkillManager) GetInstance(id string) (*Skill, error) {
// 	logrus.WithFields(logrus.Fields{"id": id}).Debug("Getting skill instance")

// 	return nil, nil
// }

// // func (sm *SkillManager) Get(id string) (*Skill, error) {
// // 	log := logrus.WithFields(logrus.Fields{"id": id})
// // 	log.Debug("Getting skill")

// // 	skill, ok := sm.Skills[id]
// // 	if !ok {
// // 		log.Error("Skill not found")
// // 		return nil, fmt.Errorf("Skill not found")
// // 	}

// // 	return skill, nil
// // }

// // func (sm *SkillManager) Add(skill *Skill) error {
// // 	logrus.WithFields(logrus.Fields{"id": skill.ID}).Debug("Adding skill")
// // 	sm.Skills[skill.ID] = skill

// // 	return nil
// // }

// // func (sm *SkillManager) Remove(id string) error {
// // 	logrus.WithFields(logrus.Fields{"id": id}).Debug("Removing skill")
// // 	delete(sm.Skills, id)

// // 	return nil
// // }

// var CoreSkills = []SkillSpec{
// 	// BODY
// 	{
// 		ID:              "diving",
// 		Name:            "Diving",
// 		Description:     "Diving brings together a wide array of actions performed underwater. This skill can be applied when diving, swimming underwater, using complex diving equipment, and holding your breath.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		LinkedAttribute: AttributeBody,
// 		Specializations: []string{
// 			"Arctic",
// 			"Cave",
// 			"Commercial",
// 			"Controlled Hyperventilation",
// 			"Liquid Breathing Apparatus",
// 			"Military",
// 			"Mixed Gas",
// 			"Oxygen Extraction",
// 			"SCUBA"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "free_fall",
// 		Name:            "Free-Fall",
// 		Description:     "This skill covers any jump from height, including leaps from a third-floor window to jumps from a plane at high altitude. If it involves any kind of attempt to slow or control your fall, this covers it, so it includes skydiving with a parachute, flying a wingsuit, or descending on a line, bungee cord, or zipline.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		LinkedAttribute: AttributeBody,
// 		Specializations: []string{
// 			"BASE Jumping",
// 			"Break-Fall",
// 			"Bungee",
// 			"HALO",
// 			"Low Altitude",
// 			"Parachute",
// 			"Rappelling",
// 			"Static Line",
// 			"Wingsuit",
// 			"Zipline"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// AGILITY
// 	{
// 		ID:              "archery",
// 		Name:            "Archery",
// 		Description:     "Archery is used to fire string-loaded projectile weapons. An archer is familiar with many different styles of bow and the multitude of arrows that can be used to maximum effect.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Bow",
// 			"Crossbow",
// 			"Non-Standard Ammunition",
// 			"Slingshot"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "automatics",
// 		Name:            "Automatics",
// 		Description:     "The Automatics skill covers a specific subset of firearms larger than handheld pistols but smaller than rifles. This category includes submachine guns and other fully automatic carbines.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		Group:           "firearms",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Assault Rifles",
// 			"Cyber-Implant",
// 			"Machine Pistols",
// 			"Submachine Guns",
// 			"Carbines"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "blades",
// 		Name:            "Blades",
// 		Description:     "Slice and dice! The Blades skill includes the use of all handheld slashing and stabbing weapons. You can use a range of edged weapons including daggers, swords, and axes.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		Group:           "close_combat",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Axes",
// 			"Knives",
// 			"Swords",
// 			"Parrying"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "clubs",
// 		Name:            "Clubs",
// 		Description:     "Clubs governs the use of all hand-held bludgeoning instruments. With this skill you can turn any blunt item, be it a baseball bat, crutch, or mace, into a weapon.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		Group:           "close_combat",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Batons",
// 			"Hammers",
// 			"Saps",
// 			"Staves",
// 			"Parrying"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "escape_artist",
// 		Name:            "Escape Artist",
// 		Description:     "Escape Artist measures the character’s ability to escape from bindings by using body contortion and manual dexterity.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Contortionism",
// 			"Cuffs",
// 			"Ropes",
// 			"Zip Ties"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "exotic_melee_weapon",
// 		Name:            "Exotic Melee Weapon",
// 		Description:     "Sometimes a regular gun or blade won’t do the job and you need something fancier. Or weirder. This skill must be taken once for each unusual ranged weapon you want to use. Some examples include blowguns, gyrojet pistols, flamethrowers, and lasers.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryCombat,
// 		Exotic:          true,
// 		LinkedAttribute: AttributeAgility,
// 		RuleSource:      RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "exotic_ranged_weapon",
// 		Name:            "Exotic Ranged Weapon",
// 		Description:     "Sometimes a regular gun or blade won’t do the job and you need something fancier. Or weirder. This skill must be taken once for each unusual ranged weapon you want to use. Some examples include blowguns, gyrojet pistols, flamethrowers, and lasers.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryCombat,
// 		Exotic:          true,
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Blowguns",
// 			"Gyrojet Pistols",
// 			"Flamethrowers",
// 			"Lasers",
// 			"Spidersilk"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "gunnery",
// 		Name:            "Gunnery",
// 		Description:     "Gunnery is used when firing any vehicle-mounted weapon, regardless of how or where the weapon is mounted. This skill extends to manual and sensor-enhanced gunnery.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryVehicle,
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Artillery",
// 			"Ballistic",
// 			"Energy",
// 			"Guided Missile",
// 			"Rocket"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "gymnastics",
// 		Name:            "Gymnastics",
// 		Description:     "Gymnastics measures your balance, general athleticism, and all-around ability to use your body.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		Group:           "athletics",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Balance",
// 			"Climbing",
// 			"Dance",
// 			"Dodging",
// 			"Leaping",
// 			"Parkour",
// 			"Rolling"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "heavy_weapons",
// 		Name:            "Heavy Weapons",
// 		Description:     "The term heavy weapon is designated for all projectile weaponry larger than an assault rifle, such as grenade launchers, machine guns, and assault cannons. This skill is exclusive to handheld and non-vehicle-mounted weaponry—if you’ve got a gun mounted on or in a vehicle, use Gunnery.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Assault Cannons",
// 			"Grenade Launchers",
// 			"Guided Missiles",
// 			"Machine Guns",
// 			"Rocket Launchers"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "locksmith",
// 		Name:            "Locksmith",
// 		Description:     "This skill covers building, repairing, and opening mechanical and electronic locks. While largely banished to antiquity, traditional mechanical locking mechanisms are still in use around the globe, often as throwbacks or backups. Electronic locks are far more common and quite susceptible to your ministrations.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Combination",
// 			"Keypad",
// 			"Maglock",
// 			"Tumbler",
// 			"Voice Recognition"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "longarms",
// 		Name:            "Longarms",
// 		Description:     "The Longarms skill is for firing extended-barrel weapons such as sporting rifles and sniper rifles. This grouping also includes weapons like shotguns that are designed to be braced against the shoulder.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		Group:           "firearms",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Extended-Range Shots",
// 			"Long-Range Shots",
// 			"Shotguns",
// 			"Sniper Rifles",
// 			"Sporting Rifles"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "palming",
// 		Name:            "Palming",
// 		Description:     "Palming is sleight-of-hand skill that gives a character the ability to snag, hide, and pass off small objects.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryPhysical,
// 		Group:           "stealth",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Legerdemain",
// 			"Pickpocket",
// 			"Pilfering"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "pistols",
// 		Name:            "Pistols",
// 		Description:     "The Pistols skill is for firing handguns. This category includes hold-out pistols, light pistols, heavy pistols, machine pistols, and revolvers.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		Group:           "firearms",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Holdouts",
// 			"Revolvers",
// 			"Semi-Automatics",
// 			"Tasers"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "sneaking",
// 		Name:            "Sneaking",
// 		Description:     "Need to get where you’re not supposed to be? This skill allows you to remain inconspicuous in various situations.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		Group:           "stealth",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Desert",
// 			"Jungle",
// 			"Urban",
// 			"Vehicle",
// 			"Wilderness"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "throwing_weapon",
// 		Name:            "Throwing Weapon",
// 		Description:     "Throwing Weapons is a broad-based attack skill that can be used for any handheld item that is thrown by the user as a weapon.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Aerodynamic",
// 			"Blades",
// 			"Non-Aerodynamic"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "unarmed_combat",
// 		Name:            "Unarmed Combat",
// 		Description:     "Unarmed Combat covers the various self-defense and attack moves that employ the body as a primary weapon. This includes a wide array of martial arts along with the use of cybernetic implant weaponry and the fighting styles that sprung up around those implants.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryCombat,
// 		Group:           "close_combat",
// 		LinkedAttribute: AttributeAgility,
// 		Specializations: []string{
// 			"Blocking",
// 			"Cyber Implants",
// 			"Subduing Combat"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// REACTION
// 	{
// 		ID:              "pilot_aerospace",
// 		Name:            "Pilot Aerospace",
// 		Description:     "Aerospace vehicles include all reduced- and zero-gravity aircraft capable of suborbital or extra-orbital flight.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryVehicle,
// 		LinkedAttribute: AttributeReaction,
// 		Specializations: []string{
// 			"Deep Space",
// 			"Launch Craft",
// 			"Remote Operation",
// 			"Semiballistic",
// 			"Suborbital"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "pilot_aircraft",
// 		Name:            "Pilot Aircraft",
// 		Description:     "This skill is used to pilot any manned or unmanned aircraft operating solely within planetary atmosphere.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryVehicle,
// 		LinkedAttribute: AttributeReaction,
// 		Specializations: []string{
// 			"Fixed-Wing",
// 			"Lighter-Than-Air",
// 			"Remote Operation",
// 			"Rotary Wing",
// 			"Tilt Wing",
// 			"Vectored Thrust"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "pilot_exotic_vehicle",
// 		Name:            "Pilot Exotic Vehicle",
// 		Description:     "Characters must take this skill one time for each specific exotic vehicle. Characters may control the vehicle remotely with this skill where possible.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryVehicle,
// 		Exotic:          true,
// 		LinkedAttribute: AttributeReaction,
// 		Specializations: []string{
// 			"Festo Pigeon 2.0"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "pilot_ground_craft",
// 		Name:            "Pilot Ground Craft",
// 		Description:     "This skill is used to pilot any ground-based vehicle, excluding legged vehicles. This skill applies whether the pilot is in the vehicle or controlling the vehicle via remote access.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryVehicle,
// 		LinkedAttribute: AttributeReaction,
// 		Specializations: []string{
// 			"Bike",
// 			"Hovercraft",
// 			"Remote Operation",
// 			"Tracked",
// 			"Wheeled"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "pilot_walker",
// 		Name:            "Pilot Walker",
// 		Description:     "Any vehicle that walks on two or more legs is piloted through this skill. Characters may control the walker physically or remotely.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryVehicle,
// 		LinkedAttribute: AttributeReaction,
// 		Specializations: []string{
// 			"Biped",
// 			"Multiped",
// 			"Quadruped",
// 			"Remote Operation"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "pilot_watercraft",
// 		Name:            "Pilot Watercraft",
// 		Description:     "This skill is used to pilot any waterborne vehicle, whether from inside it or by remote control.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryVehicle,
// 		LinkedAttribute: AttributeReaction,
// 		Specializations: []string{
// 			"Hydrofoil",
// 			"Motorboat",
// 			"Remote Operation",
// 			"Sail",
// 			"Ship",
// 			"Submarine"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// STRENGTH
// 	{
// 		ID:                     "running",
// 		Name:                   "Running",
// 		Description:            "Running, as you may guess, is about how much ground you can cover quickly.",
// 		Defaultable:            true,
// 		Category:               SkillCategoryPhysical,
// 		Group:                  "athletics",
// 		RequiresGroundMovement: true,
// 		LinkedAttribute:        AttributeStrength,
// 		Specializations: []string{
// 			"Desert",
// 			"Distance",
// 			"Sprinting",
// 			"Urban",
// 			"Wilderness"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:                   "swimming",
// 		Name:                 "Swimming",
// 		Description:          "This skill determines the character’s ability to swim in various bodies of water. The skill level affects the distance and speed at which a character can swim.",
// 		Defaultable:          true,
// 		Category:             SkillCategoryPhysical,
// 		Group:                "athletics",
// 		RequiresSwimMovement: true,
// 		LinkedAttribute:      AttributeStrength,
// 		Specializations: []string{
// 			"Dash",
// 			"Long Distance"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// WILLPOWER
// 	{
// 		ID:              "astral_combat",
// 		Name:            "Astral Combat",
// 		Description:     "Fighting in Astral Space requires the Astral Combat skill. Combat in the Astral World relies on a very different set of abilities and attributes than physical combatants.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		LinkedAttribute: AttributeWillpower,
// 		Specializations: []string{
// 			"Magicians",
// 			"Mana Barriers",
// 			"Spirits"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "survival",
// 		Name:            "Survival",
// 		Description:     "In the desert with nothing more than a tin cup, a poncho, and an iron rod? You’ll need this skill to help you get out alive. Survival is the ability to stay alive in extreme environmental conditions for extended periods of time. The skill governs a character’s ability to perform vital outdoor tasks such as start a fire, build a shelter, scrounge for food, etc. in hostile environments.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		Group:           "outdoors",
// 		LinkedAttribute: AttributeWillpower,
// 		Specializations: []string{
// 			"Desert",
// 			"Forest",
// 			"Jungle",
// 			"Mountain",
// 			"Polar",
// 			"Urban"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// CHARISMA
// 	{
// 		ID:              "animal_handling",
// 		Name:            "Animal Handling",
// 		Description:     "This skill governs the training, care, riding (if they’re big enough), and control of non-sentient animals. Competent trainers have the ability to handle multiple animals. It is even possible to approach an untrained animal and get it to trust you, or at least not eat you.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Herding",
// 			"Riding",
// 			"Training"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "con",
// 		Name:            "Con",
// 		Description:     "Con governs the ability to manipulate or fool an NPC during a social encounter. This skill covers a range of confidence games as well as the principles behind those cons.",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		Group:           "acting",
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Fast Talking",
// 			"Seduction"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "etiquette",
// 		Name:            "Etiquette",
// 		Description:     "Etiquette represents the level of understanding and awareness of proper social rituals. The skill works as a sort of social version of Sneak, allowing you to move unimpeded through various social situations. Etiquette also serves as a social safety net in case a player botches a social situation in a way a skilled character would not.",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		Group:           "influence",
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Catholic Church",
// 			"Corporate",
// 			"High Society",
// 			"Mafia",
// 			"Magic",
// 			"Matrix",
// 			"Media",
// 			"Mercenary",
// 			"Street",
// 			"Yakuza"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "impersonation",
// 		Name:            "Impersonation",
// 		Description:     "Impersonation is the ability to assume the identity of another person, including voice and physical mannerisms. The skill is limited by the physical abilities of the character. A dwarf might be able to impersonate a troll over a commlink, but the illusion shatters when he is face to face with his target.",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		Group:           "acting",
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Dwarf",
// 			"Elf",
// 			"Human",
// 			"Ork",
// 			"Troll"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "instruction",
// 		Name:            "Instruction",
// 		Description:     "Instruction governs the ability to teach people. The skill level helps determine how comfortable the instructor is delivering new material as well as how complex of a skill may be taught.",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{"[Active or Knowledge Skill Category]"},
// 		RuleSource:      RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "intimidation",
// 		Name:            "Intimidation",
// 		Description:     "Intimidation is about creating the impression that you are more menacing than another person in order to get them to do what you want. The skill may be applied multiple ways, from negotiation to interrogation. Intimidation is an Opposed Intimidation + Charisma [Social] Test against the target’s Charisma + Willpower, modified by the appropriate entries on the Social Modifiers Table. ",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Interrogation",
// 			"Mental",
// 			"Physical",
// 			"Torture"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "leadership",
// 		Name:            "Leadership",
// 		Description:     "Leadership is the ability to direct and motivate others. It’s like Con, except rather than using deception you’re using a position of authority. This skill is especially helpful in situations where the will of a teammate is shaken or someone is being asked to do something uncomfortable. The Leadership skill is not meant to replace or make up for poor teamwork. When using Leadership make an opposed test Charisma + Leadership.",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		Group:           "influence",
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Command",
// 			"Direct",
// 			"Inspire",
// 			"Rally"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "negotiation",
// 		Name:            "Negotiation",
// 		Description:     "Negotiation governs a character’s ability to apply their charisma, tactics, and knowledge of situational psychology in order to create a better position when making deals.",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		Group:           "influence",
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Bargaining",
// 			"Contracts",
// 			"Diplomacy"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "performance",
// 		Name:            "Performance",
// 		Description:     "This skill governs the ability to execute a performing art. Performance is to the arts what Artisan is to craft. The performer uses her skill to entertain or even captivate an audience.",
// 		Defaultable:     true,
// 		Category:        SkillCategorySocial,
// 		Group:           "acting",
// 		LinkedAttribute: AttributeCharisma,
// 		Specializations: []string{
// 			"Acting",
// 			"Comedy",
// 			"Dance",
// 			"Presentation"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// LOGIC // TODO: what is this one?
// 	// {
// 	// 	ID:              "academic_knowledge",
// 	// 	Name:            "Academic Knowledge",
// 	// 	Description:     "Academic knowledge is linked to Logic. This type of knowledge includes university subjects such as history, science, design, technology, magical theory, and the people and organizations with fingers in those pies. The humanities (cultures, art, philosophy, and so on) are also included in this category.",
// 	// 	Defaultable:     true,
// 	// 	LinkedAttribute: AttributeLogic,
// 	// 	Specializations: []string{"By trigger (Command, Contact, Time), by spell type (Combat Spells, Detection Spells, etc.)"},
// 	// 	RuleSource:      RuleSourceSR5Core,
// 	// },
// 	// {
// 	// 	ID:              "aeronautics_mechanics",
// 	// 	Name:            "Aeronautics Mechanics",
// 	// 	Description:     "Aeronautics mechanics have the ability to repair a variety of aerospace vehicles, provided the proper tools and parts are available.",
// 	// 	LinkedAttribute: AttributeLogic,
// 	// 	Specializations: []string{"Fixed-Wing", "Rotary-Wing", "VTOL", "Jet Engines", "Propeller Engines"},
// 	// },
// 	{
// 		ID:              "arcana",
// 		Name:            "Arcana",
// 		Description:     "Arcana governs the creation of magical formulae used to create spells, foci, and all other manner of magical manipulations. Arcana is required to understand formulae that may be purchased over the counter or discovered by other means.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryPseudoMagical,
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Focus Design",
// 			"Magical Artifacts",
// 			"Spell Design",
// 			"Spirit Formula",
// 			"Spirit Theory"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "armorer",
// 		Name:            "Armorer",
// 		Description:     "Armorer encompasses the broad array of skills required to build and maintain weapons and armor. As with all mechanics-based skills, the proper tools and equipment are required to perform any repair or build operation.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Armor",
// 			"Artillery",
// 			"Explosives",
// 			"Firearms",
// 			"Melee Weapons",
// 			"Heavy Weapons",
// 			"Weapon Accessories"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "automotive_mechanic",
// 		Name:            "Automotive Mechanic",
// 		Description:     "Automotive mechanics are tasked with fixing all types of ground-based vehicles ranging from commercial automobiles to wheeled drones to tanks. Repairs require the proper tools and time.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "engineering",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Hover",
// 			"Tracked",
// 			"Walker",
// 			"Wheeled"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "biotechnology",
// 		Name:            "Biotechnology",
// 		Description:     "Biotechnology is a wide-ranging skill primarily used by doctors and scientists to grow organic body parts. This skill is the basis for cloning as well as all forms of bioware. Provided the right equipment is available, biotechnology can be used to repair damaged bioware, clone new tissue, or detect any bioware in a subject’s body. This skill does not allow characters to install or remove bioware.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "biotech",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Bioinformatics",
// 			"Bioware",
// 			"Cloning",
// 			"Gene Therapy",
// 			"Vat Maintenance"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "chemistry",
// 		Name:            "Chemistry",
// 		Description:     "Chemistry permits the character to create chemical reactions and develop chemical compounds ranging from drugs, to perfumes, to biopolymers like NuSkin. Chemistry can also be used to analyze chemical compounds to determine what they are.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Analytical",
// 			"Biochemistry",
// 			"Drugs",
// 			"Inorganic",
// 			"Organic",
// 			"Physical",
// 			"Toxins"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "computer",
// 		Name:            "Computer",
// 		Description:     "Computer is the base skill for interacting with the Matrix. It represents the ability to use computers and other Matrix-connected devices. The Computer skill focuses on understanding multiple operating systems. It does not allow the character to exploit code (Hacking) or strip down mainframes (Hardware).",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "electronics",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Edit File",
// 			"Matrix Perception",
// 			"Matrix Search",
// 			"[Computer Matrix Action]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "cybercombat",
// 		Name:            "Cybercombat",
// 		Description:     "Cybercombat is the skill used by hackers to engage in combat on the Matrix.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "cracking",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Devices",
// 			"Grids",
// 			"IC",
// 			"Personas",
// 			"Sprites",
// 			"[Target Type]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "cybertechnology",
// 		Name:            "Cybertechnology",
// 		Description:     "Cybertechnology is the ability to create, maintain, and repair cybernetic parts. A character with the proper tools and parts may repair or even build new cybernetics. Cybertechnology is not a surgical skill. Characters cannot attach or re-attach cybernetics to organic material with this skill. This skill may be used to modify or upgrade cybernetics within cyberlimbs.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "biotech",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Bodyware",
// 			"Cyberlimbs",
// 			"Headware",
// 			"Nanoware",
// 			"Repair"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "demolitions",
// 		Name:            "Demolitions",
// 		Description:     "Demolitions is used to prepare, plant, detonate, and often defuse chemical-based explosives.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Commercial Explosives",
// 			"Defusing",
// 			"Improvised Explosives",
// 			"Plastic Explosives"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "electronic_warfare",
// 		Name:            "Electronic Warfare",
// 		Description:     "Electronic Warfare is the basis of military signals intelligence. It governs the encoding, disruption, spoofing, and decoding of communication systems. Providing the user has the proper equipment, the skill can be used to manipulate or even take over the signal of any item’s communication system.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "cracking",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Communications",
// 			"Encryption",
// 			"Jamming",
// 			"Sensor Operations",
// 			"[Electronic Warfare Matrix Action]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "first_aid",
// 		Name:            "First Aid",
// 		Description:     "First Aid is the ability to provide emergency medical assistance similar to that of a paramedic. This skill may be used to stabilize wounds and prevent characters from dying. First Aid cannot be used to perform surgery or repair damaged implants.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "biotech",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Broken Bones",
// 			"Burns",
// 			"Gunshot Wounds",
// 			"Resuscitation",
// 			"[Type of Treatment]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "forgery",
// 		Name:            "Forgery",
// 		Description:     "Forgery is used to produce counterfeit items or alter existing items to a specific purpose. Depending on the type of forgery, the forger may need specific tools or schematics to complete the task.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Counterfeiting",
// 			"Credstick Forgery",
// 			"False ID",
// 			"Image Doctoring",
// 			"Paper Forgery"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "hacking",
// 		Name:            "Hacking",
// 		Description:     "Hacking is used to discover and exploit security flaws in computers and other electronics.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "cracking",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Devices",
// 			"Files",
// 			"Hosts",
// 			"Personas"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "hardware",
// 		Name:            "Hardware",
// 		Description:     "Hardware reflects a characters ability to build and repair electronic devices. A workspace, proper materials, and sufficient build time are required to enact a repair or to build a new device. See Building & Repairing, at right.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "electronics",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Commlinks",
// 			"Cyberdecks",
// 			"Jack Out",
// 			"Smartguns",
// 			"[Hardware Type]",
// 			"[Hardware Matrix Action]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// {
// 	// 	ID:              "industrial_mechanics",
// 	// 	Name:            "Industrial Mechanics",
// 	// 	Description:     "An industrial mechanic is tasked with repairing or modifying large-scale machines, such as assembly line equipment, power generators, HVAC units, industrial robots, etc. See Building and Repairing, at right.",
// 	// 	LinkedAttribute: AttributeLogic,
// 	// 	Specializations: []string{"Assembly Line", "HVAC", "Industrial Robots", "Power Generators", "Pumps", "Water Treatment"},
// 	// 	RuleSource:      RuleSourceSR5Core,
// 	// },
// 	{
// 		ID:              "medicine",
// 		Name:            "Medicine",
// 		Description:     "Medicine is used to perform advanced medical procedures such as surgeries. It includes long-term medical support for disease and illness, and the skill can be used to diagnose a character’s medical condition. This skill is used to implant or remove cybernetics and bioware but cannot be used to repair or maintain implanted devices.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "biotech",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Cosmetic Surgery",
// 			"Extended Care",
// 			"Implant Surgery",
// 			"Magical Health",
// 			"Organ Culture",
// 			"Trauma Surgery"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// {
// 	// 	ID:              "nautical_mechanics",
// 	// 	Name:            "Nautical Mechanics",
// 	// 	Description:     "Nautical Mechanic is concerned with the maintenance and repair of watercraft. This skill is only effective if the necessary equipment and time are available.",
// 	// 	LinkedAttribute: AttributeLogic,
// 	// 	Specializations: []string{"Boats", "Submarines", "Hovercraft", "Sailboats", "Motorboats"},
// 	// 	RuleSource:      RuleSourceSR5Core,
// 	// },
// 	// {
// 	// 	ID:              "professional_knowledge",
// 	// 	Name:            "Professional Knowledge",
// 	// 	Description:     "Professional Knowledge skills deal with subjects related to normal trades, professions, and occupations, things like journalism, engineering, business, and so on. You might find them helpful when doing legwork for a run, especially those in the corporate world. All Professional Knowledge skills are linked to Logic.",
// 	// 	LinkedAttribute: AttributeLogic,
// 	// 	Specializations: []string{"By profession (Accounting, Architecture, Journalism, Law, etc.)"},
// 	// 	RuleSource:      RuleSourceSR5Core,
// 	// },
// 	{
// 		ID:              "software",
// 		Name:            "Software",
// 		Description:     "Software is the skill used to create and manipulate programming in the Matrix. It’s also what technomancers use when they create their complex forms.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "electronics",
// 		LinkedAttribute: AttributeLogic,
// 		Specializations: []string{
// 			"Data Bombs",
// 			"[Complex Form]",
// 			"[Program Type]",
// 			"[Software Matrix Action]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// INTUITION
// 	{
// 		ID:              "artisan",
// 		Name:            "Artisan",
// 		Description:     "This skill includes several different forms of artistic impression as well as the handcrafting of fine objects that would otherwise be produced on an assembly line. The world’s top artists and crafters are considered artisans.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryTechnical,
// 		LinkedAttribute: AttributeIntuition,
// 		Specializations: []string{
// 			"Carpentry",
// 			"Cooking",
// 			"Drawing",
// 			"Sculpting",
// 			"Fashion",
// 			"Writing",
// 			"[Discipline]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "assensing",
// 		Name:            "Assensing",
// 		Description:     "Assensing is a magic user’s ability to read and interpret fluctuations in the astral world. This skill allows practitioners to learn information by reading astral auras. Only characters capable of astral perception may take this skill.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		LinkedAttribute: AttributeIntuition,
// 		Specializations: []string{
// 			"Aura Reading",
// 			"Astral Signatures",
// 			"Psychometry",
// 			"[Aura Type]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "disguise",
// 		Name:            "Disguise",
// 		Description:     "Disguise covers non-magical forms of masking your identity, including makeup and enhancement.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		Group:           "stealth",
// 		LinkedAttribute: AttributeIntuition,
// 		Specializations: []string{
// 			"Camouflage",
// 			"Cosmetic",
// 			"Theatrical",
// 			"Trideo & Video"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// {
// 	// 	ID:              "intrests_knowledge",
// 	// 	Name:            "Interests Knowledge",
// 	// 	Description:     "Strange as it might sound, you might have some hobbies outside of slinging mana and bullets. Interests are the kind of Knowledge skill that describes what you know because of what you do for fun. There are no guidelines (and no limit) to the sort of interest skills you can have. Interest Knowledge skills are linked to Intuition",
// 	// 	LinkedAttribute: AttributeIntuition,
// 	// 	Specializations: []string{"By interest (Ancient Weapons, Classic Literature, Trid Shows, etc.)"},
// 	// 	RuleSource:      RuleSourceSR5Core,
// 	// },
// 	// {
// 	// 	ID:              "language",
// 	// 	Name:            "Anishinaabe",
// 	// 	Description:     "Language is the ability to converse in a specific language through written and verbal means. Characters who speak multiple languages must purchase a separate language skill for each language.",
// 	//     Defaultable:     false,
// 	//     Category:        SkillCategoryLanguage,
// 	// 	LinkedAttribute: AttributeIntuition,
// 	// 	Specializations: []string{
// 	// 		"Read/Write",
// 	// 		"Speak",
// 	// 		"[Dialect]",
// 	// 		"[Lingo]",
// 	// 		"Street",
// 	// 		"l33tspeak",
// 	// 		"Milspec",
// 	// 		"Corp",
// 	// 		"Orbital"},
// 	// 	RuleSource: RuleSourceSR5Core,
// 	// },
// 	{
// 		ID:              "navigation",
// 		Name:            "Navigation",
// 		Description:     "Navigation governs the use of technology and natural instinct to navigate through territory. This skill enables characters to read maps, use GPS devices, follow AR nav points, or follow a course by landmarks or general direction sense. Navigation applies to both AR and non-AR-enhanced environments.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryTechnical,
// 		Group:           "outdoors",
// 		LinkedAttribute: AttributeIntuition,
// 		Specializations: []string{
// 			"Augmented Reality Markers",
// 			"Celestial",
// 			"Compass",
// 			"Maps",
// 			"GPS"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "perception",
// 		Name:            "Perception",
// 		Description:     "Perception refers to the ability to spot anomalies in everyday situations, making it one of the key skills a shadowrunner needs.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		LinkedAttribute: AttributeIntuition,
// 		Specializations: []string{
// 			"Hearing",
// 			"Numinous Perception",
// 			"Scent",
// 			"Searching",
// 			"Taste",
// 			"Touch",
// 			"Visual"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// {
// 	// 	ID:              "street_knowledge",
// 	// 	Name:            "Street Knowledge",
// 	// 	Description:     "Street Knowledge is linked to Intuition. This type of Knowledge skill is about knowing the movers and shakers in an urban area, along with how things get done on the street. You know about the people who live in different neighborhoods, who to ask to get what, and where things are. The information that these skills cover tends to change rapidly, but your instincts help you keep up.",
// 	// 	LinkedAttribute: AttributeIntuition,
// 	// 	Specializations: []string{"By city (Seattle, Hong Kong, etc.)"},
// 	// 	RuleSource:      RuleSourceSR5Core,
// 	// },
// 	{
// 		ID:              "tracking",
// 		Name:            "Tracking",
// 		Description:     "This skill confers the ability to detect the passage of metahumans and other game through terrain and use those clues to follow that individual. This skill also allows you to identify unmarked trails and common game paths is various environments.",
// 		Defaultable:     true,
// 		Category:        SkillCategoryPhysical,
// 		Group:           "outdoors",
// 		LinkedAttribute: AttributeIntuition,
// 		Specializations: []string{
// 			"Desert",
// 			"Forest",
// 			"Jungle",
// 			"Mountain",
// 			"Polar",
// 			"Urban",
// 			"[Terrain]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// MAGIC
// 	{
// 		ID:              "alchemy",
// 		Name:            "Alchemy",
// 		Description:     "Alchemy is used to create substances that store spells. Alchemy is most commonly used to brew potions, distill magical reagents, and even create orichalcum.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{
// 			"Combat",
// 			"Command",
// 			"Contact",
// 			"Detection",
// 			"Health",
// 			"Illusion",
// 			"Manipulation",
// 			"Time",
// 			"[Trigger]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "artificing",
// 		Name:            "Artificing",
// 		Description:     "Artificing is the process of crafting magical foci. The skill may also be used forensically, in order to assense qualities about an existing focus’ creation and purpose.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		Group:           "enchanting",
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{
// 			"Focus Analysis",
// 			"[Focus Type]"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "banishing",
// 		Name:            "Banishing",
// 		Description:     "Banishing is used to disrupt the link between spirits and the physical world. Banished spirits are forced to return to their native plane and are no longer required to complete unfulfilled services.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		Group:           "conjuring",
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{
// 			"Spirits of Air",
// 			"Spirits of Beasts",
// 			"Spirits of Earth",
// 			"Spirits of Fire",
// 			"Spirits of Man",
// 			"Spirits of Water",
// 			"Guardian Spirit",
// 			"Guidance Spirit",
// 			"Plant Spirit",
// 			"Task Spirit"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "binding",
// 		Name:            "Binding",
// 		Description:     "Binding is used to compel a summoned spirit to perform a number of additional services.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		Group:           "conjuring",
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{
// 			"Spirits of Air",
// 			"Spirits of Beasts",
// 			"Spirits of Earth",
// 			"Spirits of Fire",
// 			"Spirits of Man",
// 			"Spirits of Water",
// 			"Guardian Spirit",
// 			"Guidance Spirit",
// 			"Plant Spirit",
// 			"Task Spirit"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "counterspelling",
// 		Name:            "Counterspelling",
// 		Description:     "Counterspelling is a defensive skill used to defend against magical attacks and dispel sustained magical spells.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		Group:           "sorcery",
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{
// 			"Combat",
// 			"Detection",
// 			"Health",
// 			"Illusion",
// 			"Manipulation"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "disenchanting",
// 		Name:            "Disenchanting",
// 		Description:     "This skill governs a character’s ability to remove the enchantment from an item.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		Group:           "enchanting",
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{"[Type]"},
// 		RuleSource:      RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "ritual_spellcasting",
// 		Name:            "Ritual Spellcasting",
// 		Description:     "Ritual spellcasting is a spellcasting skill used to cast ritual spells.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		Group:           "sorcery",
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{"[Keyword]"},
// 		RuleSource:      RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "spellcasting",
// 		Name:            "Spellcasting",
// 		Description:     "The Spellcasting skill permits the character to channel mana into effects known as spells.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryMagical,
// 		Group:           "sorcery",
// 		LinkedAttribute: AttributeMagic,
// 		Specializations: []string{
// 			"Combat",
// 			"Detection",
// 			"Health",
// 			"Illusion",
// 			"Manipulation"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	{
// 		ID:          "summoning",
// 		Name:        "Summoning",
// 		Description: "This skill is used to summon spirits.",
// 		Defaultable: false,
// 		Category:    SkillCategoryMagical,
// 		Group:       "conjuring",
// 		Specializations: []string{
// 			"Spirits of Air",
// 			"Spirits of Beasts",
// 			"Spirits of Earth",
// 			"Spirits of Fire",
// 			"Spirits of Man",
// 			"Spirits of Water",
// 			"Guardian Spirit",
// 			"Guidance Spirit",
// 			"Plant Spirit",
// 			"Task Spirit"},
// 		RuleSource: RuleSourceSR5Core,
// 	},
// 	// RESONANCE
// 	{
// 		ID:              "compiling",
// 		Name:            "Compiling",
// 		Description:     "Compiling involves the ability to translate the complex 0s and 1s of machine source language and the rhythms of the resonance into sprites.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryResonance,
// 		Group:           "tasking",
// 		LinkedAttribute: AttributeResonance,
// 		Specializations: []string{"[Sprite Type]"},
// 		RuleSource:      RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "decompiling",
// 		Name:            "Decompiling",
// 		Description:     "Decompiling is a character’s ability to effectively delete previously compiled sprites.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryResonance,
// 		Group:           "tasking",
// 		LinkedAttribute: AttributeResonance,
// 		Specializations: []string{"[Sprite Type]"},
// 		RuleSource:      RuleSourceSR5Core,
// 	},
// 	{
// 		ID:              "registering",
// 		Name:            "Registering",
// 		Description:     "This skill allows a technomancer to register sprites on the Matrix, thereby convincing the grids that they are legitimate.",
// 		Defaultable:     false,
// 		Category:        SkillCategoryResonance,
// 		Group:           "tasking",
// 		LinkedAttribute: AttributeResonance,
// 		Specializations: []string{"[Sprite Type]"},
// 		RuleSource:      RuleSourceSR5Core,
// 	},
// }
