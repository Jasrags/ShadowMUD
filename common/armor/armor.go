package armor

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	CategoryArmor                    Category = "Armor"
	CategoryClothing                 Category = "Clothing"
	CategoryCloaks                   Category = "Cloaks"
	CategoryHighFashionArmorClothing Category = "High-Fashion Armor Clothing"
	CategorySpecialtyArmor           Category = "Specialty Armor"
)

type (
	Specs map[string]*Spec
	Spec  struct {
		ID            string              `yaml:"id"`
		Name          string              `yaml:"name"`
		Hidden        bool                `yaml:"hidden"`
		Description   string              `yaml:"description"`
		Rating        int                 `yaml:"rating"`
		Category      Category            `yaml:"category"`
		Capacity      int                 `yaml:"capacity"`
		Availability  int                 `yaml:"availability"`
		Legality      shared.LegalityType `yaml:"legality"`
		Modifications []Modification      `yaml:"modifications"`
		Tags          []shared.ItemTag    `yaml:"tags"`
		Modifiers     shared.Modifiers    `yaml:"modifiers"`
		Cost          int                 `yaml:"cost"`
		RuleSource    shared.RuleSource   `yaml:"rule_source"`
	}
	Armors map[string]*Armor
	Armor  struct {
		ID            string           `yaml:"id"`
		Rating        int              `yaml:"rating"`
		Modifications []Modification   `yaml:"modifications"`
		Modifiers     shared.Modifiers `yaml:"modifiers"`
		Spec          *Spec            `yaml:"-"`
	}
	Category    string
	ModCategory string
)

