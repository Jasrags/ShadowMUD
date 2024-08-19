package screen

import (
	"io"
	"strconv"
	"strings"

	"github.com/Jasrags/ShadowMUD/common"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
)

func PromptConfirmInput(u *common.User, prompt string) (bool, error) {
	var input string

	io.WriteString(u.Session, cfmt.Sprint(prompt))
	input, err := u.Term.ReadLine()
	if err != nil {
		logrus.WithError(err).Error("Error reading input")
		return false, err
	}

	input = strings.TrimSpace(input)

	switch input {
	case "y", "Y", "yes", "YES":
		return true, nil
	case "n", "N", "no", "NO":
		return false, nil
	}

	io.WriteString(u.Session, cfmt.Sprintf("Invalid input, please try again\n"))
	return PromptConfirmInput(u, prompt)
}

func PromptUserInput(u *common.User, prompt string) (string, error) {
	var input string

	io.WriteString(u.Session, cfmt.Sprint(prompt))
	input, err := u.Term.ReadLine()
	if err != nil {
		logrus.WithError(err).Error("Error reading input")

		return "", err
	}

	if input == "" {
		io.WriteString(u.Session, cfmt.Sprintf(requiredInputMsg))
		return PromptUserInput(u, prompt)
	}

	return input, nil
}

func PromptUserPasswordInput(u *common.User, prompt string) (string, error) {
	var input string

	input, err := u.Term.ReadPassword(cfmt.Sprint(prompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading input")
		return "", err
	}

	if input == "" {
		io.WriteString(u.Session, cfmt.Sprintf(requiredInputMsg))
		return PromptUserPasswordInput(u, prompt)
	}

	return input, nil
}

type SelectOption struct {
	Text  string
	Value string
}

func PromptUserSelect(u *common.User, title, prompt string, options []SelectOption) (string, error) {
	io.WriteString(u.Session, cfmt.Sprintf(title))

	for i, option := range options {
		io.WriteString(u.Session, cfmt.Sprintf("%d. %s", i+1, option.Text))
	}

	io.WriteString(u.Session, cfmt.Sprintf(prompt))
	choice, err := u.Term.ReadLine()
	if err != nil {
		logrus.WithError(err).Error("Error reading input")
		return "", err
	}

	choice = strings.TrimSpace(choice)
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(options) {
		io.WriteString(u.Session, cfmt.Sprintf("Invalid input, please try again\n"))
		return PromptUserSelect(u, title, prompt, options)
	}

	return options[index-1].Value, nil
}
