package main

// const (
// 	ConfigFilepath = "_data/config/server.yaml"
// )

// type World struct {
// 	config config.Server
// }

// func NewWorld(osSignalChan chan os.Signal) *World {
// 	var serverConfig config.Server
// 	utils.LoadStructFromYAML(ConfigFilepath, &serverConfig)

// 	// logrus.WithField("serverConfig", serverConfig).Info("Loaded server configuration")

// 	logrusLevel, err := logrus.ParseLevel(serverConfig.LogLevel)
// 	if err != nil {
// 		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
// 		logrusLevel = logrus.InfoLevel
// 	}
// 	logrus.SetLevel(logrusLevel)
// 	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

// 	w := &World{
// 		config: serverConfig,
// 	}

// 	zones, err := utils.LoadAllFiles[common.ZoneSpec](serverConfig.Data.BaseDir + serverConfig.Data.ZonesDir)
// 	if err != nil {
// 		logrus.WithError(err).Fatal("Could not load zones")
// 	}
// 	for _, zone := range zones {
// 		logrus.WithFields(logrus.Fields{"zone_id": zone.ID, "zone_name": zone.Name}).Info("Loaded zone")
// 		// w.zones[zone.ID] = zone
// 	}
// 	// utils.LoadAllFiles[common.RoomSpec](serverConfig.Data.BaseDir + common.RoomsDataPath)

// 	return w
// }
