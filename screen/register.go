package screen

import (
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s *Screens) PromptRegisterUser() int {
	// Do we have registration disabled?
	// TODO: Fix this loop
	if !s.cfg.RegistrationEnabled {
		io.WriteString(s.user.Session, cfmt.Sprint(registrationClosedMsg))
		return StatePromptLoginUser
	}

promptRegisterUsername:
	// Collect new username
	username, errUsername := PromptUserInput(s.user, usernameNewPrompt)
	if errUsername != nil {
		logrus.WithError(errUsername).Error("Error reading username")
		io.WriteString(s.user.Session, cfmt.Sprintf(inputErrorMsg))
		return StateQuit
	}

	// Is the username in the min/max lengths?
	if len(username) < s.cfg.UsernameMinLength || len(username) > s.cfg.UsernameMaxLength {
		io.WriteString(s.user.Session,
			cfmt.Sprintf(usernameMixMaxLengthMsg, s.cfg.UsernameMinLength, s.cfg.UsernameMaxLength))
		goto promptRegisterUsername
	}

	logrus.WithField("username", username).Info("Received new username")

	// TODO: Add a check for existing usernames
	// TODO: check if username is already taken
	for _, bannedName := range s.cfg.BannedNames {
		if strings.EqualFold(username, bannedName) {
			io.WriteString(s.user.Session,
				cfmt.Sprintf(usernameBannedMsg, username))
			return StatePromptRegisterUser
		}
	}

	// Confirm the username
	usernameConfirm, errUsernameConfirm := PromptConfirmInput(s.user, cfmt.Sprintf(usernameConfirmPrompt, username))
	// usernameConfirm, errUsernameConfirm := PromptUserInput(s.user,
	// cfmt.Sprintf(usernameConfirmPrompt, username))
	if errUsernameConfirm != nil {
		logrus.WithError(errUsername).Error("Error reading username")
		io.WriteString(s.user.Session, cfmt.Sprintf(inputErrorMsg))
		return StateQuit
	}

	// Did the user confirm the username?
	if !usernameConfirm {
		goto promptRegisterUsername
	}

	// // Did the user confirm the username?
	// if !strings.EqualFold(usernameConfirm, "y") {
	// 	io.WriteString(s.user.Session,
	// 		cfmt.Sprintf(usernameDeclinedMsg, username))
	// 	goto promptRegisterUsername
	// }

promptRegisterPassword:
	// Collect new password
	password, errPassword := PromptUserPasswordInput(s.user, passwordNewPrompt)
	if errPassword != nil {
		logrus.WithError(errPassword).Error("Error reading password")
		return StateQuit
	}
	logrus.WithFields(logrus.Fields{"password": password}).Debug("Received password")

	// Is the password in the min/max lengths?
	if len(password) < s.cfg.PasswordMinLength || len(password) > s.cfg.PasswordMaxLength {
		io.WriteString(s.user.Session,
			cfmt.Sprintf(passwordMinMaxLengthMsg, s.cfg.PasswordMinLength, s.cfg.PasswordMaxLength))
		goto promptRegisterPassword
	}

	// Confirm the password
	passwordConfirm, errPasswordConfirm := PromptUserPasswordInput(s.user, passwordConfirmPrompt)
	if errPasswordConfirm != nil {
		logrus.WithError(errPasswordConfirm).Error("Error reading confirm password")
		return StateQuit
	}
	logrus.WithFields(logrus.Fields{"username": username}).Debug("Received confirm password")

	// Do the passwords match?
	if password != passwordConfirm {
		io.WriteString(s.user.Session,
			cfmt.Sprintf(passwordMismatchMsg))
		goto promptRegisterPassword
	}

	// Hash the password with bcrypt
	hashedPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(password), s.cfg.PasswordBcryptCost)
	if errHashPassword != nil {
		logrus.WithError(errHashPassword).Error("Error hashing password")
		return StateQuit
	}
	logrus.WithFields(logrus.Fields{"hashedPassword": string(hashedPassword)}).Debug("Created password hash")

	goto promptRegisterUsername

	// Add ID, Username, Password
	s.user.ID = uuid.New().String()
	s.user.Username = username
	s.user.Password = string(hashedPassword)
	s.user.CreatedAt = time.Now()

	// Save the character
	if err := s.user.Save(); err != nil {
		logrus.WithError(err).Error("Error saving user")
		return StatePromptRegisterUser
	}

	io.WriteString(s.user.Session, cfmt.Sprintf(userCreatedMsg, username))

	return StatePromptMainMenu
}
