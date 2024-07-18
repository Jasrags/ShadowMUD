package core

type ContactType string

const (
	ContactTypeFixer ContactType = "Fixer"
)

type Contact struct {
	ID          string      `yaml:"id,omitempty"`
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Type        ContactType `yaml:"type"`
	Connection  int         `yaml:"connection"`
	Loyalty     int         `yaml:"loyalty"`
	RuleSource  string      `yaml:"rule_source"`
	FileVersion string      `yaml:"file_version"`
}

var CoreContacts = []Contact{}

// TODO: Load the data from the yaml files
func LoadContacts() map[string]Contact {
	data := make(map[string]Contact)
	return data
}
