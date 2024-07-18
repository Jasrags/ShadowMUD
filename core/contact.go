package core

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/core/util"
)

type ContactType string

const (
	ContactsDataPath = "data/contacts"
	ContactFilename  = ContactsDataPath + "/%s.yaml"
)

const (
	ContactTypeFixer ContactType = "Fixer"
)

type ContactSpec struct {
	ID          string      `yaml:"id,omitempty"`
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Type        ContactType `yaml:"type"`
	RuleSource  RuleSource  `yaml:"rule_source"`
}

type Contact struct {
	ID         string      `yaml:"id,omitempty"`
	Spec       ContactSpec `yaml:"-"`
	Connection int         `yaml:"connection"`
	Loyalty    int         `yaml:"loyalty"`
}

var CoreContacts = []ContactSpec{
	{
		ID:          "brian_flannigan",
		Name:        "Brian Flannigan",
		Description: "A fixer is a person who arranges illicit goods or services for characters. Fixers are the go-to people for characters who need to buy or sell illegal goods, find a buyer for a stolen item, or hire a shadowrunner for a job.",
		Type:        ContactTypeFixer,
		RuleSource:  RuleSourceSR5Core,
	},
}

// func LoadContacts(dataPath string) map[string]Contact {
// 	data := make(map[string]Contact)

// 	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Info("Started loading contacts")
// 	files, err := os.ReadDir(dataPath)
// 	if err != nil {
// 		logrus.WithError(err).Error("Could not read contacts data directory")
// 		return data
// 	}

// 	for _, file := range files {
// 		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
// 			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

// 			var v Contact
// 			err := util.LoadStructFromYAML(filePath, &v)
// 			if err != nil {
// 				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Error("Could not load contact")
// 				return data
// 			}
// 			data[v.ID] = v
// 			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded contact file")
// 		}
// 	}

// 	logrus.WithFields(logrus.Fields{"count": len(data)}).Info("Done loading contacts")

// 	return data
// }

func SaveCoreContacts(fileDir string) error {
	for _, v := range CoreContacts {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+ContactFilename, v.ID), &v); err != nil {
			return err
		}
	}

	return nil
}
