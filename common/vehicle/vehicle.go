package vehicle

import (
	"github.com/Jasrags/ShadowMUD/common/shared"
)

const (
	CategoryBikes                 Category = "Bikes"
	CategoryCars                  Category = "Cars"
	CategoryTrucks                Category = "Trucks"
	CategoryMunicipalConstruction Category = "Municipal/Construction"
	CategoryCorpsecPoliceMilitary Category = "Corpsec/Police/Military"
	CategoryBoats                 Category = "Boats"
	CategorySubmarines            Category = "Submarines"
	CategoryFixedWingAircraft     Category = "Fixed Wing Aircraft"
	CategoryLTAV                  Category = "LTAV"
	CategoryRotocraft             Category = "Rotocraft"
	CategoryVTOLVSTOL             Category = "VTOL/VSTOL"
	CategoryDronesMicro           Category = "Drones: Micro"
	CategoryDronesMini            Category = "Drones: Mini"
	CategoryDronesSmall           Category = "Drones: Small"
	CategoryDronesMedium          Category = "Drones: Medium"
	CategoryDronesAnthro          Category = "Drones: Anthro"
	CategoryDronesLarge           Category = "Drones: Large"
	CategoryDronesHuge            Category = "Drones: Huge"
	CategoryDronesMissile         Category = "Drones: Missile"

	ModCategoryBody          ModCategory = "Body"
	ModCategoryCosmetic      ModCategory = "Cosmetic"
	ModCategoryElectromatic  ModCategory = "Electromatic"
	ModCategoryModelSpecific ModCategory = "Model-Specific"
	ModCategoryPowertrain    ModCategory = "Powertrain"
	ModCategoryProtection    ModCategory = "Protection"
	ModCategoryWeapons       ModCategory = "Weapons"
)

type (
	Category    string
	ModCategory string
	Gear        struct {
		ID        string `yaml:"id"`
		Name      string `yaml:"name"`
		Rating    int    `yaml:"rating"`
		MaxRating int    `yaml:"max_rating"`
	}
	Mod struct {
		Name string `yaml:"name"`
	}
	Spec struct {
		ID           string              `yaml:"id"`
		Name         string              `yaml:"name"`
		Description  string              `yaml:"description"`
		Accel        int                 `yaml:"accel"`
		Armor        int                 `yaml:"armor"`
		Body         int                 `yaml:"body"`
		Category     Category            `yaml:"category"`
		Handling     int                 `yaml:"handling"`
		Pilot        int                 `yaml:"pilot"`
		Sensor       int                 `yaml:"sensor"`
		Speed        int                 `yaml:"speed"`
		Gears        map[string]*Gear    `yaml:"gears"`
		Mods         map[string]*Mod     `yaml:"mods"`
		Seats        int                 `yaml:"seats"`
		Availability int                 `yaml:"availability"`
		Legality     shared.LegalityType `yaml:"legality"`
		Cost         int                 `yaml:"cost"`
		RuleSource   shared.RuleSource   `yaml:"rule_source"`
	}
	ModSpec struct {
		ID           string              `yaml:"id"`
		Name         string              `yaml:"name"`
		Description  string              `yaml:"description"`
		Category     ModCategory         `yaml:"category"`
		Rating       int                 `yaml:"rating"`
		Slots        int                 `yaml:"slots"`
		Availability int                 `yaml:"availability"`
		Legality     shared.LegalityType `yaml:"legality"`
		Cost         int                 `yaml:"cost"`
		RuleSource   shared.RuleSource   `yaml:"rule_source"`
		//    "required": {
		//       "vehicledetails": {
		//         "OR": {
		//           "name": [
		//             "Ares Paladin",
		//             "CrashCart Medicart (Large)",
		//             "Evo Falcon-EX"
		//           ]
		//         }
		//       }
	}
)
