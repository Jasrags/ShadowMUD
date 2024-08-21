package contact

// func LoadContacts(dataPath string) map[string]Contact {
// 	data := make(map[string]Contact)

// 	logrus.WithFields(logrus.Fields{"data_path": dataPath}).Info("Started loading contacts")
// 	files, err := os.ReadDir(dataPath)
// 	if err != nil {
// 		logrus.WithError(err).Error("Could not read contacts data directory")
// 		return data
// 	}

// 	for _, file := range files {
// 		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
// 			filePath := fmt.Sprintf("%s/%s", dataPath, file.Name())

// 			var v Contact
// 			err := utils.LoadStructFromYAML(filePath, &v)
// 			if err != nil {
// 				logrus.WithError(err).WithFields(logrus.Fields{"filename": file.Name()}).Error("Could not load contact")
// 				return data
// 			}
// 			data[v.ID] = v
// 			logrus.WithFields(logrus.Fields{"filename": file.Name()}).Debug("Loaded contact file")
// 		}
// 	}

// 	logrus.WithFields(logrus.Fields{"count": len(data)}).Info("Done loading contacts")

// 	return data
// }
