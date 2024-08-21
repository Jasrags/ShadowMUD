package common

// func SaveFile[T any](filename string, data T) error {
//     buffer := bytes.Buffer{}
//     encoder := gob.NewEncoder(&buffer)
//     err := encoder.Encode(data)
//     if err != nil {
//         return err
//     }
//     return nil
// }
// func Serialize[T any](data T) ([]byte, error) {
// 	buffer := bytes.Buffer{}
// 	encoder := gob.NewEncoder(&buffer)
// 	err := encoder.Encode(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buffer.Bytes(), nil
// }

//   func Deserialize[T any](b []byte) (T, error) {
//     buffer := bytes.Buffer{}
//     buffer.Write(b)
//     decoder := gob.NewDecoder(&buffer)
//     var data T
//     err := decoder.Decode(&data)
//     if err != nil {
//       return data, err
//     }
//     return data, nil
//   }

// type Loadable[K comparable] interface {
// 	Id() K // Must be a unique identifier for the data
// }

// func LoadFile[T RoomSpec](filename string) (T, error) {
// 	logrus.WithFields(logrus.Fields{"filename": filename}).Debug("Loading file")
// 	var v T

// 	if err := utils.LoadStructFromYAML(filename, v); err != nil {
// 		logrus.WithError(err).WithField("filename", filename).Error("Could not load file")
// 		return v, err
// 	}

// 	logrus.WithFields(logrus.Fields{"filename": filename}).Debug("Loaded file")
// 	return v, nil
// }

// func LoadFiles[T RoomSpec](path string) ([]T, error) {
// 	logrus.WithFields(logrus.Fields{"path": path}).Debug("Started loading files")
// 	list := []T{}

// 	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
// 		if info.IsDir() {
// 			subdir := info.Name()
// 			subfiles, err := os.ReadDir(path)
// 			if err != nil {
// 				logrus.WithError(err).WithFields(
// 					logrus.Fields{"path": path, "subdir": subdir}).
// 					Fatal("Could not read subdirectory")
// 			}
// 			for _, subfile := range subfiles {
// 				if strings.HasSuffix(subfile.Name(), ".yaml") {
// 					// name := strings.TrimSuffix(subfile.Name(), ".yaml")

// 					v, err := LoadFile[T](fmt.Sprintf("%s/%s", path, subfile.Name()))
// 					if err != nil {
// 						logrus.WithError(err).WithFields(
// 							logrus.Fields{"path": path, "subdir": subdir, "filename": subfile.Name()}).
// 							Fatal("Could not load file")
// 					}
// 					list = append(list, v)
// 					// if err := LoadRoom(fmt.Sprintf("%s/%s", subdir, name), v); err != nil {
// 					// logrus.WithFields(logrus.Fields{"filename": subfile.Name(), "subdirectory": subdir}).WithError(err).Fatal("Could not load room")
// 					// }
// 					logrus.WithFields(logrus.Fields{"filename": subfile.Name(), "subdirectory": subdir}).Debug("Loaded file")
// 				}
// 			}
// 		}
// 		return nil
// 	}); err != nil {
// 		logrus.WithFields(logrus.Fields{"path": path}).WithError(err).Fatal("Could not walk directory")
// 	}

// 	// for _, filepath := range filepaths {
// 	//     v, err := LoadFile(filepath)
// 	//     if err != nil {
// 	//         logrus.WithError(err).WithField("filepath", filepath).Fatal("Could not load file")
// 	//     }

// 	//     list[v.ID] = v
// 	//     logrus.WithFields(logrus.Fields{"filepath": filepath}).Debug("Loaded file")
// 	// }

// 	logrus.WithFields(logrus.Fields{"path": path}).WithFields(
// 		logrus.Fields{"count": len(list)}).Debug("Done loading files")

// 	return list, nil
// }
