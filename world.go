package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/sirupsen/logrus"
)

type World struct {
	*ssh.Server
	config config.Server
	progs  []*tea.Program
}

func NewWorld(serverConfig config.Server) *World {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(serverConfig.Host, serverConfig.Port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			// bubbletea.MiddlewareWithProgramHandler(a.ProgramHandler, termenv.ANSI256),
			bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				// return model.NewInitialModel(s), []tea.ProgramOption{tea.WithAltScreen()}
				return model.NewGameModel(s), []tea.ProgramOption{tea.WithAltScreen()}
			}),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		logrus.WithError(err).Error("Could not create server")
	}

	return &World{
		Server: s,
		config: serverConfig,
	}
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

// func (w *World) ProgramHandler(s ssh.Session) *tea.Program {
// 	// model := initialModel()
// 	// model.app = a
// 	// model.id = s.User()

// 	// p := tea.NewProgram(model, bubbletea.MakeOptions(s)...)
// 	// a.progs = append(a.progs, p)

// 	// return p
// }

// send dispatches a message to all running programs.
func (w *World) send(msg tea.Msg) {
	for _, p := range w.progs {
		go p.Send(msg)
	}
}

// import (
// 	"net"
// 	"sync"

// 	"github.com/google/uuid"
// )

// type World struct {
// 	connections sync.Map
// }

// func NewWorld() *World {
// 	return &World{
// 		connections: sync.Map{},
// 	}
// }

// func (w *World) AddConnection(conn net.Conn) string {
// 	id := uuid.New().String()
// 	w.connections.Store(id, conn)

// 	return id
// }
