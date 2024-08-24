package screen

// import (
// 	"io"
// 	"strconv"
// 	"strings"

// 	"github.com/Jasrags/ShadowMUD/common/user"

// 	"github.com/i582/cfmt/cmd/cfmt"
// 	"github.com/sirupsen/logrus"
// )

// type (
// 	Commands map[string]Command
// 	Command  struct {
// 		Name   string   // Command name, e.g., "say", "look"
// 		Args   []string // Arguments for the command
// 		Sender struct {
// 			ID   string
// 			Name string
// 		}
// 		Recipient struct {
// 			ID   string
// 			Name string
// 		}
// 	}
// )

// func (s *Screens) EnterGame(u *user.User) int {
// 	// c := common.NewCharacter()
// 	// c.Name = "Test Character"
// 	// c.MetatypeID = "human"
// 	// c.Metatype = w.metatypes[c.MetatypeID]
// 	// c.UserID = u.ID
// 	// c.RoomID = r.ID
// 	// c.Room = r

// 	// u.AddCharacter(c)
// 	// // s.user.Characters[c.ID] = c
// 	// // u.Character = c
// 	// u.SetActiveCharacterByID(c.ID)
// 	// c.Room.AddCharacter(c)

// 	if len(u.Characters) == 0 {
// 		io.WriteString(s.session, cfmt.Sprintf(noCharactersCreatedMsg))
// 		return StatePromptCreateCharacter
// 	}

// 	io.WriteString(s.session, cfmt.Sprintf("{{Choose a character to enter the game:}}::#00ff00\n"))

// 	i := 0
// 	choiceIdMap := make(map[string]string)
// 	for _, c := range u.Characters {
// 		choiceIdMap[strconv.Itoa(i+1)] = c.ID
// 		io.WriteString(s.session, cfmt.Sprintf(characterListOption, i+1, c.Name))
// 		i++
// 	}

// 	for k, v := range choiceIdMap {
// 		logrus.WithFields(logrus.Fields{"key": k, "value": v}).Info("Choice ID Map")
// 	}

// 	choice, errReadLine := s.term.ReadLine()
// 	if errReadLine != nil {
// 		logrus.WithError(errReadLine).Error("Error reading menu choice")

// 		return StateQuit
// 	}

// 	choice = strings.ToLower(strings.TrimSpace(choice))
// 	logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received character choice")

// 	c := u.GetCharacterByID(choiceIdMap[choice])
// 	u.SetActiveCharacter(c)

// 	// var r *common.Room
// 	if u.Character.RoomID == "" {
// 		logrus.Warn("Character has no room, adding to default room")

// 		// r := common.NewRoom(&common.CoreRooms[0])
// 		// r.AddCharacter(c)
// 	}

// 	// r.AddCharacter(c)

// 	return StateGameLoop

// 	// u.ActiveCharacter = c

// 	// i := 1
// 	// for _, c := range s.user.Characters {
// 	// 	io.WriteString(s.session, cfmt.Sprintf(characterListOption, i+1, c.Name))
// 	// 	i++
// 	// }

// 	// choice, errReadLine := u.Term.ReadLine()
// 	// if errReadLine != nil {
// 	// 	logrus.WithError(errReadLine).Error("Error reading menu choice")

// 	// 	return StateQuit
// 	// }

// 	// choice = strings.ToLower(strings.TrimSpace(choice))
// 	// logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received character choice")

// 	// return StatePromptMainMenu
// }

// // TODO: Add the AutocompleteCallback
// func (s *Screens) GameLoop() int {
// 	// s.user.Term.AutoCompleteCallback = w.AutoCompleteCallback

// 	for {
// 		io.WriteString(s.session, cfmt.Sprintf(gameLoopPrompt))
// 		line, err := s.user.Term.ReadLine()
// 		if err != nil {
// 			logrus.WithError(err).Error("Error reading line")

// 			return StateQuit

// 		}
// 		line = strings.TrimSpace(line)
// 		logrus.WithField("line", line).Debug("Received line")
// 		io.WriteString(s.session, cfmt.Sprintf(inputEchoMsg, line))

// 		// Parse the input into a command and its arguments
// 		// parts := strings.SplitN(line, " ", 2)
// 		// name := strings.ToLower(parts[0])
// 		// var cmd Command
// 		// if len(parts) == 1 {
// 		// 	cmd = Command{Name: name}
// 		// } else if len(parts) == 2 {
// 		// 	cmd = Command{Name: name, Args: strings.Split(parts[1], " ")}
// 		// }

// 		// Send the command to the input queue
// 		// w.commandQueue <- cmd
// 	}
// }
