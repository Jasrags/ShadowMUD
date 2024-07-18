package core

const (
	MeleeWeaponDataPath      = "data/items/weapons/melee"
	MeleeWeaponFilename      = MeleeWeaponDataPath + "/%s.yaml"
	MeleeWeaponileMinVersion = "0.0.1"
)

// type WeaponType string

// const (
// 	WeaponTypeMelee  WeaponType = "Melee"
// 	WeaponTypeRanged WeaponType = "Ranged"
// )

// type WeaponCategory string

// const (
// 	WeaponCategoryImprovised WeaponCategory = "Improvised"
// 	WeaponCategoryClubs      WeaponCategory = "Clubs"
// 	WeaponCategoryBlades     WeaponCategory = "Blades"
// 	WeaponCategoryExotic     WeaponCategory = "Exotic"
// 	WeaponCategoryMisc       WeaponCategory = "Misc"
// )

// type WeaponSubCategory string

// const (
// 	WeaponSubCategoryPistol        WeaponSubCategory = "Pistol"
// 	WeaponSubCategorySubmachineGun WeaponSubCategory = "Submachine Gun"
// 	WeaponSubCategoryRifle         WeaponSubCategory = "Rifle"
// 	WeaponSubCategoryShotgun       WeaponSubCategory = "Shotgun"
// 	WeaponSubCategoryMachineGun    WeaponSubCategory = "Machine Gun"
// )

// type Weapon struct {
// 	ID             string            `yaml:"id,omitempty"`
// 	Name           string            `yaml:"name"`
// 	Description    string            `yaml:"description"`
// 	Type           WeaponType        `yaml:"type"`
// 	Category       WeaponCategory    `yaml:"category,omitempty"`
// 	SubCategory    WeaponSubCategory `yaml:"sub_category,omitempty"`
// 	Concealability AttributesInfo    `yaml:"concealability"`
// 	// Tags             []WeaponTag       `yaml:"tags"`
// 	Accuracy         AttributesInfo `yaml:"accuracy"`
// 	Reach            int            `yaml:"reach,omitempty"`
// 	DamageValue      int            `yaml:"damage_value,omitempty"`
// 	DamageType       DamageType     `yaml:"damage_type,omitempty"`
// 	ArmorPenatration int            `yaml:"armor_penatration,omitempty"`
// 	Availability     int            `yaml:"availability,omitempty"`
// 	Legality         LegalityType   `yaml:"legality,omitempty"`
// 	Cost             int            `yaml:"cost,omitempty"`
// 	RuleSource       string         `yaml:"rule_source,omitempty"`
// 	FileVersion      string         `yaml:"file_version,omitempty"`
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

// // var testWeapons = []Weapon{
// // 	{
// // 		Name:             "Combat Axe",
// // 		Type:             WeaponTypeMelee,
// // 		Category:         WeaponCategoryBlades,
// // 		Accuracy:         4,
// // 		Reach:            2,
// // 		DamageValue:      5, //(STR+5)P
// // 		DamageType:       DamageTypePhysical,
// // 		ArmorPenatration: -4,
// // 		Availability:     12,
// // 		Legality:         LegalityTypeRestricted,
// // 		Cost:             4000,
// // 		RuleSource:       "SR5:Core",
// // 		Tags:             []WeaponTag{WeaponTagMelee, WeaponTagBlades, WeaponTagTwoHanded},
// // 	},
// // 	{
// // 		Name:             "Shiawase Arms Blazer",
// // 		Accuracy:         6,
// // 		DamageValue:      7,
// // 		DamageType:       DamageTypePhysical,
// // 		ArmorPenatration: 0,
// // 		// Modes:            []WeaponFiringMode{WeaponFiringModeSemiAutomatic},
// // 		// Recoil:           0,
// // 		// AmmoType:         "Regular",
// // 		// AmmoCapacity:     11,
// // 		Availability: 4,
// // 		Legality:     LegalityTypeRestricted,
// // 		Cost:         320,
// // 		RuleSource:   "SR5:Core",
// // 		Tags:         []WeaponTag{WeaponTagRanged, WeaponTagFirearm, WeaponTagFlamethrower},
// // 	},
// // }

// // type WeaponFiringMode int

// // const (
// // 	WeaponFiringModeSingleShot WeaponFiringMode = iota
// // 	WeaponFiringModeSemiAutomatic
// // 	WeaponFiringModeBurstFire
// // 	WeaponFiringModeLongBurst
// // 	WeaponFiringModeFullAuto
// // 	WeaponFiringModeSuppressiveFire
// // )

var (
	WeaponsMelee = map[string]WeaponMelee{}
)

// func LoadMeleeWeapons(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	logrus.Debug("Started loading melee weapons")

// 	files, errReadDir := os.ReadDir(MeleeWeaponDataPath)
// 	if errReadDir != nil {
// 		logrus.WithError(errReadDir).Fatal("Could not read melee weapons directory")
// 	}

// 	// Create a map to store the metatypes
// 	meleeWeapons := make(map[string]WeaponMelee, len(files))

// 	for _, file := range files {
// 		if strings.HasSuffix(file.Name(), ".yaml") {
// 			filepath := fmt.Sprintf("%s/%s", MeleeWeaponDataPath, file.Name())

// 			var meleeWeapon WeaponMelee
// 			if err := util.LoadStructFromYAML(filepath, &meleeWeapon); err != nil {
// 				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load metatype")
// 			}

// 			meleeWeapons[meleeWeapon.Name] = meleeWeapon
// 		}
// 		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded melee weapon file")
// 	}

// 	logrus.WithFields(logrus.Fields{"count": len(meleeWeapons)}).Info("Done loading melee weapons")

// 	WeaponsMelee = meleeWeapons
// }
