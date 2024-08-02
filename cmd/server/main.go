package main

import (
	"io"
	"net"
	"os"
	"strings"
	"time"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/google/uuid"
	"github.com/i582/cfmt/cmd/cfmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"

	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
)

var (
	// sigChan = make(chan os.Signal, 2)
	usernamePrompt            = "{{Username: }}::#ffffff|bold"
	passwordPrompt            = "{{Password: }}::#ffffff|bold"
	passwordConfirmPrompt     = "{{Confirm password: }}::#ffffff|bold"
	invalidLoginMsg           = "{{You have entered an invalid username or password.}}::#ff8700\n"
	loginSuccessfulMsg        = "{{Login successful.}}::#00ff00\n"
	registrationClosedMsg     = "{{Registration is currently closed.}}::#ff8700\n"
	passwordMismatchMsg       = "{{Passwords do not match.}}::#ff8700\n"
	passwordMinLengthMsg      = "{{Password must be at least %d characters.}}::#ff8700\n"
	passwordMaxLengthMsg      = "{{Password must be at most %d characters.}}::#ff8700\n"
	userCreatedMsg            = "{{User %s has been created.}}::#00ff00\n"
	newUsernamePrompt         = "{{Enter your desired username: }}::#ffffff|bold"
	confirmUsernamePrompt     = "{{Confirm username %s}}::#ffffff|bold {{(y/n)}}::#00ff00|bold{{:}}::#ffffff|bold"
	declinedUsernameMsg       = "{{Username %s was not confirmed.}}::#ff8700\n"
	menuOptionEnterGame       = "{{1.}}::#00ff00 Enter game [%s]\n"
	menuOptionCreateCharacter = "{{2.}}::#00ff00 Create character (%d/%d)\n"
	menuOptionListCharacters  = "{{3.}}::#00ff00 List characters\n"
	menuOptionDeleteCharacter = "{{4.}}::#00ff00 Delete character\n"
	menuOptionChangePassword  = "{{5.}}::#00ff00 Change password\n"
	menuOptionQuit            = "{{0.}}::#00ff00 Quit\n"
	menuPrompt                = "{{Enter the number of the option you would like to select: }}::#ffffff|bold"
	menuInvalidChoice         = "Invalid choice: %s\n"
	gameLoopPrompt            = "{{> }}::#ffffff|bold"
	inputEchoMsg              = "{{You typed:}}::#ffffff|bold %s\n"
)

const (
	ConfigFilepath = "_data/config/server.yaml"
)

const (
	StateBanner = iota
	StatePromptLoginUser
	StatePromptRegisterUser
	StatePromptMainMenu
	StateGameLoop
)

func main() {
	// get config
	var config config.Server
	utils.LoadStructFromYAML(ConfigFilepath, &config)

	logrus.WithField("config", config).Info("Loaded server configuration")

	// set up loggerig
	logrusLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)

	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

	// load data

	// Start the server
	server := &ssh.Server{
		Addr:        net.JoinHostPort(config.Host, config.Port),
		IdleTimeout: config.IdleTimeout,
	}
	defer server.Close()

	ssh.Handle(func(s ssh.Session) {
		logrus.WithFields(logrus.Fields{"user": s.User(), "remote_addr": s.RemoteAddr()}).Info("New connection")

		state := StateBanner

		for {
			switch state {
			case StateBanner:
				state = Banner(s)
			case StatePromptLoginUser:
				state = PromptLoginUser(s)
			case StatePromptRegisterUser:
				state = PromptRegisterUser(s)
			case StatePromptMainMenu:
				state = PromptMainMenu(s)
			case StateGameLoop:
				state = GameLoop(s)
			default:
				logrus.WithField("state", state).Error("Invalid state")
				state = StateBanner
			}
		}
	})

	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Error("Could not start server")
	}

	// Shutdown the server

	// block until a signal comes in
	// <-sigChan

	os.Exit(0)
}

func Banner(s ssh.Session) int {
	io.WriteString(s, cfmt.Sprint(`
{{     ::::::::  :::    :::     :::     :::::::::   ::::::::  :::       ::: ::::    ::::  :::    ::: :::::::::  }}::#ff8700
{{    :+:    :+: :+:    :+:   :+: :+:   :+:    :+: :+:    :+: :+:       :+: +:+:+: :+:+:+ :+:    :+: :+:    :+: }}::#ff5f00
{{    +:+        +:+    +:+  +:+   +:+  +:+    +:+ +:+    +:+ +:+       +:+ +:+ +:+:+ +:+ +:+    +:+ +:+    +:+ }}::#ff0000
{{    +#++:++#++ +#++:++#++ +#++:++#++: +#+    +:+ +#+    +:+ +#+  +:+  +#+ +#+  +:+  +#+ +#+    +:+ +#+    +:+ }}::#d70000
{{           +#+ +#+    +#+ +#+     +#+ +#+    +#+ +#+    +#+ +#+ +#+#+ +#+ +#+       +#+ +#+    +#+ +#+    +#+ }}::#af0000
{{    #+#    #+# #+#    #+# #+#     #+# #+#    #+# #+#    #+#  #+#+# #+#+#  #+#       #+# #+#    #+# #+#    #+# }}::#870000
{{     ########  ###    ### ###     ### #########   ########    ###   ###   ###       ###  ########  #########  }}::#5f0000

`))
	io.WriteString(s, cfmt.Sprintf("Enter your username to continue or type %s\n", "{{new}}::#ffffff|bold"))

	return StatePromptLoginUser
}

