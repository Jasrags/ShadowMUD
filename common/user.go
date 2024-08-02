package common

import (
	"fmt"
	"io"
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
		return nil
	}

	u := &User{
		Session: s,
		Pty:     pty,
		Window:  ptyWindow,
		Term:    term.NewTerminal(s, ""),
		lock:    sync.RWMutex{},

		Roles:      UserRoles{UserRoleUser},
		Characters: make(Characters),
		Logins:     Logins{},
		Bans:       Bans{},
	}

	u.AddUserRoles(UserRoleUser)

	return u
}

type Logins []Login

type Login struct {
	Time time.Time `yaml:"time"`
	IP   string    `yaml:"ip"`
}

type Bans []Ban

type Ban struct {
	CreatedAt time.Time `yaml:"created_at"`
	ExpiresAt time.Time `yaml:"time"`
	Reason    string    `yaml:"reason"`
	CreatedBy string    `yaml:"created_by"`
}

type UserRole string

type UserRoles []UserRole

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	lock            sync.RWMutex      `yaml:"-"`
	Session         ssh.Session       `yaml:"-"`
	Pty             ssh.Pty           `yaml:"-"`
	Window          <-chan ssh.Window `yaml:"-"`
	Term            *term.Terminal    `yaml:"-"`
	ActiveCharacter *Character        `yaml:"-"`

	ID                string     `yaml:"id"`
	Username          string     `yaml:"username"`
	Roles             UserRoles  `yaml:"roles"`
	Password          string     `yaml:"password"`
	ActiveCharacterID string     `yaml:"active_character_id,omitempty"`
	Characters        Characters `yaml:"characters,omitempty"`
	CreatedAt         time.Time  `yaml:"created_at"`
	UpdatedAt         time.Time  `yaml:"updated_at,omitempty"`
	DeletedAt         time.Time  `yaml:"deleted_at,omitempty"`
	Bans              Bans       `yaml:"bans"`
	Logins            Logins     `yaml:"logins"`
}

func (u *User) AddUserRoles(roles ...UserRole) {
	defer u.lock.Unlock()
	u.lock.Lock()

	u.Roles = append(u.Roles, roles...)
}

func (u *User) RemoveUserRoles(roles ...UserRole) {
	defer u.lock.Unlock()
	u.lock.Lock()

	for _, role := range roles {
		for i, r := range u.Roles {
			if r == role {
				u.Roles = append(u.Roles[:i], u.Roles[i+1:]...)
			}
		}
	}
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

	u.UpdatedAt = time.Now()

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
