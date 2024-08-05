package common

const (
	QualitiesFilepath = "_data/qualities"

	QualityTypePositive QualityType = "Positive"
	QualityTypeNegative QualityType = "Negative"
)

type (
	QualityType string
	QualitySpec struct {
		ID            string      `yaml:"id,omitempty"`
		Type          QualityType `yaml:"type"`
		Name          string      `yaml:"name"`
		MaxRating     int         `yaml:"max_rating,omitempty"`
		Description   string      `yaml:"description"`
		Prerequisites []string    `yaml:"prerequisites,omitempty"`
		Modifiers     []Modifier  `yaml:"modifiers"`
		Cost          int         `yaml:"cost"`
		RuleSource    RuleSource  `yaml:"rule_source"`
		Hidden        bool        `yaml:"hidden,omitempty"`
	}
	Qualities map[string]*QualitySpec
	Quality   struct {
		ID        string       `yaml:"id,omitempty"`
		Rating    int          `yaml:"rating,omitempty"`
		Modifiers []Modifier   `yaml:"modifiers"`
		Spec      *QualitySpec `yaml:"-"`
	}
)

func NewQuality() *Quality {
	return &Quality{}
}

var CoreQualties = []QualitySpec{
	// Racial Qualities
	{
		ID:          "low_light_vision",
		Type:        QualityTypePositive,
		Name:        "Low-Light Vision",
		Description: "The character can see in dim light as if it were normal light. The character can see twice as far as a normal human in starlight, moonlight, and similar conditions of poor illumination. This quality is common among elves and some other metatypes.",
		RuleSource:  RuleSourceSR5Core,
		Hidden:      true,
	},
	{
		ID:          "thermographic_vision",
		Type:        QualityTypePositive,
		Name:        "Thermographic Vision",
		Description: "The character can see heat sources. The character can see in the infrared spectrum, allowing him to see warm objects in the dark. This quality is common among trolls and some other metatypes.",
		RuleSource:  RuleSourceSR5Core,
		Hidden:      true,
	},
	{
		ID:          "ambidextrous",
		Type:        QualityTypePositive,
		Name:        "Ambidextrous",
		Description: "The Ambidextrous character can handle objects equally well with either hand. Without this quality, any action performed solely with the off–hand (i.e., firing a gun) suffers a –2 dice pool modifier (see Attacker Using Off-Hand Weapon, p. 178).",
		Cost:        4,
		RuleSource:  RuleSourceSR5Core,
		// No -2 DP Modifier for off-hand actions
	},
	{
		ID:          "analytical_mind",
		Type:        QualityTypePositive,
		Name:        "Analytical Mind",
		Description: "Analytical Mind describes the uncanny ability to logically analyze information, deduce solutions to problems, or separate vital information from distractions and noise. It’s useful in cracking cyphers, solving puzzles, figuring out traps, and sifting through data.",
		Cost:        5,
		RuleSource:  RuleSourceSR5Core,
		// This quality gives the character a +2 dice pool modifier to any Logic Tests involving pattern recognition, evidence analysis, clue hunting, or solving puzzles.
		// This quality also reduces the time it takes the character to solve a problem by half.
	},
	{
		ID:          "aptitude",
		Type:        QualityTypePositive,
		Name:        "Aptitude",
		Description: "The standard limit for skills is 12. Every so often, there is a character who can exceed limitations and be truly exceptional in a particular skill.",
		Cost:        14,
		RuleSource:  RuleSourceSR5Core,
		// With this particular quality, the character can have one skill rated at 7 at character creation, and may eventually build that skill up to rating 13.
		// Characters may only take the Aptitude quality once.
	},
	{
		ID:          "astral_chameleon",
		Type:        QualityTypePositive,
		Name:        "Astral Chameleon",
		Description: "The character is difficult to spot on the astral plane. The character’s aura is less distinct, making him harder to spot. The character receives a +2 dice pool modifier to Infiltration Tests made to hide his aura from astral perception.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "bilingual",
		Type:        QualityTypePositive,
		Name:        "Bilingual",
		Description: "The character speaks two languages fluently. This quality can be taken multiple times to represent additional languages.",
		Cost:        5,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "blandness",
		Type:        QualityTypePositive,
		Name:        "Blandness",
		Description: "The character is so average and unremarkable that he tends to blend into the background. The character receives a +2 dice pool modifier to Infiltration Tests made to avoid notice in a crowd.",
		Cost:        8,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "catlike",
		Type:        QualityTypePositive,
		Name:        "Catlike",
		Description: "The character has the grace and agility of a cat. The character receives a +2 dice pool modifier on all Gymnastics Tests.",
		Cost:        7,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "codeslinger",
		Type:        QualityTypePositive,
		Name:        "Codeslinger",
		Description: "The character is a master of a particular computer program. Choose a single program at character creation. The character receives a +2 dice pool modifier on all tests involving that program.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "double_jointed",
		Type:        QualityTypePositive,
		Name:        "Double-jointed",
		Description: "The character is unusually flexible. The character receives a +2 dice pool modifier on all tests involving Escape Artist.",
		Cost:        6,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "exceptional_attribute",
		Type:        QualityTypePositive,
		Name:        "Exceptional Attribute",
		Description: "The character has one attribute that is truly exceptional. The character may raise one attribute by one point above the racial maximum (to a maximum of 7).",
		Cost:        14,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "first_impression",
		Type:        QualityTypePositive,
		Name:        "First Impression",
		Description: "The character knows how to make a good first impression. The character receives a +2 dice pool modifier on all Social Tests when meeting someone for the first time.",
		Cost:        11,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "focused_concentration",
		Type:        QualityTypePositive,
		Name:        "Focused Concentration",
		Description: "The character can concentrate on a task with an intensity that borders on the supernatural. The character can take this quality multiple times, each time applying it to a different skill. The character receives a +1 dice pool modifier on all tests involving that skill. This quality can be taken up to six times for a single skill.",
		MaxRating:   6,
		Cost:        4,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "gearhead",
		Type:        QualityTypePositive,
		Name:        "Gearhead",
		Description: "The character is a natural with machines. The character receives a +2 dice pool modifier on all tests involving a specific type of vehicle or drone.",
		Cost:        11,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "guts",
		Type:        QualityTypePositive,
		Name:        "Guts",
		Description: "The character has a strong will and a strong stomach. The character receives a +2 dice pool modifier on all tests to resist fear and intimidation.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "high_pain_tolerance",
		Type:        QualityTypePositive,
		Name:        "High Pain Tolerance",
		Description: "The character can shrug off pain that would incapacitate others. The character receives a +1 dice pool modifier on all tests to resist the effects of wound modifiers.",
		MaxRating:   3,
		Cost:        7,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "home_ground",
		Type:        QualityTypePositive,
		Name:        "Home Ground",
		Description: "The character knows a particular area like the back of his hand. The character receives a +2 dice pool modifier on all Knowledge and Language skill tests when operating in his home ground.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "human_looking",
		Type:        QualityTypePositive,
		Name:        "Human-looking",
		Description: "The character looks human, even if he isn’t. The character can pass for human in most situations.",
		Cost:        6,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "indomitable",
		Type:        QualityTypePositive,
		Name:        "Indomitable",
		Description: "The character is tough and resilient. The character receives a +1 dice pool modifier on all tests to resist the effects of wound modifiers.",
		MaxRating:   3,
		Cost:        8,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "juryrigger",
		Type:        QualityTypePositive,
		Name:        "Juryrigger",
		Description: "The character can fix things in a pinch. The character receives a +2 dice pool modifier on all tests involving Juryrigging.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "lucky",
		Type:        QualityTypePositive,
		Name:        "Lucky",
		Description: "The character is just plain lucky. The character can reroll any one glitched die roll (not a critical glitch).",
		Cost:        12,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "magic_resistance",
		Type:        QualityTypePositive,
		Name:        "Magic Resistance",
		Description: "The character is naturally resistant to magic. The character receives a +1 dice pool modifier on all tests to resist the effects of spells and critter powers.",
		Cost:        6, // (MAX RATING 4) 	6 KARMA PER RATING
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "mentor_spirit",
		Type:        QualityTypePositive,
		Name:        "Mentor Spirit",
		Description: "The character has a mentor spirit. The character receives a +2 dice pool modifier on all tests involving a skill associated with the mentor spirit.",
		Cost:        5,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "natural_athlete",
		Type:        QualityTypePositive,
		Name:        "Natural Athlete",
		Description: "The character is naturally athletic. The character receives a +1 dice pool modifier on all tests involving Athletics.",
		Cost:        7,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "natural_hardening",
		Type:        QualityTypePositive,
		Name:        "Natural Hardening",
		Description: "The character is naturally tough. The character receives a +1 dice pool modifier on all tests to resist damage.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "natural_immunity",
		Type:        QualityTypePositive,
		Name:        "Natural Immunity",
		Description: "The character is naturally immune to a particular toxin or pathogen. The character receives a +1 dice pool modifier on all tests to resist the effects of that toxin or pathogen.",
		Cost:        4, // OR 10 KARMA
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "photographic_memory",
		Type:        QualityTypePositive,
		Name:        "Photographic Memory",
		Description: "The character has a photographic memory. The character can remember anything he has seen or read with perfect clarity. The character receives a +2 dice pool modifier on all Memory Tests.",
		Cost:        6,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "quick_healer",
		Type:        QualityTypePositive,
		Name:        "Quick Healer",
		Description: "The character heals quickly. The character receives a +2 dice pool modifier on all Healing Tests.",
		Cost:        3,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "resistance_to_pathogens_toxins",
		Type:        QualityTypePositive,
		Name:        "Resistance To Pathogens/Toxins",
		Description: "The character is naturally resistant to a particular toxin or pathogen. The character receives a +1 dice pool modifier on all tests to resist the effects of that toxin or pathogen.",
		Cost:        4, // OR 8 KARMA
		RuleSource:  RuleSourceSR5Core,
		// +2 dice
	},
	{
		ID:          "spirit_affinity",
		Type:        QualityTypePositive,
		Name:        "Spirit Affinity",
		Description: "The character has a natural affinity for spirits. The character receives a +2 dice pool modifier on all tests involving summoning or binding spirits.",
		Cost:        7,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "toughness",
		Type:        QualityTypePositive,
		Name:        "Toughness",
		Description: "The character is tough and resilient. The character receives a +1 dice pool modifier on all tests to resist the effects of wound modifiers.",
		Cost:        9,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "will_to_live",
		Type:        QualityTypePositive,
		Name:        "Will To Live",
		Description: "The character has a strong will to survive. The character can take this quality multiple times, each time applying it to a different condition. The character receives a +1 dice pool modifier on all tests to resist the effects of that condition. This quality can be taken up to three times for a single condition.",
		Cost:        3, // (MAX RATING 3) 	3 KARMA PER RATING
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "addiction",
		Type:        QualityTypeNegative,
		Name:        "Addiction",
		Description: "The character is addicted to a particular substance. The character must consume the substance regularly or suffer withdrawal. The character receives a –2 dice pool modifier on all tests while suffering withdrawal.",
		Cost:        4, // TO 25 KARMA
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "allergy",
		Type:        QualityTypeNegative,
		Name:        "Allergy",
		Description: "The character is allergic to a particular substance. The character suffers a –2 dice pool modifier on all tests while exposed to the substance.",
		Cost:        5, // TO 25 KARMA
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "astral_beacon",
		Type:        QualityTypeNegative,
		Name:        "Astral Beacon",
		Description: "The character is an astral beacon. The character’s astral form is unusually bright and easy to spot. The character receives a –2 dice pool modifier on all tests to hide his aura from astral perception.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "bad_luck",
		Type:        QualityTypeNegative,
		Name:        "Bad Luck",
		Description: "The character is cursed with bad luck. The character receives a –1 dice pool modifier on all tests.",
		Cost:        12,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "bad_rep",
		Type:        QualityTypeNegative,
		Name:        "Bad Rep",
		Description: "The character has a bad reputation. The character receives a –2 dice pool modifier on all Social Tests when dealing with someone who has heard of his bad reputation.",
		Cost:        7,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "code_of_honor",
		Type:        QualityTypeNegative,
		Name:        "Code Of Honor",
		Description: "The character has a personal code of conduct that he follows. The character receives a +2 dice pool modifier on all tests to resist breaking his code of honor.",
		Cost:        15,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "codeblock",
		Type:        QualityTypeNegative,
		Name:        "Codeblock",
		Description: "The character has a mental block that prevents him from performing a particular action. The character receives a –2 dice pool modifier on all tests to perform the action.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "combat_paralysis",
		Type:        QualityTypeNegative,
		Name:        "Combat Paralysis",
		Description: "The character freezes up in combat. The character receives a –2 dice pool modifier on all tests to resist fear and intimidation in combat.",
		Cost:        12,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "dependents",
		Type:        QualityTypeNegative,
		Name:        "Dependents",
		Description: "The character has dependents who rely on him for support. The character must spend time and money caring for his dependents, and must be available to help them when they need him.",
		Cost:        3,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "distinctive_style",
		Type:        QualityTypeNegative,
		Name:        "Distinctive Style",
		Description: "The character has a distinctive style that makes him stand out in a crowd. The character receives a –2 dice pool modifier on all tests to avoid notice in a crowd.",
		Cost:        5,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "elf_poser",
		Type:        QualityTypeNegative,
		Name:        "Elf Poser",
		Description: "The character pretends to be an elf. The character receives a –2 dice pool modifier on all tests to avoid notice as a poser.",
		Cost:        6,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "gremlins",
		Type:        QualityTypeNegative,
		Name:        "Gremlins",
		Description: "The character’s gear is prone to malfunction. The character’s gear is unreliable and prone to malfunction. The gamemaster may call for a Gremlins Test whenever the character uses a piece of gear. The character must roll a number of dice equal to the Gremlins rating of the gear. If any of the dice come up 1, the gear malfunctions. The gamemaster determines the nature of the malfunction.",
		Cost:        4, // (MAX 4) 	4 KARMA PER RATING
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "incompetent",
		Type:        QualityTypeNegative,
		Name:        "Incompetent",
		Description: "The character is incompetent at a particular skill. The character receives a –1 dice pool modifier on all tests involving that skill.",
		Cost:        5,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "insomnia",
		Type:        QualityTypeNegative,
		Name:        "Insomnia",
		Description: "The character has trouble sleeping. The character receives a –2 dice pool modifier on all tests while suffering from insomnia.",
		Cost:        10, // OR 15 KARMA
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "loss_of_confidence",
		Type:        QualityTypeNegative,
		Name:        "Loss Of Confidence",
		Description: "The character has lost confidence in himself. The character receives a –2 dice pool modifier on all tests to resist fear and intimidation.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "low_pain_tolerance",
		Type:        QualityTypeNegative,
		Name:        "Low Pain Tolerance",
		Description: "The character has a low pain tolerance. The character receives a –3 dice pool modifier on all tests to resist the effects of wound modifiers.",
		Cost:        9,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "ork_poser",
		Type:        QualityTypeNegative,
		Name:        "Ork Poser",
		Description: "The character pretends to be an ork. The character receives a –2 dice pool modifier on all tests to avoid notice as a poser.",
		Cost:        6,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "prejudiced",
		Type:        QualityTypeNegative,
		Name:        "Prejudiced",
		Description: "The character is prejudiced against a particular group. The character receives a –1 dice pool modifier on all tests involving that group.",
		Cost:        3, // TO 10 KARMA
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "scorched",
		Type:        QualityTypeNegative,
		Name:        "Scorched",
		Description: "The character has been burned by a powerful magical attack. The character’s essence is permanently reduced by 1.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "sensitive_system",
		Type:        QualityTypeNegative,
		Name:        "Sensitive System",
		Description: "The character’s cyberware is more susceptible to damage. The character’s cyberware Essence cost is increased by 10%.",
		Cost:        12,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "simsense_vertigo",
		Type:        QualityTypeNegative,
		Name:        "Simsense Vertigo",
		Description: "The character is prone to motion sickness. The character receives a –2 dice pool modifier on all tests while suffering from motion sickness.",
		Cost:        5,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "sinner",
		Type:        QualityTypeNegative,
		Name:        "Sinner (layered)",
		Description: "The character has a criminal SIN. The character receives a –5 dice pool modifier on all Social Tests when dealing with someone who knows about his criminal SIN.",
		Cost:        5, // TO 25 KARMA
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "social_stress",
		Type:        QualityTypeNegative,
		Name:        "Social Stress",
		Description: "The character is uncomfortable in social situations. The character receives a –2 dice pool modifier on all Social Tests.",
		Cost:        8,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "spirit_bane",
		Type:        QualityTypeNegative,
		Name:        "Spirit Bane",
		Description: "The character is particularly vulnerable to spirits. The character receives a –2 dice pool modifier on all tests to resist the effects of spirits.",
		Cost:        7,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "uncouth",
		Type:        QualityTypeNegative,
		Name:        "Uncouth",
		Description: "The character is rude and socially awkward. The character receives a –2 dice pool modifier on all Social Tests.",
		Cost:        14,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "uneducated",
		Type:        QualityTypeNegative,
		Name:        "Uneducated",
		Description: "The character is uneducated and lacks knowledge of the world. The character receives a –2 dice pool modifier on all Knowledge skill tests.",
		Cost:        8,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "unsteady_hands",
		Type:        QualityTypeNegative,
		Name:        "Unsteady Hands",
		Description: "The character has unsteady hands. The character receives a –2 dice pool modifier on all tests involving fine motor skills.",
		Cost:        7,
		RuleSource:  RuleSourceSR5Core,
	},
	{
		ID:          "weak_immune_system",
		Type:        QualityTypeNegative,
		Name:        "Weak Immune System",
		Description: "The character has a weak immune system. The character receives a –2 dice pool modifier on all tests to resist the effects of pathogens and toxins.",
		Cost:        10,
		RuleSource:  RuleSourceSR5Core,
	},
}
