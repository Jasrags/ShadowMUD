package main

import (
	"shadowrunmud/character/metatype"
	"shadowrunmud/character/quality"
	"shadowrunmud/character/skill"
	"shadowrunmud/config"
	"shadowrunmud/util"

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

	logrus.Info("Starting the server")

	// Load all the data
	logrus.Info("Loading data files")
	metatype.LoadMetatypes()
	skill.LoadLanguageSkills()
	skill.LoadActiveSkills()
	skill.LoadKnowledgeSkills()
	quality.LoadQualities()
	logrus.Info("Data files loaded")

	logrus.Info("Stopping the server")
}
