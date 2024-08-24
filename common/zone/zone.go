package zone

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Jasrags/ShadowMUD/common/shared"
	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	ZonesFilepath = "_data/zones"
)

type (
	Zones map[string]*Zone
	Zone  struct {
		sync.Mutex `yaml:"-"`
		log        *logrus.Entry `yaml:"-"`

		ID          string            `yaml:"id"`
		Name        string            `yaml:"name"`
		Description string            `yaml:"description"`
		RuleSource  shared.RuleSource `yaml:"rule_source"`
	}
)

func NewZone() *Zone {
	z := &Zone{
		ID: uuid.New().String(),
	}
	z.log = logrus.WithFields(logrus.Fields{"package": "common", "type": "zone", "zone_id": z.ID, "zone_name": z.Name})

	return z
}

func (z *Zone) Filepath() string {
	return fmt.Sprintf("%s/%s.yaml", ZonesFilepath, strings.ToLower(z.ID))
}

func (z *Zone) Validate() error {
	if z.ID == "" {
		return fmt.Errorf("id is required")
	}
	if z.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}

func LoadZone(id string, v *Zone) error {
	id = strings.ToLower(id)
	filepath := fmt.Sprintf("%s/%s.yaml", ZonesFilepath, id)

	if err := utils.LoadStructFromYAML(filepath, &v); err != nil {
		return err
	}

	return nil
}

func LoadZones() Zones {
	logrus.Info("Started loading zones")
	list := make(Zones)

	files, errReadDir := os.ReadDir(ZonesFilepath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read zones directory")
	}

	for _, file := range files {
		var v Zone
		if strings.HasSuffix(file.Name(), ".yaml") {

			name := strings.TrimSuffix(file.Name(), ".yaml")
			if err := LoadZone(name, &v); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load zone")
			}

			list[v.ID] = &v
			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded zone file")
		}
	}

	logrus.WithFields(logrus.Fields{"count": len(list)}).Info("Done loading zone")

	return list
}
