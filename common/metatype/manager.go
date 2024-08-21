package metatype

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
)

const (
	Filepath = "_data/metatypes"
)

type Manager struct {
	sync.Mutex
	log       *logrus.Entry
	Metatypes Metatypes
}

func NewManager() *Manager {
	m := &Manager{
		Metatypes: make(Metatypes),
	}

	m.log = logrus.WithFields(logrus.Fields{"package": "metatype", "type": "manager"})

	return m
}

func (m *Manager) Load() {
	logrus.Info("Loading metatypes")
	defer m.Unlock()
	m.Lock()

	files, errReadDir := os.ReadDir(Filepath)
	if errReadDir != nil {
		logrus.WithFields(logrus.Fields{"filepath": Filepath}).WithError(errReadDir).Fatal("Could not read directory")
	}

	for _, file := range files {
		var v Metatype
		if strings.HasSuffix(file.Name(), ".yaml") {
			if err := utils.LoadStructFromYAML(fmt.Sprintf("%s/%s", Filepath, file.Name()), &v); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load file to struct")
			}

			m.Metatypes[v.ID] = &v
			logrus.WithFields(logrus.Fields{"filename": file.Name(), "id": v.ID, "username": v.Name}).Debug("Loaded file")
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(m.Metatypes)}).Info("Done loading metatypes")
}

func (m *Manager) GetByID(id string) *Metatype {
	if v, ok := m.Metatypes[id]; ok {
		return v
	}

	return nil
}

func (m *Manager) GetIDs() []string {
	return maps.Keys(m.Metatypes)
}
