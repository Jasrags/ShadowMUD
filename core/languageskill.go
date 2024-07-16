package core

import (
	"fmt"
	"os"
	"strings"

	"shadowrunmud/core/util"

	"github.com/sirupsen/logrus"
)

const (
	LanguageSkillDataPath       = "data/skills/languages"
	LanguageSkillFilename       = LanguageSkillDataPath + "/%s.yaml"
	LanguageSkillFileMinVersion = "0.0.1"
)

var (
	LanguageSkills        = map[string]LanguageSkill{}
	DefaultLanguageSkills = map[string]LanguageSkill{
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

type LanguageSkill struct {
	ID          string `yaml:"id,omitempty"`
	Name        string `yaml:"name"`
	IsCommon    bool   `yaml:"is_common"`
	Rank        int    `yaml:"rank,omitempty"`
	RuleSource  string `yaml:"rule_source"`
	FileVersion string `yaml:"file_version"`
}

func LoadLanguageSkills() map[string]LanguageSkill {
	logrus.Info("Started loading language skills")

	files, errReadDir := os.ReadDir(LanguageSkillDataPath)
	if errReadDir != nil {
		logrus.WithError(errReadDir).Fatal("Could not read language skills directory")
	}

	// Create a map to store the metatypes
	languageSklls := make(map[string]LanguageSkill, len(files))

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			filepath := fmt.Sprintf("%s/%s", LanguageSkillDataPath, file.Name())

			var languageSkill LanguageSkill
			if err := util.LoadStructFromYAML(filepath, &languageSkill); err != nil {
				logrus.WithFields(logrus.Fields{"filename": file.Name()}).WithError(err).Fatal("Could not load language skills")
			}

			languageSklls[languageSkill.ID] = languageSkill
		}
		logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded language skills file")
	}

	logrus.WithFields(logrus.Fields{"count": len(languageSklls)}).Info("Done loading language skills")

	return languageSklls
}

func LoadLanguageSkill(name string) (*LanguageSkill, error) {
	var v LanguageSkill
	if err := util.LoadStructFromYAML(fmt.Sprintf(LanguageSkillFilename, name), &v); err != nil {
		return nil, err
	}

	return &v, nil
}
