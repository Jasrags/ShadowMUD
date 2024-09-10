package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/Jasrags/ShadowMUD/common/character"
	"github.com/Jasrags/ShadowMUD/common/user"
	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/gliderlabs/ssh"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

// banner displays the banner for the game
func (w *World) banner(s ssh.Session) State {
	logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "package": "main", "screen": "banner"}).Debug("Displaying banner")
	io.WriteString(s, cfmt.Sprint(`
{{     ::::::::  :::    :::     :::     :::::::::   ::::::::  :::       ::: ::::    ::::  :::    ::: :::::::::  }}::#ff8700
{{    :+:    :+: :+:    :+:   :+: :+:   :+:    :+: :+:    :+: :+:       :+: +:+:+: :+:+:+ :+:    :+: :+:    :+: }}::#ff5f00
{{    +:+        +:+    +:+  +:+   +:+  +:+    +:+ +:+    +:+ +:+       +:+ +:+ +:+:+ +:+ +:+    +:+ +:+    +:+ }}::#ff0000
{{    +#++:++#++ +#++:++#++ +#++:++#++: +#+    +:+ +#+    +:+ +#+  +:+  +#+ +#+  +:+  +#+ +#+    +:+ +#+    +:+ }}::#d70000
{{           +#+ +#+    +#+ +#+     +#+ +#+    +#+ +#+    +#+ +#+ +#+#+ +#+ +#+       +#+ +#+    +#+ +#+    +#+ }}::#af0000
{{    #+#    #+# #+#    #+# #+#     #+# #+#    #+# #+#    #+#  #+#+# #+#+#  #+#       #+# #+#    #+# #+#    #+# }}::#870000
{{     ########  ###    ### ###     ### #########   ########    ###   ###   ###       ###  ########  #########  }}::#5f0000

Enter your username to continue or type {{new}}::#ffffff|bold
`))

	return StateLoginUser
}

func (w *World) promptLoginUser(s ssh.Session) (State, *user.User) {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "package": "main", "screen": "login_user"})
	l.Debug("Prompting for login")
promptUsername:
	// Collect new username
	username, errUsername := utils.PromptUserInput(s, usernamePrompt)
	if errUsername != nil {
		l.WithError(errUsername).Error("Error reading username")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		return StateQuit, nil
	}

	if username == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		goto promptUsername
	}

	username = strings.TrimSpace(username)
	l.WithField("username", username).Debug("Received username")

	// Do we need to make a new user?
	if strings.EqualFold(username, "new") {
		return StateRegisterUser, nil
	}

promptPassword:
	// Collect new password
	password, errPassword := utils.PromptUserPasswordInput(s, passwordPrompt)
	if errPassword != nil {
		l.WithError(errPassword).Error("Error reading password")
		return StateQuit, nil
	}

	if password == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		goto promptPassword
	}

	l.WithFields(logrus.Fields{"username": username, "password": password}).Debug("Received password")

	// is login enabled?
	if !w.cfg.LoginEnabled {
		l.WithField("username", username).Warn("Login is disabled")
		io.WriteString(s, cfmt.Sprint(loginClosedMsg))
		goto promptUsername
	}

	// TODO: Load the actual user from the file
	u := w.userManager.GetByUsername(username)
	if u == nil {
		l.WithField("username", username).Error("User not found")
		io.WriteString(s, cfmt.Sprintf(invalidLoginMsg))
		goto promptUsername
	}
	// u := user.New()
	// u.Load()
	// u.Password = "$2a$10$hIWeecxoHtzuECf6z39mSetPQrTalFxm82xMj.RjU3sxMW/b5d11O"
	// Try the user from the file
	// if err := user.LoadUser(username, &u); err != nil {
	// 	logrus.WithError(err).Error("Error loading user")
	// 	return StatePromptLoginUser
	// }

	// logrus.WithFields(logrus.Fields{"username": s.user.Username, "id": s.user.ID}).Debug("Loaded user")

	// Is the user banned?
	for _, ban := range u.Bans {
		if ban.ExpiresAt.After(time.Now()) {
			l.WithFields(logrus.Fields{"username": username, "expires_at": ban.ExpiresAt}).Warn("User is banned")
			io.WriteString(s, cfmt.Sprintf(loginUserBannedMsg, ban.ExpiresAt.Format(time.RFC1123)))

			return StateLoginUser, nil
		}
	}

	// Validate the password against the hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		logrus.WithError(err).Error("Password validation error")

		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			io.WriteString(s, cfmt.Sprintf(invalidLoginMsg))
			l.WithFields(logrus.Fields{"username": username}).WithError(err).Error("Invalid password for user")
		case bcrypt.ErrHashTooShort:
		case bcrypt.ErrMismatchedHashAndPassword:
		case bcrypt.ErrPasswordTooLong:
		default:
			l.WithFields(logrus.Fields{"username": username}).WithError(err).Error("Password error")
		}

		goto promptUsername
	}

	// Add a login record
	u.Logins = append(u.Logins, user.Login{
		Time: time.Now(),
		IP:   s.RemoteAddr().String(),
	})

	// // Save the user
	// if err := u.Save(); err != nil {
	// 	logrus.WithError(err).Error("Error saving user")
	// 	return StatePromptLoginUser, nil
	// }

	// Add the user to the world users
	w.AddUser(u)

	io.WriteString(s, cfmt.Sprintf(loginSuccessfulMsg))

	return StateMainMenu, u
}

func (w *World) promptRegisterUser(s ssh.Session) (State, *user.User) {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "package": "main", "screen": "register_user"})
	l.Debug("Prompting for registration")

