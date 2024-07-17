package core

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Jasrags/ShadowMUD/core/util"

	"github.com/sirupsen/logrus"
)

const (
	MeleeWeaponDataPath      = "data/items/weapons/melee"
	MeleeWeaponFilename      = MeleeWeaponDataPath + "/%s.yaml"
	MeleeWeaponileMinVersion = "0.0.1"
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

type Weapon struct {
	ID               string            `yaml:"id,omitempty"`
	Name             string            `yaml:"name"`
	Description      string            `yaml:"description"`
	Type             WeaponType        `yaml:"type"`
	Category         WeaponCategory    `yaml:"category,omitempty"`
	SubCategory      WeaponSubCategory `yaml:"sub_category,omitempty"`
	Concealability   AttributesInfo    `yaml:"concealability"`
	Tags             []WeaponTag       `yaml:"tags"`
	Accuracy         AttributesInfo    `yaml:"accuracy"`
	Reach            int               `yaml:"reach,omitempty"`
	DamageValue      int               `yaml:"damage_value,omitempty"`
	DamageType       DamageType        `yaml:"damage_type,omitempty"`
	ArmorPenatration int               `yaml:"armor_penatration,omitempty"`
	Availability     int               `yaml:"availability,omitempty"`
	Legality         LegalityType      `yaml:"legality,omitempty"`
	Cost             int               `yaml:"cost,omitempty"`
	RuleSource       string            `yaml:"rule_source,omitempty"`
	FileVersion      string            `yaml:"file_version,omitempty"`
}

func (w *Weapon) GetDamageValue() error {
	/*
	   (STR)P
	   (STR+3)P
	   9S(e)
	   6P(fire)
	   6S(e)
	   8P(f)
	   (STR+3)S / 12P
	   (STR+1 / 3)P
	   12P
	   (Rating+2)P
	*/
	return nil
}

// var testWeapons = []Weapon{
// 	{
// 		Name:             "Combat Axe",
// 		Type:             WeaponTypeMelee,
// 		Category:         WeaponCategoryBlades,
// 		Accuracy:         4,
// 		Reach:            2,
// 		DamageValue:      5, //(STR+5)P
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -4,
// 		Availability:     12,
// 		Legality:         LegalityTypeRestricted,
// 		Cost:             4000,
// 		RuleSource:       "SR5:Core",
// 		Tags:             []WeaponTag{WeaponTagMelee, WeaponTagBlades, WeaponTagTwoHanded},
// 	},
// 	{
// 		Name:             "Shiawase Arms Blazer",
// 		Accuracy:         6,
// 		DamageValue:      7,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: 0,
// 		// Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// 		// Recoil:           0,
// 		// AmmoType:         "Regular",
// 		// AmmoCapacity:     11,
// 		Availability: 4,
// 		Legality:     LegalityTypeRestricted,
// 		Cost:         320,
// 		RuleSource:   "SR5:Core",
// 		Tags:         []WeaponTag{WeaponTagRanged, WeaponTagFirearm, WeaponTagFlamethrower},
// 	},
// }

// type WeaponFiringMode int

// const (
// 	WeaponFiringModeSingleShot WeaponFiringMode = iota
// 	WeaponFiringModeSemiAutomatic
// 	WeaponFiringModeBurstFire
// 	WeaponFiringModeLongBurst
// 	WeaponFiringModeFullAuto
// 	WeaponFiringModeSuppressiveFire
// )

var (
	// MeleeWeapons = map[string]WeaponMelee{}
	WeaponsMelee = map[string]WeaponMelee{}
)

func LoadMeleeWeapons(wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Debug("Started loading melee weapons")

	files, errReadDir := os.ReadDir(MeleeWeaponDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read melee weapons directory")
	}

	// Create a map to store the metatypes
	meleeWeapons := make(map[string]WeaponMelee, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", MeleeWeaponDataPath, file.Name())

			var meleeWeapon WeaponMelee
			if err := util.LoadStructFromYAML(filepath, &meleeWeapon); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load metatype")
			}

			meleeWeapons[meleeWeapon.Name] = meleeWeapon
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded melee weapon file")
	}

	logrus.WithFields(logrus.Fields{"count": len(meleeWeapons)}).Info("Done loading melee weapons")

	WeaponsMelee = meleeWeapons
}

