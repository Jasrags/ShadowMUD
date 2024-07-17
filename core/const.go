package core

const (
	InitiativeDicePhysical        = 1
	InitiativeDiceAstral          = 2
	InitiativeDiceMatrixAR        = 1
	InitiativeDiceMatrixVRHotSim  = 4
	InitiativeDiceMatrixVRColdSim = 3
	InitiativeDiceRiggerAR        = 1
)

type RuleSource string

const (
	RuleSourceSR5Core RuleSource = "SR5:Core"
	RuleSourceSR5CF   RuleSource = "SR5:ChromeFlesh"
	RuleSourceSR5RG   RuleSource = "SR5:RunAndGun"
	RuleSourceSR5SG   RuleSource = "SR5:StreetGrimoire"
	RuleSourceSR5HT   RuleSource = "SR5:HardTargets"
	RuleSourceSR5R5   RuleSource = "SR5:Rigger5"
	RuleSourceSR5DT   RuleSource = "SR5:DataTrails"
	RuleSourceSR5CA   RuleSource = "SR5:CuttingAces"
	RuleSourceSR5SASS RuleSource = "SR5:SailAwaySweetSister"
	RuleSourceSR5GH3  RuleSource = "SR5:GunH(e)aven3"
	RuleSourceSR5BB   RuleSource = "SR5:BulletsAndBandages"
)

type ActionType string

const (
	ActionFree    ActionType = "Free"
	ActionSimple  ActionType = "Simple"
	ActionComplex ActionType = "Complex"
)

type Attribute string

const (
	AttributeBody      Attribute = "Body"
	AttributeAgility   Attribute = "Agility"
	AttributeReaction  Attribute = "Reaction"
	AttributeStrength  Attribute = "Strength"
	AttributeWillpower Attribute = "Willpower"
	AttributeCharisma  Attribute = "Charisma"
	AttributeLogic     Attribute = "Logic"
	AttributeIntuition Attribute = "Intuition"
	AttributeMagic     Attribute = "Magic"
	AttributeResonance Attribute = "Resonance"
	AttributeEssence   Attribute = "Essence"
)

type LegalityType string

const (
	LegalityTypeLegal      LegalityType = "Legal"
	LegalityTypeRestricted LegalityType = "Restricted"
	LegalityTypeForbidden  LegalityType = "Forbidden"
)

type EnvironmentType string

const (
	EnvironmentTypeUrban   EnvironmentType = "Universal"
	EnvironmentTypeHeat    EnvironmentType = "Heat"
	EnvironmentTypeCold    EnvironmentType = "Cold"
	EnvironmentTypeAquatic EnvironmentType = "Aquatic"
	EnvironmentTypeSpace   EnvironmentType = "Space"
)

type (
	EnvironmenModifier string
)

const (
	EnvironmenModifierClear         EnvironmenModifier = "Clear"
	EnvironmenModifierLightRain     EnvironmenModifier = "LightRain"
	EnvironmenModifierMediumRain    EnvironmenModifier = "MediumRain"
	EnvironmenModifierHeavyRain     EnvironmenModifier = "HeavyRain"
	EnvironmenModifierLightFog      EnvironmenModifier = "LightFog"
	EnvironmenModifierMediumFog     EnvironmenModifier = "MediumFog"
	EnvironmenModifierHeavyFog      EnvironmenModifier = "HeavyFog"
	EnvironmenModifierLightSmoke    EnvironmenModifier = "LightSmoke"
	EnvironmenModifierMediumSmoke   EnvironmenModifier = "MediumSmoke"
	EnvironmenModifierHeavySmoke    EnvironmenModifier = "HeavySmoke"
	EnvironmenModifierFullLight     EnvironmenModifier = "FullLight"
	EnvironmenModifierPartialLight  EnvironmenModifier = "PartialLight"
	EnvironmenModifierDimLight      EnvironmenModifier = "DimLight"
	EnvironmenModifierTotalDarkness EnvironmenModifier = "TotalDarkness"
	EnvironmenModifierNoGlare       EnvironmenModifier = "NoGlare"
	EnvironmenModifierWeakGlare     EnvironmenModifier = "WeakGlare"
	EnvironmenModifierModerateGlare EnvironmenModifier = "ModerateGlare"
	EnvironmenModifierBlidingGlare  EnvironmenModifier = "BlindingGlare"
	EnvironmenModifierNoWind        EnvironmenModifier = "NoWind"
	EnvironmenModifierLightBreeze   EnvironmenModifier = "LightBreeze"
	EnvironmenModifierLightWind     EnvironmenModifier = "LightWind"
	EnvironmenModifierModerateWind  EnvironmenModifier = "ModerateWind"
	EnvironmenModifierStrongWind    EnvironmenModifier = "StrongWind"
	EnvironmenModifierShortRange    EnvironmenModifier = "ShortRange"
	EnvironmenModifierMediumRange   EnvironmenModifier = "MediumRange"
	EnvironmenModifierLongRange     EnvironmenModifier = "LongRange"
	EnvironmenModifierExtremeRange  EnvironmenModifier = "ExtremeRange"
)

func GetEnvironmenModifierValue(modifier EnvironmenModifier) int {
	var modifierValue int

	switch modifier {
	case EnvironmenModifierClear, EnvironmenModifierFullLight,
		EnvironmenModifierNoGlare, EnvironmenModifierNoWind,
		EnvironmenModifierLightBreeze, EnvironmenModifierShortRange:
		modifierValue = 0
	case EnvironmenModifierLightRain, EnvironmenModifierLightFog,
		EnvironmenModifierLightSmoke, EnvironmenModifierPartialLight,
		EnvironmenModifierWeakGlare, EnvironmenModifierLightWind,
		EnvironmenModifierMediumRange:
		modifierValue = -1
	case EnvironmenModifierMediumRain, EnvironmenModifierMediumFog,
		EnvironmenModifierMediumSmoke, EnvironmenModifierDimLight,
		EnvironmenModifierModerateGlare, EnvironmenModifierModerateWind,
		EnvironmenModifierLongRange:
		modifierValue = -2
	case EnvironmenModifierHeavyRain, EnvironmenModifierHeavyFog,
		EnvironmenModifierHeavySmoke, EnvironmenModifierTotalDarkness,
		EnvironmenModifierBlidingGlare, EnvironmenModifierStrongWind,
		EnvironmenModifierExtremeRange:
		modifierValue = -6
		// TODO: Combination of two or more conditions at the –6 level row -10
	}

	return modifierValue
}

type DamageType string

const (
	DamageTypePhysical DamageType = "Physical"
	DamageTypeStun     DamageType = "Stun"
)

type Modifier struct {
	Type   string `yaml:"type"`
	Effect string `yaml:"effect"`
	Value  int    `yaml:"value"`
}
