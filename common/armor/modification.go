package armor

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	ModCategoryCustomizedBallisticMask          ModCategory = "Customized Ballistic Mask"
	ModCategoryFullBodyArmorMods                ModCategory = "Full Body Armor Mods"
	ModCategoryGeneral                          ModCategory = "General"
	ModCategoryGlobetrotterClothingLiners       ModCategory = "Globetrotter Clothing Liners"
	ModCategoryGlobetrotterJacketLiners         ModCategory = "Globetrotter Jacket Liners"
	ModCategoryGlobetrotterVestLiners           ModCategory = "Globetrotter Vest Liners"
	ModCategoryNightshadeIR                     ModCategory = "Nightshade IR"
	ModCategoryRapidTransitDetailing            ModCategory = "Rapid Transit Detailing"
	ModCategoryUrbanExplorerJumpsuitAccessories ModCategory = "Urban Explorer Jumpsuit Accessories"
	ModCategoryVictoryLiners                    ModCategory = "Victory Liners"
)

type (
	ModificationSpecs map[string]ModificationSpec
	ModificationSpec  struct {
		ID           string              `yaml:"id"`
		Name         string              `yaml:"name"`
		Description  string              `yaml:"description"`
		ArmorRating  int                 `yaml:"armor_rating"`
		Rating       int                 `yaml:"rating"`
		Cost         int                 `yaml:"cost"`
		CapacityCost int                 `yaml:"capacity_cost"`
		Availability int                 `yaml:"availability"`
		Legality     shared.LegalityType `yaml:"legality"`
		ItemTags     []shared.ItemTag    `yaml:"tags"`
		Modifiers    shared.Modifiers    `yaml:"modifiers"`
		RuleSource   shared.RuleSource   `yaml:"rule_source"`
	}
)
