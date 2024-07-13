package skills

type (
	AttributeIdx   int
	ActiveSkillIdx int
	RuleSourceIdx  int
)

const (
	AttributeBody AttributeIdx = iota
	AttributeAgility
	AttributeReaction
	AttributeStrength
	AttributeWillpower
	AttributeCharisma
	AttributeLogic
	AttributeIntuition
	AttributeMagic
	AttributeResonance
	AttributeEssence
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

const (
	RuleSourceCore RuleSourceIdx = iota
)

type ActiveSkill struct {
	ID              int
	Name            string
	IsDefaultable   bool
	Description     string
	RuleSource      RuleSourceIdx
	LinkedAttribute AttributeIdx
	Specializations []string
}

var activeSkills = map[ActiveSkillIdx]ActiveSkill{
	ActiveSkillAutomatics: {
		ID:              1,
		Name:            "Automatics",
		IsDefaultable:   true,
		LinkedAttribute: AttributeAgility,
		Specializations: []string{"Assault Rifles", "Cyber-Implant", "Machine Pistols", "Submachine Guns"},
	},
	ActiveSkillArchery: {
		ID:              1,
		Name:            "Archery",
		IsDefaultable:   true,
		LinkedAttribute: AttributeAgility,
		RuleSource:      RuleSourceCore,
		Specializations: []string{"Assault Rifles", "Cyber-Implant", "Machine Pistols", "Submachine Guns"},
	},
}
