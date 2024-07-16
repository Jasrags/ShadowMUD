package core

import (
	"fmt"
	"os"
	"strings"

	"shadowrunmud/core/util"

	"github.com/sirupsen/logrus"
)

const (
	KnowledgeSkillDataPath       = "data/skills/knowledge"
	KnowledgeSkillFilename       = KnowledgeSkillDataPath + "/%s.yaml"
	KnowledgeSkillFileMinVersion = "0.0.1"
)

var (
	KnowledgeSkills = map[string]KnowledgeSkill{}
)

type KnowledgeSkill struct {
	ID          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	IsCommon    bool   `yaml:"is_common"`
	Rank        int    `yaml:"rank,omitempty"`
	RuleSource  string `yaml:"rule_source"`
	FileVersion string `yaml:"file_version"`
}

func LoadKnowledgeSkills() {
	logrus.Info("Started loading knowledge skills")

	files, errReadDir := os.ReadDir(KnowledgeSkillDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read knowledge skills directory")
	}

	// Create a map to store the metatypes
	knowledgeSkills := make(map[string]KnowledgeSkill, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", KnowledgeSkillDataPath, file.Name())

			var knowledgeSkill KnowledgeSkill
			if err := util.LoadStructFromYAML(filepath, &knowledgeSkill); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load knowledge skills")
			}

			knowledgeSkills[knowledgeSkill.Name] = knowledgeSkill
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded knowledge skills file")
	}

	logrus.WithFields(logrus.Fields{"count": len(knowledgeSkills)}).Info("Done loading knowledge skills")

	KnowledgeSkills = knowledgeSkills
}

func LoadKnowledgeSkill(name string) (*KnowledgeSkill, error) {
	var v KnowledgeSkill
	if err := util.LoadStructFromYAML(fmt.Sprintf(KnowledgeSkillFilename, name), &v); err != nil {
		return nil, err
	}

	return &v, nil
}
