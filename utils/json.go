package utils

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

func LoadStructFromJSON[T any](filename string, data T) error {
	log := logrus.WithFields(logrus.Fields{"filename": filename})
	log.Debug("Loading from JSON")

	file, err := os.Open(filename)
	if err != nil {
		log.WithError(err).Error("Could not open file")
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		log.WithError(err).Error("Could not decode file")
		return err
	}

	log.Debug("Loaded from JSON")

	return nil
}

// func LoadJSON[T any](filePath string, v T) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		logrus.WithError(err).Error("Failed to open file")
// 		return err
// 	}
// 	defer file.Close()

// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(v); err != nil {
// 		logrus.WithError(err).Error("Failed to decode JSON")
// 		return err
// 	}

// 	return nil
// }
