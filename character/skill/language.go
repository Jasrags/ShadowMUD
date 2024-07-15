package skill

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"shadowrunmud/util"

	"github.com/sirupsen/logrus"
)

const (
	LanguageSkillDataPath       = "data/skills/languages"
	LanguageSkillFilename       = LanguageSkillDataPath + "/%s.yaml"
	LanguageSkillFileMinVersion = "0.0.1"
)

type ()

var (
	LanguageSkills        = map[string]LanguageSkill{}
	DefaultLanguageSkills = map[string]languageSkill{
		"English":    {Name: "English", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Spanish":    {Name: "Spanish", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Lakota":     {Name: "Lakota", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Dakota":     {Name: "Dakota", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Dine":       {Name: "Diné (Navajo)", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Russian":    {Name: "Russian", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"French":     {Name: "French", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Italian":    {Name: "Italian", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"German":     {Name: "German", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Aztlaner":   {Name: "Aztlaner Spanish", IsCommon: true, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Sperethiel": {Name: "Sperethiel", IsCommon: false, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Orzet":      {Name: "Or’zet", IsCommon: false, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Japanese":   {Name: "Japanese", IsCommon: false, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
		"Mandarin":   {Name: "Mandarin", IsCommon: false, RuleSource: "SR5:Core", FileVersion: "0.0.1"},
	}
)

type LanguageSkill interface {
	GetID() string
	SetID(string)
	GetName() string
	GetIsCommon() bool
	GetIsNative() bool
	SetIsNative(bool)
	GetRank() int
	SetRank(int)
	GetRuleSource() string
	GetFileVersion() string
}

type languageSkill struct {
	ID          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	IsCommon    bool   `yaml:"is_common"`
	IsNative    bool   `yaml:"is_native,omitempty"`
	Rank        int    `yaml:"rank,omitempty"`
	RuleSource  string `yaml:"rule_source"`
	FileVersion string `yaml:"file_version"`
}

func (ls *languageSkill) GetID() string {
	return ls.ID
}

func (ls *languageSkill) SetID(id string) {
	ls.ID = id
}

func (ls *languageSkill) GetName() string {
	return ls.Name
}

func (ls *languageSkill) GetIsCommon() bool {
	return ls.IsCommon
}

func (ls *languageSkill) GetIsNative() bool {
	return ls.IsNative
}

func (ls *languageSkill) SetIsNative(isNative bool) {
	ls.IsNative = isNative
}

func (ls *languageSkill) GetRank() int {
	return ls.Rank
}

func (ls *languageSkill) SetRank(rank int) {
	ls.Rank = rank
}

func (ls *languageSkill) GetRuleSource() string {
	return ls.RuleSource
}

func (ls *languageSkill) GetFileVersion() string {
	return ls.FileVersion
}

func LoadLanguageSkills(wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Debug("Started loading language skills")

	files, errReadDir := os.ReadDir(LanguageSkillDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read language skills directory")
	}

	// Create a map to store the metatypes
	languageSklls := make(map[string]LanguageSkill, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", LanguageSkillDataPath, file.Name())

			var languageSkll LanguageSkill
			if err := util.LoadStructFromYAML(filepath, &languageSkll); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load language skills")
			}

			languageSklls[languageSkll.GetName()] = languageSkll
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Info("Loaded language skills file")
	}

	logrus.WithFields(logrus.Fields{"count": len(languageSklls)}).Info("Done loading language skills")

	LanguageSkills = languageSklls
}
