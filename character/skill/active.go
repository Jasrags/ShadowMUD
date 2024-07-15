package skill

import (
	"fmt"
	"os"
	"shadowrunmud/util"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	ActiveSkillDataPath       = "data/skills/active"
	ActiveSkillFilename       = ActiveSkillDataPath + "/%s.yaml"
	ActiveSkillFileMinVersion = "0.0.1"
)

var (
	ActiveSkills = map[string]ActiveSkill{}
)

type ActiveSkill interface {
	GetID() string
	SetID(string)
	GetName() string
	GetIsDefaultable() bool
	GetDescription() string
	GetLinkedAttribute() AttributeIdx
	GetSkillGroup() string
	GetSpecializations() []string
	GetSelectedSpecialization() string
	SetSelectedSpecialization(string)
	GetRank() int
	SetRank(int)
	GetRuleSource() string
	GetFileVersion() string
}

type activeSkill struct {
	ID                     string       `yaml:"id,omitempty"`
	Name                   string       `yaml:"name"`
	IsDefaultable          bool         `yaml:"is_defaultable"`
	Description            string       `yaml:"description"`
	LinkedAttribute        AttributeIdx `yaml:"linked_attribute"`
	SkillGroup             string       `yaml:"skill_group,omitempty"`
	Specializations        []string     `yaml:"specializations"`
	SelectedSpecialization string       `yaml:"selected_specialization,omitempty"`
	Rank                   int          `yaml:"rank,omitempty"`
	RuleSource             string       `yaml:"rule_source"`
	FileVersion            string       `yaml:"file_version"`
}

func (s activeSkill) GetID() string {
	return s.ID
}

func (s *activeSkill) SetID(id string) {
	s.ID = id
}

func (s activeSkill) GetName() string {
	return s.Name
}

func (s activeSkill) GetIsDefaultable() bool {
	return s.IsDefaultable
}

func (s activeSkill) GetDescription() string {
	return s.Description
}

func (s activeSkill) GetLinkedAttribute() AttributeIdx {
	return s.LinkedAttribute
}

func (s activeSkill) GetSkillGroup() string {
	return s.SkillGroup
}

func (s activeSkill) GetSpecializations() []string {
	return s.Specializations
}

func (s activeSkill) GetSelectedSpecialization() string {
	return s.SelectedSpecialization
}

func (s *activeSkill) SetSelectedSpecialization(spec string) {
	s.SelectedSpecialization = spec
}

func (s activeSkill) GetRank() int {
	return s.Rank
}

func (s *activeSkill) SetRank(rank int) {
	s.Rank = rank
}

func (s activeSkill) GetRuleSource() string {
	return s.RuleSource
}

func (s activeSkill) GetFileVersion() string {
	return s.FileVersion
}

var (
	DefaultActiveSkills = map[string]activeSkill{
		"automatics": {
			Name:            "Automatics",
			IsDefaultable:   true,
			LinkedAttribute: AttributeAgility,
			Specializations: []string{"Assault Rifles", "Cyber-Implant", "Machine Pistols", "Submachine Guns"},
			RuleSource:      "SR5:Core",
			FileVersion:     "0.0.1",
		},
		"archery": {
			Name:            "Archery",
			IsDefaultable:   true,
			LinkedAttribute: AttributeAgility,
			Specializations: []string{"Bow", "Crossbow", "Non-Standard Ammunition", "Slingshot"},
			RuleSource:      "SR5:Core",
			FileVersion:     "0.0.1",
		},
	}
)

type (
	AttributeIdx   string
	ActiveSkillIdx int
	RuleSourceIdx  int
)

const (
	AttributeBody      AttributeIdx = "Body"
	AttributeAgility   AttributeIdx = "Agility"
	AttributeReaction  AttributeIdx = "Reaction"
	AttributeStrength  AttributeIdx = "Strength"
	AttributeWillpower AttributeIdx = "Willpower"
	AttributeCharisma  AttributeIdx = "Charisma"
	AttributeLogic     AttributeIdx = "Logic"
	AttributeIntuition AttributeIdx = "Intuition"
	AttributeMagic     AttributeIdx = "Magic"
	AttributeResonance AttributeIdx = "Resonance"
	AttributeEssence   AttributeIdx = "Essence"
)

