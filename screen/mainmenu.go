package screen

import (
	"io"
	"strings"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
)

func (s *Screens) PromptMainMenu() int {
	io.WriteString(s.user.Session, cfmt.Sprintf(mainMenuTitle))
	io.WriteString(s.user.Session, cfmt.Sprintf(menuOptionEnterGame))
	io.WriteString(s.user.Session, cfmt.Sprintf(menuOptionCreateCharacter, len(s.user.Characters), s.cfg.UserCharacterMaxCount))
	io.WriteString(s.user.Session, cfmt.Sprintf(menuOptionListCharacters))
	io.WriteString(s.user.Session, cfmt.Sprintf(menuOptionDeleteCharacter))
	io.WriteString(s.user.Session, cfmt.Sprintf(menuOptionChangePassword))
	io.WriteString(s.user.Session, cfmt.Sprintf(menuOptionQuit))
	io.WriteString(s.user.Session, cfmt.Sprint(menuPrompt))

	menuChoice, errReadLine := s.user.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading menu choice")

		return StateQuit
	}

	menuChoice = strings.ToLower(strings.TrimSpace(menuChoice))
	logrus.WithFields(logrus.Fields{"choice": menuChoice}).Info("Received menu choice")

	switch menuChoice {
	case "1":
		return StateEnterGame
	case "2":
		return StatePromptCreateCharacter
	case "3":
		return StatePromptListCharacters
	case "4":
		return StatePromptDeleteCharacter
	case "5":
		return StatePromptChangePassword
	case "0":
		return StateQuit
	default:
		io.WriteString(s.user.Session, cfmt.Sprintf(menuInvalidChoice, menuChoice))
		return StatePromptMainMenu
	}
}
