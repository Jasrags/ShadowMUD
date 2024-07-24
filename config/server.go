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
	DataFiles struct {
		Version          string `yaml:"version"`
		BaseDir          string `yaml:"base_dir"`
		TestDir          string `yaml:"test_dir"`
		CharacerFilesDir string `yaml:"character_files_dir"`

		// Characters         string `yaml:"characters"`
		// Config             string `yaml:"config"`
		// Contacts           string `yaml:"contacts"`
		// Armor              string `yaml:"armor"`
		// ArmorModifications string `yaml:"armor_modifications"`
	} `yaml:"data_directories"`
	BannedNames []string `yaml:"banned_names"`
}
