package main

import (
	"io"
	"strings"
	"sync"
	"time"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/google/uuid"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/gliderlabs/ssh"
)

type World struct {
	lock      sync.Mutex
	cfg       *config.Server
	users     map[string]*common.User
	sessions  map[string]*ssh.Session
	metatypes common.Metatypes
}

func NewWorld(cfg *config.Server) *World {
	w := &World{
		cfg:      cfg,
		users:    make(map[string]*common.User),
		sessions: make(map[string]*ssh.Session),
	}
	return w
}

const (
	StateBanner = iota
	StatePromptLoginUser
	StatePromptRegisterUser
	StatePromptMainMenu
	StateEnterGame
	StatePromptCreateCharacter
	StatePromptListCharacters
	StatePromptDeleteCharacter
	StatePromptChangePassword
	StateMOTD
	StateGameLoop
	StateQuit
)

var (
	usernamePrompt                  = "{{Username: }}::#ffffff|bold"
	passwordPrompt                  = "{{Password: }}::#ffffff|bold"
	passwordConfirmPrompt           = "{{Confirm password: }}::#ffffff|bold"
	invalidLoginMsg                 = "{{You have entered an invalid username or password.}}::#ff8700\n"
	loginSuccessfulMsg              = "{{Login successful.}}::#00ff00\n"
	loginUserBannedMsg              = "{{You are banned until %s.}}::#ff8700\n"
	registrationClosedMsg           = "{{Registration is currently closed.}}::#ff8700\n"
	registrationUsernameBanned      = "{{Username %s is not allowed.}}::#ff8700\n"
	loginClosedMsg                  = "{{Login is currently closed.}}::#ff8700\n"
	passwordMismatchMsg             = "{{Passwords do not match.}}::#ff8700\n"
	passwordMinLengthMsg            = "{{Password must be at least %d characters.}}::#ff8700\n"
	passwordMaxLengthMsg            = "{{Password must be at most %d characters.}}::#ff8700\n"
	userCreatedMsg                  = "{{User %s has been created.}}::#00ff00\n"
	newUsernamePrompt               = "{{Enter your desired username: }}::#ffffff|bold"
	confirmUsernamePrompt           = "{{Confirm username %s}}::#ffffff|bold {{(y/n)}}::#00ff00|bold{{:}}::#ffffff|bold"
	declinedUsernameMsg             = "{{Username %s was not confirmed.}}::#ff8700\n"
	mainMenuTitle                   = "\n{{Main Menu}}::#00ff00|bold\n"
	menuOptionEnterGame             = "{{1.}}::#00ff00 Enter game [%s]\n"
	menuOptionCreateCharacter       = "{{2.}}::#00ff00 Create character (%d/%d)\n"
	menuOptionListCharacters        = "{{3.}}::#00ff00 List characters\n"
	menuOptionDeleteCharacter       = "{{4.}}::#00ff00 Delete character\n"
	menuOptionChangePassword        = "{{5.}}::#00ff00 Change password\n"
	menuOptionQuit                  = "{{0.}}::#00ff00 Quit\n"
	menuPrompt                      = "{{Enter the number of the option you would like to select: }}::#ffffff|bold"
	menuInvalidChoice               = "Invalid choice: %s\n"
	gameLoopPrompt                  = "{{> }}::#ffffff|bold"
	inputEchoMsg                    = "{{You typed:}}::#ffffff|bold %s\n"
	createCharacterMenuTitle        = "\n{{Character Creation}}::#00ff00|bold\n"
	createCharacterMenuOptionPregen = "{{1.}}::#00ff00 Choose an pre-generated character\n"
	createCharacterMenuOptionCustom = "{{2.}}::#00ff00 Create an custom character\n"
	createCharacterMenuOptionLearn  = "{{3.}}::#00ff00 Learn more about Shadowrun characters\n"
	createCharacterMenuOptionReturn = "{{4.}}::#00ff00 Return to the main menu\n"
	noCharactersCreatedMsg          = "\n{{You have no characters created, let's make one now!}}::#ff8700\n"
	characterListOption             = "{{%d.}}::#00ff00 %s\n"
)

