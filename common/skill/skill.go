package skill

import (
	"github.com/Jasrags/ShadowMUD/common/shared"
)

const (
	SkillsFilepath = "_data/skills"

	SkillTypeActive    SkillType = "Active"
	SkillTypeKnowledge SkillType = "Knowledge"
	SkillTypeGroup     SkillType = "Group"

	// Skill Categories
	SkillCategoryCombat        Category = "Combat Active"
	SkillCategoryPhysical      Category = "Physical Active"
	SkillCategorySocial        Category = "Social Active"
	SkillCategoryMagical       Category = "Magical Active"
	SkillCategoryPseudoMagical Category = "Pseudo-Magical Active"
	SkillCategoryResonance     Category = "Resonance Active"
	SkillCategoryTechnical     Category = "Technical Active"
	SkillCategoryVehicle       Category = "Vehicle Active"
	// Knowledge Skills
	SkillCategoryAcademic     Category = "Academic"
	SkillCategoryInterest     Category = "Interest"
	SkillCategoryLanguage     Category = "Language"
	SkillCategoryProfessional Category = "Professional"
	SkillCategoryStreet       Category = "Street"
)

type (
	// SkillID       string
	SkillType string
	Category  string
	Specs     map[string]*Spec
	Spec      struct {
		ID                     string               `yaml:"id"`
		Name                   string               `yaml:"name"`
		Hidden                 bool                 `yaml:"hidden"`
		Category               Category             `yaml:"category"`
		Description            string               `yaml:"description"`
		Defaultable            bool                 `yaml:"defaultable"`
		Exotic                 bool                 `yaml:"exotic"`
		RequiresGroundMovement bool                 `yaml:"requires_ground_movement"`
		RequiresSwimMovement   bool                 `yaml:"requires_swim_movement"`
		RequiresFlyMovement    bool                 `yaml:"requires_fly_movement"`
		LinkedAttribute        shared.AttributeType `yaml:"linked_attribute"`
		Group                  string               `yaml:"group"`
		Specializations        []string             `yaml:"specializations"`
		RuleSource             shared.RuleSource    `yaml:"rule_source"`
	}
	// Skills map[string]*Skill
	// Skill  struct {
	// 	ID             string `yaml:"id"`
	// 	Specialization string `yaml:"specialization"`
	// 	Rating         int    `yaml:"rating"`
	// 	// Modifiers              Modifiers  `yaml:"modifiers"`
	// 	Spec *Spec `yaml:"-"`
	// }
)