const (
	ActiveSkillDriving ActiveSkillIdx = iota
	ActiveSkillFreeFall
	ActiveSkillArchery
	ActiveSkillAutomatics
	ActiveSkillBlades
	ActiveSkillClubs
	ActiveSkillEscapeArtist
	ActiveSkillGunnery
	ActiveSkillHeavyWeapons
	ActiveSkillGymnastics
	ActiveSkillLocksmith
	ActiveSkillLongarms
	ActiveSkillPistols
	ActiveSkillPalming
	ActiveSkillSneaking
	ActiveSkillThrowingWeapon
	ActiveSkillUnarmedCombat
	ActiveSkillExoticMeleeWeapon
	ActiveSkillExoticRangedWeapon
	ActiveSkillPilotAerospace
	ActiveSkillPilotWalker
	ActiveSkillPilotAircraft
	ActiveSkillPilotExoticVehicle
	ActiveSkillPilotGroundCraft
	ActiveSkillPilotWatercraft
	ActiveSkillRunning
	ActiveSkillSwimming
	ActiveSkillAstralCombat
	ActiveSkillSurvival
	ActiveSkillCon
	ActiveSkillInstruction
	ActiveSkillLeadership
	ActiveSkillPerformance
	ActiveSkillAnimalHandling
	ActiveSkillEtiquette
	ActiveSkillIntimidation
	ActiveSkillNegotiation
	ActiveSkillImpersonation
	ActiveSkillAcademicKnowledge
	ActiveSkillArcana
	ActiveSkillAutomotiveMechanic
	ActiveSkillChemistry
	ActiveSkillCybertechnology
	ActiveSkillDemolitions
	ActiveSkillFirstAid
	ActiveSkillHacking
	ActiveSkillMedicine
	ActiveSkillProfessionalKnowledge
	ActiveSkillForgery
	ActiveSkillAeronauticsMechanics
	ActiveSkillArmorer
	ActiveSkillBiotechnology
	ActiveSkillComputer
	ActiveSkillCybercombat
	ActiveSkillElectronicWarfare
	ActiveSkillIndustrialMechanics
	ActiveSkillHardware
	ActiveSkillNauticalMechanics
	ActiveSkillSoftware
	ActiveSkillArtisan
	ActiveSkillDisguise
	ActiveSkillLanguage
	ActiveSkillPerception
	ActiveSkillTracking
	ActiveSkillAssensing
	ActiveSkillNavigation
	ActiveSkillInterestsKnowledge
	ActiveSkillStreetKnowledge
	ActiveSkillAlchemy
	ActiveSkillBinding
	ActiveSkillRitualSpellcasting
	ActiveSkillSummoning
	ActiveSkillDisenchanting
	ActiveSkillBanishing
	ActiveSkillCounterspelling
	ActiveSkillSpellcasting
	ActiveSkillEnchanting
	ActiveSkillCompiling
	ActiveSkillRegistering
	ActiveSkillDecompiling
)

func LoadActiveSkills(wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Debug("Started loading active skills")

	files, errReadDir := os.ReadDir(ActiveSkillDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read active skills directory")
	}

	// Create a map to store the metatypes
	activeSkills := make(map[string]ActiveSkill, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", ActiveSkillDataPath, file.Name())

			var activeSkill ActiveSkill
			if err := util.LoadStructFromYAML(filepath, &activeSkill); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load active skills")
			}

			activeSkills[activeSkill.GetName()] = activeSkill
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded active skills file")
	}

	logrus.WithFields(logrus.Fields{"count": len(activeSkills)}).Info("Done loading active skills")

	ActiveSkills = activeSkills
}