func (w *World) LoadData() {
	logrus.Info("Started loading data")

	w.metatypes = common.LoadMetatypes()

	logrus.Info("Finished loading data")
}

// Handle incoming SSH sessions
func (w *World) Handler(s ssh.Session) {
	logrus.WithFields(logrus.Fields{"user": s.User(), "remote_addr": s.RemoteAddr()}).Info("New connection")

	u := common.NewUser(s)

	state := StateBanner

	for {
		switch state {
		case StateBanner:
			state = w.Banner(u)
		case StatePromptLoginUser:
			state = w.PromptLoginUser(u)
		case StatePromptRegisterUser:
			state = w.PromptRegisterUser(u)
		case StatePromptMainMenu:
			state = w.PromptMainMenu(u)
		case StateEnterGame:
			state = w.EnterGame(u)
			// TODO: If character has not been selected, prompt for character selection
			// TODO: Enter game loop
		case StatePromptCreateCharacter:
			state = w.PromptCreateCharacter(u)
			// TODO: Prompt for using archtype or creating a new character
		case StatePromptListCharacters:
			state = w.PromptListCharacters(u)
			// TODO: List out all the characters for the user with short details
		case StatePromptDeleteCharacter:
			state = w.PromptDeleteCharacter(u)
			// TODO: List out all the characters for the user
			// TODO: Prompt for the character to delete
			// TODO: Confirm the deletion
			// TODO: Delete the character
		case StatePromptChangePassword:
			state = w.PromptChangePassword(u)
			// TODO: Prompt for new password
			// TODO: Prompt for confirm password
			// TODO: Confirm the new passwords match
			// TODO: Hash the new password
			// TODO: Save the new password
		case StateMOTD:
			state = w.MOTD(u)
		case StateGameLoop:
			state = w.GameLoop(u)
		case StateQuit:
			u.Save()
			io.WriteString(u.Session, cfmt.Sprint("Goodbye!\n"))
			w.RemoveSession(u.ID)
			w.RemoveUser(u.Username)
			u.Session.Close()
			return
		default:
			logrus.WithField("state", state).Error("Invalid state")
			state = StateBanner
		}
	}

}

func (w *World) Banner(u *common.User) int {
	io.WriteString(u.Session, cfmt.Sprint(`
{{     ::::::::  :::    :::     :::     :::::::::   ::::::::  :::       ::: ::::    ::::  :::    ::: :::::::::  }}::#ff8700
{{    :+:    :+: :+:    :+:   :+: :+:   :+:    :+: :+:    :+: :+:       :+: +:+:+: :+:+:+ :+:    :+: :+:    :+: }}::#ff5f00
{{    +:+        +:+    +:+  +:+   +:+  +:+    +:+ +:+    +:+ +:+       +:+ +:+ +:+:+ +:+ +:+    +:+ +:+    +:+ }}::#ff0000
{{    +#++:++#++ +#++:++#++ +#++:++#++: +#+    +:+ +#+    +:+ +#+  +:+  +#+ +#+  +:+  +#+ +#+    +:+ +#+    +:+ }}::#d70000
{{           +#+ +#+    +#+ +#+     +#+ +#+    +#+ +#+    +#+ +#+ +#+#+ +#+ +#+       +#+ +#+    +#+ +#+    +#+ }}::#af0000
{{    #+#    #+# #+#    #+# #+#     #+# #+#    #+# #+#    #+#  #+#+# #+#+#  #+#       #+# #+#    #+# #+#    #+# }}::#870000
{{     ########  ###    ### ###     ### #########   ########    ###   ###   ###       ###  ########  #########  }}::#5f0000

`))
	io.WriteString(u.Session, cfmt.Sprintf("Enter your username to continue or type %s\n", "{{new}}::#ffffff|bold"))

	return StatePromptLoginUser
}