// type (
// WeaponMeleeIdx int
// )

// const (
// WeaponMeleeCombatAxe WeaponMeleeIdx = iota
// )

// var WeaponsMelee = map[string]WeaponMelee{
// 	"Combat Axe": {
// 		Name:             "Combat Axe",
// 		Description:      "A two-handed tungsten-alloy monster, available single- or double-bladed. A hardened thrusting point can be concealed, spring-loaded, in the handle. While it may seem barbaric or primitive, the physical augmentations of 2075 can make an old-school weapon like a combat axe even deadlier than a firearm in the right hands. So go ahead, go medieval.",
// 		Accuracy:         4,
// 		Reach:            2,
// 		DamageValue:      5,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -4,
// 		Availability:     "12R",
// 		Cost:             4000,
// 		RuleSource:       "SR5:Core",
// 	},
// 	"Combat Knife": {
// 		Name:             "Combat Knife",
// 		Description:      "A long, KA-BAR-style fighting knife with a blacked-out blade and a chisel point for punching through armor.",
// 		Accuracy:         6,
// 		DamageValue:      2,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -3,
// 		Availability:     "4",
// 		Cost:             300,
// 		RuleSource:       "SR5:Core",
// 	},
// 	"Forearm Snap-Blades": {
// 		Name:             "Forearm Snap-Blades",
// 		Description:      "These are essentially like spurs, only they’re external—painless to install or remove. A forearm sheath conceals three blades that can be extended or retracted via wireless link or muscle movement commands.",
// 		Accuracy:         4,
// 		DamageValue:      2,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -2,
// 		Availability:     "7R",
// 		Cost:             200,
// 		RuleSource:       "SR5:Core",
// 	},
// 	"Katana": {
// 		Name:             "Katana",
// 		Description:      "The legendary two-handed sword of the samurai. The katana has become synonymous not just with street samurai, but with shadowrunners in general, at least in the trids; they’ve got everyone packing a katana, from the decker to the mage. But being a silly cliché doesn’t make this blade any less dangerous, or less helpful in a fight.",
// 		Accuracy:         7,
// 		Reach:            1,
// 		DamageValue:      3,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -3,
// 		Availability:     "9R",
// 		Cost:             1000,
// 		RuleSource:       "SR5:Core",
// 	},
// 	"Knife": {
// 		Name:             "Knife",
// 		Description:      "Your basic, all-purpose street cutter. It comes in a bewildering array of styles, colors, and flavors, but the statistics don’t vary from one to the next. Ordinary knives are used by the poor, the desperate, or the cautious as backup weapons.",
// 		Accuracy:         5,
// 		DamageValue:      1,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -1,
// 		Availability:     "4",
// 		Cost:             10,
// 		RuleSource:       "SR5:Core",
// 	},
// 	"Pole Arm": {
// 		Name:             "Pole Arm",
// 		Description:      "This is, at its very essence, a blade on the end of a very long stick. It usually features an axe-head, glaive, spear point, or similar maiming implement. It’s not easy to handle and it’s just about impossible to conceal, but it’s popular both with trolls and with melee combatants looking to safely keep said large trolls at more than arm’s length.",
// 		Accuracy:         5,
// 		Reach:            3,
// 		DamageValue:      3,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -2,
// 		Availability:     "6R",
// 		Cost:             1000,
// 		RuleSource:       "SR5:Core",
// 	},
// 	"Survival Knife": {
// 		Name:             "Survival Knife",
// 		Description:      "A fine quality blade—smooth on one edge, serrated on the other—with several accessories, including a GPS monitor, mini-multitool, micro-lighter, and a hidden compartment in the handle. The sides of the steel are coated with a non-toxic chemical that blacks out the blade when inert to prevent unwanted reflection, but can be activated to provide two hours of phosphorescent light. All knives can cut flesh, but a survival knife is better at cutting rope and wood, or otherwise being used as a tool. The survival knife is the kind of gadget that no professional should be without.",
// 		Accuracy:         5,
// 		DamageValue:      2,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -1,
// 		Availability:     "4",
// 		Cost:             100,
// 		RuleSource:       "SR5:Core",
// 	},
// 	"Sword": {
// 		Name:             "Sword",
// 		Description:      "It’s sharp, it’s heavy, and it will fragging cut you wide open. Available in a wide variety of styles (wakizashi, seax, scimitar, jian, machete, and so on and so forth), this one-handed blade is not as formidable as a katana but is substantially easier to hide.",
// 		Accuracy:         6,
// 		Reach:            1,
// 		DamageValue:      3,
// 		DamageType:       DamageTypePhysical,
// 		ArmorPenatration: -2,
// 		Availability:     "5R",
// 		Cost:             500,
// 		RuleSource:       "SR5:Core",
// 	},
// }

