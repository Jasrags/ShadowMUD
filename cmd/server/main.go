package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"shadowrunmud/character/metatype"
	"shadowrunmud/character/quality"
	"shadowrunmud/character/skill"
	"shadowrunmud/config"
	"shadowrunmud/util"
	"syscall"
	"time"

	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"
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

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(serverConfig.Host, serverConfig.Port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			// bubbletea.MiddlewareWithProgramHandler(s.ProgramHandler, termenv.ANSI256),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		logrus.WithError(err).Error("Could not create server")
	}

	logrus.WithFields(logrus.Fields{"host": serverConfig.Host, "port": serverConfig.Port}).Info("Starting SSH server")
	go func() {
		if err = s.ListenAndServe(); err != nil {
			logrus.WithError(err).Error("Could not start server")
			done <- nil
		}
	}()

	<-done

	logrus.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Could not stop server")
	}

	logrus.Info("Stopping the server")
}
