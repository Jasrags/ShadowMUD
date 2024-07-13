package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/google/uuid"
	"github.com/muesli/termenv"
)

const (
	host = "localhost"
	port = "23234"
)

// app contains a wish server and the list of running programs.
type app struct {
	*ssh.Server
	progs sync.Map
}

// send dispatches a message to all running programs.
func (a *app) send(msg tea.Msg) {
	a.progs.Range(func(key, value interface{}) bool {
		go value.(*tea.Program).Send(msg)

		return true
	})
}

func (a *app) sendToId(id string, msg tea.Msg) {
	prog, ok := a.progs.Load(id)
	if !ok {
		return
	}

	go prog.(*tea.Program).Send(msg)
}

func newApp() *app {
	a := new(app)

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.MiddlewareWithProgramHandler(a.ProgramHandler, termenv.ANSI256),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	a.Server = s

	return a
}

func (a *app) Start() {
	var err error
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = a.ListenAndServe(); err != nil {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := a.Shutdown(ctx); err != nil {
		log.Error("Could not stop server", "error", err)
	}
}

func (a *app) ProgramHandler(s ssh.Session) *tea.Program {
	model := initialChatModel(s)
	model.app = a
	model.id = uuid.New().String()
	log.Info("New program", "id", model.id)

	p := tea.NewProgram(model, bubbletea.MakeOptions(s)...)
	a.progs.Store(model.id, p)

	a.sendToId(model.id, chatMsg{id: model.id, text: "Welcome to the chat room!"})

	return p
}

func main() {
	app := newApp()
	app.Start()
}

type (
	errMsg  error
	chatMsg struct {
		id   string
		text string
	}
)