promptUsername:
	// Is registration enabled?
	if !w.cfg.RegistrationEnabled {
		l.Warn("Registration is disabled")
		io.WriteString(s, cfmt.Sprint(registrationClosedMsg))
		return StateBanner, nil
	}

	// Collect new username
	username, errUsername := utils.PromptUserInput(s, usernameNewPrompt)
	if errUsername != nil {
		l.WithError(errUsername).Error("Error reading username")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		goto promptUsername
	}

	if username == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		goto promptUsername
	}

	if len(username) < w.cfg.UsernameMinLength || len(username) > w.cfg.UsernameMaxLength {
		io.WriteString(s, cfmt.Sprintf(usernameMixMaxLengthMsg, w.cfg.UsernameMinLength, w.cfg.UsernameMaxLength))
		goto promptUsername
	}

	l.WithField("username", username).Debug("Received username")

	// Check if the username is banned
	for _, ban := range w.cfg.BannedNames {
		if strings.EqualFold(username, ban) {
			l.WithField("username", username).Warn("Username is banned")
			io.WriteString(s, cfmt.Sprintf(usernameBannedMsg, username))
			goto promptUsername
		}
	}

	confirmed, errUsernameConfirm := utils.PromptConfirmInput(s, cfmt.Sprintf(usernameConfirmPrompt, username))
	if errUsernameConfirm != nil {
		l.WithError(errUsernameConfirm).Error("Error reading confirm username")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		goto promptUsername
	}

	if !confirmed {
		io.WriteString(s, cfmt.Sprintf(usernameDeclinedMsg, username))
		goto promptUsername
	}

	l.WithField("username", username).Debug("Confirmed username")

promptPassword:
	// Collect new password
	password, errPassword := utils.PromptUserPasswordInput(s, passwordNewPrompt)
	if errPassword != nil {
		l.WithError(errPassword).Error("Error reading password")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		goto promptPassword
	}

	if password == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		goto promptPassword
	}

	l.WithFields(logrus.Fields{"username": username, "password": password}).Debug("Received password")

	// Is the password in the min/max lengths?
	if len(password) < w.cfg.PasswordMinLength || len(password) > w.cfg.PasswordMaxLength {
		io.WriteString(s, cfmt.Sprintf(passwordMinMaxLengthMsg, w.cfg.PasswordMinLength, w.cfg.PasswordMaxLength))
		goto promptPassword
	}

	// Confirm the password
	passwordConfirm, errPasswordConfirm := utils.PromptUserPasswordInput(s, passwordConfirmPrompt)
	if errPasswordConfirm != nil {
		l.WithError(errPasswordConfirm).Error("Error reading confirm password")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		return StateQuit, nil
	}

	if passwordConfirm == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		goto promptPassword
	}

	l.WithFields(logrus.Fields{"username": username, "password": password}).Debug("Received confirm password")

	// Do the passwords match?
	if password != passwordConfirm {
		io.WriteString(s, cfmt.Sprintf(passwordMismatchMsg))
		goto promptPassword
	}

	// Hash the password with bcrypt
	hashedPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(password), w.cfg.PasswordBcryptCost)
	if errHashPassword != nil {
		l.WithError(errHashPassword).Error("Error hashing password")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		return StateQuit, nil
	}

	l.WithFields(logrus.Fields{"username": username, "password": string(hashedPassword)}).Debug("Created password hash")

	// Create the user
	u := user.New()
	u.Username = username
	u.Password = string(hashedPassword)
	u.AddUserRoles(user.RoleUser)

	// Save the user
	if err := u.Save(); err != nil {
		l.WithError(err).Error("Error saving user")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		return StateQuit, nil
	}

	// Add the user to the world users
	w.AddUser(u)
	w.userManager.Add(u)

	io.WriteString(s, cfmt.Sprintf(userCreatedMsg, username))

	return StateMainMenu, u
}

func (w *World) promptMainMenu(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "package": "main", "user_id": u.ID, "username": u.Username, "screen": "main_menu"})
	l.Debug("Prompting for main menu")

promptMainMenu:
	io.WriteString(s, cfmt.Sprintf(mainMenuTitle))
	io.WriteString(s, cfmt.Sprintf(menuOptionEnterGame))
	io.WriteString(s, cfmt.Sprintf(menuOptionCreateCharacter, len(u.Characters), w.cfg.UserCharacterMaxCount))
	io.WriteString(s, cfmt.Sprintf(menuOptionListCharacters))
	io.WriteString(s, cfmt.Sprintf(menuOptionDeleteCharacter))
	io.WriteString(s, cfmt.Sprintf(menuOptionChangePassword))
	io.WriteString(s, cfmt.Sprintf(menuOptionQuit))

	t := term.NewTerminal(s, cfmt.Sprint(menuPrompt))
	menuChoice, errReadLine := t.ReadLine()
	if errReadLine != nil {
		l.WithError(errReadLine).Error("Error reading menu choice")

		goto promptMainMenu
	}

	menuChoice = strings.ToLower(strings.TrimSpace(menuChoice))
	l.WithFields(logrus.Fields{"choice": menuChoice}).Info("Received menu choice")

	switch menuChoice {
	case "1":
		return StateEnterGame
	case "2":
		return StatePromptCreateCharacterMenu
	case "3":
		return StatePromptListCharacters
	case "4":
		return StatePromptDeleteCharacter
	case "5":
		return StatePromptChangePassword
	case "0":
		return StateQuit
	default:
		io.WriteString(s, cfmt.Sprintf(menuInvalidChoice, menuChoice))
		goto promptMainMenu
	}
}

