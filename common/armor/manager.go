package armor

import (
	"sync"

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
	Modifications Modifications
}

func NewManager() *Manager {
	m := &Manager{
		Specs:         make(Specs),
		Modifications: make(Modifications),
	}

	return m
}
