package common

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/gliderlabs/ssh"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

const (
	UserFilepath = "_data/users"
)

func PromptUsername(s ssh.Session) (string, error) {
	t := term.NewTerminal(s, "")

	io.WriteString(s, cfmt.Sprint("{{Username: }}::#ffffff|bold"))
	username, errReadLine := t.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")
		return "", errReadLine
	}

	username = strings.TrimSpace(username)
	logrus.WithField("username", username).Info("Received username")

	return username, nil
}

func IsNewUsername(username string) bool {
	if _, err := os.Stat(fmt.Sprintf("%s/%s.yaml", UserFilepath, strings.ToLower(username))); os.IsNotExist(err) {
		logrus.WithError(err).WithFields(logrus.Fields{"username": username}).Debug("Username does not exist")
		return true
	}

	logrus.WithFields(logrus.Fields{"username": username}).Debug("Username exists")

	return false
}

func PromptPassword(s ssh.Session) (string, error) {
	t := term.NewTerminal(s, "")

	// Collect password without echoing
	cfmt.Sprint("{{Password: }}::#ffffff|bold")
	passwordBytes, err := t.ReadPassword(cfmt.Sprint("{{Password: }}::#ffffff|bold"))
	if err != nil {
		log.Println("Error reading password:", err)
		return "", err
	}

	password := strings.TrimSpace(string(passwordBytes))
	logrus.WithField("password", password).Info("Received password")

	return password, nil
}

func AuthenticateUser(s ssh.Session) error {
	t := term.NewTerminal(s, "")

	io.WriteString(s, cfmt.Sprint("{{Username: }}::#ffffff|bold"))
	username, errReadLine := t.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading username")
		return errReadLine
	}

	username = strings.TrimSpace(username)
	logrus.WithField("username", username).Info("Received username")

	if new := IsNewUsername(username); new {
		logrus.WithField("new", new).Info("Username is new")
	}

	// Collect password without echoing
	cfmt.Sprint("{{Password: }}::#ffffff|bold")
	passwordBytes, err := t.ReadPassword(cfmt.Sprint("{{Password: }}::#ffffff|bold"))
	if err != nil {
		log.Println("Error reading password:", err)
		return err
	}

	password := strings.TrimSpace(string(passwordBytes))
	logrus.WithField("password", password).Info("Received password")

	// if u, err := LoadUser(username); err != nil {
	// 	logrus.WithError(err).Error("Error loading user")
	// }

	return nil
}

// LoadUser loads a user from the filesystem
func LoadUser(username string, u *User) error {
	username = strings.ToLower(username)
	filepath := fmt.Sprintf("%s/%s.yaml", UserFilepath, username)

	// Check if the user file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return err
	}

	if err := utils.LoadStructFromYAML(filepath, &u); err != nil {
		return err
	}

	return nil
}

func NewUser(s ssh.Session) *User {
	pty, ptyWindow, isActive := s.Pty()
	if !isActive {
		logrus.Error("Session is not active")
	}

	return &User{
		Session: s,
		Pty:     pty,
		Window:  ptyWindow,
		Term:    term.NewTerminal(s, ""),
		lock:    sync.RWMutex{},
		// log:        logrus.WithFields(logrus.Fields{"id": u.ID, "username": u.Username}),
		Characters: make(map[string]*Character),
	}
}

type Login struct {
	Time time.Time `yaml:"time"`
	IP   string    `yaml:"ip"`
}

type User struct {
	lock            sync.RWMutex      `yaml:"-"`
	Session         ssh.Session       `yaml:"-"`
	Pty             ssh.Pty           `yaml:"-"`
	Window          <-chan ssh.Window `yaml:"-"`
	Term            *term.Terminal    `yaml:"-"`
	ActiveCharacter *Character        `yaml:"-"`

	ID                string                `yaml:"id,omitempty"`
	Username          string                `yaml:"username"`
	Password          string                `yaml:"password"`
	ActiveCharacterID string                `yaml:"active_character_id"`
	Characters        map[string]*Character `yaml:"characters"`
	Logins            []Login               `yaml:"logins"`
	// TODO: Add in support for user bans with a duration and a reason
}

func (u *User) ChangePassword(password string) {
	defer u.lock.Unlock()
	u.lock.Lock()

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password")
		return
	}

	u.Password = string(bcryptPassword)
}

func (u *User) SetActiveCharacter(c *Character) {
	defer u.lock.Unlock()
	u.lock.Lock()

	u.ActiveCharacter = c
	u.ActiveCharacterID = c.ID
}

func (u *User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}

	if u.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil

}

func (u *User) Filepath() string {
	return fmt.Sprintf("%s/%s.yaml", UserFilepath, strings.ToLower(u.Username))

}

func (u *User) Save() error {
	logrus.Debug("Saving character")

	defer u.lock.Unlock()
	u.lock.Lock()

	if err := utils.SaveStructToYAML(u.Filepath(), u); err != nil {
		logrus.WithError(err).Error("Error saving user")
		return err
	}

	logrus.Debug("Saved character")

	return nil
}

func (u *User) GameLoop() error {
	io.WriteString(u.Session, cfmt.Sprintf("{{> }}::#ffffff|bold"))
	line, err := u.Term.ReadLine()
	if err != nil {
		return err
	}
	logrus.WithField("line", line).Debug("Received line")
	io.WriteString(u.Session, cfmt.Sprintf("{{You typed:}}::#ffffff|bold %s\n", line))

	return nil
}