func (w *World) enterGame(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "enter_game"})
	l.Debug("Entering game")

	// If no characters have been created, prompt the user to create one
	if len(u.Characters) == 0 {
		l.Debug("No characters created")
		io.WriteString(s, cfmt.Sprintf(noCharactersCreatedMsg))
		return StatePromptCreateCharacterMenu
	}

	// If only one character has been created, set it as the active character
	if len(u.Characters) == 1 {
		l.Debug("One character created")
		u.SetActiveCharacter(u.GetCharacterByID(u.CharacterIDs[0]))
		return StateGameLoop
	}

promptChooseCharacter:
	// Otherwise, prompt the user to choose a character
	io.WriteString(s, cfmt.Sprintf(chooseCharacterPrompt))

	i := 0
	choiceIdMap := make(map[string]string)
	for _, c := range u.Characters {
		choiceIdMap[strconv.Itoa(i+1)] = c.ID
		io.WriteString(s, cfmt.Sprintf(characterListOption, i+1, c.Name))
		i++
	}

	for k, v := range choiceIdMap {
		logrus.WithFields(logrus.Fields{"key": k, "value": v}).Info("Choice ID Map")
	}

	t := term.NewTerminal(s, cfmt.Sprint(menuPrompt))
	choice, errReadLine := t.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading menu choice")
		goto promptChooseCharacter
	}

	if choice == "" {
		goto promptChooseCharacter
	}

	choice = strings.ToLower(strings.TrimSpace(choice))
	logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received character choice")

	c := u.GetCharacterByID(choiceIdMap[choice])
	u.SetActiveCharacter(c)

	// // var r *common.Room
	// if u.Character.RoomID == "" {
	// 	logrus.Warn("Character has no room, adding to default room")

	// 	// r := common.NewRoom(&common.CoreRooms[0])
	// 	// r.AddCharacter(c)
	// }

	// r.AddCharacter(c)

	// u.ActiveCharacter = c

	// i := 1
	// for _, c := range s.user.Characters {
	// 	io.WriteString(s, cfmt.Sprintf(characterListOption, i+1, c.Name))
	// 	i++
	// }

	// choice, errReadLine := u.Term.ReadLine()
	// if errReadLine != nil {
	// 	logrus.WithError(errReadLine).Error("Error reading menu choice")

	// 	return StateQuit
	// }

	// choice = strings.ToLower(strings.TrimSpace(choice))
	// logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received character choice")

	// return StatePromptMainMenu
	return StateMOTD
}

func (w *World) promptMOTD(s ssh.Session) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "package": "main", "screen": "motd"})
	l.Debug("Prompting for MOTD")
	io.WriteString(s, cfmt.Sprint(featureNotImplementedMsg))
	return StateEnterGame
}

// func (w *World) promptCreateCharacter(s ssh.Session, u *user.User) State {
// 	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "create_character"})
// 	l.Debug("Prompting for character creation")

// 	// promptCreateCharacterMenu:
// 	// If the user has reached the maximum number of characters, return to the main menu
// 	if len(u.Characters) >= w.cfg.UserCharacterMaxCount {
// 		io.WriteString(s, cfmt.Sprintf(characterMaxCharactersMsg))
// 		return StateMainMenu
// 	}

// state := StateCreateCharacterMenu
// for {
// 	switch state {
// 	case StateCreateCharacterMenu:
// 		state = w.promptCreateCharacterMenu(s, u)
// 	case StateCreateCharacterName:
// 		state = w.promptCreateCharacterName(s, u)
// 	case StateCreateCharacterMetatype:
// 		// state = w.promptCreateCharacterMetatype(s, u)
// 	case StateCreateCharacterMagic:
// 		// state = w.promptCreateCharacterMagic(s, u)
// 	case StateCreateCharacterAttributes:
// 		// state = w.promptCreateCharacterAttributes(s, u)
// 	case StateCreateCharacterSkills:
// 		// state = w.promptCreateCharacterSkills(s, u)
// 	case StateCreateCharacterQualities:
// 		// state = w.promptCreateCharacterQualities(s, u)
// 	case StateCreateCharacterNuyen:
// 		// state = w.promptCreateCharacterNuyen(s, u)
// 	case StateCreateCharacterSpend:
// 		// state = w.promptCreateCharacterSpend(s, u)
// 	case StateCreateCharacterQuit:
// 		return StateMainMenu
// 	default:
// 		io.WriteString(s, cfmt.Sprint(featureNotImplementedMsg))
// 		return StateMainMenu
// 	}
// }

// // Prompt the user to choose a character creation method
// io.WriteString(s, cfmt.Sprintf(createCharacterMenuTitle))
// io.WriteString(s, cfmt.Sprintf(createCharacterMenuOptionPregen))
// io.WriteString(s, cfmt.Sprintf(createCharacterMenuOptionCustom))
// // io.WriteString(s, cfmt.Sprintf(createCharacterMenuOptionLearn))
// // io.WriteString(s, cfmt.Sprintf(createCharacterMenuOptionReturn))
// t := term.NewTerminal(s, cfmt.Sprint(menuPrompt))
// choice, errReadLine := t.ReadLine()
// if errReadLine != nil {
// 	l.WithError(errReadLine).Error("Error reading menu choice")
// 	goto promptCreateCharacterMenu
// }

// choice = strings.ToLower(strings.TrimSpace(choice))
// l.WithFields(logrus.Fields{"choice": choice}).Info("Received character choice")

