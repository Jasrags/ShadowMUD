package main

import (
	"net"
	"os"

	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/fatih/structs"

	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
)

const (
	ConfigFilepath = "_data/config/server.yaml"
)

var (
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

	// set up logging
	logrusLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)
	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

	w := NewWorld(cfg)
	w.LoadData()

	// Start the server
	server := &ssh.Server{
		Addr:        net.JoinHostPort(cfg.Host, cfg.Port),
		IdleTimeout: cfg.IdleTimeout,
		ConnectionFailedCallback: func(conn net.Conn, err error) {
			logrus.WithError(err).Error("Connection failed")
			conn.Close()
		},
	}
	defer server.Close()
	// defer close(w.commandQueue)

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
