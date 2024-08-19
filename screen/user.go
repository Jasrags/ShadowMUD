package screen

import (
	"io"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s *Screens) PromptChangePassword() int {
promptChangePassword:
	// Collect new password
	password, errPassword := PromptUserPasswordInput(s.user, passwordNewPrompt)
	if errPassword != nil {
		s.log.WithError(errPassword).Error("Error reading password")
		return StateQuit
	}
	s.log.WithFields(logrus.Fields{"user": s.user.Username, "id": s.user.ID, "password": password}).
		Debug("Received password")

	// Is the password in the min/max lengths?
	if len(password) < s.cfg.PasswordMinLength || len(password) > s.cfg.PasswordMaxLength {
		io.WriteString(s.user.Session,
			cfmt.Sprintf(passwordMinMaxLengthMsg, s.cfg.PasswordMinLength, s.cfg.PasswordMaxLength))
		goto promptChangePassword
	}

	// Confirm the password
	passwordConfirm, errPasswordConfirm := PromptUserPasswordInput(s.user, passwordConfirmPrompt)
	if errPasswordConfirm != nil {
		s.log.WithFields(logrus.Fields{"user": s.user.Username, "id": s.user.ID}).
			WithError(errPasswordConfirm).Error("Error reading confirm password")
		return StateQuit
	}
	s.log.WithFields(logrus.Fields{"user": s.user.Username, "id": s.user.ID, "password": password}).
		Debug("Received confirm password")

	// Do the passwords match?
	if password != passwordConfirm {
		io.WriteString(s.user.Session, cfmt.Sprintf(passwordMismatchMsg))
		s.log.WithFields(logrus.Fields{"user": s.user.Username, "id": s.user.ID, "password": password}).
			Warn("Passwords do not match")
		goto promptChangePassword
	}

	// Hash the password with bcrypt
	hashedPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(password), s.cfg.PasswordBcryptCost)
	if errHashPassword != nil {
		s.log.WithFields(logrus.Fields{"user": s.user.Username, "id": s.user.ID}).
			WithError(errHashPassword).
			Error("Error hashing password")
		return StateQuit
	}
	s.log.WithFields(logrus.Fields{"user": s.user.Username, "id": s.user.ID, "hashedPassword": string(hashedPassword)}).
		Debug("Created password hash")

	// Save the new password
	s.user.Password = string(hashedPassword)
	s.user.Save()

	io.WriteString(s.user.Session, cfmt.Sprintf(passwordChangedMsg))

	return StatePromptMainMenu
}
