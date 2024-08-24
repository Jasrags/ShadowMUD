package main

import (
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/utils"

	_ "embed"

	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
)

const (
	ConfigFilepath = "_data/config/server.yaml"
)

var (
	w   *World
	cfg *config.Server
)

func main() {
	if err := utils.LoadStructFromYAML(ConfigFilepath, &cfg); err != nil {
		logrus.WithError(err).Fatal("Could not load server configuration")
	}

	logrus.WithField("config", cfg).Info("Loaded server configuration")

	logrus.Info("==================================")
	for k, v := range structs.Map(cfg) {
		logrus.Infof("%24s: %v", k, v)
	}
	logrus.Info("==================================")

	logrusLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)
	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

	w = NewWorld(cfg)
	w.LoadData()
	w.StartServer()
	defer w.StopServer()
}
