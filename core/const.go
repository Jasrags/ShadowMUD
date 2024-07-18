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

type (
	ModifierType   string
	ModifierEffect string
)

type Modifier struct {
	Type   string `yaml:"type,omitempty"`
	Effect string `yaml:"effect,omitempty"`
	Value  int    `yaml:"value,omitempty"`
}

type ItemTag string

const (
	ItemTagWeapon                 ItemTag = "Weapon"
	ItemTagMelee                  ItemTag = "Melee"
	ItemTagImprovised             ItemTag = "Improvised"
	ItemTagClub                   ItemTag = "Club"
	ItemTagBlade                  ItemTag = "Blade"
	ItemTagExotic                 ItemTag = "Exotic"
	ItemTagMiscellaneous          ItemTag = "Miscellaneous"
	ItemTagThrowing               ItemTag = "Throwing"
	ItemTagWhip                   ItemTag = "Whip"
	ItemTagMonofiliament          ItemTag = "Monofiliament"
	ItemTagBaton                  ItemTag = "Baton"
	ItemTagKatana                 ItemTag = "Katana"
	ItemTagAxe                    ItemTag = "Axe"
	ItemTagKnife                  ItemTag = "Knife"
	ItemTagSword                  ItemTag = "Sword"
	ItemTagPolearm                ItemTag = "Polearm"
	ItemTagChain                  ItemTag = "Chain"
	ItemTagStaff                  ItemTag = "Staff"
	ItemTagBallistic              ItemTag = "Ballistic"
	ItemTagCrossbow               ItemTag = "Crossbow"
	ItemTagBow                    ItemTag = "Bow"
	ItemTagProjectile             ItemTag = "Projectile"
	ItemTagFlamethrower           ItemTag = "Flamethrower"
	ItemTagTaser                  ItemTag = "Taser"
	ItemTagRanged                 ItemTag = "Ranged"
	ItemTagFirearm                ItemTag = "Firearm"
	ItemTagPistol                 ItemTag = "Pistol"
	ItemTagHoldOutPistol          ItemTag = "Hold-Out Pistol"
	ItemTagLightPistol            ItemTag = "Light Pistol"
	ItemTagHeavyPistol            ItemTag = "Heavy Pistol"
	ItemTagMachinePistol          ItemTag = "Machine Pistol"
	ItemTagSubmachineGun          ItemTag = "Submachine Gun"
	ItemTagRifle                  ItemTag = "Rifle"
	ItemTagAssaultRifle           ItemTag = "Assault Rifle"
	ItemTagSniperRifle            ItemTag = "Sniper Rifle"
	ItemTagSportingRifle          ItemTag = "Sporting Rifle"
	ItemTagShotgun                ItemTag = "Shotgun"
	ItemTagLightMachineGun        ItemTag = "Light Machine Gun"
	ItemTagMediumMachineGun       ItemTag = "Medium Machine Gun"
	ItemTagHeavyMachineGun        ItemTag = "Heavy Machine Gun"
	ItemTagLaser                  ItemTag = "Laser"
	ItemTagLargeCaliberProjectile ItemTag = "Large Caliber Projectile"
	ItemTagGrenadeLauncher        ItemTag = "Grenade Launcher"
	ItemTagMissileLauncher        ItemTag = "Missile Launcher"
	ItemTagImplantMeleeWeapon     ItemTag = "Implant Melee Weapon"
	ItemTagImplantFirearm         ItemTag = "Implant Firearm"
	ItemTagAmmunition             ItemTag = "Ammunition"
	ItemTagBallisticProjectile    ItemTag = "Ballistic Projectile"
	ItemTagGrenade                ItemTag = "Grenade"
	ItemTagRocket                 ItemTag = "Rocket"
	ItemTagMissile                ItemTag = "Missile"
	ItemTagChainsaw               ItemTag = "Chainsaw"
	ItemTagClothing               ItemTag = "Clothing"
	ItemTagArmor                  ItemTag = "Armor"
	ItemTagHelmet                 ItemTag = "Helmet"
	ItemTagCoat                   ItemTag = "Coat"
	ItemTagShield                 ItemTag = "Shield"
)