func PromptLoginUser(s ssh.Session) int {
	t := term.NewTerminal(s, "")

	// registrationClosedMsg := "{{Registration is currently closed.}}::#ff8700"

	// Collect username
	io.WriteString(s, cfmt.Sprint(usernamePrompt))
	username, errReadLine := t.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")

		return StatePromptLoginUser
	}

	username = strings.TrimSpace(username)
	logrus.WithField("username", username).Debug("Received username")

	// Do we need to make a new user?
	if strings.EqualFold(username, "new") {
		PromptRegisterUser(s)

		return StatePromptRegisterUser
	}

	password, err := t.ReadPassword(cfmt.Sprint(passwordPrompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading password")
		return StatePromptLoginUser
	}
	logrus.WithFields(logrus.Fields{"username": username}).Debug("Received password")

	// TODO: check if username is banned
	// TODO: check if user is banned

	var u common.User
	if err := common.LoadUser(username, &u); err != nil {
		logrus.WithError(err).Error("Error loading user")

		return StatePromptLoginUser
	}

	// Validate the password against the hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		logrus.WithError(err).Error("Password validation error")

		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			io.WriteString(s, cfmt.Sprintf(invalidLoginMsg))
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
		IP:   s.RemoteAddr().String(),
	})

	if err := u.Save(); err != nil {
		logrus.WithError(err).Error("Error saving user")
		return StatePromptLoginUser
	}

	io.WriteString(s, cfmt.Sprintf(loginSuccessfulMsg))

	return StatePromptMainMenu
}

func PromptRegisterUser(s ssh.Session) int {
	t := term.NewTerminal(s, "")

	io.WriteString(s, cfmt.Sprint(newUsernamePrompt))
	username, errReadLine := t.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")

		return StatePromptRegisterUser
	}

	logrus.WithField("username", username).Info("Received new username")

	// TODO: check if username is already taken
	// TODO: check if username is banned

	io.WriteString(s, cfmt.Sprintf(confirmUsernamePrompt, username))
	confirmUsername, errReadLine := t.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")

		return StatePromptRegisterUser
	}

	if !strings.EqualFold(confirmUsername, "y") {
		io.WriteString(s, cfmt.Sprintf(declinedUsernameMsg, username))

		return StatePromptRegisterUser
	}

	password, err := t.ReadPassword(cfmt.Sprint(passwordPrompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading password")
		return StatePromptLoginUser
	}
	logrus.WithFields(logrus.Fields{"username": username}).Debug("Received password")

	passwordConfirm, err := t.ReadPassword(cfmt.Sprint(passwordConfirmPrompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading confirm password")
		return StatePromptLoginUser
	}
	logrus.WithFields(logrus.Fields{"username": username}).Debug("Received confirm password")

	if password != passwordConfirm {
		io.WriteString(s, cfmt.Sprintf(passwordMismatchMsg))

		return StatePromptRegisterUser
	}

	// TODO: make this use config.PasswordMinLength
	if len(password) < 6 {
		io.WriteString(s, cfmt.Sprintf(passwordMinLengthMsg, 6))

		return StatePromptRegisterUser
	}

	// TODO: make this use config.PasswordMaxLength
	if len(password) > 72 {
		io.WriteString(s, cfmt.Sprintf(passwordMaxLengthMsg, 72))

		return StatePromptRegisterUser
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password")
		return StatePromptRegisterUser
	}

	u := common.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
	}

	if err := u.Save(); err != nil {
		logrus.WithError(err).Error("Error saving user")
		return StatePromptRegisterUser
	}

	io.WriteString(s, cfmt.Sprintf(userCreatedMsg, username))

	return StatePromptMainMenu
}

func PromptMainMenu(s ssh.Session) int {
	t := term.NewTerminal(s, "")

	io.WriteString(s, cfmt.Sprintf(menuOptionEnterGame, "Character Name"))
	io.WriteString(s, cfmt.Sprintf(menuOptionCreateCharacter, 1, 3))
	io.WriteString(s, cfmt.Sprintf(menuOptionListCharacters))
	io.WriteString(s, cfmt.Sprintf(menuOptionDeleteCharacter))
	io.WriteString(s, cfmt.Sprintf(menuOptionChangePassword))
	io.WriteString(s, cfmt.Sprintf(menuOptionQuit))
	io.WriteString(s, cfmt.Sprint(menuPrompt))

	menuChoice, errReadLine := t.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading menu choice")

		return StatePromptRegisterUser
	}

	menuChoice = strings.ToLower(strings.TrimSpace(menuChoice))
	logrus.WithFields(logrus.Fields{"choice": menuChoice}).Info("Received menu choice")

	switch menuChoice {
	case "1":
		return StateGameLoop
	case "2":
		// Create character
		return StatePromptMainMenu
	case "3":
		// List characters
		return StatePromptMainMenu
	case "4":
		// Delete character
		return StatePromptMainMenu
	case "5":
		// Change password
		return StatePromptMainMenu
	case "0":
		// Quit
		return StatePromptMainMenu
	default:
		io.WriteString(s, cfmt.Sprintf(menuInvalidChoice, menuChoice))
		return StatePromptMainMenu
	}
}

func GameLoop(s ssh.Session) int {
	t := term.NewTerminal(s, "")

	for {
		io.WriteString(s, cfmt.Sprintf(gameLoopPrompt))
		line, err := t.ReadLine()
		if err != nil {
			logrus.WithError(err).Error("Error reading line")
		}
		line = strings.TrimSpace(line)
		logrus.WithField("line", line).Debug("Received line")
		io.WriteString(s, cfmt.Sprintf(inputEchoMsg, line))
	}
}
