package contact

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	ContactsFilepath = "_data/contacts"

	TypeFixer Type = "Fixer"
)

type (
	Type  string
	Specs map[string]*Spec
	Spec  struct {
		ID          string            `yaml:"id"`
		Name        string            `yaml:"name"`
		Description string            `yaml:"description"`
		Type        Type              `yaml:"type"`
		RuleSource  shared.RuleSource `yaml:"rule_source"`
	}
	Contacts map[string]*Contact
	Contact  struct {
		ID         string `yaml:"id"`
		Connection int    `yaml:"connection"`
		Loyalty    int    `yaml:"loyalty"`
		Spec       *Spec  `yaml:"-"`
	}
)

var CoreContacts = []Spec{
	{
		ID:          "brian_flannigan",
		Name:        "Brian Flannigan",
		Description: "A fixer is a person who arranges illicit goods or services for characters. Fixers are the go-to people for characters who need to buy or sell illegal goods, find a buyer for a stolen item, or hire a shadowrunner for a job.",
		Type:        TypeFixer,
		RuleSource:  shared.RuleSourceSR5Core,
	},
}
