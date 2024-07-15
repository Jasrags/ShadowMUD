package quality

import (
	"fmt"
	"os"
	"shadowrunmud/util"
	"strings"
	"sync"

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

type Quality interface {
	GetID() string
	SetID(string)
	GetName() string
	GetDescription() string
	GetPrerequisites() string
	GetCost() int
	GetRating() int
	GetRuleSource() string
}

func NewQuality() Quality {
	return &qual{}
}

type qual struct {
	ID            string `yaml:"id,omitempty"`
	Name          string `yaml:"name"`
	Description   string `yaml:"description"`
	Prerequisites string `yaml:"prerequisites,omitempty"`
	Cost          int    `yaml:"cost"`
	Rating        int    `yaml:"rating,omitempty"`
	RuleSource    string `yaml:"rule_source"`
	FileVersion   string `yaml:"file_version"`
}

func (q *qual) GetID() string {
	return q.ID
}

func (q *qual) SetID(id string) {
	q.ID = id
}

func (q *qual) GetName() string {
	return q.Name
}

func (q *qual) GetDescription() string {
	return q.Description
}

func (q *qual) GetPrerequisites() string {
	return q.Prerequisites
}

func (q *qual) GetCost() int {
	return q.Cost
}

func (q *qual) GetRating() int {
	return q.Rating
}

func (q *qual) GetRuleSource() string {
	return q.RuleSource
}

func LoadQualities(wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Debug("Started loading qualities")

	files, errReadDir := os.ReadDir(QualityDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read metatype directory")
	}

	// Create a map to store the metatypes
	qualities := make(map[string]qual, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", QualityDataPath, file.Name())

			var quality Quality
			if err := util.LoadStructFromYAML(filepath, &quality); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load metatype")
			}

			qualities[quality.GetName()] = quality
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded metatype file")
	}

	logrus.WithFields(logrus.Fields{"count": len(qualities)}).Info("Done loading metatypes")

	Qualities = qualities
}

func LoadQuality(name string) (Quality, error) {
	var m Quality
	if err := util.LoadStructFromYAML(fmt.Sprintf(QualityFilename, name), &m); err != nil {
		return nil, err
	}

	return m, nil
}
