package config

import "time"

type Server struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	LogLevel       string `yaml:"log_level"`
	MaxConnections int    `yaml:"max_connections"`
	Timeouts       struct {
		Idle time.Duration `yaml:"idle"`
	}
	Data struct {
		BaseDir       string `yaml:"base_dir"`
		TestDir       string `yaml:"test_dir"`
		CharactersDir string `yaml:"characters_dir"`
		ZonesDir      string `yaml:"zones_dir"`
		RoomsDir      string `yaml:"rooms_dir"`
	} `yaml:"data"`
	BannedNames []string `yaml:"banned_names"`
}
