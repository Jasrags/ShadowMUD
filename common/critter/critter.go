package critter

import (
	"github.com/Jasrags/ShadowMUD/common/power"
	"github.com/Jasrags/ShadowMUD/common/shared"
	"github.com/Jasrags/ShadowMUD/common/skill"
)

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
	Category                   string
	Attribute[T int | float64] struct {
		Min    T `yaml:"min"`
		Max    T `yaml:"max"`
		AugMax T `yaml:"aug_max"`
	}
	Attributes struct {
		Body       Attribute[int]     `yaml:"body"`
		Agility    Attribute[int]     `yaml:"agility"`
		Reaction   Attribute[int]     `yaml:"reaction"`
		Strength   Attribute[int]     `yaml:"strength"`
		Willpower  Attribute[int]     `yaml:"willpower"`
		Logic      Attribute[int]     `yaml:"logic"`
		Intuition  Attribute[int]     `yaml:"intuition"`
		Charisma   Attribute[int]     `yaml:"charisma"`
		Edge       Attribute[int]     `yaml:"edge"`
		Initiative Attribute[int]     `yaml:"initiative"`
		Essence    Attribute[float64] `yaml:"essence"`
		Magic      Attribute[int]     `yaml:"magic"`
		Resonance  Attribute[int]     `yaml:"resonance"`
	}
	Powers map[string]*Power
	Power  struct {
		ID     string      `yaml:"id"`
		Rating int         `yaml:"rating"`
		Spec   *power.Spec `yaml:"-"`
	}
	Skills map[string]*Skill
	Skill  struct {
		ID             string      `yaml:"id"`
		Specialization string      `yaml:"specialization"`
		Rating         int         `yaml:"rating"`
		Spec           *skill.Spec `yaml:"-"`
	}
	Spec struct {
		ID          string     `yaml:"id"`
		Name        string     `yaml:"name"`
		Category    Category   `yaml:"category"`
		Description string     `yaml:"description"`
		Attributes  Attributes `yaml:"attributes"`
		Powers      Powers     `yaml:"powers"`
		Skills      Skills     `yaml:"skills"`
		// Bonus
		RuleSource shared.RuleSource `yaml:"rule_source"`
	}
)
