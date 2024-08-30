package gear

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	GearFilePath = "_data/items/gear"

	BlackmarketCategoryMagic       = "Magic"
	BlackmarketCategoryWeapons     = "Weapons"
	BlackmarketCategoryArmor       = "Armor"
	BlackmarketCategoryElectronics = "Electronics"
	BlackmarketCategorySoftware    = "Software"
	BlackmarketCategoryDrugs       = "Drugs"

	CategoryAlchemicalTools                  Category = "Alchemical Tools"
	CategoryAmmunition                       Category = "Ammunition"
	CategoryArmorEnhancements                Category = "Armor Enhancements"
	CategoryAudioDevices                     Category = "Audio Devices"
	CategoryAudioEnhancements                Category = "Audio Enhancements"
	CategoryAutosofts                        Category = "Autosofts"
	CategoryBiotech                          Category = "Biotech"
	CategoryBreakingAndEnteringGear          Category = "Breaking and Entering Gear"
	CategoryBTLs                             Category = "BTLs"
	CategoryChemicals                        Category = "Chemicals"
	CategoryCommlinks                        Category = "Commlinks"
	CategoryCommlinkCyberdeckFormFactors     Category = "Commlink/Cyberdeck Form Factors"
	CategoryCommlinkAccessories              Category = "Commlink Accessories"
	CategoryCommlinkApps                     Category = "Commlink Apps"
	CategoryCommonPrograms                   Category = "Common Programs"
	CategoryCommunicationsAndCountermeasures Category = "Communications and Countermeasures"
	CategoryContractsUpkeep                  Category = "Contracts/Upkeep"
	CategoryCritterGear                      Category = "Critter Gear"
	CategoryCurrency                         Category = "Currency"
	CategoryCustom                           Category = "Custom"
	CategoryCustomCyberdeckAttributes        Category = "Custom Cyberdeck Attributes"
	CategoryCustomDrug                       Category = "Custom Drug"
	CategoryCyberdeckModules                 Category = "Cyberdeck Modules"
	CategoryCyberdecks                       Category = "Cyberdecks"
	CategoryCyberterminals                   Category = "Cyberterminals"
	CategoryDisguises                        Category = "Disguises"
	CategoryDrugs                            Category = "Drugs"
	CategoryElectronicsAccessories           Category = "Electronics Accessories"
	CategoryElectronicModification           Category = "Electronic Modification"
	CategoryElectronicParts                  Category = "Electronic Parts"
	CategoryEntertainment                    Category = "Entertainment"
	CategoryExplosives                       Category = "Explosives"
	CategoryExtractionDevices                Category = "Extraction Devices"
	CategoryFoci                             Category = "Foci"
	CategoryFood                             Category = "Food"
	CategoryFormulae                         Category = "Formulae"
	CategoryGrappleGun                       Category = "Grapple Gun"
	CategoryHackingPrograms                  Category = "Hacking Programs"
	CategoryHousewares                       Category = "Housewares"
	CategoryIDCredsticks                     Category = "ID/Credsticks"
	CategoryMagicalCompounds                 Category = "Magical Compounds"
	CategoryMagicalSupplies                  Category = "Magical Supplies"
	CategoryMetatypeSpecific                 Category = "Metatype-Specific"
	CategoryMiscellany                       Category = "Miscellany"
	CategoryMusicalInstruments               Category = "Musical Instruments"
	CategoryNanogear                         Category = "Nanogear"
	CategoryPaydata                          Category = "Paydata"
	CategoryPITac                            Category = "PI-Tac"
	CategoryPrinting                         Category = "Printing"
	CategoryReporterGear                     Category = "Reporter Gear"
	CategoryRFIDTags                         Category = "RFID Tags"
	CategoryRiggerCommandConsoles            Category = "Rigger Command Consoles"
	CategorySecurityDevices                  Category = "Security Devices"
	CategorySensors                          Category = "Sensors"
	CategorySensorFunctions                  Category = "Sensor Functions"
	CategorySensorHousings                   Category = "Sensor Housings"
	CategoryServices                         Category = "Services"
	CategorySkillsofts                       Category = "Skillsofts"
	CategorySoftware                         Category = "Software"
	CategorySoftwareTweaks                   Category = "Software Tweaks"
	CategorySurvivalGear                     Category = "Survival Gear"
	CategoryTailoredPerfumeCologne           Category = "Tailored Perfume/Cologne"
	CategoryTools                            Category = "Tools"
	CategoryToolsOfTheTrade                  Category = "Tools of the Trade"
	CategoryToxins                           Category = "Toxins"
	CategoryVisionDevices                    Category = "Vision Devices"
	CategoryVisionEnhancements               Category = "Vision Enhancements"
	CategoryMatrixAccessories                Category = "Matrix Accessories"
	CategoryBoosterChips                     Category = "Booster Chips"
	CategoryAppearanceModification           Category = "Appearance Modification"
	CategoryDrugGrades                       Category = "Drug Grades"
	// Armor,Bioware,Cyberware,Drugs,Electronics,Geneware,Magic,Nanoware,Software,Vehicles,Weapons
)

type (
	Category            string
	BlackmarketCategory string
	Specs               map[string]*Spec
	Spec                struct {
		ID           string              `yaml:"id"`
		Name         string              `yaml:"name"`
		Category     Category            `yaml:"category"`
		Availability int                 `yaml:"availability"`
		Legality     shared.LegalityType `yaml:"legality"`
		Cost         int                 `yaml:"cost"`
		Modifiers    shared.Modifiers    `yaml:"modifiers"`
		RuleSource   shared.RuleSource   `yaml:"rule_source"`
	}
	// Gears map[string]*Gear
	// Gear  struct {
	// 	ID        string           `yaml:"id"`
	// 	Rating    int              `yaml:"rating"`
	// 	Modifiers shared.Modifiers `yaml:"modifiers"`
	// 	Spec      *Spec            `yaml:"-"`
	// }
)

var CoreGear = []Spec{
	{},
}
