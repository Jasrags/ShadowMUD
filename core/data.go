package core

type CoreDataSkills struct {
	Language  map[string]LanguageSkill
	Active    map[string]ActiveSkill
	Knowledge map[string]KnowledgeSkill
	Groups    map[string]SkillGroup
}

type CoreDataWeapons struct {
	Melee         map[string]WeaponMelee
	Ranged        map[string]WeaponRanged
	Ammunition    map[string]WeaponAmunition
	Modifications map[string]WeaponModification
}

type CoreDataArmor struct {
	Armor        map[string]Armor
	Modificatons map[string]ArmorModification
}

type CoreData struct {
	Metatypes map[string]Metatype
	Skills    *CoreDataSkills
	Weapons   *CoreDataWeapons
	Armor     *CoreDataArmor
	Qualities map[string]Quality
	Cyberware map[string]Cyberware
	Bioware   map[string]Bioware
	Contacts  map[string]Contact
	// TODO: Gear?
	// TODO: Item?
}

func NewCoreData() *CoreData {
	return &CoreData{
		Metatypes: make(map[string]Metatype),
		Skills: &CoreDataSkills{
			Language:  LoadLanguageSkills(),
			Active:    LoadActiveSkills(),
			Knowledge: LoadKnowledgeSkills(),
			Groups:    LoadSkillGroups(),
		},
		Weapons: &CoreDataWeapons{
			Melee:         LoadMeleeWeapons(),
			Ranged:        LoadRangedWeapons(),
			Ammunition:    LoadWeaponAmunition(),
			Modifications: LoadWeaponModifications(),
		},
		Armor: &CoreDataArmor{
			Armor:        LoadArmor(),
			Modificatons: LoadArmorModificatons(),
		},
		Qualities: LoadQualities(),
		Cyberware: LoadCyberware(),
		Bioware:   LoadBioware(),
		Contacts:  LoadContacts(),
	}
}

var LoadedCoreData = NewCoreData()

func LoadCoreData() {
	// cd := NewCoreData()
	// LoadedCoreData.Skills.Language = LoadLanguageSkills()
	// LoadedCoreData.Skills.Active = LoadActiveSkills()
	// LoadedCoreData.Skills.Knowledge = LoadKnowledgeSkills()
	// Loa
	// LoadedCoreData.Metatypes = LoadMetatypes()
	// LoadedCoreData.Qualities = LoadQualities()
	// LoadedCoreData.CyberwareList = LoadCyberware()
}
