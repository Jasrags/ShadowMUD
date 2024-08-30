package critter

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	CategoryAI                   Category = "A.I.s"
	CategoryDracoforms           Category = "Dracoforms"
	CategoryEntropicSprites      Category = "Entropic Sprites"
	CategoryFey                  Category = "Fey"
	CategoryGhostsAndHaunts      Category = "Ghosts and Haunts"
	CategoryHarbingers           Category = "Harbingers"
	CategoryImps                 Category = "Imps"
	CategoryInfected             Category = "Infected"
	CategoryInsectSpirits        Category = "Insect Spirits"
	CategoryMundaneCritters      Category = "Mundane Critters"
	CategoryMutantCritters       Category = "Mutant Critters"
	CategoryParanormalCritters   Category = "Paranormal Critters"
	CategoryPrimordialSpirits    Category = "Primordial Spirits"
	CategoryProtosapients        Category = "Protosapients"
	CategoryRitual               Category = "Ritual"
	CategoryShadowSpirits        Category = "Shadow Spirits"
	CategoryShedim               Category = "Shedim"
	CategorySpirits              Category = "Spirits"
	CategorySprites              Category = "Sprites"
	CategoryTechnocritters       Category = "Technocritters"
	CategoryToxicCritters        Category = "Toxic Critters"
	CategoryToxicSpirits         Category = "Toxic Spirits"
	CategoryWarforms             Category = "Warforms"
	CategoryExtraplanarTravelers Category = "Extraplanar Travelers"
	CategoryNecroSpirits         Category = "Necro Spirits"
)

type (
	Category string
	Spec     struct {
		ID          string   `yaml:"id"`
		Name        string   `yaml:"name"`
		Category    Category `yaml:"category"`
		Description string   `yaml:"description"`
		// Bonus
		// Powers
		// Skills
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
)