func (w *World) PromptLoginUser(u *common.User) int {
	// Collect username
	io.WriteString(u.Session, cfmt.Sprint(usernamePrompt))
	username, errReadLine := u.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")

		return StatePromptLoginUser
	}

	username = strings.TrimSpace(username)
	logrus.WithField("username", username).Debug("Received username")

	// Do we need to make a new user?
	if strings.EqualFold(username, "new") {
		w.PromptRegisterUser(u)

		return StatePromptRegisterUser
	}

	password, err := u.Term.ReadPassword(cfmt.Sprint(passwordPrompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading password")
		return StatePromptLoginUser
	}
	logrus.WithFields(logrus.Fields{"username": username}).Debug("Received password")

	// is login enabled?
	if !w.cfg.LoginEnabled {
		io.WriteString(u.Session, cfmt.Sprint(loginClosedMsg))

		return StatePromptLoginUser
	}

	// Try  the user from the file
	if err := common.LoadUser(username, u); err != nil {
		logrus.WithError(err).Error("Error loading user")

		return StatePromptLoginUser
	}

	logrus.WithFields(logrus.Fields{"username": u.Username, "id": u.ID}).Debug("Loaded user")

	// Is the user banned?
	for _, ban := range u.Bans {
		if ban.ExpiresAt.After(time.Now()) {
			io.WriteString(u.Session, cfmt.Sprintf(loginUserBannedMsg, ban.ExpiresAt.Format(time.RFC1123)))

			return StatePromptLoginUser
		}
	}

	// Validate the password against the hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		logrus.WithError(err).Error("Password validation error")

		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			io.WriteString(u.Session, cfmt.Sprintf(invalidLoginMsg))
			logrus.WithFields(logrus.Fields{"username": username}).WithError(err).Error("Invalid password for user")

		case bcrypt.ErrHashTooShort:
		case bcrypt.ErrMismatchedHashAndPassword:
		case bcrypt.ErrPasswordTooLong:
		default:
			logrus.WithFields(logrus.Fields{"username": username}).WithError(err).Error("Password error")
		}

		return StatePromptLoginUser
	}

	// Add a login record
	u.Logins = append(u.Logins, common.Login{
		Time: time.Now(),
		IP:   u.Session.RemoteAddr().String(),
	})

	// Save the user
	if err := u.Save(); err != nil {
		logrus.WithError(err).Error("Error saving user")
		return StatePromptLoginUser
	}

	io.WriteString(u.Session, cfmt.Sprintf(loginSuccessfulMsg))

	return StatePromptMainMenu
}

