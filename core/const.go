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

type WeaponFiringMode string

const (
	WeaponFiringModeSingleShot    WeaponFiringMode = "Single-Shot"
	WeaponFiringModeSemiAutomatic WeaponFiringMode = "Semi-Automatic"
	WeaponFiringModeBurstFire     WeaponFiringMode = "Burst Fire"
	WeaponFiringModeLongBurst     WeaponFiringMode = "Long Burst"
	WeaponFiringModeFullAuto      WeaponFiringMode = "Full Auto"
)

type WeaponType string

const (
	WeaponTypeMelee  WeaponType = "Melee"
	WeaponTypeRanged WeaponType = "Ranged"
)

type WeaponCategory string

const (
	WeaponCategoryImprovised WeaponCategory = "Improvised"
	WeaponCategoryClubs      WeaponCategory = "Clubs"
	WeaponCategoryBlades     WeaponCategory = "Blades"
	WeaponCategoryExotic     WeaponCategory = "Exotic"
	WeaponCategoryMisc       WeaponCategory = "Misc"
)

type WeaponSubCategory string

const (
	WeaponSubCategoryPistol        WeaponSubCategory = "Pistol"
	WeaponSubCategorySubmachineGun WeaponSubCategory = "Submachine Gun"
	WeaponSubCategoryRifle         WeaponSubCategory = "Rifle"
	WeaponSubCategoryShotgun       WeaponSubCategory = "Shotgun"
	WeaponSubCategoryMachineGun    WeaponSubCategory = "Machine Gun"
)

type WeaponTag string
type ItemTag string

const (
	WeaponTagMelee            WeaponTag = "Melee"
	WeaponTagImprovised       WeaponTag = "Improvised"
	WeaponTagClubs            WeaponTag = "Clubs"
	WeaponTagBlades           WeaponTag = "Blades"
	WeaponTagExotic           WeaponTag = "Exotic"
	WeaponTagMisc             WeaponTag = "Misc"
	WeaponTagThrowing         WeaponTag = "Throwing"
	WeaponTagBallistic        WeaponTag = "Ballistic"
	WeaponTagFlamethrower     WeaponTag = "Flamethrower"
	WeaponTagTaser            WeaponTag = "Taser"
	WeaponTagRanged           WeaponTag = "Ranged"
	WeaponTagFirearm          WeaponTag = "Firearm"
	WeaponTagPistol           WeaponTag = "Pistol"
	WeaponTagHoldOutPistol    WeaponTag = "Hold-Out Pistol"
	WeaponTagLightPistol      WeaponTag = "Light Pistol"
	WeaponTagHeavyPistol      WeaponTag = "Heavy Pistol"
	WeaponTagMachinePistol    WeaponTag = "Machine Pistol"
	WeaponTagSubmachineGun    WeaponTag = "Submachine Gun"
	WeaponTagRifle            WeaponTag = "Rifle"
	WeaponTagAssaultRifle     WeaponTag = "Assault Rifle"
	WeaponTagSniperRifle      WeaponTag = "Sniper Rifle"
	WeaponTagSportingRifle    WeaponTag = "Sporting Rifle"
	WeaponTagShotgun          WeaponTag = "Shotgun"
	WeaponTagMachineGun       WeaponTag = "Machine Gun"
	WeaponTagLightMachineGun  WeaponTag = "Light Machine Gun"
	WeaponTagMediumMachineGun WeaponTag = "Medium Machine Gun"
	WeaponTagHeavyMachineGun  WeaponTag = "Heavy Machine Gun"
	WeaponTagExoticFirearm    WeaponTag = "Exotic Firearm"
	WeaponTagLaser            WeaponTag = "Laser"
	WeaponTagLargeCaliber     WeaponTag = "Large-Caliber"
	WeaponTagAssaultCannon    WeaponTag = "Assault Cannon"
	WeaponTagGrenadeLauncher  WeaponTag = "Grenade Launcher"
	WeaponTagMissileLauncher  WeaponTag = "Missile Launcher"
	WeaponTagImplant          WeaponTag = "Implant"
	WeaponTagImplantMelee     WeaponTag = "Implant Melee"
	WeaponTagImplantFirearm   WeaponTag = "Implant Firearm"
	WeaponTagTwoHanded        WeaponTag = "Two-Handed"
)
