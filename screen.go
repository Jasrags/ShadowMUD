package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/Jasrags/ShadowMUD/common/user"
	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/gliderlabs/ssh"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

var (
	// Prompt/Message strings
	// Character Creation
	// Game loop
	// inputEchoMsg        = "{{You typed:}}::#ffffff|bold %s\n"
	// gameLoopPrompt      = "{{> }}::#ffffff|bold"
	characterListOption = "{{%d.}}::#00ff00 %s\n"
	// Login
	loginClosedMsg     = "{{Login is currently closed.}}::#ff8700\n"
	usernamePrompt     = "{{Username: }}::#ffffff|bold"
	passwordPrompt     = "{{Password: }}::#ffffff|bold"
	invalidLoginMsg    = "{{You have entered an invalid username or password.}}::#ff8700\n"
	loginSuccessfulMsg = "{{Login successful.}}::#00ff00\n"
	loginUserBannedMsg = "{{You are banned until %s.}}::#ff8700\n"

	inputErrorMsg    = "{{An error occurred while reading your response.}}::#ff8700\n"
	requiredInputMsg = "{{You must enter a value.}}::#ff8700\n"

	// Registration
	registrationClosedMsg   = "{{Registration is currently closed.}}::#ff8700\n"
	usernameBannedMsg       = "{{Username '%s' is not allowed.}}::#ff8700\n"
	usernameNewPrompt       = "{{Enter your desired username: }}::#ffffff|bold"
	usernameConfirmPrompt   = "{{Confirm username '%s'}}::#ffffff|bold {{(y/n)}}::#00ff00|bold{{:}}::#ffffff|bold"
	usernameMixMaxLengthMsg = "{{Username must be between %d and %d characters.}}::#ff8700\n"
	usernameDeclinedMsg     = "{{Username '%s' was not confirmed.}}::#ff8700\n"
	passwordNewPrompt       = "{{Enter new password: }}::#ffffff|bold"
	passwordConfirmPrompt   = "{{Confirm password: }}::#ffffff|bold"
	passwordMismatchMsg     = "{{Passwords do not match.}}::#ff8700\n"
	passwordMinMaxLengthMsg = "{{Password must be between %d and %d characters.}}::#ff8700\n"
	userCreatedMsg          = "{{User '%s' has been created.}}::#00ff00\n"

	// Menu
	menuPrompt                = "{{Enter the number of the option you would like to select: }}::#ffffff|bold"
	menuInvalidChoice         = "Invalid choice: %s\n"
	mainMenuTitle             = "\n{{Main Menu}}::#00ff00|bold\n"
	menuOptionEnterGame       = "{{1.}}::#00ff00 Enter game\n"
	menuOptionCreateCharacter = "{{2.}}::#00ff00 Create character (%d/%d)\n"
	menuOptionListCharacters  = "{{3.}}::#00ff00 List characters\n"
	menuOptionDeleteCharacter = "{{4.}}::#00ff00 Delete character\n"
	menuOptionChangePassword  = "{{5.}}::#00ff00 Change password\n"
	menuOptionQuit            = "{{0.}}::#00ff00 Quit\n"

	// Character Creation
	characterNamePrompt              = "{{Enter the name of your character: }}::#ffffff|bold"
	noCharactersCreatedMsg           = "\n{{You have no characters created, let's make one now!}}::#ff8700\n"
	characterNameDeclinedMsg         = "{{Character name '%s' was not confirmed.}}::#ff8700\n"
	characterNoneCreatedMsg          = "{{You have no characters created.}}::#ff8700\n"
	characterMaxCharactersMsg        = "{{You have reached the maximum number of characters allowed.}}::#ff8700\n"
	characterNameMixMaxLengthMsg     = "{{Character name must be between %d and %d characters.}}::#ff8700\n"
	createCharacterMenuTitle         = "\n{{Character Creation}}::#00ff00|bold\n"
	createCharacterMenuOptionPregen  = "{{1.}}::#00ff00 Choose an pre-generated character\n"
	createCharacterMenuOptionCustom  = "{{2.}}::#00ff00 Create an custom character\n"
	createCharacterMenuOptionLearn   = "{{3.}}::#00ff00 Learn more about Shadowrun characters\n"
	createCharacterMenuOptionReturn  = "{{4.}}::#00ff00 Return to the main menu\n"
	createCharacterNamePrompt        = "{{Enter the name of your character: }}::#ffffff|bold"
	createCharacterConfirmNamePrompt = "{{Confirm character name '%s'}}::#ffffff|bold {{(y/n)}}::#00ff00|bold{{:}}::#ffffff|bold"
	chooseCharacterPrompt            = "{{Choose a character to enter the game:}}::#00ff00\n"

	//     "\nCharacter Creation Menu:"
	// Set your character's name
	// Select Metatype
	// Choose magic or resonance
	// Purchase qualties
	// Purchase skills
	// Purchase spells/powers/complex forms
	// Purchase Nuyen
	// Spending Nuyen

	//     "1. Allocate Attribute Points"
	//     "2. Allocate Skill Points"
	//     "3. Allocate Quality Points"
	//     "4. View Character"
	//     "5. Finalize Character"
	//     "6. Exit"
	//    "Choose an option: "

	quitMsg                  = "{{Goodbye!}}::#00ff00\n"
	passwordChangedMsg       = "{{Password has been changed.}}::#0000ff\n"
	featureNotImplementedMsg = "{{Feature not implemented}}::#ff0000\n"
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
		return StatePromptCreateCharacter
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
		return StatePromptCreateCharacter
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

func (w *World) promptCreateCharacter(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "create_character"})
	l.Debug("Prompting for character creation")

	// promptCreateCharacterMenu:
	// If the user has reached the maximum number of characters, return to the main menu
	if len(u.Characters) >= w.cfg.UserCharacterMaxCount {
		io.WriteString(s, cfmt.Sprintf(characterMaxCharactersMsg))
		return StateMainMenu
	}

	state := StateCreateCharacterMenu
	for {
		switch state {
		case StateCreateCharacterMenu:
			state = w.promptCreateCharacterMenu(s, u)
		case StateCreateCharacterName:
			state = w.promptCreateCharacterName(s, u)
		case StateCreateCharacterMetatype:
			// state = w.promptCreateCharacterMetatype(s, u)
		case StateCreateCharacterMagic:
			// state = w.promptCreateCharacterMagic(s, u)
		case StateCreateCharacterAttributes:
			// state = w.promptCreateCharacterAttributes(s, u)
		case StateCreateCharacterSkills:
			// state = w.promptCreateCharacterSkills(s, u)
		case StateCreateCharacterQualities:
			// state = w.promptCreateCharacterQualities(s, u)
		case StateCreateCharacterNuyen:
			// state = w.promptCreateCharacterNuyen(s, u)
		case StateCreateCharacterSpend:
			// state = w.promptCreateCharacterSpend(s, u)
		case StateCreateCharacterQuit:
			return StateMainMenu
		}
	}

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
}

func (w *World) promptCreateCharacterMenu(s ssh.Session, u *user.User) State {
	l := logrus.WithFields(logrus.Fields{"remote_addr": s.RemoteAddr(), "user_id": u.ID, "username": u.Username, "package": "main", "screen": "create_character_menu"})
	l.Debug("Prompting for character creation menu")
	io.WriteString(s, cfmt.Sprint(featureNotImplementedMsg))

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

	return StateCreateCharacterMenu
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
