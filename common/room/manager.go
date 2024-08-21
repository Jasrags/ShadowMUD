package room

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/sirupsen/logrus"
)

func LoadRoom(id string, v *Spec) error {
	id = strings.ToLower(id)
	filepath := fmt.Sprintf("%s/%s.yaml", RoomsFilepath, id)

	if err := utils.LoadStructFromYAML(filepath, &v); err != nil {
		return err
	}

	return nil
}

func LoadRooms() Rooms {
	logrus.Info("Started loading rooms")
	list := make(Rooms)

	if err := filepath.Walk(RoomsFilepath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			subdir := info.Name()
			subfiles, err := os.ReadDir(path)
			if err != nil {
				logrus.WithError(err).WithField("subdir", subdir).Fatal("Could not read subdirectory")
			}
			for _, subfile := range subfiles {
				var v Spec
				if strings.HasSuffix(subfile.Name(), ".yaml") {
					name := strings.TrimSuffix(subfile.Name(), ".yaml")
					if err := LoadRoom(fmt.Sprintf("%s/%s", subdir, name), &v); err != nil {
						logrus.WithFields(logrus.Fields{"filename": subfile.Name(), "subdir": subdir}).WithError(err).Fatal("Could not load room")
					}
					r := NewRoom(&v)
					// list[v.ID] = NewRoom(&v)
					list[fmt.Sprintf("%s:%s", r.Spec.ZoneID, r.ID)] = r
					logrus.WithFields(logrus.Fields{"filename": subfile.Name(), "subdir": subdir}).Debug("Loaded room file")
				}
			}
		}
		return nil
	}); err != nil {
		logrus.WithError(err).Fatal("Could not walk rooms directory")
	}

	logrus.WithFields(logrus.Fields{"count": len(list)}).Info("Done loading rooms")

	return list
}
