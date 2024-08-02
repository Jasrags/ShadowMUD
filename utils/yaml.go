package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type LoadableSimple interface {
	Validate() error  // General validation (or none)
	Filepath() string // Relative file path to some base directory - can include subfolders
}

func LoadAllFiles[T LoadableSimple](basePath string) ([]T, error) {
	loadedData := make([]T, 0, 128)

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".yaml" {
			return nil
		}

		loaded, err := LoadFlatFile[T](path)
		if err != nil {
			return err
		}

		loadedData = append(loadedData, loaded)

		return nil
	})

	return loadedData, err
}

func LoadFlatFile[T LoadableSimple](path string) (T, error) {
	var loaded T

	path = filepath.FromSlash(path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		return loaded, fmt.Errorf(`filepath: %s, %w`, path, err)
	}

	if fileInfo.IsDir() {
		return loaded, fmt.Errorf(`filepath: %s, %w`, path, errors.New(`is a directory`))
	}

	if filepath.Ext(path) != ".yaml" {
		bytes, err := os.ReadFile(path)
		if err != nil {
			return loaded, fmt.Errorf(`filepath: %s, %w`, path, err)
		}

		err = yaml.Unmarshal(bytes, &loaded)
		if err != nil {
			return loaded, fmt.Errorf(`filepath: %s, %w`, path, err)
		}
	}

	// Make sure the Filepath it claims is correct in case we need to save it later
	if !strings.HasSuffix(path, loaded.Filepath()) {
		return loaded, fmt.Errorf(`filepath: %s, %w`, path, fmt.Errorf(`filesystem path "%s" did not end in Filepath() "%s" for type %T`, path, loaded.Filepath(), loaded))
	}

	// validate the structure
	if err := loaded.Validate(); err != nil {
		return loaded, fmt.Errorf(`filepath: %s, %w`, path, err)
	}

	return loaded, nil
}

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
