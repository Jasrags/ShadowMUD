package skill

import (
	"fmt"
	"os"
	"shadowrunmud/util"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	KnowledgeSkillDataPath                             = "data/skills/knowledge"
	KnowledgeSkillFilename                             = KnowledgeSkillDataPath + "/%s.yaml"
	KnowledgeSkillFileMinVersion                       = "0.0.1"
	KnowledgeSkillSeattleStreetGangs KnowledgeSkillIdx = iota
)

type (
	KnowledgeSkillIdx int
)

var (
	KnowledgeSkills = map[string]KnowledgeSkill{}
)

type KnowledgeSkill interface {
	GetID() string
	SetID(string)
	GetName() string
	GetDescription() string
	GetIsCommon() bool
	GetRank() int
	SetRank(int)
	GetRuleSource() string
	GetFileVersion() string
}

type knowledgeSkill struct {
	ID          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	IsCommon    bool   `yaml:"is_common"`
	Rank        int    `yaml:"rank,omitempty"`
	RuleSource  string `yaml:"rule_source"`
	FileVersion string `yaml:"file_version"`
}

func (s knowledgeSkill) GetID() string {
	return s.ID
}

func (s *knowledgeSkill) SetID(id string) {
	s.ID = id
}

func (s knowledgeSkill) GetName() string {
	return s.Name
}

func (s knowledgeSkill) GetDescription() string {
	return s.Description
}

func (s knowledgeSkill) GetIsCommon() bool {
	return s.IsCommon
}

func (s knowledgeSkill) GetRank() int {
	return s.Rank
}

func (s *knowledgeSkill) SetRank(rank int) {
	s.Rank = rank
}

func (s knowledgeSkill) GetRuleSource() string {
	return s.RuleSource
}

func (s knowledgeSkill) GetFileVersion() string {
	return s.FileVersion
}

func LoadKnowledgeSkills(wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Debug("Started loading knowledge skills")

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

			knowledgeSkills[knowledgeSkill.GetName()] = knowledgeSkill
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded knowledge skills file")
	}

	logrus.WithFields(logrus.Fields{"count": len(knowledgeSkills)}).Info("Done loading knowledge skills")

	KnowledgeSkills = knowledgeSkills
}
