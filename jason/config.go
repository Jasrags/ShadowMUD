package jason

import (
	"regexp"
)

type Config struct {
	BuildPointKarma        int
	CharacterNameRegex     *regexp.Regexp
	CharacterNameMinLength int
	CharacterNameMaxLength int
	CharacterNames         []string
	BannedNames            []string
}

var DefaultConfig = Config{
	BuildPointKarma:        800,
	CharacterNameRegex:     regexp.MustCompile("^[a-zA-Z0-9]*$"),
	CharacterNameMinLength: 3,
	CharacterNameMaxLength: 16,
	CharacterNames:         []string{"Jason", "John", "Jane"},
	BannedNames:            []string{"Admin", "Moderator"},
}
