package command

import "strings"

var commandMap = CommandMap{
	"say": {
		Name: "say",
		Func: func(args []string) string {
			return "You said: " + strings.Join(args, " ")
		},
	},
	"look": {
		Name: "look",
		Func: func(args []string) string {
			return "You look around"
		},
	},
}

type CommandFunc func(args []string) string

type UserCommand struct {
	Name    string
	Aliases []string
	Func    CommandFunc
	IsAdmin bool
}

// CommandMap maps command words to functions
type CommandMap map[string]UserCommand

// RegisterCommand registers a command word with a corresponding function
func (m CommandMap) RegisterCommand(word string, fn CommandFunc) {
	// m[word] = fn
}

// ExecuteCommand executes a command based on the command word and returns the result
func (h *CommandHandler) ExecuteCommand(cmdWord string) CommandFunc {
	fn, ok := commandMap[cmdWord]
	if !ok {
		return nil
	}
	return fn.Func
}

// Command interface
type Command interface {
	Execute() string
	SetArgs(args []string)
}

// // SayCommand struct
// type SayCommand struct {
// 	Arguments []string
// }

// // Execute method for SayCommand
// func (c *SayCommand) Execute() string {
// 	message := strings.Join(c.Arguments, " ")
// 	return message
// }

// // SetArgs method for SayCommand
// func (c *SayCommand) SetArgs(args []string) {
// 	c.Arguments = args
// }

// // LookCommand struct
// type LookCommand struct {
// 	Arguments []string
// }

// // Execute method for LookCommand
// func (c *LookCommand) Execute() string {
// 	// Logic to look around in the game, ignoring arguments for simplicity
// 	return "Looking around..."
// }

// // SetArgs method for LookCommand
// func (c *LookCommand) SetArgs(args []string) {
// 	c.Arguments = args
// }

// CommandHandler to execute commands
type CommandHandler struct {
	// Can store a history of commands or other relevant data
}

// // ExecuteCommand executes a given Command with arguments and returns the result
// func (h *CommandHandler) ExecuteCommand(cmd Command, args []string) string {
// 	cmd.SetArgs(args)
// 	return cmd.Execute()
// }

// // Example usage
// func main() {
//     sayCommand := &SayCommand{}
//     lookCommand := &LookCommand{}

//     handler := CommandHandler{}

//     // Execute commands with arguments and print the results
//     fmt.Println(handler.ExecuteCommand(sayCommand, []string{"Hello,", "world!"}))
//     fmt.Println(handler.ExecuteCommand(lookCommand, []string{})) // LookCommand ignores arguments
// }
