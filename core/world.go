package core

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jasrags/ShadowMUD/config"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"
	"github.com/sirupsen/logrus"
)

type World struct {
	*ssh.Server
	config config.Server
	progs  []*tea.Program
}

func NewWorld(serverConfig config.Server) *World {
	w := &World{
		config: serverConfig,
	}

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(serverConfig.Host, serverConfig.Port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.MiddlewareWithProgramHandler(w.ProgramHandler, termenv.ANSI256),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		logrus.WithError(err).Error("Could not create server")
	}

	w.Server = s

	return w
}

func (w *World) Start() {
	var err error
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logrus.WithFields(logrus.Fields{"host": w.config.Host, "port": w.config.Port}).Info("Starting SSH server")
	go func() {
		if err = w.Server.ListenAndServe(); err != nil {
			logrus.WithError(err).Error("Could not start server")
			done <- nil
		}
	}()

	<-done

	logrus.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer func() { cancel() }()
	if err := w.Server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Could not stop server")
	}

	logrus.Info("Stopping the server")
}

func (w *World) ProgramHandler(s ssh.Session) *tea.Program {
	// m := NewInitialModel(s)
	m := NewGameModel(s)
	// model := initialModel()
	// model.app = a
	// model.id = s.User()

	p := tea.NewProgram(m, bubbletea.MakeOptions(s)...)
	w.progs = append(w.progs, p)

	return p
}

// send dispatches a message to all running programs.
func (w *World) send(msg tea.Msg) {
	for _, p := range w.progs {
		go p.Send(msg)
	}
}
