package main

import (
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/utils"
)

const (
	ConfigFilepath = "_data/config/server.yaml"
)

func main() {
	var serverConfig config.Server
	utils.LoadStructFromYAML(ConfigFilepath, &serverConfig)

	w := core.NewWorld(serverConfig)
	w.Start()
}
