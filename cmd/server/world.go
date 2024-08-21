package main

import (
	"io"
	"sync"

	"github.com/Jasrags/ShadowMUD/common/armor"
	"github.com/Jasrags/ShadowMUD/common/metatype"
	"github.com/Jasrags/ShadowMUD/common/room"
	"github.com/Jasrags/ShadowMUD/common/user"
	"github.com/Jasrags/ShadowMUD/common/zone"
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/screen"

	"github.com/gliderlabs/ssh"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
)

type (
	Sessions map[string]*ssh.Session
	World    struct {
		lock sync.Mutex
		cfg  *config.Server

		// users    common.Users
		sessions Sessions

		// Loaded data
		// zones     common.Zones
		// rooms     common.Rooms
		// metatypes *common.MetatypeManager
		// pregens   common.Pregens

		zones zone.Zones
		rooms room.Rooms

		userManager     *user.Manager
		armorManager    *armor.Manager
		metatypeManager *metatype.Manager
		// pregens   common.Pregens

		broadcast    chan string
		userChannels map[string]chan string
		// commandQueue chan Command
	}
)

func NewWorld(cfg *config.Server) *World {
	w := &World{
		cfg:             cfg,
		metatypeManager: metatype.NewManager(),
		userManager:     user.NewManager(),
		armorManager:    armor.NewManager(),

		sessions:     make(Sessions),
		broadcast:    make(chan string),
		userChannels: make(map[string]chan string),

		// commandQueue: make(chan Command),
	}
	go w.broadcastMessages()
	go w.handleCommands()

	return w
}
func (w *World) handleCommands() {
	// for cmd := range w.commandQueue {
	// w.ProcessCommand(cmd)
	// }
}

func (w *World) broadcastMessages() {
	for msg := range w.broadcast {
		w.lock.Lock()
		for _, userChan := range w.userChannels {
			userChan <- msg
		}
		w.lock.Unlock()
	}
}

func (w *World) LoadData() {
	logrus.Info("Started loading data")

	// w.zones = common.LoadZones()
	// w.rooms = common.LoadRooms()
	w.userManager.Load()
	w.metatypeManager.Load()
	w.armorManager.Load()
	// w.metatypes = common.LoadMetatypes()
	// // w.pregens = common.LoadPregens()
	// sm := common.NewSkillManager()
	// if err := sm.LoadSkills(); err != nil {
	// 	logrus.WithError(err).Error("Error loading skills")
	// }
	logrus.Info("Finished loading data")
}

func (w *World) RoundTick() {
}

func (w *World) TurnTick() {
}

// Handle incoming SSH sessions
func (w *World) Handler(s ssh.Session) {
	logrus.WithFields(logrus.Fields{"user": s.User(), "remote_addr": s.RemoteAddr()}).Info("New connection")

	u := user.New(s)
	screens := screen.New(u, w.cfg)
	state := screen.StateBanner
	for {
		switch state {
		case screen.StateBanner:
			state = screens.Banner()
		case screen.StatePromptLoginUser:
			state = screens.PromptLoginUser()
		case screen.StatePromptRegisterUser:
			state = screens.PromptRegisterUser()
		case screen.StatePromptMainMenu:
			state = screens.PromptMainMenu()
		case screen.StateEnterGame:
			state = screens.EnterGame()
			// TODO: If character has not been selected, prompt for character selection
			// TODO: Enter game loop
		case screen.StatePromptCreateCharacter:
			state = screens.PromptCreateCharacter()
			// TODO: Prompt for using archtype or creating a new character
		case screen.StatePromptListCharacters:
			state = screens.PromptListCharacters()
			// TODO: List out all the characters for the user with short details
		case screen.StatePromptDeleteCharacter:
			state = screens.PromptDeleteCharacter()
			// TODO: List out all the characters for the user
			// TODO: Prompt for the character to delete
			// TODO: Confirm the deletion
			// TODO: Delete the character
		case screen.StatePromptChangePassword:
			state = screens.PromptChangePassword()
		case screen.StateMOTD:
			state = screens.MOTD()
		case screen.StateGameLoop:
			state = screens.GameLoop()
		case screen.StateQuit:
			u.Save()
			io.WriteString(u.Session, cfmt.Sprint("Goodbye!\n"))
			w.RemoveSession(u.ID)
			w.userManager.RemoveByID(u.ID)
			// w.RemoveUser(u.Username)
			u.Session.Close()
			return
		default:
			logrus.WithField("state", state).Error("Invalid state")
			state = screen.StateBanner
		}
	}
}

// TODO: Match against commands, users, etc...
// TODO: Cycle through the list of available commands when we have more than one
func (w *World) AutoCompleteCallback(line string, pos int, key rune) (string, int, bool) {
	logrus.WithFields(logrus.Fields{"line": line, "pos": pos, "key": key, "key_string": string(key)}).Debug("AutoCompleteCallback")

	return "", pos, false // Return the current result as the auto-completed text
}

// ProcessCommand simulates processing a command in the world
// func (w *World) ProcessCommand(cmd Command) {
// 	logrus.WithFields(logrus.Fields{"command": cmd.Name, "args": cmd.Args}).Info("Processing command")
// 	switch cmd.Name {
// 	case "say":
// 	case "tell":
// 		w.userChannels[cmd.Recipient.Name] <- cfmt.Sprintf("{{%s}} tells you: %s\n", cmd.Sender.Name, cmd.Args[0])
// 	}
// }

func (w *World) SendMessageToAllUsers(message string) {
	w.broadcast <- cfmt.Sprintf("Brodcast: {{%s}}::#00ff00\n", message)
}

// func (w *World) GetUser(username string) *user.User {
// 	w.lock.Lock()
// 	defer w.lock.Unlock()
// 	return w.users[username]
// }

// func (w *World) AddUser(u *user.User) {
// 	w.lock.Lock()
// 	defer w.lock.Unlock()
// 	w.users[u.Username] = u
// 	w.userChannels[u.Username] = make(chan string)
// }

// func (w *World) RemoveUser(username string) {
// 	w.lock.Lock()
// 	defer w.lock.Unlock()
// 	delete(w.users, username)
// 	delete(w.userChannels, username)
// }

func (w *World) GetSession(id string) *ssh.Session {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.sessions[id]
}

func (w *World) AddSession(id string, s *ssh.Session) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.sessions[id] = s
}

func (w *World) RemoveSession(id string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	delete(w.sessions, id)
}

// func NewWorld(osSignalChan chan os.Signal) *World {
// 	var serverConfig config.Server
// 	utils.LoadStructFromYAML(ConfigFilepath, &serverConfig)

// 	// logrus.WithField("serverConfig", serverConfig).Info("Loaded server configuration")

// 	logrusLevel, err := logrus.ParseLevel(serverConfig.LogLevel)
// 	if err != nil {
// 		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
// 		logrusLevel = logrus.InfoLevel
// 	}
// 	logrus.SetLevel(logrusLevel)
// 	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

// 	w := &World{
// 		config: serverConfig,
// 	}

// 	zones, err := utils.LoadAllFiles[common.ZoneSpec](serverConfig.Data.BaseDir + serverConfig.Data.ZonesDir)
// 	if err != nil {
// 		logrus.WithError(err).Fatal("Could not load zones")
// 	}
// 	for _, zone := range zones {
// 		logrus.WithFields(logrus.Fields{"zone_id": zone.ID, "zone_name": zone.Name}).Info("Loaded zone")
// 		// w.zones[zone.ID] = zone
// 	}
// 	// utils.LoadAllFiles[common.RoomSpec](serverConfig.Data.BaseDir + common.RoomsDataPath)

// 	return w
// }
