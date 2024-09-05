package weapon

import (
	"github.com/Jasrags/ShadowMUD/common/shared"
)

const (
	WeaponsFilePath = "_data/items/weapons"

	WeaponTypeMelee  Type = "Melee"
	WeaponTypeRanged Type = "Ranged"
	// Melee
	WeaponGroupImprovised Group = "Improvised"
	WeaponGroupClubs      Group = "Clubs"
	WeaponGroupBlades     Group = "Blades"
	WeaponGroupExotic     Group = "Exotic"
	WeaponGroupMisc       Group = "Misc"
	// Ranged
	WeaponGroupThrowingWeapons         Group = "Throwing Weapons"
	WeaponGroupBallisticProjectiles    Group = "Ballistic Projectiles"
	WeaponGroupFlamethrowers           Group = "Flamethrowers"
	WeaponGroupExoticRangedWeapons     Group = "Exotic Ranged Weapons"
	WeaponGroupTasers                  Group = "Tasers"
	WeaponGroupFirearms                Group = "Firearms"
	WeaponGroupLasers                  Group = "Lasers"
	WeaponGroupLargeCaliberProjectiles Group = "Large-Caliber Projectiles"
	WeaponGroupImplantWeapons          Group = "Implant Weapons"
	// Ranged
	WeaponCategoryPistol         Category = "Pistol"
	WeaponCategorySubmachineGun  Category = "Submachine Gun"
	WeaponCategoryRifle          Category = "Rifle"
	WeaponCategoryShotgun        Category = "Shotgun"
	WeaponCategoryMachineGun     Category = "Machine Gun"
	WeaponCategoryExoticFirearms Category = "Exotic Firearms"
	// Melee
	WeaponCategoryClub   Category = "Club"
	WeaponCategoryStaff  Category = "Staff"
	WeaponCategoryKnife  Category = "Knife"
	WeaponCategoryAxe    Category = "Axe"
	WeaponCategorySword  Category = "Sword"
	WeaponCategoryKatana Category = "Katana"
	WeaponCategoryDagger Category = "Dagger"
	// Ranged
	WeaponSubCategoryHoldOutPistol      SubCategory = "Hold-Out Pistol"
	WeaponSubCategoryLightPistol        SubCategory = "Light Pistol"
	WeaponSubCategoryHeavyPistol        SubCategory = "Heavy Pistol"
	WeaponSubCategoryMachinePistol      SubCategory = "Machine Pistol"
	WeaponSubCategoryAssaultRifle       SubCategory = "Assault Rifle"
	WeaponSubCategorySniperRifle        SubCategory = "Sniper Rifle"
	WeaponSubCategorySportingRifle      SubCategory = "Sporting Rifle"
	WeaponSubCategoryLightMachineGun    SubCategory = "Light Machine Gun"
	WeaponSubCategoryMediumMachineGun   SubCategory = "Medium Machine Gun"
	WeaponSubCategoryHeavyMachineGun    SubCategory = "Heavy Machine Gun"
	WeaponSubCategoryAssaultCannon      SubCategory = "Assault Cannon"
	WeaponSubCategoryGrenadeLauncher    SubCategory = "Grenade Launcher"
	WeaponSubCategoryMissileLauncher    SubCategory = "Missile Launcher"
	WeaponSubCategoryImplantMeleeWeapon SubCategory = "Implant Melee Weapon"
	WeaponSubCategoryImplantFirearm     SubCategory = "Implant Firearm"

	WeaponRangedReloadBreakAction        ReloadType = "b"
	WeaponRangedReloadDetachableMagazine ReloadType = "c"
	WeaponRangedReloadDrum               ReloadType = "d"
	WeaponRangedReloadMuzzleLoader       ReloadType = "ml"
	WeaponRangedReloadInternalMagazine   ReloadType = "m"
	WeaponRangedReloadCylinder           ReloadType = "cy"
	WeaponRangedReloadBelt               ReloadType = "belt"

	WeaponFiringModeSingleShot    FiringMode = "Single-Shot"
	WeaponFiringModeSemiAutomatic FiringMode = "Semi-Automatic"
	WeaponFiringModeBurstFire     FiringMode = "Burst Fire"
	WeaponFiringModeLongBurst     FiringMode = "Long Burst"
	WeaponFiringModeFullAuto      FiringMode = "Full Auto"
)

type (
	Type        string
	Group       string
	Category    string
	SubCategory string
	ReloadType  string
	FiringMode  string
	Specs       map[string]*Spec
	Spec        struct {
		ID               string              `yaml:"id"`
		Name             string              `yaml:"name"`
		Description      string              `yaml:"description"`
		Type             Type                `yaml:"type"`
		Group            Group               `yaml:"group"`
		Category         Category            `yaml:"category"`
		SubCategory      SubCategory         `yaml:"sub_category"`
		Concealability   int                 `yaml:"concealability"`
		Accuracy         int                 `yaml:"accuracy"`
		Reach            int                 `yaml:"reach"`
		DamageValue      int                 `yaml:"damage_value"`
		DamageType       shared.DamageType   `yaml:"damage_type"`
		ArmorPenatration int                 `yaml:"armor_penatration"`
		Recoil           int                 `yaml:"recoil"`
		Reload           ReloadType          `yaml:"reload"`
		FiringModes      []FiringMode        `yaml:"firing_modes"`
		AmmoCapacity     int                 `yaml:"ammo_capacity"`
		Availability     int                 `yaml:"availability"`
		Legality         shared.LegalityType `yaml:"legality"`
		Tags             []shared.ItemTag    `yaml:"tags"`
		Modifications    []ModificationSpec  `yaml:"modifications"`
		Modifiers        shared.Modifiers    `yaml:"modifiers"`
		Cost             int                 `yaml:"cost"`
		RuleSource       shared.RuleSource   `yaml:"rule_source"`
	}
)

// func NewWeapon(spec *Spec) *Weapon {
// 	w := &Weapon{
// 		ID:   uuid.New().String(),
// 		Spec: spec,
// 	}
// 	w.log = logrus.WithFields(logrus.Fields{"package": "common", "type": "weapon", "weapon_id": w.ID, "weapon_name": w.Spec.Name})

// 	return w
// }

// func (w *Weapon) GetDamageValue() error {
// 	/*
// 	   (STR)P
// 	   (STR+3)P
// 	   9S(e)
// 	   6P(fire)
// 	   6S(e)
// 	   8P(f)
// 	   (STR+3)S / 12P
// 	   (STR+1 / 3)P
// 	   12P
// 	   (Rating+2)P
// 	*/
// 	return nil
// }
