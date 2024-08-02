package main

import (
	"net"
	"os"

	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
)

const (
	ConfigFilepath = "_data/config/server.yaml"
)

var (
	// w   *World
	cfg *config.Server
)

func main() {
	utils.LoadStructFromYAML(ConfigFilepath, &cfg)

	logrus.WithField("config", cfg).Info("Loaded server configuration")

	// set up logging
	logrusLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)

	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

	w := NewWorld(cfg)
	// load data (zones, rooms, items, etc...)

	// Start the server
	server := &ssh.Server{
		Addr:        net.JoinHostPort(cfg.Host, cfg.Port),
		IdleTimeout: cfg.IdleTimeout,
	}
	defer server.Close()

	// handle connections
	ssh.Handle(w.Handler)

	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Error("Could not start server")
	}

	// Shutdown the server

	// block until a signal comes in
	// <-sigChan

	os.Exit(0)
}
