package bioware

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/sirupsen/logrus"
)

const (
	Filepath = "_data/items/bioware"
)

type Manager struct {
	sync.Mutex
	log   *logrus.Entry
	Specs Specs
}

func NewManager() *Manager {
	m := &Manager{
		Specs: make(Specs),
	}

	m.log = logrus.WithFields(logrus.Fields{"package": "bioware", "type": "manager"})

	return m
}

func (m *Manager) Load() error {
	m.Lock()
	defer m.Unlock()

	if err := m.loadSpecs(); err != nil {
		return err
	}

	return nil
}

func (m *Manager) loadSpecs() error {
	list, err := utils.LoadStructsFromDir[Spec](Filepath)
	if err != nil {
		return err
	}

	for _, item := range list {
		m.Specs[item.ID] = item
	}

	return nil
}

func test[T any](filepath string) ([]T, error) {
	var list []T

	files, errReadDir := os.ReadDir(filepath)
	if errReadDir != nil {
		return list, errReadDir
	}

	for _, file := range files {
		var v T
		if strings.HasSuffix(file.Name(), ".yaml") {
			if err := utils.LoadStructFromYAML(fmt.Sprintf("%s/%s", filepath, file.Name()), &v); err != nil {
				return list, err
			}

			list = append(list, v)
		}
	}

	return list, nil
}
