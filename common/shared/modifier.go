package shared

const (
	ModifierTypeBody             ModifierType = "Body"
	ModifierTypeAgility          ModifierType = "Agility"
	ModifierTypeReaction         ModifierType = "Reaction"
	ModifierTypeStrength         ModifierType = "Strength"
	ModifierTypeWillpower        ModifierType = "Willpower"
	ModifierTypeCharisma         ModifierType = "Charisma"
	ModifierTypeLogic            ModifierType = "Logic"
	ModifierTypeIntuition        ModifierType = "Intuition"
	ModifierTypeMagic            ModifierType = "Magic"
	ModifierTypeInitiative       ModifierType = "Initiative"
	ModifierTypeInitiativeDice   ModifierType = "InitiativeDice"
	ModifierTypeSocialLimit      ModifierType = "SocialLimit"
	ModifierTypePhysicalLimit    ModifierType = "PhysicalLimit"
	ModifierTypeMentalLimit      ModifierType = "MentalLimit"
	ModifierTypeArmorRating      ModifierType = "ArmorRating"
	ModifierTypeArmorValue       ModifierType = "ArmorValue"
	ModifierTypeArmorPenetration ModifierType = "ArmorPenetration"
	ModifierTypeDamageValue      ModifierType = "DamageValue"
	ModifierTypeDamageType       ModifierType = "DamageType"
	ModifierTypeAccuracy         ModifierType = "Accuracy"
	ModifierTypeRange            ModifierType = "Range"
	ModifierTypeCost             ModifierType = "Cost"
	ModifierTypeAvailability     ModifierType = "Availability"
	ModifierTypeElemental        ModifierType = "Elemental"

	ModifierEffectAdd      ModifierEffect = "Add"
	ModifierEffectSubtract ModifierEffect = "Subtract"
	ModifierEffectMultiply ModifierEffect = "Multiply"
	ModifierEffectDivide   ModifierEffect = "Divide"
	ModifierEffectSet      ModifierEffect = "Set"
	ModifierEffectAttack   ModifierEffect = "Attack"
)

type (
	ModifierType   string
	ModifierEffect string
	Modifiers      []Modifier
	Modifier       struct {
		Type   ModifierType   `yaml:"type,omitempty"`
		Effect ModifierEffect `yaml:"effect,omitempty"`
		Value  int            `yaml:"value,omitempty"`
	}
)
