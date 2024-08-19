package screen

import (
	"io"
	"strings"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
)

func (s *Screens) PromptCreateCharacter() int {
	if len(s.user.Characters) >= s.cfg.UserCharacterMaxCount {
		io.WriteString(s.user.Session, cfmt.Sprintf(characterMaxCharactersMsg))
		return StatePromptMainMenu
	}

	// TODO: Add a blurb about the character creation process
	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuTitle))
	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionPregen))
	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionCustom))
	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionLearn))
	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionReturn))
	io.WriteString(s.user.Session, cfmt.Sprint(menuPrompt))

	choice, errReadLine := s.user.Term.ReadLine()
	if errReadLine != nil {
		logrus.WithError(errReadLine).Error("Error reading menu choice")

		return StateQuit
	}

	choice = strings.TrimSpace(strings.ToLower(strings.TrimSpace(choice)))
	logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received menu choice")

	switch choice {
	case "1":
		io.WriteString(s.user.Session, cfmt.Sprintf(featureNotImplementedMsg))
		// TODO list our pre-generated characters
		return StatePromptCreateCharacter
	case "2":
		io.WriteString(s.user.Session, cfmt.Sprintf(featureNotImplementedMsg))
		// TODO: Prompt for custom character creation
		// Choose name
		// Choose metatype
		// Choose attributes
		return StatePromptCreateCharacter
	case "3":
		io.WriteString(s.user.Session, cfmt.Sprintf(featureNotImplementedMsg))
		// TODO: Show longer description about shadowrun characters, archtypes, etc...
		return StatePromptCreateCharacter
	case "4":
		return StatePromptMainMenu
	default:
		io.WriteString(s.user.Session, cfmt.Sprintf(menuInvalidChoice, choice))
		return StatePromptCreateCharacter
	}

	// TODO: Check that the name is not banned
	// TODO: Check that the name is not already taken
}

func (s *Screens) PromptChooseCharacterName(u *common.User) int {
promptInputCharacterName:
	name, err := PromptUserInput(u, cfmt.Sprintf("Enter a name for your character: "))
	if err != nil {
		io.WriteString(s.user.Session, cfmt.Sprintf(inputErrorMsg))
		return StateQuit
	}

	name = strings.TrimSpace(name)
	if name == "" {
		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
		return StatePromptCreateCharacter
	}

promptConfirmCharacterName:
	// TODO: Check that the name is not banned
	// TODO: Check that the name is not already taken
	// TODO: Check that the name is not too long or too short
	// TODO: Check that the name does not contain invalid characters

	confirmName, err := PromptUserInput(u, cfmt.Sprintf("Confirm character name '%s' (y/n): ", name))
	if err != nil {
		io.WriteString(s.user.Session, cfmt.Sprintf(inputErrorMsg))
		return StateQuit
	}

	if confirmName == "" {
		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
		goto promptConfirmCharacterName
	}

	confirmName = strings.TrimSpace(confirmName)
	if !strings.EqualFold(confirmName, "y") {
		io.WriteString(s.user.Session,
			cfmt.Sprintf("Character name '%s' was not confirmed.", name))
		goto promptInputCharacterName
	}

	return StatePromptCreateCharacter
}
