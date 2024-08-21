package shared

type (
	ElementType        string
	RuleSource         string
	ActionType         string
	LegalityType       string
	EnvironmentType    string
	EnvironmenModifier string
	DamageType         string
	ItemTag            string
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