// switch choice {
// case "1":
// 	return StatePromptCreatePregenCharacter
// case "2":
// 	return StatePromptCreateCustomCharacter
// // case "3":
// // return StatePromptCreateCharacterLearn
// // case "4":
// default:
// 	io.WriteString(s, cfmt.Sprintf(menuInvalidChoice, choice))
// 	goto promptCreateCharacterMenu
// }
// }

// type BuildCommand struct {
// 	s ssh.Session
// 	u *user.User
// 	w *World
// }

// func NewBuildCommand(s ssh.Session, u *user.User, w *World) *BuildCommand {
// 	return &BuildCommand{s: s, u: u, w: w}
// }

// // CommandHandler type
// type BuildCommandHandler func(args []string) string

// // Main function to parse and execute commands
// func (c BuildCommand) parseAndExecuteCommand(input string) string {
// 	var commands = map[string]BuildCommandHandler{
// 		"help": c.handleHelpCommand,
// 		"set":  c.handleSetCommand,
// 		"show": c.handleShowCommand,
// 		"list": c.handleListCommand,
// 	}

// 	parts := strings.Fields(input)
// 	if len(parts) == 0 {
// 		return "No command provided"
// 	}

// 	command := parts[0]
// 	args := parts[1:]

// 	if handler, found := commands[command]; found {
// 		return handler(args)
// 	}

// 	return cfmt.Sprintf("Unknown command: %s", command)
// }

// // Handler for 'help' command
// func (c BuildCommand) handleHelpCommand(args []string) string {
// 	if err := c.w.templates.ExecuteTemplate(c.s, "character_builder_help.tmpl", nil); err != nil {
// 		logrus.WithError(err).Error("Error executing template")
// 		io.WriteString(c.s, cfmt.Sprintf("An error occurred while displaying help"))

// 		return ""
// 	}

// 	if err := utils.PromptPressEnterInput(c.s); err != nil {
// 		logrus.WithError(err).Error("Error reading input")
// 	}

// 	return ""
// }

// // Handler for 'set' command
// func (c BuildCommand) handleSetCommand(args []string) string {
// 	if len(args) < 2 {
// 		return "Usage: set <attribute> <value>"
// 	}

// 	attribute := args[0]
// 	value := strings.Join(args[1:], " ")

// 	switch attribute {
// 	case "name":

// 		// 			// Check if the name is between the min and max lengths
// 		// 			if len(name) < w.cfg.CharacterNameMinLength || len(name) > w.cfg.CharacterNameMaxLength {
// 		// 				io.WriteString(s, cfmt.Sprintf(characterNameMixMaxLengthMsg, w.cfg.CharacterNameMinLength, w.cfg.CharacterNameMaxLength))
// 		// 				goto loop
// 		// 			}

// 		// 			// TODO: Check if the name is already taken

// 		// 			// Check if the name is banned
// 		// 			for _, ban := range w.cfg.BannedNames {
// 		// 				if strings.EqualFold(name, ban) {
// 		// 					io.WriteString(s, cfmt.Sprintf(usernameBannedMsg, name))
// 		// 					goto loop
// 		// 				}
// 		// 			}

// 		// 			// Check if the name contains only alphabetic characters
// 		// 			if !regexp.MustCompile("^[a-zA-Z]+$").MatchString(name) {
// 		// 				io.WriteString(s, cfmt.Sprintf("{{Invalid character name. Only alphabetic characters are allowed.}}::#ff0000\n"))
// 		// 				goto loop
// 		// 			}

// 		// c.Name = name

// 		// Set the character's name
// 		// Example: character.Name = value
// 		return cfmt.Sprintf("Name set to '%s'", value)
// 	case "metatype":
// 		// Set the character's metatype
// 		// Example: character.Metatype = value
// 		return cfmt.Sprintf("Metatype set to '%s'", value)
// 	default:
// 		return cfmt.Sprintf("{{Unknown attribute:}}::#ff8700 '%s'", attribute)
// 	}
// }

// // Handler for 'show' command
// func (c BuildCommand) handleShowCommand(args []string) string {
// 	if len(args) < 1 {
// 		return "Usage: show <item>"
// 	}

// 	item := args[0]

// 	switch item {
// 	case "metatype":
// 		// Show available metatypes
// 		// Example: return list of metatypes
// 		return "Available metatypes: Human, Elf, Dwarf, Ork, Troll"
// 	default:
// 		return cfmt.Sprintf("Unknown item: %s", item)
// 	}
// }

// // Handler for 'list' command
// func (c BuildCommand) handleListCommand(args []string) string {
// 	if len(args) < 1 {
// 		return "Usage: list <item>"
// 	}

// 	item := args[0]

// 	switch item {
// 	case "metatypes":
// 		// Show available metatypes
// 		// Example: return list of metatypes
// 		return "Available metatypes: Human, Elf, Dwarf, Ork, Troll"
// 	default:
// 		return cfmt.Sprintf("{{Unknown item:}}::#ff8700 '%s'\n", item)
// 	}
// }

func (w *World) promptCreateCharacterMenu(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "create_character_menu"})
	l.Debug("Prompting for character creation menu")

	if len(u.Characters) >= w.cfg.UserCharacterMaxCount {
		io.WriteString(s, cfmt.Sprintf(characterMaxCharactersMsg))
		return StateMainMenu
	}

	// Create the new character
	c := character.New(w.cfg)
	c.State = character.StateIncomplete

	io.WriteString(s, cfmt.Sprintf("\n{{Character Creation}}::#00ff00|bold\n"))

