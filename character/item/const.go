package item

type LegalityType string

const (
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
