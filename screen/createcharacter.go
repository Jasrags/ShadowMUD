package screen

// // Constants for menu options
// // const (
// // 	characterMaxCharactersMsg       = "You have reached the maximum number of characters."
// // 	createCharacterMenuTitle        = "=== Create Character ===\n"
// // 	createCharacterMenuOptionPregen = "1. Use a pre-generated character\n"
// // 	createCharacterMenuOptionCustom = "2. Create a custom character\n"
// // 	createCharacterMenuOptionLearn  = "3. Learn about character creation\n"
// // 	createCharacterMenuOptionReturn = "4. Return to main menu\n"
// // 	menuPrompt                      = "Please choose an option: "
// // )

// // Priority categories
// const (
// 	Metatype = iota
// 	Attributes
// 	Magic
// 	Skills
// 	Resources
// )

// // Priority options
// var priorityOptions = []string{"A", "B", "C", "D", "E"}

// // TODO: Prompt for character name
// // TODO: Prompt for character concept (if using pregen)
// // TODO: Prompt for character metatype

// func (s *Screens) PromptCreateCharacter() int {
// 	if len(s.user.Characters) >= s.cfg.UserCharacterMaxCount {
// 		io.WriteString(s.user.Session, cfmt.Sprintf(characterMaxCharactersMsg))
// 		return StatePromptMainMenu
// 	}

// 	// TODO: Add a blurb about the character creation process
// 	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuTitle))
// 	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionPregen))
// 	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionCustom))
// 	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionLearn))
// 	io.WriteString(s.user.Session, cfmt.Sprintf(createCharacterMenuOptionReturn))
// 	io.WriteString(s.user.Session, cfmt.Sprint(menuPrompt))

// 	choice, errReadLine := s.user.Term.ReadLine()
// 	if errReadLine != nil {
// 		logrus.WithError(errReadLine).Error("Error reading menu choice")

// 		return StateQuit
// 	}

// 	choice = strings.TrimSpace(strings.ToLower(strings.TrimSpace(choice)))
// 	logrus.WithFields(logrus.Fields{"choice": choice}).Info("Received menu choice")

// 	switch choice {
// 	case "1":
// 		io.WriteString(s.user.Session, cfmt.Sprintf(featureNotImplementedMsg))
// 		// TODO list our pre-generated characters
// 		return s.handlePregenCharacter()
// 		// return StatePromptCreateCharacter
// 	case "2":
// 		// io.WriteString(s.user.Session, cfmt.Sprintf(featureNotImplementedMsg))

// 		// TODO: Prompt for custom character creation
// 		// Choose name
// 		// Choose metatype
// 		// Choose attributes
// 		// return StatePromptCreateCharacter
// 		return s.handleCustomCharacter()
// 	case "3":
// 		io.WriteString(s.user.Session, cfmt.Sprintf(featureNotImplementedMsg))
// 		// TODO: Show longer description about shadowrun characters, archtypes, etc...
// 		return StatePromptCreateCharacter
// 	case "4":
// 		return StatePromptMainMenu
// 	default:
// 		io.WriteString(s.user.Session, cfmt.Sprintf(menuInvalidChoice, choice))
// 		return StatePromptCreateCharacter
// 	}

// }

// func (s *Screens) PromptChooseCharacterName() int {
// 	// TODO: Check that the name is not banned
// 	// TODO: Check that the name is not already taken
// promptInputCharacterName:
// 	name, err := s.PromptUserInput(cfmt.Sprintf("Enter a name for your character: "))
// 	if err != nil {
// 		io.WriteString(s.user.Session, cfmt.Sprintf(inputErrorMsg))
// 		return StateQuit
// 	}

// 	name = strings.TrimSpace(name)
// 	if name == "" {
// 		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
// 		return StatePromptCreateCharacter
// 	}

// promptConfirmCharacterName:
// 	// TODO: Check that the name is not banned
// 	// TODO: Check that the name is not already taken
// 	// TODO: Check that the name is not too long or too short
// 	// TODO: Check that the name does not contain invalid characters

// 	confirmName, err := s.PromptUserInput(cfmt.Sprintf("Confirm character name '%s' (y/n): ", name))
// 	if err != nil {
// 		io.WriteString(s.user.Session, cfmt.Sprintf(inputErrorMsg))
// 		return StateQuit
// 	}

