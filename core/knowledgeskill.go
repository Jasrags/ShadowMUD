package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jasrags/ShadowMUD/core/util"

	"github.com/sirupsen/logrus"
)

const (
	KnowledgeSkillDataPath       = "data/skills/knowledge"
	KnowledgeSkillFilename       = KnowledgeSkillDataPath + "/%s.yaml"
	KnowledgeSkillFileMinVersion = "0.0.1"
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

func LoadKnowledgeSkills() map[string]KnowledgeSkill {
	logrus.Info("Started loading knowledge skills")

	files, errReadDir := os.ReadDir(KnowledgeSkillDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read knowledge skills directory")
	}

	// Create a map to store the metatypes
	list := make(map[string]KnowledgeSkill, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", KnowledgeSkillDataPath, file.Name())

			var v KnowledgeSkill
			if err := util.LoadStructFromYAML(filepath, &v); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load knowledge skills")
			}

			list[v.ID] = v
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded knowledge skills file")
	}

	logrus.WithFields(logrus.Fields{"count": len(list)}).Info("Done loading knowledge skills")

	return list
}

func LoadKnowledgeSkill(name string) (*KnowledgeSkill, error) {
	var v KnowledgeSkill
	if err := util.LoadStructFromYAML(fmt.Sprintf(KnowledgeSkillFilename, name), &v); err != nil {
		return nil, err
	}

	return &v, nil
}
