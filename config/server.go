package config

type Server struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	LogLevel string `yaml:"log_level"`
}
