package contact

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	ContactsFilepath = "_data/contacts"

	TypeAmerindianTribesperson       Type = "Amerindian Tribesperson"
	TypeAntiquitiesAndOdditiesDealer Type = "Antiquities and Oddities Dealer"
	TypeArmorer                      Type = "Armorer"
	TypeBartender                    Type = "Bartender"
	TypeBeatCop                      Type = "Beat Cop"
	TypeBlogger                      Type = "Blogger"
	TypeBodyguard                    Type = "Bodyguard"
	TypeBookie                       Type = "Bookie"
	TypeCleaner                      Type = "Cleaner"
	TypeClubHopper                   Type = "Club Hopper"
	TypeClubOwner                    Type = "Club Owner"
	TypeCompanyMan                   Type = "Company Man"
	TypeCorporateHeadhunter          Type = "Corporate Headhunter"
	TypeCorporateManager             Type = "Corporate Manager"
	TypeCorporateScientist           Type = "Corporate Scientist"
	TypeCorporateSecretary           Type = "Corporate Secretary"
	TypeCorporateWageSlave           Type = "Corporate Wage Slave"
	TypeCoyote                       Type = "Coyote"
	TypeCultMember                   Type = "Cult Member"
	TypeDockworker                   Type = "Dockworker"
	TypeFence                        Type = "Fence"
	TypeFirefighter                  Type = "Firefighter"
	TypeFixer                        Type = "Fixer"
	TypeFreedomFighter               Type = "Freedom Fighter"
	TypeForensicsExpert              Type = "Forensics Expert"
	TypeGambler                      Type = "Gambler"
	TypeGangLeader                   Type = "Gang Leader"
	TypeGoGanger                     Type = "Go-Ganger"
	TypeGrassrootsPolitician         Type = "Grassroots Politician"
	TypeHateGroupMember              Type = "Hate Group Member"
	TypeHermeticAcademic             Type = "Hermetic Academic"
	TypeHighStakesNegotiator         Type = "High Stakes Negotiator"
	TypeIDManufacturer               Type = "ID Manufacturer"
	TypeInfoBroker                   Type = "Infobroker"
	TypeJanitor                      Type = "Janitor"
	TypeMafiaConsiglieri             Type = "Mafia Consiglieri"
	TypeMechanic                     Type = "Mechanic"
	TypeMercenaryAlchemist           Type = "Mercenary Alchemist"
	TypeMrJohnson                    Type = "Mr. Johnson"
	TypeNomad                        Type = "Nomad"
	TypeOrkNationOrganizer           Type = "Ork Nation Organizer"
	TypeParabotanist                 Type = "Parabotanist"
	TypeParabiologist                Type = "Parabiologist"
	TypeParamed                      Type = "Paramed"
	TypeParamedShaman                Type = "Paramed Shaman"
	TypeParasecurityExpert           Type = "Parasecurity Expert"
	TypePawnBroker                   Type = "Pawn Broker"
	TypePimp                         Type = "Pimp"
	TypePirate                       Type = "Pirate"
	TypePoliceChief                  Type = "Police Chief"
	TypePoliticalIntern              Type = "Political Intern"
	TypeRadical                      Type = "Radical"
	TypeRentACop                     Type = "Rent-a-Cop"
	TypeSharkLawyer                  Type = "Shark Lawyer"
	TypeSimsenseStar                 Type = "Simsense Star"
	TypeSlumlord                     Type = "Slumlord"
	TypeSmuggler                     Type = "Smuggler"
	TypeSnitch                       Type = "Snitch"
	TypeSpider                       Type = "Spider"
	TypeSquatter                     Type = "Squatter"
	TypeStreetDoc                    Type = "Street Doc"
	TypeStreetVendor                 Type = "Street Vendor"
	TypeStripper                     Type = "Stripper"
	TypeSupplySergeant               Type = "Supply Sergeant"
	TypeTalentScout                  Type = "Talent Scout"
	TypeTalislegger                  Type = "Talislegger"
	TypeTalismonger                  Type = "Talismonger"
	TypeTamanousMember               Type = "Tamanous Member"
	TypeTaxiDriver                   Type = "Taxi Driver"
	TypeTerraFirstActivist           Type = "TerraFirst! Activist"
	TypeTridPirate                   Type = "Trid Pirate"
	TypeTrollStreetDealer            Type = "Troll Street Dealer"
	TypeUrbanAnthropologist          Type = "Urban Anthropologist"
	TypeVoryShestiorka               Type = "Vory Shestiorka"
	TypeWizKidGanger                 Type = "Wiz Kid Ganger"

	GenderFemale  Gender = "Female"
	GenderMale    Gender = "Male"
	GenderUnknown Gender = "Unknown"

	AgeYoung      Age = "Young"
	AgeMiddleAged Age = "Middle-Aged"
	AgeOld        Age = "Old"
	AgeUnknown    Age = "Unknown"

	PersonalLifeDivorced   PersonalLife = "Divorced"
	PersonalLifeFamilial   PersonalLife = "Familial Relationship"
	PersonalLifeInRelation PersonalLife = "In Relationship"
	PersonalLifeNone       PersonalLife = "None of Your Damn Business"
	PersonalLifeSingle     PersonalLife = "Single"
	PersonalLifeWidowed    PersonalLife = "Widowed"
	PersonalLifeUnknown    PersonalLife = "Unknown"

	InvestigationLegwork        Investigation = "Legwork"
	InvestigationNetworking     Investigation = "Networking"
	InvestigationSwag           Investigation = "Swag"
	InvestigationShadowServices Investigation = "Shadow Services"
	InvestigationPersonalFavors Investigation = "Personal Favors"
	InvestigationSupport        Investigation = "Support"

	PreferredPaymentBarterEasyToSellItems  PreferredPayment = "Barter (Easy-to-Sell Items)"
	PreferredPaymentBarterHobbyViceItems   PreferredPayment = "Barter (Hobby/Vice Items)"
	PreferredPaymentBarterProfessionItems  PreferredPayment = "Barter (Profession Items)"
	PreferredPaymentCashCorpScrip          PreferredPayment = "Cash (Corp Scrip)"
	PreferredPaymentCashCredstick          PreferredPayment = "Cash (Credstick)"
	PreferredPaymentCashHardCurrency       PreferredPayment = "Cash (Hard Currency)"
	PreferredPaymentCashECC                PreferredPayment = "Cash (ECC)"
	PreferredPaymentServiceDrekJobs        PreferredPayment = "Service (Drek Jobs)"
	PreferredPaymentServiceFreeLaborJobs   PreferredPayment = "Service (Free-Labor Jobs)"
	PreferredPaymentServiceShadowrunnerJob PreferredPayment = "Service (Shadowrunner Job)"

	HobbyViceAnimalsParacritters                  HobbyVice = "Animals (Paracritters)"
	HobbyViceBadHabitDreamChips                   HobbyVice = "Bad Habit (Dream Chips)"
	HobbyViceBadHabitNovacoke                     HobbyVice = "Bad Habit (Novacoke)"
	HobbyViceBadHabitTripChips                    HobbyVice = "Bad Habit (Trip Chips)"
	HobbyViceEntertainmentTridShows               HobbyVice = "Entertainment (Trid Shows)"
	HobbyViceEntertainmentMovies                  HobbyVice = "Entertainment (Movies)"
	HobbyViceEntertainmentMusic                   HobbyVice = "Entertainment (Music)"
	HobbyViceEntertainmentTridShowOddCoven        HobbyVice = "Entertainment (Trid Show 'Odd Coven')"
	HobbyViceEntertainmentTridRealityShows        HobbyVice = "Entertainment (Trid Reality Shows)"
	HobbyViceEntertainmentRPGsARLARPGraphicNovels HobbyVice = "Entertainment (RPGs, ARLARP, Graphic Novels)"
	HobbyViceEntertainmentArtwork                 HobbyVice = "Entertainment (Artwork)"
	HobbyViceEntertainmentActionTrideos           HobbyVice = "Entertainment (Action Trideos)"
	HobbyViceFamilyObligationsBrother             HobbyVice = "Family Obligations (Brother)"
	HobbyViceFamilyObligationsSister              HobbyVice = "Family Obligations (Sister)"
	HobbyViceFamilyObligationsKids                HobbyVice = "Family Obligations (Kids)"
	HobbyViceFamilyObligationsParents             HobbyVice = "Family Obligations (Parents)"
	HobbyViceGamblingCards                        HobbyVice = "Gambling (Cards)"
	HobbyViceGamblingHorses                       HobbyVice = "Gambling (Horses)"
	HobbyViceNothingOfInterest                    HobbyVice = "Nothing of Interest"
	HobbyVicePersonalGroomingClothes              HobbyVice = "Personal Grooming (Clothes)"
	HobbyVicePersonalGroomingFashion              HobbyVice = "Personal Grooming (Fashion)"
	HobbyVicePersonalGroomingShoes                HobbyVice = "Personal Grooming (Shoes)"
	HobbyViceSocialHabitAlcohol                   HobbyVice = "Social Habit (Alcohol)"
	HobbyViceSocialHabitCigarettes                HobbyVice = "Social Habit (Cigarettes)"
	HobbyViceSocialHabitCigars                    HobbyVice = "Social Habit (Cigars)"
	HobbyViceSocialHabitElvenWines                HobbyVice = "Social Habit (Elven Wines)"
	HobbyViceVehiclesCars                         HobbyVice = "Vehicles (Cars)"
	HobbyViceVehiclesDrones                       HobbyVice = "Vehicles (Drones)"
	HobbyViceVehiclesAntiqueCars                  HobbyVice = "Vehicles (Antique Cars)"
	HobbyViceVehiclesSportsCars                   HobbyVice = "Vehicles (Sports Cars)"
	HobbyViceWeaponsGuns                          HobbyVice = "Weapons (Guns)"
	HobbyViceWeaponsBlades                        HobbyVice = "Weapons (Blades)"
	HobbyViceWeaponsMilitary                      HobbyVice = "Weapons (Military)"
)

type (
	Type             string
	Gender           string
	Age              string
	PersonalLife     string
	Investigation    string
	PreferredPayment string
	HobbyVice        string

	Specs map[string]*Spec
	Spec  struct {
		ID                string             `yaml:"id"`
		Name              string             `yaml:"name"`
		Description       string             `yaml:"description"`
		Type              Type               `yaml:"type"`
		Gender            Gender             `yaml:"gender"`
		Age               Age                `yaml:"age"`
		PersonalLife      PersonalLife       `yaml:"personal_life"`
		Investigations    []Investigation    `yaml:"investigations"`
		PreferredPayments []PreferredPayment `yaml:"preferred_payments"`
		HobbyVices        []HobbyVice        `yaml:"hobby_vices"`
		RuleSource        shared.RuleSource  `yaml:"rule_source"`
	}
)

var CoreContacts = []Spec{
	{
		ID:          "brian_flannigan",
		Name:        "Brian Flannigan",
		Description: "A fixer is a person who arranges illicit goods or services for characters. Fixers are the go-to people for characters who need to buy or sell illegal goods, find a buyer for a stolen item, or hire a shadowrunner for a job.",
		Type:        TypeFixer,
		RuleSource:  shared.RuleSourceSR5Core,
	},
}
