package shared

import "fmt"

var (
	// General input errors
	ErrNotImplemented     = fmt.Errorf("not_implemented")
	ErrInvalidInput       = fmt.Errorf("invalid_input")
	ErrInvalidChoice      = fmt.Errorf("invalid_choice")
	ErrCommandNotProvided = fmt.Errorf("command_not_provided")
	ErrCommandUnknown     = fmt.Errorf("command_unknown")
	// Login
	ErrInvalidNameOrPassword = fmt.Errorf("invalid_name_or_password")
	// Name
	ErrNameNotAllowed      = fmt.Errorf("name_not_allowed")
	ErrNameLength          = fmt.Errorf("name_length")
	ErrNameExists          = fmt.Errorf("name_exists")
	ErrNameNotAlphanumeric = fmt.Errorf("name_not_alphanumeric")
	// Password
	ErrPasswordLength   = fmt.Errorf("password_length")
	ErrPasswordMismatch = fmt.Errorf("password_mismatch")
	// Character
	ErrCharacterMaxCount = fmt.Errorf("character_max_count")
)
