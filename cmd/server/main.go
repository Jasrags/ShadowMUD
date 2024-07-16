package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"
	"github.com/Jasrags/ShadowMUD/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
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
	logrus.Info("Loading core data")
	core.LoadCoreData()
	logrus.Info("Core data loaded")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(serverConfig.Host, serverConfig.Port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			// bubbletea.MiddlewareWithProgramHandler(a.ProgramHandler, termenv.ANSI256),
			bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				return model.NewInitialModel(s), []tea.ProgramOption{tea.WithAltScreen()}
			}),
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
