package config

type Server struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	LogLevel string `yaml:"log_level"`
	// DataDirectories struct {
	// Characters         string `yaml:"characters"`
	// Config             string `yaml:"config"`
	// Contacts           string `yaml:"contacts"`
	// Armor              string `yaml:"armor"`
	// ArmorModifications string `yaml:"armor_modifications"`
	// } `yaml:"data_directories"`
}
