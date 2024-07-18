package core

const (
	WeaponModificationsDataPath = "data/items/weapons/modifications"
	WeaponModificationFilename  = WeaponModificationsDataPath + "/%s.yaml"
)

type WeaponMountPoint string

const (
	WeaponMountPointUnderBarrel WeaponMountPoint = "Under-Barrel"
	WeaponMountPointBarrel      WeaponMountPoint = "Barrel"
	WeaponMountPointStock       WeaponMountPoint = "Stock"
	WeaponMountPointTop         WeaponMountPoint = "Top"
	WeaponMountPointSide        WeaponMountPoint = "Side"
	WeaponMountPointInternal    WeaponMountPoint = "Internal"
)

type WeaponModificationSpec struct {
	ID          string             `yaml:"id,omitempty"`
	Name        string             `yaml:"name,omitempty"`
	Description string             `yaml:"description,omitempty"`
	MountPoints []WeaponMountPoint `yaml:"mount_points"`
	// ArmorRating  int          `yaml:"armor_rating,omitempty"`
	Cost         int          `yaml:"cost,omitempty"`
	Capacity     int          `yaml:"capacity,omitempty"`
	Availability int          `yaml:"availability,omitempty"`
	Legality     LegalityType `yaml:"legality,omitempty"`
	ItemTags     []ItemTag    `yaml:"tags"`
	Modifiers    []Modifier   `yaml:"modifiers"`
	RuleSource   RuleSource   `yaml:"rule_source,omitempty"`
}

type WeaponModification struct {
	ID        string                 `yaml:"id,omitempty"`
	Rating    int                    `yaml:"rating,omitempty"`
	ItemTags  []ItemTag              `yaml:"tags"`
	Modifiers []Modifier             `yaml:"modifiers"`
	Spec      WeaponModificationSpec `yaml:"_"`
}

