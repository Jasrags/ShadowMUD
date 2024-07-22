package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/sirupsen/logrus"
)

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
	Skills    CoreDataSkills
	Weapons   CoreDataWeapons
	Armor     CoreDataArmor
	Qualities map[string]Quality
	Cyberware map[string]Cyberware
	Bioware   map[string]Bioware
	Contacts  map[string]Contact
	// TODO: Gear?
	// TODO: Item?
}

var LoadedCoreData CoreData

func (cd *CoreData) Load() {
	logrus.Info("Started loading game data")
	cd.Metatypes = cd.LoadMetatypes(MetatypeDataPath)
	cd.Skills.Language = cd.LoadLanguageSkills(LanguageSkillDataPath)
	cd.Skills.Active = cd.LoadActiveSkills(ActiveSkillDataPath)
	cd.Skills.Knowledge = cd.LoadKnowledgeSkills(KnowledgeSkillDataPath)
	cd.Skills.Groups = cd.LoadSkillGroups(SkillGroupsDataPath)
	// cd.Weapons.Melee = cd.LoadMeleeWeapons(MeleeWeaponDataPath)
	// cd.Weapons.Ranged = cd.LoadRangedWeapons(WeaponRangedDataPath)
	cd.Weapons.Ammunition = cd.LoadWeaponAmunition(WeaponAmunitionDataPath)
	cd.Weapons.Modifications = cd.LoadWeaponModifications(WeaponModificationsDataPath)
	cd.Armor.Armor = cd.LoadArmor(ArmorDataPath)
	cd.Armor.Modificatons = cd.LoadArmorModificatons(ArmorModificationsDataPath)
	cd.Qualities = cd.LoadQualities(QualityDataPath)
	cd.Cyberware = cd.LoadCyberware(CyberwareDataPath)
	cd.Bioware = cd.LoadBioware(BiowareDataPath)
	cd.Contacts = cd.LoadContacts(ContactsDataPath)
	logrus.Info("Done loading game data")
}

func (cd CoreData) LoadMetatypes(dataPath string) map[string]Metatype {
	name := "metatypes"
	data := make(map[string]Metatype)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v Metatype
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadLanguageSkills(dataPath string) map[string]LanguageSkill {
	name := "language skills"
	data := make(map[string]LanguageSkill)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v LanguageSkill
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadActiveSkills(dataPath string) map[string]ActiveSkill {
	name := "active skills"
	data := make(map[string]ActiveSkill)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v ActiveSkill
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadKnowledgeSkills(dataPath string) map[string]KnowledgeSkill {
	name := "knowledge skills"
	data := make(map[string]KnowledgeSkill)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v KnowledgeSkill
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadSkillGroups(dataPath string) map[string]SkillGroup {
	name := "skill groups"
	data := make(map[string]SkillGroup)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v SkillGroup
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadMeleeWeapons(dataPath string) map[string]WeaponMelee {
	name := "melee weapons"
	data := make(map[string]WeaponMelee)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v WeaponMelee
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadRangedWeapons(dataPath string) map[string]WeaponRanged {
	name := "ranged weapons"
	data := make(map[string]WeaponRanged)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v WeaponRanged
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadWeaponAmunition(dataPath string) map[string]WeaponAmunition {
	name := "weapon ammunition"
	data := make(map[string]WeaponAmunition)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v WeaponAmunition
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadWeaponModifications(dataPath string) map[string]WeaponModification {
	name := "weapon modifications"
	data := make(map[string]WeaponModification)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v WeaponModification
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadArmor(dataPath string) map[string]Armor {
	name := "armor"
	data := make(map[string]Armor)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v Armor
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)
	return data
}

func (cd CoreData) LoadArmorModificatons(dataPath string) map[string]ArmorModification {
	name := "armor modifications"
	data := make(map[string]ArmorModification)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v ArmorModification
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)

	return data
}

func (cd CoreData) LoadQualities(dataPath string) map[string]Quality {
	name := "qualities"
	data := make(map[string]Quality)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v Quality
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)

	return data
}

func (cd CoreData) LoadCyberware(dataPath string) map[string]Cyberware {
	name := "cyberware"
	data := make(map[string]Cyberware)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v Cyberware
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)

	return data
}

func (cd CoreData) LoadBioware(dataPath string) map[string]Bioware {
	name := "bioware"
	data := make(map[string]Bioware)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v Bioware
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)

	return data
}

func (cd CoreData) LoadContacts(dataPath string) map[string]Contact {
	name := "contacts"
	data := make(map[string]Contact)

	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Infof("Started loading %s", name)
	files, err := os.ReadDir(dataPath)
	if err != nil {
		logrus.WithError(err).Errorf("Could not read %s data directory", name)
		return data
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

			var v Contact
			if err := utils.LoadStructFromYAML(filePath, &v); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Errorf("Could not load file")
				return data
			}
			data[v.ID] = v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debugf("Loaded %s file", name)
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(data)}).Infof("Done loading %s", name)

	return data
}
