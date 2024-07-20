package core

type ModifierType string

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
)

type ModifierEffect string

const (
	ModifierEffectAdd      ModifierEffect = "Add"
	ModifierEffectSubtract ModifierEffect = "Subtract"
	ModifierEffectMultiply ModifierEffect = "Multiply"
	ModifierEffectDivide   ModifierEffect = "Divide"
	ModifierEffectSet      ModifierEffect = "Set"
)

type Modifier struct {
	Type   ModifierType   `yaml:"type,omitempty"`
	Effect ModifierEffect `yaml:"effect,omitempty"`
	Value  int            `yaml:"value,omitempty"`
}

func (c *Character) ApplyModifiers(m []Modifier) {
	// Apply a modifier to a character
}
