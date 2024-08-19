package screen

var (
	// Prompt/Message strings
	// Character Creation
	// Game loop
	inputEchoMsg        = "{{You typed:}}::#ffffff|bold %s\n"
	gameLoopPrompt      = "{{> }}::#ffffff|bold"
	characterListOption = "{{%d.}}::#00ff00 %s\n"
	// Login
	loginClosedMsg     = "{{Login is currently closed.}}::#ff8700\n"
	usernamePrompt     = "{{Username: }}::#ffffff|bold"
	passwordPrompt     = "{{Password: }}::#ffffff|bold"
	invalidLoginMsg    = "{{You have entered an invalid username or password.}}::#ff8700\n"
	loginSuccessfulMsg = "{{Login successful.}}::#00ff00\n"
	loginUserBannedMsg = "{{You are banned until %s.}}::#ff8700\n"

	inputErrorMsg    = "{{An error occurred while reading your response.}}::#ff8700\n"
	requiredInputMsg = "{{You must enter a value.}}::#ff8700\n"

	// Registration
	registrationClosedMsg   = "{{Registration is currently closed.}}::#ff8700\n"
	usernameBannedMsg       = "{{Username '%s' is not allowed.}}::#ff8700\n"
	usernameNewPrompt       = "{{Enter your desired username: }}::#ffffff|bold"
	usernameConfirmPrompt   = "{{Confirm username '%s'}}::#ffffff|bold {{(y/n)}}::#00ff00|bold{{:}}::#ffffff|bold"
	usernameMixMaxLengthMsg = "{{Username must be between %d and %d characters.}}::#ff8700\n"
	// usernameDeclinedMsg     = "{{Username '%s' was not confirmed.}}::#ff8700\n"
	passwordNewPrompt       = "{{Enter new password: }}::#ffffff|bold"
	passwordConfirmPrompt   = "{{Confirm password: }}::#ffffff|bold"
	passwordMismatchMsg     = "{{Passwords do not match.}}::#ff8700\n"
	passwordMinMaxLengthMsg = "{{Password must be between %d and %d characters.}}::#ff8700\n"
	userCreatedMsg          = "{{User '%s' has been created.}}::#00ff00\n"

	// Menu
	menuPrompt                = "{{Enter the number of the option you would like to select: }}::#ffffff|bold"
	menuInvalidChoice         = "Invalid choice: %s\n"
	mainMenuTitle             = "\n{{Main Menu}}::#00ff00|bold\n"
	menuOptionEnterGame       = "{{1.}}::#00ff00 Enter game\n"
	menuOptionCreateCharacter = "{{2.}}::#00ff00 Create character (%d/%d)\n"
	menuOptionListCharacters  = "{{3.}}::#00ff00 List characters\n"
	menuOptionDeleteCharacter = "{{4.}}::#00ff00 Delete character\n"
	menuOptionChangePassword  = "{{5.}}::#00ff00 Change password\n"
	menuOptionQuit            = "{{0.}}::#00ff00 Quit\n"

	// Character Creation
	noCharactersCreatedMsg          = "\n{{You have no characters created, let's make one now!}}::#ff8700\n"
	characterNoneCreatedMsg         = "{{You have no characters created.}}::#ff8700\n"
	characterMaxCharactersMsg       = "{{You have reached the maximum number of characters allowed.}}::#ff8700\n"
	characterNameMixMaxLengthMsg    = "{{Character name must be between %d and %d characters.}}::#ff8700\n"
	createCharacterMenuTitle        = "\n{{Character Creation}}::#00ff00|bold\n"
	createCharacterMenuOptionPregen = "{{1.}}::#00ff00 Choose an pre-generated character\n"
	createCharacterMenuOptionCustom = "{{2.}}::#00ff00 Create an custom character\n"
	createCharacterMenuOptionLearn  = "{{3.}}::#00ff00 Learn more about Shadowrun characters\n"
	createCharacterMenuOptionReturn = "{{4.}}::#00ff00 Return to the main menu\n"

	passwordChangedMsg       = "{{Password has been changed.}}::#0000ff\n"
	featureNotImplementedMsg = "{{Feature not implemented}}::#ff0000\n"
)