var CoreWeaponModifications = []WeaponModificationSpec{
	{
		//Under-Barrel
		ID:           "bipod",
		Name:         "Bipod",
		Description:  "A bipod is a two-legged support that attaches to the barrel of a firearm, allowing the shooter to rest the weapon on the ground or another surface to improve stability.",
		MountPoints:  []WeaponMountPoint{WeaponMountPointUnderBarrel},
		Availability: 2,
		Legality:     LegalityTypeLegal,
		Cost:         200,
		RuleSource:   RuleSourceSR5Core,
		Modifiers: []Modifier{
			{
				Type:   "RecoilCompensation",
				Effect: "Add",
				Value:  2,
			},
		},
		// Attach Bipod
		//     One Minute
		// Fold/Deploy Bipod
		//     Simple Action
		// Remove Bipod
		//     Complex Action
		// Wireless
		//     Folding up or deploying the bipod is a Free Action.
		// A bipod can be attached to the underbarrel mount of a weapon and provides 2 points of Recoil Compensation when properly deployed.
		// Attaching a bipod takes one minute. Folding up or deploying a bipod is a Simple Action. Removing it is a Complex Action.
	},
	{
		ID:           "concealable_holster",
		Name:         "Concealable Holster",
		Description:  "A holster that is designed to be easily concealed under clothing.",
		Availability: 2,
		Legality:     LegalityTypeLegal,
		Cost:         150,
		RuleSource:   RuleSourceSR5Core,
		Modifiers: []Modifier{
			{
				Type:   "Concealability",
				Effect: "Add",
				Value:  -1,
			},
		},
		// Wireless
		//  Wireless sensors and a smart-fabric coated weave allow the holster to alter color and texture in real time adding an additional –1 to the item’s Concealability.
		// The concealable holster adds –1 to the item’s Concealability.
		// Only pistols and tasers fit in a Concealable
	},
	{
		ID:           "gas_vent_system",
		Name:         "Gas Vent System",
		Description:  "Gas-vent recoil compensation systems are barrel-mounted accessories that vent a weapon’s barrel gases at a specific vector to counter muzzle climb. Once installed, a gas-vent cannot be removed.",
		MountPoints:  []WeaponMountPoint{WeaponMountPointBarrel},
		Availability: 4, // (Rating×3)R
		Legality:     LegalityTypeRestricted,
		Cost:         200, // Rating×200¥
		RuleSource:   RuleSourceSR5Core,
		Modifiers: []Modifier{
			{
				Type:   "RecoilCompensation",
				Effect: "Add",
				Value:  2, // Rating
			},
		},
		//  Gas-vent systems provide a number of points of Recoil Compensation equal to their rating.
	},
	{
		ID:           "gyro_mount",
		Name:         "Gyro Mount",
		Description:  "This heavy upper-body harness features an attached, articulated, motorized gyro-stabilized arm that mounts an assault rifle or a heavy weapon.",
		MountPoints:  []WeaponMountPoint{WeaponMountPointUnderBarrel},
		Availability: 7,
		Legality:     LegalityTypeLegal,
		Cost:         1400,
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
		// Attach/Remove Weapon
		//     Simple Action
		// Put On Gyro-Mount
		//     5 Minutes
		// Remove Gyro-Mount
		//     Complex Action
		// Wireless
		//     Activating the harness’s quick-release with a wireless signal to exit the harness is a Free Action.
		//
		// The system neutralizes up to 6 points of recoil and movement modifiers.
		// Attaching or removing a weapon from the mount takes a Simple Action.
		// Putting on a gyro-mount harness takes about five minutes, while the quick-release allows you to get out of it with a Complex Action
	},
	{
		ID:           "hidden_arm_slide",
		Name:         "Hidden Arm Slide",
		Description:  "A hidden arm slide is a concealed holster that allows the wearer to draw a weapon from under a sleeve or pant leg.",
		Availability: 4,
		Legality:     LegalityTypeRestricted,
		Cost:         350,
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
		// Wireless
		//  You can ready the weapon in the slide as a Free Action.
		//If you quick draw the weapon in this slide, the threshold for the quick draw is 2.
		// It also gives the weapon a –1 Concealability modifier.
	},
	{
		ID:           "imaging_scope",
		Name:         "Imaging Scope",
		Description:  "An imaging scope is a telescopic sight that uses a digital camera to provide a magnified image of the target. The image is displayed on a screen in the eyepiece of the scope, allowing the shooter to see the target in low-light conditions or through smoke or fog.",
		Availability: 2,
		Legality:     LegalityTypeLegal,
		Cost:         300,
		MountPoints:  []WeaponMountPoint{WeaponMountPointTop},
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "laser_sight",
		Name:         "Laser Sight",
		Description:  "A laser sight is a small, battery-powered laser that attaches to the barrel of a firearm. The laser projects a red or green dot on the target, allowing the shooter to aim the weapon more accurately.",
		Availability: 2,
		Legality:     LegalityTypeLegal,
		Cost:         125,
		MountPoints:  []WeaponMountPoint{WeaponMountPointTop, WeaponMountPointUnderBarrel},
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "periscope",
		Name:         "Periscope",
		Description:  "A periscope is a device that allows the user to see around corners or over obstacles without exposing themselves to enemy fire. The periscope consists of a tube with mirrors at each end that reflect light from the target to the eyepiece of the periscope.",
		Availability: 3,
		Legality:     LegalityTypeLegal,
		Cost:         70,
		MountPoints:  []WeaponMountPoint{WeaponMountPointTop},
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "quick_draw_holster",
		Name:         "Quick-Draw Holster",
		Description:  "A quick-draw holster is a holster that is designed to allow the wearer to draw a weapon quickly and easily. The holster is usually made of leather or nylon and has a snap or Velcro closure to keep the weapon secure.",
		Availability: 4,
		Legality:     LegalityTypeLegal,
		Cost:         175,
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "shock_pad",
		Name:         "Shock Pad",
		Description:  "A shock pad is a device that attaches to the grip of a firearm and delivers an electric shock to the shooter if the weapon is taken from them. The shock pad is designed to prevent unauthorized use of the weapon and can be activated or deactivated with a switch on the grip.",
		Availability: 2,
		Legality:     LegalityTypeLegal,
		Cost:         50,
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "suppressor",
		Name:         "Suppressor",
		Description:  "A suppressor is a device that attaches to the barrel of a firearm and reduces the noise and muzzle flash produced by the weapon. The suppressor is usually made of metal or plastic and contains baffles that slow and cool the escaping gases from the barrel.",
		Availability: 9,
		Legality:     LegalityTypeForbidden,
		Cost:         500,
		MountPoints:  []WeaponMountPoint{WeaponMountPointBarrel},
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "smart_firing_platform",
		Name:         "Smart Firing Platform",
		Description:  "A smart firing platform is a computerized system that attaches to a firearm and assists the shooter in aiming the weapon. The platform uses sensors and cameras to track the target and adjust the aim of the weapon to compensate for movement and environmental conditions.",
		Availability: 12,
		Legality:     LegalityTypeForbidden,
		Cost:         2500,
		MountPoints:  []WeaponMountPoint{WeaponMountPointUnderBarrel},
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "smartgun_system_external",
		Name:         "Smartgun System, External",
		Description:  "A smartgun system is a computerized aiming system that attaches to a firearm and assists the shooter in aiming the weapon. The system uses sensors and cameras to track the target and adjust the aim of the weapon to compensate for movement and environmental conditions.",
		Availability: 4,
		Legality:     LegalityTypeRestricted,
		Cost:         200,
		MountPoints:  []WeaponMountPoint{WeaponMountPointTop, WeaponMountPointUnderBarrel},
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "smartgun_system_internal",
		Name:         "Smartgun System, Internal",
		Description:  "A smartgun system is a computerized aiming system that is built into a firearm and assists the shooter in aiming the weapon. The system uses sensors and cameras to track the target and adjust the aim of the weapon to compensate for movement and environmental conditions.",
		Availability: 6,
		Legality:     LegalityTypeRestricted,
		Cost:         2, // (2×Weapon Cost)¥
		MountPoints:  []WeaponMountPoint{WeaponMountPointInternal},
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "spare_clip",
		Name:         "Spare Clip",
		Description:  "A spare clip comes unloaded but can hold the maximum rounds for the weapon.",
		Availability: 4,
		Legality:     LegalityTypeLegal,
		Cost:         5,
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "speed_loader",
		Name:         "Speed Loader",
		Description:  "A speed loader is a device that attaches to the cylinder of a revolver and allows the shooter to reload the weapon quickly and easily. The speed loader is usually made of metal or plastic and contains a spring-loaded mechanism that pushes the rounds into the cylinder when the loader is twisted.",
		Availability: 2,
		Legality:     LegalityTypeLegal,
		Cost:         25,
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
	{
		ID:           "tripod",
		Name:         "Tripod",
		Description:  "A tripod is a three-legged support that attaches to the barrel of a firearm, allowing the shooter to rest the weapon on the ground or another surface to improve stability.",
		MountPoints:  []WeaponMountPoint{WeaponMountPointUnderBarrel},
		Availability: 4,
		Legality:     LegalityTypeLegal,
		Cost:         500,
		Modifiers:    []Modifier{},
		RuleSource:   RuleSourceSR5Core,
	},
}
