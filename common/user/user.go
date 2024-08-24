package user

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Jasrags/ShadowMUD/common/character"
	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type (
	Logins []Login
	Login  struct {
		Time time.Time `yaml:"time"`
		IP   string    `yaml:"ip"`
	}
	Bans []Ban
	Ban  struct {
		CreatedAt time.Time `yaml:"created_at"`
		ExpiresAt time.Time `yaml:"time"`
		Reason    string    `yaml:"reason"`
		CreatedBy string    `yaml:"created_by"`
	}
	Role  string
	Roles []Role

	Users map[string]*User
	User  struct {
		sync.Mutex `yaml:"-"`
		log        *logrus.Entry        `yaml:"-"`
		Character  *character.Character `yaml:"-"` // Set to the active character

		// Saved fields
		ID           string               `yaml:"id"`
		Username     string               `yaml:"username"`
		Roles        Roles                `yaml:"roles"`
		Password     string               `yaml:"password"`
		Characters   character.Characters `yaml:"characters"`
		CharacterIDs []string             `yaml:"character_ids"`
		Bans         Bans                 `yaml:"bans"`
		Logins       Logins               `yaml:"logins"`
		CreatedAt    time.Time            `yaml:"created_at"`
		UpdatedAt    time.Time            `yaml:"updated_at,omitempty"`
		DeletedAt    time.Time            `yaml:"deleted_at,omitempty"`
	}
)

func New() *User {
	u := &User{
		Roles:      Roles{},
		Characters: character.Characters{},
		// Characters: make(character.Characters),
		Logins: Logins{},
		Bans:   Bans{},
	}

	u.log = logrus.WithFields(logrus.Fields{
		"package": "user",
		"type":    "user",
	})

	return u
}

func (u *User) Load() error {
	u.log.Debug("Loading user")
	u.ID = uuid.New().String()
	u.Username = utils.GenerateRandomUsername()

	return nil
}

func (u *User) AddUserRoles(roles ...Role) {
	defer u.Unlock()
	u.Lock()

	u.Roles = append(u.Roles, roles...)
}

func (u *User) RemoveUserRoles(roles ...Role) {
	defer u.Unlock()
	u.Lock()

	for _, role := range roles {
		for i, r := range u.Roles {
			if r == role {
				u.Roles = append(u.Roles[:i], u.Roles[i+1:]...)
			}
		}
	}
}

func (u *User) HasRole(role Role) bool {
	defer u.Unlock()
	u.Lock()

	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}

	return false
}

func (u *User) SetActiveCharacterByID(id string) {
	defer u.Unlock()
	u.Lock()

	if c, ok := u.Characters[id]; ok {
		u.Character = c

		return
	}

	logrus.WithFields(logrus.Fields{"id": id}).Error("Character not found")
}

func (u *User) AddCharacter(c *character.Character) {
	logrus.WithFields(logrus.Fields{"id": c.ID, "name": c.Name}).Debug("Adding character to user")
	defer u.Unlock()
	u.Lock()

	u.Characters[c.ID] = c
	logrus.Debug("Character added to user")
}

func (u *User) RemoveCharacterByID(id string) {
	logrus.WithField("id", id).Debug("Removing character by ID")
	defer u.Unlock()
	u.Lock()

	if _, ok := u.Characters[id]; !ok {
		logrus.WithField("id", id).Error("Character not found")

		return
	}

	delete(u.Characters, id)
	logrus.WithField("id", id).Debug("Character removed")
}

func (u *User) GetCharacterByID(id string) *character.Character {
	logrus.WithField("id", id).Debug("Getting character by ID")
	defer u.Unlock()
	u.Lock()

	if c, ok := u.Characters[id]; ok {
		logrus.WithField("id", id).Debug("Character found")
		return c
	}

	logrus.WithField("id", id).Error("Character not found")

	return nil
}

func (u *User) ChangePassword(password string) {
	defer u.Unlock()
	u.Lock()

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password")

		return
	}

	u.Password = string(bcryptPassword)
}

func (u *User) SetActiveCharacter(c *character.Character) {
	logrus.WithFields(logrus.Fields{"id": c.ID, "name": c.Name}).Debug("Setting active character")
	defer u.Unlock()
	u.Lock()

	u.Character = c
	logrus.WithFields(logrus.Fields{"id": c.ID, "name": c.Name}).Debug("Active character set")
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
	return fmt.Sprintf("%s/%s.yaml", Filepath, strings.ToLower(u.Username))
}

func (u *User) Save() error {
	u.log.Debug("Saving user")

	defer u.Unlock()
	u.Lock()

	u.UpdatedAt = time.Now()

	if err := utils.SaveStructToYAML(u.Filepath(), u); err != nil {
		u.log.WithError(err).Error("Error saving user")
		return err
	}

	u.log.Debug("Saved user")

	return nil
}

// func (u *User) GameLoop() error {
// 	io.WriteString(u.Session, cfmt.Sprintf("{{> }}::#ffffff|bold"))
// 	line, err := u.Term.ReadLine()
// 	if err != nil {
// 		return err
// 	}
// 	logrus.WithField("line", line).Debug("Received line")
// 	io.WriteString(u.Session, cfmt.Sprintf("{{You typed:}}::#ffffff|bold %s\n", line))

// 	return nil
// }
