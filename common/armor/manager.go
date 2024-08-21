package armor

import (
	"sync"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/sirupsen/logrus"
)

const (
	Filepath              = "_data/items/armor"
	ModificationsFilepath = "_data/items/armor/modifications"
)

type Manager struct {
	sync.Mutex
	log           *logrus.Entry
	Specs         Specs
	Modifications ModificationSpecs
}

func NewManager() *Manager {
	m := &Manager{
		Specs:         make(Specs),
		Modifications: make(ModificationSpecs),
	}

	m.log = logrus.WithFields(logrus.Fields{"package": "armor", "type": "manager"})

	return m
}

func (m *Manager) Load() error {
	m.Lock()
	defer m.Unlock()

	if err := m.loadSpecs(); err != nil {
		return err
	}

	if err := m.loadModifications(); err != nil {
		return err
	}

	return nil
}

func (m *Manager) loadSpecs() error {
	m.log.Info("Loading armor specs")
	list, err := utils.LoadStructsFromDir[Spec](Filepath)
	if err != nil {
		m.log.WithError(err).Fatal("Could not load armor specs")
		return err
	}

	for _, item := range list {
		m.Specs[item.ID] = item
	}
	m.log.WithFields(logrus.Fields{"count": len(m.Specs)}).Info("Done loading armor specs")

	return nil
}

func (m *Manager) loadModifications() error {
	m.log.Info("Loading armor modifications")
	list, err := utils.LoadStructsFromDir[ModificationSpec](Filepath)
	if err != nil {
		m.log.WithError(err).Fatal("Could not load modifications")
		return err
	}

	for _, item := range list {
		m.Modifications[item.ID] = item
	}
	m.log.WithFields(logrus.Fields{"count": len(m.Modifications)}).Info("Done loading armor modifications")
	return nil
}