// 	if confirmName == "" {
// 		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
// 		goto promptConfirmCharacterName
// 	}

// 	confirmName = strings.TrimSpace(confirmName)
// 	if !strings.EqualFold(confirmName, "y") {
// 		io.WriteString(s.user.Session,
// 			cfmt.Sprintf("Character name '%s' was not confirmed.", name))
// 		goto promptInputCharacterName
// 	}

// 	return StatePromptCreateCharacter
// }

// func (s *Screens) handlePregenCharacter() int {
// 	// TODO: Implement pre-generated character selection
// 	io.WriteString(s.user.Session, "Pre-generated character selection is not implemented yet.\n")
// 	return StatePromptMainMenu
// }

// func (s *Screens) handleCustomCharacter() int {
// 	// Prompt for character
// promptCharacterName:
// 	name, errName := s.PromptUserInput(cfmt.Sprint(createCharacterNamePrompt))
// 	// io.WriteString(s.user.Session, "Enter character name: ")
// 	// name, errName := s.user.Term.ReadLine()
// 	if errName != nil {
// 		logrus.WithError(errName).Error("Error reading character name")
// 		return StateQuit
// 	}
// 	name = strings.TrimSpace(name)

// 	conrimed, errConfirm := s.PromptConfirmInput(createCharacterConfirmNamePrompt)
// 	if errConfirm != nil {
// 		logrus.WithError(errConfirm).Error("Error reading confirming character name")
// 		return StateQuit
// 	}

// 	if !conrimed {
// 		goto promptCharacterName
// 	}

// 	// Prompt for character metatype
// 	io.WriteString(s.user.Session, "Enter character metatype: ")
// 	metatype, err := s.user.Term.ReadLine()
// 	if err != nil {
// 		logrus.WithError(err).Error("Error reading character metatype")
// 		return StateQuit
// 	}
// 	metatype = strings.TrimSpace(metatype)

// 	// Allocate priorities
// 	priorities := make(map[int]string)
// 	for i, category := range []string{"Metatype", "Attributes", "Magic", "Skills", "Resources"} {
// 		io.WriteString(s.user.Session, fmt.Sprintf("Allocate priority for %s (A, B, C, D, E): ", category))
// 		priority, err := s.user.Term.ReadLine()
// 		if err != nil {
// 			logrus.WithError(err).Error("Error reading priority")
// 			return StateQuit
// 		}
// 		priority = strings.TrimSpace(strings.ToUpper(priority))
// 		if !contains(priorityOptions, priority) {
// 			io.WriteString(s.user.Session, "Invalid priority. Please try again.\n")
// 			i--
// 			continue
// 		}
// 		priorities[i] = priority
// 	}

// 	// Display and confirm choices
// 	io.WriteString(s.user.Session, "Character creation summary:\n")
// 	io.WriteString(s.user.Session, fmt.Sprintf("Name: %s\n", name))
// 	io.WriteString(s.user.Session, fmt.Sprintf("Metatype: %s\n", metatype))
// 	for i, category := range []string{"Metatype", "Attributes", "Magic", "Skills", "Resources"} {
// 		io.WriteString(s.user.Session, fmt.Sprintf("%s: %s\n", category, priorities[i]))
// 	}
// 	io.WriteString(s.user.Session, "Confirm character creation? (yes/no): ")
// 	confirm, err := s.user.Term.ReadLine()
// 	if err != nil {
// 		logrus.WithError(err).Error("Error reading confirmation")
// 		return StateQuit
// 	}
// 	confirm = strings.TrimSpace(strings.ToLower(confirm))
// 	if confirm != "yes" {
// 		io.WriteString(s.user.Session, "Character creation cancelled.\n")
// 		return StatePromptMainMenu
// 	}

// 	// TODO: Save the character details
// 	io.WriteString(s.user.Session, "Character created successfully!\n")
// 	return StatePromptMainMenu
// }

// func (s *Screens) displayCharacterCreationInfo() int {
// 	// TODO: Implement character creation information display
// 	io.WriteString(s.user.Session, "Character creation information is not implemented yet.\n")
// 	return StatePromptMainMenu
// }

// // Helper function to check if a slice contains a string
// func contains(slice []string, item string) bool {
// 	for _, s := range slice {
// 		if s == item {
// 			return true
// 		}
// 	}
// 	return false
// }