loop:
	for {
		if err := w.templates.ExecuteTemplate(s, "character_sheet.tmpl", c); err != nil {
			logrus.WithError(err).Error("Error executing template")
			io.WriteString(s, cfmt.Sprintf("An error occurred while displaying help"))
			break
		}

		// Prompt
		input, err := utils.PromptUserInput(s, gameLoopPrompt)
		if err != nil {
			l.WithError(err).Error("Error reading input")
			break
		}

		// input = strings.TrimSpace(input)
		args := strings.Fields(strings.TrimSpace(input))
		command := strings.ToLower(args[0])
		args = args[1:]
		// command := strings.ToLower(fields[0])
		// args := fields[1:]
		l.WithFields(logrus.Fields{"input": input, "command": command, "args": args}).Debug("Received input")

		if input == "" {
			goto loop
		}

		switch command {
		case "help":
			if err := w.templates.ExecuteTemplate(s, "character_builder_help.tmpl", nil); err != nil {
				logrus.WithError(err).Error("Error executing template")
				io.WriteString(s, cfmt.Sprintf("An error occurred while displaying help"))
				break
			}
		case "save":
			io.WriteString(s, featureNotImplementedMsg)
			io.WriteString(s, "Character saved.\n")
		case "finalize":
			// Implement finalize logic
			io.WriteString(s, featureNotImplementedMsg)
			io.WriteString(s, "Character finalized.\n")
		case "discard":
			// Implement discard logic
			io.WriteString(s, featureNotImplementedMsg)
			io.WriteString(s, "Character discarded.\n")
			return StateMainMenu
		case "restart":
			// Implement restart logic
			io.WriteString(s, "Character restarted.\n")
		case "exit":
			io.WriteString(s, featureNotImplementedMsg)
			io.WriteString(s, "Exiting character builder.\n")
			return StateMainMenu
		case "set":
			if len(args) < 2 {
				io.WriteString(s, "Usage: set <property> <value>\n")
				continue
			}
			property := strings.ToLower(args[0])
			value := strings.Join(args[1:], " ")
			logrus.WithFields(logrus.Fields{"property": property, "value": value}).Debug("Setting property")
			switch property {
			case "name":
				if err := c.SetName(value); err != nil {
					logrus.WithError(err).Error("Error setting character name")
				}
				io.WriteString(s, "Character name set to "+value+"\n")
			case "metatype":
				// Implement set metatype logic
				io.WriteString(s, featureNotImplementedMsg)
				// io.WriteString(s, "Character metatype set to "+value+"\n")
			case "magic":
				// Implement set magic logic
				io.WriteString(s, featureNotImplementedMsg)
				// io.WriteString(s, "Character magic type set to "+value+"\n")
			case "attribute":
				if len(args) < 4 {
					io.WriteString(s, "Usage: set attribute <attribute> <value>\n")
					continue
				}
				io.WriteString(s, featureNotImplementedMsg)
				attribute := args[2]
				attributeValue := args[3]
				// Implement set attribute logic
				io.WriteString(s, "Character attribute "+attribute+" set to "+attributeValue+"\n")
			case "skill":
				if len(args) < 4 {
					io.WriteString(s, "Usage: set skill <skill> <value>\n")
					continue
				}
				io.WriteString(s, featureNotImplementedMsg)
				skill := args[2]
				skillValue := args[3]
				// Implement set skill logic
				io.WriteString(s, "Character skill "+skill+" set to "+skillValue+"\n")
			case "quality":
				// Implement set quality logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Character quality set to "+value+"\n")
			case "nuyen":
				// Implement set nuyen logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Character nuyen set to "+value+"\n")
			default:
				io.WriteString(s, "Unknown property: "+property+"\n")
			}
		case "list":
			if len(args) < 2 {
				io.WriteString(s, "Usage: list <item>\n")
				continue
			}
			item := args[1]
			switch item {
			case "metatypes":
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Available metatypes: Human, Elf, Dwarf, Ork, Troll\n")
			case "magic":
				// Implement list magic logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Available magic types: Adept, Magician, Mystic Adept\n")
			case "attributes":
				// Implement list attributes logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Available attributes: Strength, Agility, Willpower, Logic, Charisma\n")
			case "skills":
				// Implement list skills logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Available skills: Hacking, Shooting, Negotiation, Stealth\n")
			case "qualities":
				// Implement list qualities logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Available qualities: Toughness, Quick Healer, Analytical Mind\n")
			default:
				io.WriteString(s, "Unknown item: "+item+"\n")
			}
		case "show":
			if len(args) < 3 {
				io.WriteString(s, "Usage: show <item> <name>\n")
				continue
			}
			item := args[1]
			name := args[2]
			switch item {
			case "metatype":
				// Implement show metatype logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Details of metatype: "+name+"\n")
			case "magic":
				// Implement show magic logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Details of magic type: "+name+"\n")
			case "attribute":
				// Implement show attribute logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Details of attribute: "+name+"\n")
			case "skill":
				// Implement show skill logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Details of skill: "+name+"\n")
			case "quality":
				// Implement show quality logic
				io.WriteString(s, featureNotImplementedMsg)
				io.WriteString(s, "Details of quality: "+name+"\n")
			default:
				io.WriteString(s, "Unknown item: "+item+"\n")
			}
		default:
			io.WriteString(s, "Unknown command: "+command+"\n")
		}

	}
	return StateMainMenu
	// input = strings.ToLower(input)
	// inputFields := strings.Fields(input)
	// cmd := inputFields[0]
	// 	switch cmd {
	// 	case "help":
	// 		io.WriteString(s, cfmt.Sprint(buildHelp))
	// 	// case "save":
	// 	// case "finalize":
	// 	// case "discard":
	// 	// case "restart":
	// 	case "exit":
	// 		return StateMainMenu
	// 	case "set":
	// 		if len(inputFields) < 2 {
	// 			io.WriteString(s, cfmt.Sprintf("{{Usage: set <field> <value>}}::#ff8700\n"))
	// 			goto loop
	// 		}
	// 		switch inputFields[1] {
	// 		case "name":
	// 			name := inputFields[2]

	// 			// Check if the name is between the min and max lengths
	// 			if len(name) < w.cfg.CharacterNameMinLength || len(name) > w.cfg.CharacterNameMaxLength {
	// 				io.WriteString(s, cfmt.Sprintf(characterNameMixMaxLengthMsg, w.cfg.CharacterNameMinLength, w.cfg.CharacterNameMaxLength))
	// 				goto loop
	// 			}

	// 			// TODO: Check if the name is already taken

	// 			// Check if the name is banned
	// 			for _, ban := range w.cfg.BannedNames {
	// 				if strings.EqualFold(name, ban) {
	// 					io.WriteString(s, cfmt.Sprintf(usernameBannedMsg, name))
	// 					goto loop
	// 				}
	// 			}

	// 			// Check if the name contains only alphabetic characters
	// 			if !regexp.MustCompile("^[a-zA-Z]+$").MatchString(name) {
	// 				io.WriteString(s, cfmt.Sprintf("{{Invalid character name. Only alphabetic characters are allowed.}}::#ff0000\n"))
	// 				goto loop
	// 			}

	// 			c.Name = name

	// 			io.WriteString(s, cfmt.Sprintf("{{Character name set to '%s'}}::#00ff00\n", name))
	// 		case "metatype":
	// 		case "magic":
	// 		case "attributes":
	// 		case "skills":
	// 		case "qualities":
	// 		}
	// 	case "list":
	// 		switch inputFields[1] {
	// 		case "metatypes":
	// 			io.WriteString(s, cfmt.Sprintf("{{Metatypes}}::#00ff00\n"))
	// 			for _, m := range metatype.CoreMetatypes {
	// 				io.WriteString(s, cfmt.Sprintf(buildMetatypeInfo,
	// 					m.Name, m.Category, m.Description,
	// 					m.Attributes.Body.Min, m.Attributes.Body.Max, m.Attributes.Body.AugMax,
	// 					m.Attributes.Agility.Min, m.Attributes.Agility.Max, m.Attributes.Agility.AugMax,
	// 					m.Attributes.Reaction.Min, m.Attributes.Reaction.Max, m.Attributes.Reaction.AugMax,
	// 					m.Attributes.Strength.Min, m.Attributes.Strength.Max, m.Attributes.Strength.AugMax,
	// 					m.Attributes.Willpower.Min, m.Attributes.Willpower.Max, m.Attributes.Willpower.AugMax,
	// 					m.Attributes.Logic.Min, m.Attributes.Logic.Max, m.Attributes.Logic.AugMax,
	// 					m.Attributes.Intuition.Min, m.Attributes.Intuition.Max, m.Attributes.Intuition.AugMax,
	// 					m.Attributes.Charisma.Min, m.Attributes.Charisma.Max, m.Attributes.Charisma.AugMax,
	// 					m.Attributes.Edge.Min, m.Attributes.Edge.Max, m.Attributes.Edge.AugMax,
	// 					m.Attributes.Initiative.Min, m.Attributes.Initiative.Max, m.Attributes.Initiative.AugMax,
	// 					m.Attributes.Essence.Min, m.Attributes.Essence.Max, m.Attributes.Essence.AugMax,
	// 					m.Attributes.Magic.Min, m.Attributes.Magic.Max, m.Attributes.Magic.AugMax,
	// 					m.Attributes.Resonance.Min, m.Attributes.Resonance.Max, m.Attributes.Resonance.AugMax,
	// 					m.Qualities, m.QualityRestrictions))
	// 			}
	// 		case "magic":
	// 		case "attributes":
	// 		case "skills":
	// 		case "qualities":
	// 		}
	// 	case "show":
	// 		switch inputFields[1] {
	// 		case "metatype":
	// 		case "magic":
	// 		case "attribute":
	// 		case "skill":
	// 		case "qualitie":
	// 		}
	// 	}
	// }

	// return StateMainMenu

	// 	type Option struct {
	// 		Text  string
	// 		State State
	// 	}

	// 	options := map[string]Option{
	// 		"1": {Text: "Set your character's name", State: StateCreateCharacterName},
	// 		"0": {Text: "Exit character creation", State: StateMainMenu},
	// 	}

	// promptCreateCharacterMenu:
	// io.WriteString(s, cfmt.Sprintf(\n{{Character Creation}}::#00ff00|bold\n))
	// 	for k, o := range options {
	// 		io.WriteString(s, cfmt.Sprintf("{{%s.}}::#00ff00 %s\n", k, o.Text))
	// 	}
	// 	// io.WriteString(s, cfmt.Sprintf("1. Set your character's name\n"))
	// 	// io.WriteString(s, cfmt.Sprintf("0. Exit character creation\n"))

	// 	menuChoice, errReadLine := utils.PromptUserInput(s, menuPrompt)
	// 	if errReadLine != nil {
	// 		l.WithError(errReadLine).Error("Error reading menu choice")

	// 		goto promptCreateCharacterMenu
	// 	}

	// 	menuChoice = strings.ToLower(strings.TrimSpace(menuChoice))
	// 	l.WithFields(logrus.Fields{"choice": menuChoice}).Info("Received menu choice")

	// 	if o, ok := options[menuChoice]; ok {
	// 		return o.State
	// 	} else {
	// 		io.WriteString(s, cfmt.Sprintf(menuInvalidChoice, menuChoice))
	// 		goto promptCreateCharacterMenu
	// 	}
	// 	// switch menuChoice {
	// 	// case "1":
	// 	// 	return StateCreateCharacterName
	// 	// case "0":
	// 	// 	return StateMainMenu
	// 	// default:
	// 	// 	io.WriteString(s, cfmt.Sprintf(menuInvalidChoice, menuChoice))
	// 	// 	goto promptCreateCharacterMenu
	// 	// }

	return StateMainMenu
}

