package metatype

import (
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

func (m *Manager) Load() error {
	m.Lock()
	defer m.Unlock()

	if err := m.loadMetatypes(); err != nil {
		return err
	}

	return nil
}

func (m *Manager) loadMetatypes() error {
	m.log.Info("Loading metatypes")
	list, err := utils.LoadStructsFromDir[Metatype](Filepath)
	if err != nil {
		m.log.WithError(err).Fatal("Could not load metatypes")
		return err
	}

	for _, item := range list {
		m.Metatypes[item.ID] = item
	}
	m.log.WithFields(logrus.Fields{"count": len(m.Metatypes)}).Info("Done loading metatypes")
	return nil
}

// func (m *Manager) GetByID(id string) *Metatype {
// 	if v, ok := m.Metatypes[id]; ok {
// 		return v
// 	}

// 	return nil
// }

func (m *Manager) GetIDs() []string {
	return maps.Keys(m.Metatypes)
}
