package config

import "time"

type Server struct {
	Host                     string        `yaml:"host"`
	Port                     string        `yaml:"port"`
	LogLevel                 string        `yaml:"log_level"`
	MaxConnections           int           `yaml:"max_connections"`
	LoginEnabled             bool          `yaml:"login_enabled"`
	LoginMaxAttempts         int           `yaml:"login_max_attempts"`
	UsernameMinLength        int           `yaml:"username_min_length"`
	UsernameMaxLength        int           `yaml:"username_max_length"`
	CharacterNameMinLength   int           `yaml:"character_name_min_length"`
	CharacterNameMaxLength   int           `yaml:"character_name_max_length"`
	PasswordMinLength        int           `yaml:"password_min_length"`
	PasswordMaxLength        int           `yaml:"password_max_length"`
	PasswordBcryptCost       int           `yaml:"password_bcrypt_cost"`
	RegistrationEnabled      bool          `yaml:"registration_enabled"`
	CharacterCreationEnabled bool          `yaml:"character_creation_enabled"`
	UserCharacterMaxCount    int           `yaml:"user_character_max_count"`
	IdleTimeout              time.Duration `yaml:"idle_timeout"`
	BannedNames              []string      `yaml:"banned_names"`
}
