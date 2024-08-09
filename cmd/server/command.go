package main

type (
	Commands map[string]Command
	Command  struct {
		Name   string   // Command name, e.g., "say", "look"
		Args   []string // Arguments for the command
		Sender struct {
			ID   string
			Name string
		}
		Recipient struct {
			ID   string
			Name string
		}
	}
)
