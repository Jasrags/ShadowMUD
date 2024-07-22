package utils

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func FormatFilename(filename string) string {
	r := strings.NewReplacer(" ", "_", "/", "_", "\\", "_", ":", "_", "*", "_", "?", "_", "\"", "_", "<", "_", ">", "_", "|", "_")
	return strings.ToLower(r.Replace(filename))
}

func SaveStructToYAML(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func LoadStructFromYAML(filename string, data interface{}) error {
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
