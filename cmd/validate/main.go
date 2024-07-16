package main

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := ValidateLanguageSkillFiles(); err != nil {
		logrus.WithError(err).Fatal("Error validating language skills")
	}
}

func ValidateLanguageSkillFiles() error {
	logrus.Info("Started validating language skills")

	skills := core.LoadLanguageSkills()

	for k, v := range skills {
		if v.ID == "" {
			logrus.WithFields(logrus.Fields{"key": k, "name": v.Name}).Error("Language skill does not have an ID, using Name as ID")
			v.ID = util.FormatFilename(v.Name)
		}

		v.ID = util.FormatFilename(k)
		fileName := fmt.Sprintf(core.LanguageSkillFilename, v.ID)

		logrus.WithFields(logrus.Fields{"id": v.ID, "file_name": fileName}).Info("Validating language skill")
		if err := util.SaveStructToYAML(fmt.Sprintf(core.LanguageSkillFilename, v.ID), v); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{"id": v.ID, "file_name": fileName}).Error("Unable to save language skill")
		}
	}

	return nil
}
