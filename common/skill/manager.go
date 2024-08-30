package skill

import "github.com/sirupsen/logrus"

type SkillManager struct {
	Active    Specs
	Knowledge Specs
}

func NewSkillManager() *SkillManager {
	return &SkillManager{
		Active:    make(Specs),
		Knowledge: make(Specs),
	}
}

func (sm *SkillManager) LoadSkill(id string) (*Spec, error) {
	logrus.WithFields(logrus.Fields{"id": id}).Debug("Loading skill")
	return nil, nil
}

func (sm *SkillManager) LoadSkills() error {
	logrus.Info("Loading skills")

	for _, skill := range CoreSkills {
		switch skill.Category {
		case SkillCategoryCombat, SkillCategoryPhysical, SkillCategorySocial, SkillCategoryMagical, SkillCategoryPseudoMagical, SkillCategoryResonance, SkillCategoryTechnical, SkillCategoryVehicle:
			sm.Active[skill.ID] = &skill
		case SkillCategoryAcademic, SkillCategoryInterest, SkillCategoryLanguage, SkillCategoryProfessional, SkillCategoryStreet:
			sm.Knowledge[skill.ID] = &skill
		}
	}

	logrus.Infof("Loaded %d active skills", len(sm.Active))
	logrus.Infof("Loaded %d knowledge skills", len(sm.Knowledge))

	// sm.Active = Core
	// var list []Skill
	// if err := utils.LoadStructFromYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "active"), &list); err != nil {
	// 	logrus.WithError(err).Error("Error loading active skills")
	// 	return err
	// }
	// for _, skill := range sl.Skill {
	// 	logrus.WithFields(logrus.Fields{"id": skill.ID, "name": skill.Name}).Debug("Loaded active skill")
	// }
	// if err := utils.LoadStructFromYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "knowledge"), &sm.Knowledge); err != nil {
	// 	logrus.WithError(err).Error("Error loading knowledge skills")
	// 	return err
	// }
	// for _, skill := range sm.Knowledge {
	// 	logrus.WithFields(logrus.Fields{"id": skill.ID, "name": skill.Name}).Debug("Loaded knowledge skill")
	// }

	return nil
}

func (sm *SkillManager) SaveSkill(skill *Spec) error {
	logrus.WithFields(logrus.Fields{"id": skill.ID}).Debug("Saving skill")
	return nil
}

func (sm *SkillManager) SaveSkills() error {
	logrus.Debug("Saving skills")
	// // Save Active Skills
	// if err := utils.SaveStructToYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "active"), sm.Active); err != nil {
	// 	logrus.WithError(err).Error("Error saving active skills")
	// 	return err
	// }
	// // Save Knowledge Skills
	// if err := utils.SaveStructToYAML(fmt.Sprintf("%s/%s.yaml", SkillsFilepath, "knowledge"), sm.Knowledge); err != nil {
	// 	logrus.WithError(err).Error("Error saving knowledge skills")
	// 	return err
	// }

	return nil
}

func (sm *SkillManager) GetInstance(id string) (*Spec, error) {
	logrus.WithFields(logrus.Fields{"id": id}).Debug("Getting skill instance")

	return nil, nil
}

// func (sm *SkillManager) Get(id string) (*Skill, error) {
// 	log := logrus.WithFields(logrus.Fields{"id": id})
// 	log.Debug("Getting skill")

// 	skill, ok := sm.Skills[id]
// 	if !ok {
// 		log.Error("Skill not found")
// 		return nil, fmt.Errorf("Skill not found")
// 	}

// 	return skill, nil
// }

// func (sm *SkillManager) Add(skill *Skill) error {
// 	logrus.WithFields(logrus.Fields{"id": skill.ID}).Debug("Adding skill")
// 	sm.Skills[skill.ID] = skill

// 	return nil
// }

// func (sm *SkillManager) Remove(id string) error {
// 	logrus.WithFields(logrus.Fields{"id": id}).Debug("Removing skill")
// 	delete(sm.Skills, id)

// 	return nil
// }
