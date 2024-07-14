package util

import (
	"os"

	"gopkg.in/yaml.v3"
)

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
