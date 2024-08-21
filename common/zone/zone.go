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

var CoreZones = []*Zone{
	{
		ID:          "the_void",
		Name:        "The Void",
		Description: "The Void is a place of nothingness. It is a place where the laws of physics do not apply. It is a place where time and space are meaningless. It is a place where the mind can wander and the soul can rest. It is a place where the universe is born and where it dies. It is a place where the gods dwell and where the demons lurk. It is a place where the past, present, and future are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one. It is a place where the alpha and the omega are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one. It is a place where the alpha and the omega are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one. It is a place where the alpha and the omega are one. It is a place where the light and the dark are one. It is a place where the living and the dead are one. It is a place where the mind and the body are one. It is a place where the self and the other are one. It is a place where the dream and the reality are one. It is a place where the truth and the lie are one. It is a place where the beginning and the end are one.",
	},
	{
		ID:          "seattle",
		Name:        "Seattle",
		Description: "Seattle is a city in the UCAS, located in the Pacific Northwest. It is the largest city in the UCAS and the largest metroplex in the world. Seattle is a major hub for trade, commerce, and cult ure, and is home to a number of megacorporations, including Ares Macrotechnology, Aztechnology, and NeoNET.",
		RuleSource:  shared.RuleSourceSR5Core,
	},
}
