package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func FormatFilename(filename string) string {
	r := strings.NewReplacer(" ", "_", "/", "_", "\\", "_", ":", "_", "*", "_", "?", "_", "\"", "_", "<", "_", ">", "_", "|", "_")
	return strings.ToLower(r.Replace(filename))
}

func SaveStructToYAML[T any](filename string, data T) error {
	log := logrus.WithFields(logrus.Fields{"filename": filename})
	log.Debug("Saving to YAML")

	file, err := os.Create(filename)
	if err != nil {
		log.WithError(err).Error("Could not create file")
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(data); err != nil {
		log.WithError(err).Error("Could not encode file")
		return err
	}

	log.Debug("Saved to YAML")

	return nil
}

func LoadStructFromYAML[T any](filename string, data T) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}

func LoadStructsFromDir[T any](filepath string) ([]T, error) {
	var list []T

	files, errReadDir := os.ReadDir(filepath)
	if errReadDir != nil {
		return list, errReadDir
	}

	for _, file := range files {
		var v T
		if strings.HasSuffix(file.Name(), ".yaml") {
			if err := LoadStructFromYAML(fmt.Sprintf("%s/%s", filepath, file.Name()), &v); err != nil {
				return list, err
			}

			list = append(list, v)
		}
	}

	return list, nil
}
