package item

type (
	DamageTypeIdx int
)

const (
	DamageTypePhysical DamageTypeIdx = iota
	DamageTypeStun
)

type WeaponMelee struct {
	ID               int
	Name             string
	Description      string
	Accuracy         int
	Reach            int
	DamageValue      int
	DamageType       DamageTypeIdx
	ArmorPenatration int
	Availability     string
	Cost             int
	RuleSource       string
}

type WeaponRanged struct {
	ID               int
	Name             string
	Description      string
	Accuracy         int
	DamageValue      int
	DamageType       DamageTypeIdx
	ArmorPenatration int
	Modes            []WeaponFiringModeIdx
	Recoil           int
	AmmoType         string
	AmmoCapacity     int
	Availability     string
	Cost             int
	RuleSource       string
}

type WeaponFiringModeIdx int

const (
	WeaponFiringModeSingleShot WeaponFiringModeIdx = iota
	WeaponFiringModeSemiAutomatic
	WeaponFiringModeBurstFire
	WeaponFiringModeLongBurst
	WeaponFiringModeFullAuto
	WeaponFiringModeSuppressiveFire
)

// type (
// 	WeaponCategoryIdx    int
// 	WeaponSubcategoryIdx int
// )

const (
// WeaponMelee WeaponCategoryIdx = iota
// WeaponRanged
// WeaponFirearm
// WeaponLargeCaliber
// WeaponMisc
)

// const (
// 	WeaponMeleeBlades WeaponSubcategoryIdx = iota
// 	WeaponMeleeClubs
// 	WeaponMeleeExotic
// 	WeaponMeleeImplant
// 	WeaponMeleeImprovised
// 	WeaponMeleeMisc
// )

type (
	WeaponMeleeIdx int
)

const (
	WeaponMeleeCombatAxe WeaponMeleeIdx = iota
)

