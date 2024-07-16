package main

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"

	"github.com/sirupsen/logrus"
)

func main() {
	var serverConfig config.Server
	util.LoadStructFromYAML("data/config/server.yaml", &serverConfig)

	logrusLevel, err := logrus.ParseLevel(serverConfig.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)

	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

	logrus.Info("Started validating data files")

	if err := ValidateMetatypeFiles(); err != nil {
		logrus.WithError(err).Fatal("Error validating metatypes")
	}
	// if err := ValidateLanguageSkillFiles(); err != nil {
	// 	logrus.WithError(err).Fatal("Error validating language skills")
	// }
}

func ValidateMetatypeFiles() error {
	logrus.Info("Started validating metatypes")

	metatypes := core.LoadMetatypes()
	// logrus.Info(metatypes)

	for k, v := range metatypes {
		if v.ID == "" {
			logrus.WithFields(logrus.Fields{"key": k, "name": v.Name}).Error("Metatype does not have an ID, using Name as ID")
			v.ID = util.FormatFilename(v.Name)
		}

		v.ID = util.FormatFilename(k)
		fileName := fmt.Sprintf(core.MetatypeFilename, v.ID)

		logrus.WithFields(logrus.Fields{"id": v.ID, "file_name": fileName}).Info("Validating metatype")
		// if err := util.SaveStructToYAML(fmt.Sprintf(core.MetatypeFilename, v.ID), v); err != nil {
		// 	logrus.WithError(err).WithFields(logrus.Fields{"id": v.ID, "file_name": fileName}).Error("Unable to save metatype")
		// }
	}

	return nil

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