func (w *World) promptCreateCharacterName(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "create_character_name"})
	l.Debug("Prompting for character name")

promptCharacterName:
	name, errUsername := utils.PromptUserInput(s, characterNamePrompt)
	if errUsername != nil {
		l.WithError(errUsername).Error("Error reading character name")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		goto promptCharacterName
	}

	if name == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		goto promptCharacterName
	}

	if len(name) < w.cfg.CharacterNameMinLength || len(name) > w.cfg.CharacterNameMaxLength {
		io.WriteString(s, cfmt.Sprintf(characterNameMixMaxLengthMsg, w.cfg.CharacterNameMinLength, w.cfg.CharacterNameMaxLength))
		goto promptCharacterName
	}

	// Check if the character name is already taken

	confirm, errConfirmName := utils.PromptConfirmInput(s, cfmt.Sprintf(createCharacterConfirmNamePrompt, name))
	if errConfirmName != nil {
		l.WithError(errConfirmName).Error("Error reading confirm character name")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		goto promptCharacterName
	}

	if !confirm {
		io.WriteString(s, cfmt.Sprintf(characterNameDeclinedMsg, name))
		goto promptCharacterName
	}

	l.WithField("name", name).Debug("Received character name")

	return StatePromptCreateCharacterMenu
}

func (w *World) promptCreatePregenCharacter(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "create_pregen_character"})
	l.Debug("Prompting for pre-generated character creation")
	io.WriteString(s, cfmt.Sprint(featureNotImplementedMsg))

	return StateMainMenu
}

