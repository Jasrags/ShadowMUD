package screen

// func (s *Screens) PromptLoginUser() int {
// promptUsername:
// 	// Collect new username
// 	username, errUsername := s.PromptUserInput(usernamePrompt)
// 	if errUsername != nil {
// 		logrus.WithError(errUsername).Error("Error reading username")
// 		io.WriteString(s.user.Session, cfmt.Sprintf(inputErrorMsg))
// 		return StateQuit
// 	}

// 	if username == "" {
// 		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
// 		goto promptUsername
// 	}

// 	username = strings.TrimSpace(username)
// 	logrus.WithField("username", username).Debug("Received username")

// 	// Do we need to make a new user?
// 	if strings.EqualFold(username, "new") {
// 		return StatePromptRegisterUser
// 	}

// promptPassword:
// 	// Collect new password
// 	password, errPassword := s.PromptUserPasswordInput(passwordPrompt)
// 	if errPassword != nil {
// 		logrus.WithError(errPassword).Error("Error reading password")
// 		return StateQuit
// 	}

// 	if password == "" {
// 		io.WriteString(s.user.Session, cfmt.Sprintf(requiredInputMsg))
// 		goto promptPassword
// 	}

// 	logrus.WithFields(logrus.Fields{"username": username, "password": password}).Debug("Received password")

// 	// is login enabled?
// 	if !s.cfg.LoginEnabled {
// 		io.WriteString(s.user.Session, cfmt.Sprint(loginClosedMsg))
// 		goto promptUsername
// 	}

// 	// Try the user from the file
// 	// if err := common.LoadUser(username, s.user); err != nil {
// 	// 	logrus.WithError(err).Error("Error loading user")
// 	// 	return StatePromptLoginUser
// 	// }

// 	logrus.WithFields(logrus.Fields{"username": s.user.Username, "id": s.user.ID}).Debug("Loaded user")

// 	// Is the user banned?
// 	for _, ban := range s.user.Bans {
// 		if ban.ExpiresAt.After(time.Now()) {
// 			io.WriteString(s.user.Session, cfmt.Sprintf(loginUserBannedMsg, ban.ExpiresAt.Format(time.RFC1123)))

// 			return StatePromptLoginUser
// 		}
// 	}

// 	// Validate the password against the hash
// 	if err := bcrypt.CompareHashAndPassword([]byte(s.user.Password), []byte(password)); err != nil {
// 		logrus.WithError(err).Error("Password validation error")

// 		switch err {
// 		case bcrypt.ErrMismatchedHashAndPassword:
// 			io.WriteString(s.user.Session, cfmt.Sprintf(invalidLoginMsg))
// 			logrus.WithFields(logrus.Fields{"username": username}).WithError(err).Error("Invalid password for user")
// 		case bcrypt.ErrHashTooShort:
// 		case bcrypt.ErrMismatchedHashAndPassword:
// 		case bcrypt.ErrPasswordTooLong:
// 		default:
// 			logrus.WithFields(logrus.Fields{"username": username}).WithError(err).Error("Password error")
// 		}

// 		return StatePromptLoginUser
// 	}

// 	// Add a login record
// 	s.user.Logins = append(s.user.Logins, user.Login{
// 		Time: time.Now(),
// 		IP:   s.user.Session.RemoteAddr().String(),
// 	})

// 	// Save the user
// 	if err := s.user.Save(); err != nil {
// 		logrus.WithError(err).Error("Error saving user")
// 		return StatePromptLoginUser
// 	}

// 	io.WriteString(s.user.Session, cfmt.Sprintf(loginSuccessfulMsg))

// 	return StatePromptMainMenu
// }