func (w *World) PromptRegisterUser(u *common.User) int {
	// Do we have registration disabled?
	// TODO: Fix this loop
	if !w.cfg.RegistrationEnabled {
		io.WriteString(u.Session, cfmt.Sprint(registrationClosedMsg))

		return StatePromptLoginUser
	}

	io.WriteString(u.Session, cfmt.Sprint(newUsernamePrompt))
	username, errReadLine := u.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")

		return StatePromptRegisterUser
	}

	logrus.WithField("username", username).Info("Received new username")

	// TODO: check if username is already taken
	for _, bannedName := range w.cfg.BannedNames {
		if strings.EqualFold(username, bannedName) {
			io.WriteString(u.Session, cfmt.Sprintf(registrationUsernameBanned, username))

			return StatePromptRegisterUser
		}
	}

	io.WriteString(u.Session, cfmt.Sprintf(confirmUsernamePrompt, username))
	confirmUsername, errReadLine := u.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")

		return StatePromptRegisterUser
	}

	if !strings.EqualFold(confirmUsername, "y") {
		io.WriteString(u.Session, cfmt.Sprintf(declinedUsernameMsg, username))

		return StatePromptRegisterUser
	}

	password, err := u.Term.ReadPassword(cfmt.Sprint(passwordPrompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading password")

		return StatePromptRegisterUser
	}
	logrus.WithFields(logrus.Fields{"username": username}).Debug("Received password")

	passwordConfirm, err := u.Term.ReadPassword(cfmt.Sprint(passwordConfirmPrompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading confirm password")

		return StatePromptRegisterUser
	}
	logrus.WithFields(logrus.Fields{"username": username}).Debug("Received confirm password")

	if password != passwordConfirm {
		io.WriteString(u.Session, cfmt.Sprintf(passwordMismatchMsg))

		return StatePromptRegisterUser
	}

	// Is the password in the min/max lengths?
	passwordLen := len(password)
	if passwordLen < w.cfg.PasswordMinLength {
		io.WriteString(u.Session, cfmt.Sprintf(passwordMinLengthMsg, w.cfg.PasswordMinLength))

		return StatePromptRegisterUser
	}

	// TODO: make this use config.PasswordMaxLength
	if passwordLen > w.cfg.PasswordMaxLength {
		io.WriteString(u.Session, cfmt.Sprintf(passwordMaxLengthMsg, w.cfg.PasswordMaxLength))

		return StatePromptRegisterUser
	}

	// Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), w.cfg.PasswordBcryptCost)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password")
		return StatePromptRegisterUser
	}

	// Add ID, Username, Password
	u.ID = uuid.New().String()
	u.Username = username
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()

	// Save the character
	if err := u.Save(); err != nil {
		logrus.WithError(err).Error("Error saving user")
		return StatePromptRegisterUser
	}

	io.WriteString(u.Session, cfmt.Sprintf(userCreatedMsg, username))

	return StatePromptMainMenu
}

func (w *World) PromptMainMenu(u *common.User) int {
	io.WriteString(u.Session, cfmt.Sprintf(mainMenuTitle))
	io.WriteString(u.Session, cfmt.Sprintf(menuOptionEnterGame, "Character Name"))
	io.WriteString(u.Session, cfmt.Sprintf(menuOptionCreateCharacter, len(u.Characters), w.cfg.UserCharacterMaxCount))
	io.WriteString(u.Session, cfmt.Sprintf(menuOptionListCharacters))
	io.WriteString(u.Session, cfmt.Sprintf(menuOptionDeleteCharacter))
	io.WriteString(u.Session, cfmt.Sprintf(menuOptionChangePassword))
	io.WriteString(u.Session, cfmt.Sprintf(menuOptionQuit))
	io.WriteString(u.Session, cfmt.Sprint(menuPrompt))

	menuChoice, errReadLine := u.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading menu choice")

		return StatePromptMainMenu
	}

	menuChoice = strings.ToLower(strings.TrimSpace(menuChoice))
	logrus.WithFields(logrus.Fields{"choice": menuChoice}).Info("Received menu choice")

	switch menuChoice {
	case "1":
		return StateEnterGame
	case "2":
		return StatePromptCreateCharacter
	case "3":
		return StatePromptListCharacters
	case "4":
		return StatePromptListCharacters
	case "5":
		return StatePromptChangePassword
	case "0":
		return StateQuit
	default:
		io.WriteString(u.Session, cfmt.Sprintf(menuInvalidChoice, menuChoice))
		return StatePromptMainMenu
	}
}

func (w *World) EnterGame(u *common.User) int {
	if len(u.Characters) == 0 {
		io.WriteString(u.Session, cfmt.Sprintf(noCharactersCreatedMsg))

		return StatePromptCreateCharacter
	}

	i := 1
	for _, c := range u.Characters {
		io.WriteString(u.Session, cfmt.Sprintf(characterListOption, i+1, c.Name))
		i++
	}

	choice, errReadLine := u.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading menu choice")

		return StatePromptCreateCharacter
	}

	choice = strings.ToLower(strings.TrimSpace(choice))
	logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received character choice")

	return StatePromptMainMenu
}