func (w *World) promptCreateCustomCharacter(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "create_custom_character"})
	l.Debug("Prompting for custom character creation")
promptCharacterName:
	name, errUsername := utils.PromptUserInput(s, characterNamePrompt)
	if errUsername != nil {
		l.WithError(errUsername).Error("Error reading character name")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		goto promptCharacterName
	}

	if name == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		goto promptCharacterName
	}

	if len(name) < w.cfg.CharacterNameMinLength || len(name) > w.cfg.CharacterNameMaxLength {
		io.WriteString(s, cfmt.Sprintf(characterNameMixMaxLengthMsg, w.cfg.CharacterNameMinLength, w.cfg.CharacterNameMaxLength))
		goto promptCharacterName
	}

	// Check if the character name is already taken

	confirm, errConfirmName := utils.PromptConfirmInput(s, cfmt.Sprintf(createCharacterConfirmNamePrompt, name))
	if errConfirmName != nil {
		l.WithError(errConfirmName).Error("Error reading confirm character name")
		io.WriteString(s, cfmt.Sprintf(inputErrorMsg))
		goto promptCharacterName
	}

	if !confirm {
		io.WriteString(s, cfmt.Sprintf(characterNameDeclinedMsg, name))
		goto promptCharacterName
	}

	l.WithField("name", name).Debug("Received character name")

	// TODO: Custome creation

	return StateMainMenu
}

func (w *World) promptListCharacters(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "list_characters"})
	l.Debug("Prompting for character list")
	io.WriteString(s, cfmt.Sprint(featureNotImplementedMsg))
	return StateMainMenu
}

func (w *World) promptDeleteCharacter(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "delete_character"})
	l.Debug("Prompting for character deletion")
	io.WriteString(s, cfmt.Sprint(featureNotImplementedMsg))
	return StateMainMenu
}

func (w *World) promptChangePassword(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "change_password"})
	l.Debug("Prompting for password change")

promptChangePassword:
	// Collect new password
	password, errPassword := utils.PromptUserPasswordInput(s, passwordNewPrompt)
	if errPassword != nil {
		l.WithError(errPassword).Error("Error reading password")
		goto promptChangePassword
	}
	l.WithFields(logrus.Fields{"user": u.Username, "id": u.ID, "password": password}).Debug("Received password")

	// Is the password in the min/max lengths?
	if len(password) < w.cfg.PasswordMinLength || len(password) > w.cfg.PasswordMaxLength {
		io.WriteString(s, cfmt.Sprintf(passwordMinMaxLengthMsg, w.cfg.PasswordMinLength, w.cfg.PasswordMaxLength))
		goto promptChangePassword
	}

	// Confirm the password
	passwordConfirm, errPasswordConfirm := utils.PromptUserPasswordInput(s, passwordConfirmPrompt)
	if errPasswordConfirm != nil {
		l.WithFields(logrus.Fields{"user": u.Username, "id": u.ID}).WithError(errPasswordConfirm).Error("Error reading confirm password")
		goto promptChangePassword
	}
	l.WithFields(logrus.Fields{"user": u.Username, "id": u.ID, "password": password}).Debug("Received confirm password")

	// Do the passwords match?
	if password != passwordConfirm {
		io.WriteString(s, cfmt.Sprintf(passwordMismatchMsg))
		l.WithFields(logrus.Fields{"user": u.Username, "id": u.ID, "password": password}).Warn("Passwords do not match")
		goto promptChangePassword
	}

	// Hash the password with bcrypt
	hashedPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(password), w.cfg.PasswordBcryptCost)
	if errHashPassword != nil {
		l.WithFields(logrus.Fields{"user": u.Username, "id": u.ID}).WithError(errHashPassword).Error("Error hashing password")
		goto promptChangePassword
	}
	l.WithFields(logrus.Fields{"user": u.Username, "id": u.ID, "hashedPassword": string(hashedPassword)}).Debug("Created password hash")

	// Save the new password
	u.Password = string(hashedPassword)
	u.Save()

	io.WriteString(s, cfmt.Sprintf(passwordChangedMsg))

	return StateMainMenu
}

