package user

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Jasrags/ShadowMUD/utils"
	"golang.org/x/exp/maps"

	"github.com/sirupsen/logrus"
)

const (
	Filepath = "_data/users"
)

type Manager struct {
	sync.Mutex
	log *logrus.Entry

	Users         Users
	UsernameIDMap map[string]string
}

func NewManager() *Manager {
	m := &Manager{
		Users:         make(Users),
		UsernameIDMap: make(map[string]string),
	}

	m.log = logrus.WithFields(logrus.Fields{"package": "user", "type": "manager"})

	return m
}

func (m *Manager) Load() {
	logrus.Info("Loading users")
	defer m.Unlock()
	m.Lock()

	files, errReadDir := os.ReadDir(Filepath)
	if errReadDir != nil {
		logrus.WithFields(logrus.Fields{"filepath": Filepath}).WithError(errReadDir).Fatal("Could not read directory")
	}

	for _, file := range files {
		var v User
		if strings.HasSuffix(file.Name(), ".yaml") {
			if err := utils.LoadStructFromYAML(fmt.Sprintf("%s/%s", Filepath, file.Name()), &v); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load file to struct")
			}

			m.Users[v.ID] = &v
			m.UsernameIDMap[strings.ToLower(v.Username)] = v.ID
			logrus.WithFields(logrus.Fields{"filename": file.Name(), "id": v.ID, "username": v.Username}).Debug("Loaded file")
		}
	}

	usernames := maps.Keys(m.UsernameIDMap)
	logrus.WithField("usernames", usernames).Info("Loaded usernames")

	logrus.WithFields(logrus.Fields{"count": len(m.Users)}).Info("Done loading users")
}

func (m *Manager) Add(u *User) {
	defer m.Unlock()
	m.Lock()

	m.Users[u.ID] = u
	m.UsernameIDMap[strings.ToLower(u.Username)] = u.ID
}

func (m *Manager) GetByID(id string) *User {
	defer m.Unlock()
	m.Lock()

	return m.Users[id]
}

func (m *Manager) GetByUsername(username string) *User {
	defer m.Unlock()
	m.Lock()

	id, ok := m.UsernameIDMap[username]
	if !ok {
		return nil
	}

	return m.Users[id]
}

func (m *Manager) RemoveByID(id string) {
	defer m.Unlock()
	m.Lock()

	delete(m.Users, id)
	delete(m.UsernameIDMap, m.Users[id].Username)
}

func (m *Manager) GetCount() int {
	defer m.Unlock()
	m.Lock()

	return len(m.Users)
}

func (m *Manager) GetUsers() Users {
	defer m.Unlock()
	m.Lock()

	return m.Users
}

func (m *Manager) GetIdByUsername(username string) string {
	defer m.Unlock()
	m.Lock()

	id, ok := m.UsernameIDMap[strings.ToLower(username)]
	if !ok {
		return ""
	}

	return id
}