func (w *World) PromptCreateCharacter(u *common.User) int {
	if len(u.Characters) >= w.cfg.UserCharacterMaxCount {
		io.WriteString(u.Session, cfmt.Sprintf("{{You have reached the maximum number of characters allowed.}}::#ff8700\n"))
		return StatePromptMainMenu
	}

	// TODO: Add a blurb about the character creation process
	io.WriteString(u.Session, cfmt.Sprintf(createCharacterMenuTitle))
	io.WriteString(u.Session, cfmt.Sprintf(createCharacterMenuOptionPregen))
	io.WriteString(u.Session, cfmt.Sprintf(createCharacterMenuOptionCustom))
	io.WriteString(u.Session, cfmt.Sprintf(createCharacterMenuOptionLearn))
	io.WriteString(u.Session, cfmt.Sprintf(createCharacterMenuOptionReturn))
	io.WriteString(u.Session, cfmt.Sprint(menuPrompt))

	choice, errReadLine := u.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading menu choice")

		return StatePromptMainMenu
	}

	choice = strings.ToLower(strings.TrimSpace(choice))
	logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received menu choice")

	switch choice {
	case "1":
		// TODO list our pre-generated characters
		return StatePromptCreateCharacter
	case "2":
		// TODO: Prompt for custom character creation
		return StatePromptCreateCharacter
	case "3":
		// TODO: Show longer description about shadowrun characters, archtypes, etc...
		return StatePromptCreateCharacter
	case "4":
		return StatePromptMainMenu
	default:
		io.WriteString(u.Session, cfmt.Sprintf(menuInvalidChoice, choice))
		return StatePromptCreateCharacter
	}

	// io.WriteString(u.Session, cfmt.Sprintf("{{Enter the name of your new character: }}::#ffffff|bold"))
	// name, err := u.Term.ReadLine()
	// if err != nil {
	// 	logrus.WithError(err).Error("Error reading character name")
	// 	return StatePromptCreateCharacter
	// }
	// logrus.WithFields(logrus.Fields{"name": name}).Info("Received character name")

	// TODO: Check that the name is not banned
	// TODO: Check that the name is not already taken
}

func (w *World) PromptListCharacters(u *common.User) int {
	return StatePromptMainMenu
}

func (w *World) PromptDeleteCharacter(u *common.User) int {
	return StatePromptMainMenu
}

func (w *World) PromptChangePassword(u *common.User) int {
	return StatePromptMainMenu
}

func (w *World) MenuCreateCharacter(u *common.User) int {
	return StatePromptMainMenu
}

func (w *World) MOTD(u *common.User) int {
	return StateGameLoop
}

// TODO: Add the AutocompleteCallback
func (w *World) GameLoop(u *common.User) int {
	u.Term.AutoCompleteCallback = w.AutoCompleteCallback

	for {
		io.WriteString(u.Session, cfmt.Sprintf(gameLoopPrompt))
		line, err := u.Term.ReadLine()
		if err != nil {
			logrus.WithError(err).Error("Error reading line")
		}
		line = strings.TrimSpace(line)
		logrus.WithField("line", line).Debug("Received line")
		io.WriteString(u.Session, cfmt.Sprintf(inputEchoMsg, line))
	}
}

// TODO: Match against commands, users, etc...
// TODO: Cycle through the list of available commands when we have more than one
func (w *World) AutoCompleteCallback(line string, pos int, key rune) (string, int, bool) {
	logrus.WithFields(logrus.Fields{"line": line, "pos": pos, "key": key, "key_string": string(key)}).Debug("AutoCompleteCallback")

	return "", pos, false // Return the current result as the auto-completed text
}

func (w *World) GetUser(username string) *common.User {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.users[username]
}

func (w *World) AddUser(u *common.User) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.users[u.Username] = u
}

func (w *World) RemoveUser(username string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	delete(w.users, username)
}

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
