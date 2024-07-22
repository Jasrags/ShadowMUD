package core

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Jasrags/ShadowMUD/config"
	"github.com/google/uuid"

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
	// progs     []*tea.Program
	syncProgs sync.Map
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
	id := uuid.New().String()
	logrus.WithFields(logrus.Fields{"id": id, "user": s.User(), "remote_addr": s.RemoteAddr()}).Info("Creating new program")

	// m := NewInitialModel(s)
	m := NewGameModel(s)
	m.World = w
	m.id = id

	p := tea.NewProgram(m, bubbletea.MakeOptions(s)...)
	// w.progs = append(w.progs, p)
	w.syncProgs.Store(m.id, p)

	return p
}

// send dispatches a message to all running programs.
func (w *World) send(msg tea.Msg) {
	// for _, p := range w.progs {
	// 	go p.Send(msg)
	// }
	w.syncProgs.Range(func(key, value interface{}) bool {
		p := value.(*tea.Program)
		go p.Send(msg)
		return true
	})
}
