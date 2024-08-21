package main

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common/chummer"

	"github.com/sirupsen/logrus"
)

var (
// filesnames = []string{"skills"}
// skills     *chummer.Skills
)

func main() {
	logrus.Info("Starting to load chummer data")
	LoadSkills()
	// if err := chummer.LoadFromJSON("skills", skills); err != nil {
	// 	logrus.WithError(err).Error("Error loading file")
	// 	return
	// }

	// fmt.Println("Loaded skills:", skills)
}

func LoadSkills() {
	var skills chummer.Skills

	if err := chummer.LoadFromJSON("skills", &skills); err != nil {
		logrus.WithError(err).Error("Error loading file")
		return
	}

	fmt.Printf("Loaded skills: %v\n", skills)

	// for _, skill := range skills.Skills.Skill {
	// logrus.WithFields(logrus.Fields{"id": skill.ID, "name": skill.Name}).Debug("Loaded active skill")
	// }
}
