package quality

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type (
	Manager struct {
		sync.Mutex
		log *logrus.Entry

		Specs     Specs
		Qualities Qualities
	}
)

func (m *Manager) Load() {
	// Load specs
	for _, spec := range CoreQualties {
		m.Specs[spec.ID] = spec
	}

	// Load qualities
}

func (m *Manager) GetSpec(id string) (*Spec, error) {
	s, ok := m.Specs[id]
	if !ok {
		return nil, fmt.Errorf("spec '%s' not found", id)
	}

	return &s, nil
}

func (m *Manager) GetQuality(id string) (*Quality, error) {
	s, ok := m.Specs[id]
	if !ok {
		return nil, fmt.Errorf("quality '%s' not found", id)
	}

	return NewQuality(&s), nil
}

func NewManager() *Manager {
	m := &Manager{
		Specs: make(Specs),
	}

	m.log = logrus.WithFields(logrus.Fields{"package": "quality", "type": "manager"})

	return m
}