var WeaponsMelee = map[string]WeaponMelee{
	"Combat Axe": {
		Name:             "Combat Axe",
		Description:      "A two-handed tungsten-alloy monster, available single- or double-bladed. A hardened thrusting point can be concealed, spring-loaded, in the handle. While it may seem barbaric or primitive, the physical augmentations of 2075 can make an old-school weapon like a combat axe even deadlier than a firearm in the right hands. So go ahead, go medieval.",
		Accuracy:         4,
		Reach:            2,
		DamageValue:      5,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -4,
		Availability:     "12R",
		Cost:             4000,
		RuleSource:       "SR5:Core",
	},
	"Combat Knife": {
		Name:             "Combat Knife",
		Description:      "A long, KA-BAR-style fighting knife with a blacked-out blade and a chisel point for punching through armor.",
		Accuracy:         6,
		DamageValue:      2,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -3,
		Availability:     "4",
		Cost:             300,
		RuleSource:       "SR5:Core",
	},
	"Forearm Snap-Blades": {
		Name:             "Forearm Snap-Blades",
		Description:      "These are essentially like spurs, only they’re external—painless to install or remove. A forearm sheath conceals three blades that can be extended or retracted via wireless link or muscle movement commands.",
		Accuracy:         4,
		DamageValue:      2,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     "7R",
		Cost:             200,
		RuleSource:       "SR5:Core",
	},
	"Katana": {
		Name:             "Katana",
		Description:      "The legendary two-handed sword of the samurai. The katana has become synonymous not just with street samurai, but with shadowrunners in general, at least in the trids; they’ve got everyone packing a katana, from the decker to the mage. But being a silly cliché doesn’t make this blade any less dangerous, or less helpful in a fight.",
		Accuracy:         7,
		Reach:            1,
		DamageValue:      3,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -3,
		Availability:     "9R",
		Cost:             1000,
		RuleSource:       "SR5:Core",
	},
	"Knife": {
		Name:             "Knife",
		Description:      "Your basic, all-purpose street cutter. It comes in a bewildering array of styles, colors, and flavors, but the statistics don’t vary from one to the next. Ordinary knives are used by the poor, the desperate, or the cautious as backup weapons.",
		Accuracy:         5,
		DamageValue:      1,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -1,
		Availability:     "4",
		Cost:             10,
		RuleSource:       "SR5:Core",
	},
	"Pole Arm": {
		Name:             "Pole Arm",
		Description:      "This is, at its very essence, a blade on the end of a very long stick. It usually features an axe-head, glaive, spear point, or similar maiming implement. It’s not easy to handle and it’s just about impossible to conceal, but it’s popular both with trolls and with melee combatants looking to safely keep said large trolls at more than arm’s length.",
		Accuracy:         5,
		Reach:            3,
		DamageValue:      3,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     "6R",
		Cost:             1000,
		RuleSource:       "SR5:Core",
	},
	"Survival Knife": {
		Name:             "Survival Knife",
		Description:      "A fine quality blade—smooth on one edge, serrated on the other—with several accessories, including a GPS monitor, mini-multitool, micro-lighter, and a hidden compartment in the handle. The sides of the steel are coated with a non-toxic chemical that blacks out the blade when inert to prevent unwanted reflection, but can be activated to provide two hours of phosphorescent light. All knives can cut flesh, but a survival knife is better at cutting rope and wood, or otherwise being used as a tool. The survival knife is the kind of gadget that no professional should be without.",
		Accuracy:         5,
		DamageValue:      2,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -1,
		Availability:     "4",
		Cost:             100,
		RuleSource:       "SR5:Core",
	},
	"Sword": {
		Name:             "Sword",
		Description:      "It’s sharp, it’s heavy, and it will fragging cut you wide open. Available in a wide variety of styles (wakizashi, seax, scimitar, jian, machete, and so on and so forth), this one-handed blade is not as formidable as a katana but is substantially easier to hide.",
		Accuracy:         6,
		Reach:            1,
		DamageValue:      3,
		DamageType:       DamageTypePhysical,
		ArmorPenatration: -2,
		Availability:     "5R",
		Cost:             500,
		RuleSource:       "SR5:Core",
	},
}

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
		Modes:            []WeaponFiringModeIdx{WeaponFiringModeSemiAutomatic},
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
		Modes:            []WeaponFiringModeIdx{WeaponFiringModeSemiAutomatic},
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
		Modes:            []WeaponFiringModeIdx{WeaponFiringModeSemiAutomatic, WeaponFiringModeBurstFire},
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
		Modes:            []WeaponFiringModeIdx{WeaponFiringModeSemiAutomatic},
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
		Modes:            []WeaponFiringModeIdx{WeaponFiringModeSemiAutomatic},
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
		Modes:            []WeaponFiringModeIdx{WeaponFiringModeSemiAutomatic},
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
		Modes:            []WeaponFiringModeIdx{WeaponFiringModeSingleShot},
		Recoil:           0,
		AmmoType:         "Heavy Pistol Rounds",
		AmmoCapacity:     6,
		Availability:     "3R",
		Cost:             300,
		RuleSource:       "SR5:Core",
	},
}

//     1 Melee Weapons
//         1.1 Improvised Weapons
//         1.2 Clubs
//         1.3 Blades
//         1.4 Exotic Melee Weapons
//         1.5 Misc. Melee Weapons
//     2 Throwing Weapons
//     3 Ballistic Projectiles
//     4 Flamethrowers
//     5 Exotic Ranged Weapons
//     6 Tasers
//     7 Firearms
//         7.1 Pistols
//             7.1.1 Hold-Out Pistols
//             7.1.2 Light Pistols
//             7.1.3 Heavy Pistols
//             7.1.4 Machine Pistols
//         7.2 Submachine Guns
//         7.3 Rifles
//             7.3.1 Assault Rifles
//             7.3.2 Sniper Rifles
//             7.3.3 Sporting Rifles
//         7.4 Shotguns
//         7.5 Machine Guns
//             7.5.1 Light Machine Guns
//             7.5.2 Medium Machine Guns
//             7.5.3 Heavy Machine Guns
//         7.6 Exotic Firearms
//     8 Lasers
//     9 Large-Caliber Projectiles
//         9.1 Assault Cannons
//         9.2 Grenade Launchers
//         9.3 Missile Launchers
//     10 Implant Weapons
//         10.1 Implant Melee Weapons
//         10.2 Implant Firearms
//     11 Weapon Consumables
//         11.1 Ammunition
//         11.2 Ballistic Projectile Consumables
//         11.3 Grenades
//         11.4 Rockets & Missiles
//     12 Firearm Accessories