type (
	AmmoTypeIdx int
)

const (
	AmmoTypeRegular AmmoTypeIdx = iota
	AmmoTypeLightPistolRounds
	AmmoTypeHeavyPistolRounds
)

var WeaponsRanged = map[string]WeaponRanged{
	"Ares Light Fire 70": {
		Name:             "Ares Light Fire 70",
		Description:      "",
		Accuracy:         7,
		DamageValue:      8,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: 0,
		Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
		Recoil:           0,
		AmmoType:         "Regular",
		AmmoCapacity:     16,
		Availability:     "3R",
		Cost:             200,
		RuleSource:       "SR5:Core",
	},
	"Ares Light Fire 75": {
		Name:             "Ares Light Fire 75",
		Description:      "",
		Accuracy:         6,
		DamageValue:      6,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: 0,
		Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
		Recoil:           0,
		AmmoType:         "Regular",
		AmmoCapacity:     16,
		Availability:     "6F",
		Cost:             1250,
		RuleSource:       "SR5:Core",
	},
	"Beretta 201T": {
		Name:             "Beretta 201T",
		Description:      "",
		Accuracy:         6,
		DamageValue:      6,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: 0,
		Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic, WeaponFiringModeBurstFire},
		Recoil:           1,
		AmmoType:         "Regular",
		AmmoCapacity:     21,
		Availability:     "7R",
		Cost:             210,
		RuleSource:       "SR5:Core",
	},
	"Colt America L36": {
		Name:             "Colt America L36",
		Description:      "",
		Accuracy:         7,
		DamageValue:      7,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: 0,
		Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
		Recoil:           0,
		AmmoType:         "Regular",
		AmmoCapacity:     11,
		Availability:     "4R",
		Cost:             320,
		RuleSource:       "SR5:Core",
	},
	"Fichetti Security 600": {
		Name:             "Fichetti Security 600",
		Description:      "",
		Accuracy:         6,
		DamageValue:      7,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: 0,
		Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
		Recoil:           0,
		AmmoType:         "Regular",
		AmmoCapacity:     30,
		Availability:     "6R",
		Cost:             350,
		RuleSource:       "SR5:Core",
	},
	"Taurus Omni-6, light pistol rounds": {
		Name:             "Taurus Omni-6, light pistol rounds",
		Description:      "",
		Accuracy:         5,
		DamageValue:      6,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: 0,
		Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
		Recoil:           0,
		AmmoType:         "Light Pistol Rounds",
		AmmoCapacity:     6,
		Availability:     "3R",
		Cost:             300,
		RuleSource:       "SR5:Core",
	},
	"Taurus Omni-6, heavy pistol rounds": {
		Name:             "Taurus Omni-6, heavy pistol rounds",
		Description:      "",
		Accuracy:         5,
		DamageValue:      7,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -1,
		Modes:            []WeaponFiringMode{WeaponFiringModeSingleShot},
		Recoil:           0,
		AmmoType:         "Heavy Pistol Rounds",
		AmmoCapacity:     6,
		Availability:     "3R",
		Cost:             300,
		RuleSource:       "SR5:Core",
	},
}
