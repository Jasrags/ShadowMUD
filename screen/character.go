package screen

import (
	"io"

	"github.com/i582/cfmt/cmd/cfmt"
)

func (s *Screens) PromptListCharacters() int {
	if len(s.user.Characters) == 0 {
		io.WriteString(s.user.Session, cfmt.Sprintf(characterNoneCreatedMsg))
		return StatePromptMainMenu
	}
	for _, c := range s.user.Characters {
		io.WriteString(s.user.Session, cfmt.Sprintf("{{%s [%s]}}::#00ff00\n", c.Name, c.Metatype.Name))
	}
	return StatePromptMainMenu
}

func (s *Screens) PromptDeleteCharacter() int {
	io.WriteString(s.user.Session, cfmt.Sprintf(featureNotImplementedMsg))
	if len(s.user.Characters) == 0 {
		io.WriteString(s.user.Session, cfmt.Sprintf(characterNoneCreatedMsg))
		return StatePromptMainMenu
	}
	return StatePromptMainMenu
}
