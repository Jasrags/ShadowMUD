package screen

import (
	"io"
	"strconv"
	"strings"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
)

func (s *Screens) PromptConfirmInput(prompt string) (bool, error) {
	var input string

	io.WriteString(s.user.Session, cfmt.Sprint(prompt))
	input, err := s.user.Term.ReadLine()
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

	io.WriteString(s.user.Session, cfmt.Sprintf("Invalid input, please try again\n"))
	return s.PromptConfirmInput(prompt)
}

func (s *Screens) PromptUserInput(prompt string) (string, error) {
	var input string

	io.WriteString(s.user.Session, cfmt.Sprint(prompt))
	input, err := s.user.Term.ReadLine()
	if err != nil {
		logrus.WithError(err).Error("Error reading input")

		return "", err
	}

	if input == "" {
		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
		return s.PromptUserInput(prompt)
	}

	return input, nil
}

func (s *Screens) PromptUserPasswordInput(prompt string) (string, error) {
	var input string

	input, err := s.user.Term.ReadPassword(cfmt.Sprint(prompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading input")
		return "", err
	}

	if input == "" {
		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
		return s.PromptUserPasswordInput(prompt)
	}

	return input, nil
}

type SelectOption struct {
	Text  string
	Value string
}

func (s *Screens) PromptUserSelect(title, prompt string, options []SelectOption) (string, error) {
	io.WriteString(s.user.Session, cfmt.Sprintf(title))

	for i, option := range options {
		io.WriteString(s.user.Session, cfmt.Sprintf("%d. %s", i+1, option.Text))
	}

	io.WriteString(s.user.Session, cfmt.Sprintf(prompt))
	choice, err := s.user.Term.ReadLine()
	if err != nil {
		logrus.WithError(err).Error("Error reading input")
		return "", err
	}

	choice = strings.TrimSpace(choice)
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(options) {
		io.WriteString(s.user.Session, cfmt.Sprintf("Invalid input, please try again\n"))
		return s.PromptUserSelect(title, prompt, options)
	}

	return options[index-1].Value, nil
}
