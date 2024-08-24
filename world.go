package main

import (
	"io"
	"net"
	"strings"
	"sync"

	_ "embed"

	"github.com/Jasrags/ShadowMUD/common/user"
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/google/uuid"

	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
)

type InputMessage struct {
	FromSessionID string
	Message       string
}

type ResponseMessage struct {
	ToSessionID string
	Message     string
}

type World struct {
	sync.Mutex

	cfg *config.Server
	srv *ssh.Server

	userManager *user.Manager

	sessions        map[string]*ssh.Session
	inputChan       chan InputMessage
	messageChannels map[string]chan string
}

func NewWorld(cfg *config.Server) *World {
	w := &World{
		cfg: cfg,

		userManager: user.NewManager(),

		sessions:        make(map[string]*ssh.Session),
		inputChan:       make(chan InputMessage),
		messageChannels: make(map[string]chan string),
	}

	go w.processInput()

	return w
}

func (w *World) LoadData() {
	w.userManager.Load()
}

func (w *World) StartServer() {
	w.srv = &ssh.Server{
		Addr:        net.JoinHostPort(cfg.Host, cfg.Port),
		IdleTimeout: cfg.IdleTimeout,
		ConnectionFailedCallback: func(conn net.Conn, err error) {
			logrus.WithError(err).Error("Connection failed")
			conn.Close()
		},
		Handler: w.sshHandler,
	}

	logrus.WithFields(logrus.Fields{"host": cfg.Host, "port": cfg.Port}).Info("Starting server")
	if err := w.srv.ListenAndServe(); err != nil {
		logrus.WithError(err).Error("Error starting server")
	}
}

func (w *World) StopServer() {
	logrus.Info("Stopping server")
	w.srv.Close()
	// TODO: Cleanup any open connections and channels
}

func (w *World) AddMessageChannel(id string, ch chan string) {
	logrus.WithField("id", id).Debug("Adding message channel")
	w.Lock()
	defer w.Unlock()

	w.messageChannels[id] = ch
}

func (w *World) RemoveMessageChannel(id string) {
	logrus.WithField("id", id).Debug("Removing message channel")
	w.Lock()
	defer w.Unlock()

	delete(w.messageChannels, id)
}

func (w *World) AddSession(id string, s *ssh.Session) {
	logrus.WithField("id", id).Debug("Adding session")
	w.Lock()
	defer w.Unlock()

	w.sessions[id] = s
}

func (w *World) GetSession(id string) *ssh.Session {
	logrus.WithField("id", id).Debug("Getting session")
	w.Lock()
	defer w.Unlock()

	return w.sessions[id]
}

func (w *World) RemoveSession(id string) {
	logrus.WithField("id", id).Debug("Removing session")
	w.Lock()
	defer w.Unlock()

	delete(w.sessions, id)
}

type State int

const (
	StateBanner State = iota
	StateLoginUser
	StateRegisterUser
	StateMainMenu
	StateEnterGame
	StateGameLoop
	StatePromptCreateCharacter
	StatePromptListCharacters
	StatePromptDeleteCharacter
	StatePromptChangePassword
	StateMOTD
	StateQuit
)

// sshHandler is the handler for incoming SSH connections
func (w *World) sshHandler(s ssh.Session) {
	sessionID := uuid.New().String()
	w.AddSession(sessionID, &s)
	defer w.RemoveSession(sessionID)

	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "session_id": sessionID})
	l.Info("New session")

	var u *user.User
	state := StateBanner
	for {
		switch state {
		case StateBanner:
			state = w.banner(s)
		case StateLoginUser:
			state, u = w.promptLoginUser(s)
		case StateMainMenu:
			state = w.promptMainMenu(s, u)
		case StateEnterGame:
			state = w.enterGame(s, u)
		case StatePromptCreateCharacter:
			state = w.promptCreateCharacter(s, u)
		case StatePromptListCharacters:
			state = w.promptListCharacters(s, u)
		case StatePromptDeleteCharacter:
			state = w.promptDeleteCharacter(s, u)
		case StatePromptChangePassword:
			state = w.promptChangePassword(s, u)
		case StateMOTD:
			state = w.promptMOTD(s, u)
		case StateGameLoop:
			state = w.gameLoop(s, u)
		case StateQuit:
			io.WriteString(s, "Goodbye\n")
			w.RemoveMessageChannel(u.ID)
			w.RemoveSession(sessionID)
			return
		default:
			l.WithField("state", state).Error("Unknown state")
		}
	}
}

// processInput is a goroutine that processes input messages from clients
func (w *World) processInput() {
	for inputMsg := range w.inputChan {
		logrus.WithFields(logrus.Fields{"from_session_id": inputMsg.FromSessionID, "message": inputMsg.Message}).Info("Received input")
		var processedMessage string

		// Pull off the first word of inputMsg.Message
		command := strings.ToLower(strings.Fields(inputMsg.Message)[0])
		args := strings.Fields(inputMsg.Message)[1:]

		// Use the command name to determine the action
		switch command {
		case "say":
			processedMessage = "You said: " + strings.Join(args, " ")
		default:
			processedMessage = "Unknown command: " + command
		}
		w.Lock()
		if ch, ok := w.messageChannels[inputMsg.FromSessionID]; ok {
			ch <- processedMessage
		}
		w.Unlock()

		logrus.WithFields(logrus.Fields{"from_session_id": inputMsg.FromSessionID, "message": inputMsg.Message}).Info("Sent message")
	}
}