// gameLoop is the main game loop for a connected client
func (w *World) gameLoop(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "screen": "game_loop"})
	l.Debug("Entering game loop")

	responseChannel := make(chan string)
	w.AddMessageChannel(u.ID, responseChannel)

	t := term.NewTerminal(s, "")
	io.WriteString(s, fmt.Sprintf("Hello %s [%s]\n", u.Username, u.ID))

	for {
		t.SetPrompt("> ")
		input, err := t.ReadLine()
		if err != nil {
			l.WithError(err).Error("Error reading input")
			break
		}

		if input == "exit" {
			break
		}

		io.WriteString(s, fmt.Sprintf("You entered: %s\n", input))

		inputMsg := InputMessage{
			FromSessionID: u.ID,
			Message:       input,
		}

		l.WithFields(logrus.Fields{"from_session_id": inputMsg.FromSessionID, "message": inputMsg.Message}).Info("Sending message to world")
		w.inputChan <- inputMsg

		// Wait for the response
		msg := <-responseChannel

		io.WriteString(s, fmt.Sprintf("Received message: %s\n", msg))
	}

	return StateQuit
}

var buildHelp = `
{{Commands:}}::#00ff00|bold
    help                                - Display this help message
    save                                - Save the character
    finalize                            - Finalize the character
    discard                             - Discard the character
    restart                             - Restart the character
    exit                                - Exit the charcter builder

    set name <name>                     - Set the character's name

    set metatype <metatype>             - Set the metatype 
    list metatypes                      - List the metatypes
    show metatype <metatype>            - Show the metatype

    set magic <magic>                   - Set the magic type 
    list magic                          - List the magic types
    show magic <magic>                  - Show the magic type

    set attribute <attribute> <value>   - Set the attribute to the value
    list atrributes                     - List the attributes
    show attribute <attribute>          - Show the attribute

    set skill <skill> <value>           - Set the skill to the value
    list skills                         - List the skills
    show skill <skill>                  - Show the skill

    set quality <quality>               - Set the quality
    list qualities                      - List the qualities
    show quality <quality>              - Show the quality

    set nuyen <nuyen>                   - Set the nuyen
`

var buildCharInfo = `
{{"Name:" | fg.Green }} {{ if .Name }}{{.Name}}{{ else }}set name <value>{{end}}
{{"Metatype:" | fg.Green }} {{ if .Metatype }}{{.Metatype.Name}} ({{.Metatype.Category}}){{ else }}set metatype <value>{{end}}
`

var buildCharInfo2 = `
MagicType: {{.MagicType}}
Build Points (Total/Remaining): %d/%d
Attributes: Current (Min/Max/Augmented Max)
    Body:       %d (%d/%d/%d)    Agility:    %d (%d/%d/%d)
    Reaction:   %d (%d/%d/%d)    Strength:   %d (%d/%d/%d)
    Willpower:  %d (%d/%d/%d)    Logic:      %d (%d/%d/%d)
    Intuition:  %d (%d/%d/%d)    Charisma:   %d (%d/%d/%d)
    Edge:       %d (%d/%d/%d)    Initiative: %d (%d/%d/%d)
    Essence:    %f (%f/%f/%f)    Magic:      %d (%d/%d/%d)
    Resonance:  %d (%d/%d/%d)
Skills:
    Active: %s (%d) [%s], Unarmed Combat 1, Pistols 1, Perception 1
    Knowledge: %s (%d) [%s], Security Procedures 1
    Language: %s (%d) [%s], English 13, Japanese 1
Qualities:
    Positive: %s, Ambidextrous, Biocompatibility (Cyberware), Code Slinger, First Impression, Guts, Juryrigger, Lucky, Natural Athlete, Quick Healer, Toughness
    Negative: %s, Bad Luck
Weapons: Ares Predator V, 100 rounds of ammo
Armor: Armor Jacket
Gear: %s [Rating %d] (%d)
Certified Credstick, Fake SIN [Rating 4], Medkit [Rating 6], Trauma Patch (3)
Nuyen: %d

- help for more commands
> 
`

var buildMetatypeList = `
Name:           %s
Category:       %s
Description:
%s
`

var buildMetatypeInfo = `
Name:           %s
Category:       %s
Description:
%s

Attributes (Min/Max/Augmented Max)
Body:       %d/%d/%d    Agility:    %d/%d/%d
Reaction:   %d/%d/%d    Strength:   %d/%d/%d
Willpower:  %d/%d/%d    Logic:      %d/%d/%d
Intuition:  %d/%d/%d    Charisma:   %d/%d/%d
Edge:       %d/%d/%d    Initiative: %d/%d/%d
Essence:    %f/%f/%f    Magic:      %d/%d/%d
Resonance:  %d/%d/%d

Qualities: %s
Quality Restrictions: %s
`
