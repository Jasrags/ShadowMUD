package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jasrags/ShadowMUD/core/util"

	"github.com/sirupsen/logrus"
)

const (
	QualityDataPath                   = "data/qualities"
	QualityFilename                   = QualityDataPath + "/%s.yaml"
	QualityFileMinVersion             = "0.0.1"
	QualityTypePositive   QualityType = "Positive"
	QualityTypeNegative   QualityType = "Negative"
)

type (
	QualityType string
)

var (
	Qualities = map[string]Quality{}
)

func NewQuality() *Quality {
	return &Quality{}
}

type Quality struct {
	ID            string   `yaml:"id,omitempty"`
	Name          string   `yaml:"name"`
	Description   string   `yaml:"description"`
	Prerequisites []string `yaml:"prerequisites,omitempty"`
	Cost          int      `yaml:"cost"`
	Rating        int      `yaml:"rating,omitempty"`
	RuleSource    string   `yaml:"rule_source"`
	FileVersion   string   `yaml:"file_version"`
}

func LoadQualities() {
	logrus.Info("Started loading qualities")

	files, errReadDir := os.ReadDir(QualityDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read qualities directory")
	}

	// Create a map to store the metatypes
	qualities := make(map[string]Quality, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", QualityDataPath, file.Name())

			var quality Quality
			if err := util.LoadStructFromYAML(filepath, &quality); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load qualities")
			}

			qualities[quality.Name] = quality
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded qualities file")
	}

	logrus.WithFields(logrus.Fields{"count": len(qualities)}).Info("Done loading qualities")

	Qualities = qualities
}

func LoadQuality(name string) (*Quality, error) {
	var v Quality
	if err := util.LoadStructFromYAML(fmt.Sprintf(QualityFilename, name), &v); err != nil {
		return nil, err
	}

	return &v, nil
}
