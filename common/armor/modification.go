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
	Modifications map[string]*Modification
	Modification  struct {
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

var CoreArmorModifications = []Modification{
	{
		ID:           "chemical_protection",
		Name:         "Chemical Protection",
		Description:  "Water-resistant, non-porous, impermeable materials, and a coating of neutralizing agents protect the wearer against contact-vector chemical attacks.",
		CapacityCost: 1, // [Rating]
		Availability: 6,
		Legality:     shared.LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   "SR5:Core",
		// Add the rating of the Chemical Protection modification to tests made to resist contact-vector toxin attacks.
	},
	{
		ID:           "chemical_seal",
		Name:         "Chemical Seal",
		Description:  "It provides complete protection against contact and inhalation vector chemicals, but can only be used for a total of an hour (the limit of the air supply) at a time.",
		CapacityCost: 6,
		Availability: 12,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         3000,
		RuleSource:   "SR5:Core",
		// Activate Seal
		//  Complex Action
		// Air Supply
		//  1 Hour
		// Wireless
		//  Activating the chemical seal is a Free Action.
		// Available only to full body armor that includes a helmet, the chemical seal is an airtight environmental control that takes a Complex Action to activate
	},
	{
		ID:           "fire_resistance",
		Name:         "Fire Resistance",
		Description:  "Fire-retardant, nonflammable materials protect the wearer against Fire damage.",
		CapacityCost: 6, // [Rating]
		Availability: 6,
		Legality:     shared.LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   "SR5:Core",
		// Add the full rating of the Fire Resistance modification to the Armor value when resisting Fire attacks or checking if the armor catches fire.
	},
	{
		ID:           "insulation",
		Name:         "Insulation",
		Description:  "Thermal fibers and heat-retentive materials protect the wearer against Cold damage.",
		CapacityCost: 6, // [Rating]
		Availability: 6,
		Legality:     shared.LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   "SR5:Core",
		// Add the full rating of the Insulation modification to the Armor value when resisting Cold attacks.
	},
	{
		ID:           "nonconductivity",
		Name:         "Nonconductivity",
		Description:  "Electrical insulation and non-conductive materials protect the wearer against Electricity damage.",
		CapacityCost: 6, // [Rating]
		Availability: 6,
		Legality:     shared.LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   "SR5:Core",
	},
	{
		ID:           "shock_frills",
		Name:         "Shock Frills",
		Description:  "These strips of “fur” are electrically charged when activated, standing on end and inflicting Electricity damage to anyone that touches you.",
		CapacityCost: 2,
		Availability: 6,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         250,
		RuleSource:   "SR5:Core",
		// Unarmed Combat
		//     Attack with Frills
		// Activate Frills
		//     Complex Action
		// Charges
		//     10 Charges
		// Recharge Rate
		//     One Charge Per 10 Seconds
		// Wireless
		//     The shock frills can be activated or deactivated as a Free Action. They can also recharge by induction, recharging one charge per hour.
		// Use Unarmed Combat to attack with the frills. The frills hold 10 charges; when attached to a power point, they recharge one charge per 10 seconds.
	},
	{
		ID:           "thermal_damping",
		Name:         "Thermal Damping",
		Description:  "Designed to reduce your thermal signature, these inner layers capture or bleed heat, so the outer layers maintain a surface temperature equal to the surrounding air.",
		CapacityCost: 6, // [Rating]
		Availability: 10,
		Legality:     shared.LegalityTypeRestricted,
		Cost:         500, // 500¥ * Rating
		RuleSource:   "SR5:Core",
		// Wireless
		//  The suit uses the extra information about your surroundings and also gives you its rating as a dice pool bonus to Sneaking tests against heat-based detection.
		// Add the rating to your limit on any Sneaking test against thermographic vision or thermal sensors.
	},
}

// TODO: Load the data from the yaml files
func LoadArmorModificatons() map[string]Modification {
	data := make(map[string]Modification)
	return data
}