var (
	CoreArmor = []Spec{
		{
			ID:          "clothing",
			Name:        "Clothing",
			Description: "Basic clothing.",
			Legality:    shared.LegalityTypeLegal,
			Cost:        20,
			RuleSource:  shared.RuleSourceSR5Core,
			Tags:        []shared.ItemTag{shared.ItemTagClothing},
		},
		{
			ID:           "feedback_clothing",
			Name:         "Feedback Clothing",
			Description:  "This haptic clothing allows for a tactile component to an augmented reality experience. ",
			Availability: 8,
			Cost:         500,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:          "sync_leather",
			Name:        "(Synth)Leather",
			Description: "Synthetic leather.",
			Rating:      4,
			Capacity:    4,
			Cost:        200,
			RuleSource:  shared.RuleSourceSR5Core,
			Tags:        []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:           "actioneer_business_clothes",
			Name:         "Actioneer Business Clothes",
			Description:  "These are the top-of-the-line in business wear, made by Actioneer. They are made of the finest materials and are designed to be stylish and functional. They are available in a variety of styles and colors, and are always in fashion.",
			Rating:       8,
			Capacity:     8,
			Availability: 8,
			Cost:         1500,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:           "armor_clothing",
			Name:         "Armor Clothing",
			Description:  "Lightweight ballistic fiber weave makes these garments almost impossible to detect as armor. It doesn’t provide as much protection as real armor, but it’s available in a wide variety of styles.",
			Rating:       6,
			Capacity:     6,
			Availability: 2,
			Cost:         450,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:           "armor_jacket",
			Name:         "Armor Jacket",
			Description:  "The most popular armor solution on the streets comes in all styles imaginable. It offers good protection without catching too much attention. But don’t think of wearing one to a dinner party.",
			Rating:       12,
			Capacity:     12,
			Availability: 2,
			Cost:         1000,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:           "armor_vest",
			Name:         "Armor Vest",
			Description:  "This modern flexible-wrap vest is designed to be worn under regular clothing without displaying any bulk. A popular and cost-effective option.",
			Rating:       9,
			Capacity:     9,
			Availability: 4,
			Cost:         500,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:           "chameleon_suit",
			Name:         "Chameleon Suit",
			Description:  "This suit is made of a special material that changes color to match the surroundings. It is designed to help the wearer blend in with the environment and avoid detection.",
			Rating:       9,
			Capacity:     9,
			Availability: 10,
			Legality:     shared.LegalityTypeRestricted,
			Cost:         1700,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
			// Add 2 to your limit when you make Sneaking tests to hide. A chameleon suit is also armored for the wearer’s protection.
			// Wireless
			//  The suit uses the extra information about your surroundings and also gives you a +2 dice pool bonus to Sneaking Tests for hiding.
		},
		{
			ID:           "full_body_armor",
			Name:         "Full Body Armor",
			Description:  "This is a full body suit of armor that provides the wearer with the maximum amount of protection. It is designed to protect the wearer from head to toe and is made of the most advanced materials available.",
			Rating:       15,
			Capacity:     15,
			Availability: 14,
			Legality:     shared.LegalityTypeRestricted,
			Cost:         2000,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:           "full_body_armor_helmet",
			Name:         "Full Body Armor, Helmet",
			Description:  "",
			Rating:       3,
			Capacity:     6,
			Availability: 14,
			Legality:     shared.LegalityTypeRestricted,
			Cost:         500,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing, shared.ItemTagHelmet},
			// Full Body Armor, Helmet | +3 | +3 | - | +500¥ | Core |
		},
		{
			ID:           "full_body_armor_chemical_seal",
			Name:         "Full Body Armor, Chemical Seal",
			Availability: 6,
			Cost:         6000,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
			// Full Body Armor, Chemical Seal | - | - | +6 | +6,000¥ | Core |
		},
		{
			ID:           "full_body_armor_environmental_adaptation",
			Name:         "Full Body Armor, Environmental Adaptation",
			Availability: 3,
			Cost:         1000,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
			// Full Body Armor, Environmental Adaptation | - | - | +3 | +1,000¥ | Core |
		},
		{
			ID:           "lined_coat",
			Name:         "Lined Coat",
			Description:  "A lined coat is a long coat that has been lined with a layer of ballistic cloth. It provides good protection against most small arms fire and is a popular choice for shadowrunners who want to blend in with the crowd.",
			Rating:       9,
			Capacity:     9,
			Availability: 4,
			Legality:     shared.LegalityTypeLegal,
			Cost:         900,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing, shared.ItemTagCoat},
		},
		{
			ID:           "urban_explorer_jumpsuit",
			Name:         "Urban Explorer Jumpsuit",
			Description:  "This jumpsuit is designed for urban exploration. It is made of a durable material that provides good protection against the elements and is designed to be comfortable to wear for long periods of time.",
			Rating:       9,
			Capacity:     9,
			Availability: 8,
			Legality:     shared.LegalityTypeLegal,
			Cost:         650,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing},
		},
		{
			ID:          "urban_explorer_jumpsuit_helmet",
			Name:        "Urban Explorer Jumpsuit, Helmet",
			Description: "",
			Rating:      2,
			Capacity:    2,
			Legality:    shared.LegalityTypeLegal,
			Cost:        100,
			RuleSource:  shared.RuleSourceSR5Core,
			Tags:        []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagClothing, shared.ItemTagHelmet},
		},
		{
			ID:           "helmet",
			Name:         "Helmet",
			Description:  "A helmet provides protection for the head.",
			Rating:       2,
			Capacity:     2,
			Availability: 2,
			Legality:     shared.LegalityTypeLegal,
			Cost:         100,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagHelmet},
		},
		{
			ID:           "ballistic_shield",
			Name:         "Ballistic Shield",
			Description:  "A ballistic shield is a portable, hand-held shield that is designed to protect the user from gunfire. It is made of a lightweight material that is resistant to bullets and other projectiles.",
			Rating:       6,
			Capacity:     6,
			Availability: 12,
			Legality:     shared.LegalityTypeRestricted,
			Cost:         1200,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagShield},
		},
		{
			ID:           "riot_shield",
			Name:         "Riot Shield",
			Description:  "A riot shield is a large, hand-held shield that is designed to protect the user from projectiles and other attacks. It is made of a lightweight material that is resistant to bullets and other projectiles.",
			Rating:       6,
			Capacity:     6,
			Availability: 10,
			Legality:     shared.LegalityTypeRestricted,
			Cost:         1000,
			RuleSource:   shared.RuleSourceSR5Core,
			Tags:         []shared.ItemTag{shared.ItemTagArmor, shared.ItemTagShield},
		},
	}
)
