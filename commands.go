package main

// import (
// 	"io"
// 	"strings"

// 	"github.com/Jasrags/ShadowMUD/common/character"
// 	"github.com/Jasrags/ShadowMUD/common/shared"
// 	"github.com/i582/cfmt/cmd/cfmt"
// 	"github.com/sirupsen/logrus"

// 	"github.com/gliderlabs/ssh"
// )

// type (
// 	BuildCommandFunc func(s ssh.Session, c *character.Character, args []string) error
// 	BuildCommand     struct {
// 		s ssh.Session
// 		c *character.Character
// 	}
// 	GameCommandFunc    func() error
// 	GameCommandMessage struct {
// 		FromID  string
// 		ToID    string
// 		Command string
// 		Args    []string
// 	}
// 	GameCommandResponse struct {
// 		err error
// 	}
// 	GameCommand struct {
// 		s               ssh.Session
// 		c               *character.Character
// 		ResponseChannel chan GameCommandResponse
// 		InputChannel    chan GameCommandMessage
// 	}
// )

// func NewBuildCommand(s ssh.Session, c *character.Character) *BuildCommand {
// 	bc := &BuildCommand{
// 		s: s,
// 		c: c,
// 	}

// 	return bc
// }

// func (bc *BuildCommand) parseAndExecuteCommand(command string) error {
// 	args := strings.Fields(command)
// 	if len(args) == 0 {
// 		return shared.ErrCommandNotProvided
// 	}

// 	cmd, ok := BuildCommands[strings.ToLower(args[0])]
// 	if !ok {
// 		return shared.ErrCommandUnknown
// 	}

// 	return cmd(bc.s, bc.c, args[1:])
// }

// func NewGameCommand(responseChannel chan string, inputChannel chan GameCommandMessage) *GameCommand {
// 	gc := &GameCommand{
// 		ResponseChannel: make(chan GameCommandResponse),
// 		InputChannel:    inputChannel,
// 	}

// 	return gc
// }

// var (
// 	// BuildCommands is a map of build commands
// 	BuildCommands = map[string]BuildCommandFunc{
// 		"help": w.buildCommandDisplayHelp,
// 		// "help": func(s ssh.Session, c *character.Character, args []string) error {
// 		// 	return shared.ErrNotImplemented
// 		// },
// 		"set": func(s ssh.Session, c *character.Character, args []string) error {
// 			return shared.ErrNotImplemented
// 		},
// 	}

// 	// GameCommands is a map of game commands
// 	GameCommands = map[string]GameCommandFunc{
// 		"echo": func() error {
// 			return shared.ErrNotImplemented
// 		},
// 	}
// )

// func (w *World) buildCommandDisplayHelp(s ssh.Session, c *character.Character, args []string) error {
// 	if err := w.templates.ExecuteTemplate(s, "character_builder_help.tmpl", nil); err != nil {
// 		logrus.WithError(err).Error("Error executing template")
// 		io.WriteString(s, cfmt.Sprintf("An error occurred while displaying help"))
// 		return err
// 	}
// 	return nil
// }
