package common

const (
	ArmorModificationsDataPath      = "data/items/armor/modifications"
	ArmorModificationFilename       = ArmorModificationsDataPath + "/%s.yaml"
	ArmorModificationFileMinVersion = "0.0.1"
)

type ArmorModification struct {
	ID           string       `yaml:"id,omitempty"`
	Name         string       `yaml:"name,omitempty"`
	Description  string       `yaml:"description,omitempty"`
	ArmorRating  int          `yaml:"armor_rating,omitempty"`
	Rating       int          `yaml:"rating,omitempty"`
	Cost         int          `yaml:"cost,omitempty"`
	CapacityCost int          `yaml:"capacity_cost,omitempty"`
	Availability int          `yaml:"availability,omitempty"`
	Legality     LegalityType `yaml:"legality,omitempty"`
	ItemTags     []ItemTag    `yaml:"tags"`
	RuleSource   RuleSource   `yaml:"rule_source,omitempty"`
	FileVersion  string       `yaml:"file_version,omitempty"`
}

var CoreArmorModifications = []ArmorModification{
	{
		ID:           "chemical_protection",
		Name:         "Chemical Protection",
		Description:  "Water-resistant, non-porous, impermeable materials, and a coating of neutralizing agents protect the wearer against contact-vector chemical attacks.",
		CapacityCost: 1, // [Rating]
		Availability: 6,
		Legality:     LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   RuleSourceSR5Core,
		// Add the rating of the Chemical Protection modification to tests made to resist contact-vector toxin attacks.
	},
	{
		ID:           "chemical_seal",
		Name:         "Chemical Seal",
		Description:  "It provides complete protection against contact and inhalation vector chemicals, but can only be used for a total of an hour (the limit of the air supply) at a time.",
		CapacityCost: 6,
		Availability: 12,
		Legality:     LegalityTypeRestricted,
		Cost:         3000,
		RuleSource:   RuleSourceSR5Core,
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
		Legality:     LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   RuleSourceSR5Core,
		// Add the full rating of the Fire Resistance modification to the Armor value when resisting Fire attacks or checking if the armor catches fire.
	},
	{
		ID:           "insulation",
		Name:         "Insulation",
		Description:  "Thermal fibers and heat-retentive materials protect the wearer against Cold damage.",
		CapacityCost: 6, // [Rating]
		Availability: 6,
		Legality:     LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   RuleSourceSR5Core,
		// Add the full rating of the Insulation modification to the Armor value when resisting Cold attacks.
	},
	{
		ID:           "nonconductivity",
		Name:         "Nonconductivity",
		Description:  "Electrical insulation and non-conductive materials protect the wearer against Electricity damage.",
		CapacityCost: 6, // [Rating]
		Availability: 6,
		Legality:     LegalityTypeLegal,
		Cost:         250, // 250¥ * Rating
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "shock_frills",
		Name:         "Shock Frills",
		Description:  "These strips of “fur” are electrically charged when activated, standing on end and inflicting Electricity damage to anyone that touches you.",
		CapacityCost: 2,
		Availability: 6,
		Legality:     LegalityTypeRestricted,
		Cost:         250,
		RuleSource:   RuleSourceSR5Core,
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
		Legality:     LegalityTypeRestricted,
		Cost:         500, // 500¥ * Rating
		RuleSource:   RuleSourceSR5Core,
		// Wireless
		//  The suit uses the extra information about your surroundings and also gives you its rating as a dice pool bonus to Sneaking tests against heat-based detection.
		// Add the rating to your limit on any Sneaking test against thermographic vision or thermal sensors.
	},
}

// TODO: Load the data from the yaml files
func LoadArmorModificatons() map[string]ArmorModification {
	data := make(map[string]ArmorModification)
	return data
}
