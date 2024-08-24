package utils

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gliderlabs/ssh"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
	"golang.org/x/term"
)

func GenerateRandomUsername() string {
	randomSource := rand.NewSource(uint64(time.Now().UnixNano()))
	randomGenerator := rand.New(randomSource)
	randomNumber := randomGenerator.Intn(101)

	return fmt.Sprintf("Test%d", randomNumber)
}

var (
	requiredInputMsg = "{{You must enter a value.}}::#ff8700\n"
	invalidInputMsg  = "{{Invalid input, please try again.}}::#ff8700\n"
)

func PromptConfirmInput(s ssh.Session, prompt string) (bool, error) {
	var input string

	t := term.NewTerminal(s, cfmt.Sprint(prompt))
	input, err := t.ReadLine()
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

	io.WriteString(s, cfmt.Sprintf(invalidInputMsg))

	return PromptConfirmInput(s, prompt)
}

func PromptUserInput(s ssh.Session, prompt string) (string, error) {
	var input string

	t := term.NewTerminal(s, cfmt.Sprint(prompt))
	input, err := t.ReadLine()
	if err != nil {
		logrus.WithError(err).Error("Error reading input")

		return "", err
	}

	if input == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		return PromptUserInput(s, prompt)
	}

	return input, nil
}

func PromptUserPasswordInput(s ssh.Session, prompt string) (string, error) {
	var input string
	t := term.NewTerminal(s, "")

	input, err := t.ReadPassword(cfmt.Sprint(prompt))
	if err != nil {
		logrus.WithError(err).Error("Error reading input")
		return "", err
	}

	if input == "" {
		io.WriteString(s, cfmt.Sprintf(requiredInputMsg))
		return PromptUserPasswordInput(s, prompt)
	}

	return input, nil
}
