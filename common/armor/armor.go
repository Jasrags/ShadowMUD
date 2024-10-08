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
	Category    string
	ModCategory string
	Specs       map[string]Spec
	Spec        struct {
		ID            string              `yaml:"id"`
		Name          string              `yaml:"name"`
		Hidden        bool                `yaml:"hidden"`
		Description   string              `yaml:"description"`
		Rating        int                 `yaml:"rating"`
		Category      Category            `yaml:"category"`
		Capacity      int                 `yaml:"capacity"`
		Availability  int                 `yaml:"availability"`
		Legality      shared.LegalityType `yaml:"legality"`
		Modifications []ModificationSpec  `yaml:"modifications"`
		Tags          []shared.ItemTag    `yaml:"tags"`
		Modifiers     shared.Modifiers    `yaml:"modifiers"`
		Cost          int                 `yaml:"cost"`
		RuleSource    shared.RuleSource   `yaml:"rule_source"`
	}
)
