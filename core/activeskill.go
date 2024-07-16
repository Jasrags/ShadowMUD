package core

import (
	"fmt"
	"os"
	"strings"

	"shadowrunmud/core/util"

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

type ActiveSkill struct {
	ID                     string    `yaml:"id,omitempty"`
	Name                   string    `yaml:"name"`
	IsDefaultable          bool      `yaml:"is_defaultable"`
	Description            string    `yaml:"description"`
	LinkedAttribute        Attribute `yaml:"linked_attribute"`
	SkillGroup             string    `yaml:"skill_group,omitempty"`
	Specializations        []string  `yaml:"specializations"`
	SelectedSpecialization string    `yaml:"selected_specialization,omitempty"`
	Rank                   int       `yaml:"rank,omitempty"`
	RuleSource             string    `yaml:"rule_source"`
	FileVersion            string    `yaml:"file_version"`
}

var (
	DefaultActiveSkills = map[string]ActiveSkill{
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

// type (
// 	ActiveSkillIdx int
// )

// const (
// 	ActiveSkillDriving ActiveSkillIdx = iota
// 	ActiveSkillFreeFall
// 	ActiveSkillArchery
// 	ActiveSkillAutomatics
// 	ActiveSkillBlades
// 	ActiveSkillClubs
// 	ActiveSkillEscapeArtist
// 	ActiveSkillGunnery
// 	ActiveSkillHeavyWeapons
// 	ActiveSkillGymnastics
// 	ActiveSkillLocksmith
// 	ActiveSkillLongarms
// 	ActiveSkillPistols
// 	ActiveSkillPalming
// 	ActiveSkillSneaking
// 	ActiveSkillThrowingWeapon
// 	ActiveSkillUnarmedCombat
// 	ActiveSkillExoticMeleeWeapon
// 	ActiveSkillExoticRangedWeapon
// 	ActiveSkillPilotAerospace
// 	ActiveSkillPilotWalker
// 	ActiveSkillPilotAircraft
// 	ActiveSkillPilotExoticVehicle
// 	ActiveSkillPilotGroundCraft
// 	ActiveSkillPilotWatercraft
// 	ActiveSkillRunning
// 	ActiveSkillSwimming
// 	ActiveSkillAstralCombat
// 	ActiveSkillSurvival
// 	ActiveSkillCon
// 	ActiveSkillInstruction
// 	ActiveSkillLeadership
// 	ActiveSkillPerformance
// 	ActiveSkillAnimalHandling
// 	ActiveSkillEtiquette
// 	ActiveSkillIntimidation
// 	ActiveSkillNegotiation
// 	ActiveSkillImpersonation
// 	ActiveSkillAcademicKnowledge
// 	ActiveSkillArcana
// 	ActiveSkillAutomotiveMechanic
// 	ActiveSkillChemistry
// 	ActiveSkillCybertechnology
// 	ActiveSkillDemolitions
// 	ActiveSkillFirstAid
// 	ActiveSkillHacking
// 	ActiveSkillMedicine
// 	ActiveSkillProfessionalKnowledge
// 	ActiveSkillForgery
// 	ActiveSkillAeronauticsMechanics
// 	ActiveSkillArmorer
// 	ActiveSkillBiotechnology
// 	ActiveSkillComputer
// 	ActiveSkillCybercombat
// 	ActiveSkillElectronicWarfare
// 	ActiveSkillIndustrialMechanics
// 	ActiveSkillHardware
// 	ActiveSkillNauticalMechanics
// 	ActiveSkillSoftware
// 	ActiveSkillArtisan
// 	ActiveSkillDisguise
// 	ActiveSkillLanguage
// 	ActiveSkillPerception
// 	ActiveSkillTracking
// 	ActiveSkillAssensing
// 	ActiveSkillNavigation
// 	ActiveSkillInterestsKnowledge
// 	ActiveSkillStreetKnowledge
// 	ActiveSkillAlchemy
// 	ActiveSkillBinding
// 	ActiveSkillRitualSpellcasting
// 	ActiveSkillSummoning
// 	ActiveSkillDisenchanting
// 	ActiveSkillBanishing
// 	ActiveSkillCounterspelling
// 	ActiveSkillSpellcasting
// 	ActiveSkillEnchanting
// 	ActiveSkillCompiling
// 	ActiveSkillRegistering
// 	ActiveSkillDecompiling
// )

func LoadActiveSkills() {
	logrus.Info("Started loading active skills")

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

			activeSkills[activeSkill.Name] = activeSkill
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded active skills file")
	}

	logrus.WithFields(logrus.Fields{"count": len(activeSkills)}).Info("Done loading active skills")

	ActiveSkills = activeSkills
}

func LoadActiveSkill(name string) (*ActiveSkill, error) {
	var v ActiveSkill
	if err := util.LoadStructFromYAML(fmt.Sprintf(ActiveSkillFilename, name), &v); err != nil {
		return nil, err
	}

	return &v, nil
}
