package core

// const (
// 	ConfigFilepath = "_data/config/server.yaml"
// )

// type World struct {
// 	config     config.Server
// 	characters common.Charcters
// 	zones      common.Zones
// 	rooms      common.Rooms
// }

// func NewWorld() *World {
// 	var serverConfig config.Server
// 	utils.LoadStructFromYAML(ConfigFilepath, &serverConfig)

// 	logrus.WithField("serverConfig", serverConfig).Info("Loaded server configuration")

// 	logrusLevel, err := logrus.ParseLevel(serverConfig.LogLevel)
// 	if err != nil {
// 		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
// 		logrusLevel = logrus.InfoLevel
// 	}
// 	logrus.SetLevel(logrusLevel)

// 	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

// 	w := &World{
// 		config:     serverConfig,
// 		characters: make(common.Charcters),
// 		zones:      make(common.Zones),
// 		rooms:      make(common.Rooms),
// 	}

// 	return w
// }

// func (w *World) Start() {
// 	logrus.WithFields(logrus.Fields{"host": w.config.Host, "port": w.config.Port}).Info("Starting server")
// 	ssh.Handle(w.Handler)

// 	server := &ssh.Server{
// 		Handler:                  w.Handler,
// 		ConnectionFailedCallback: w.ConnectionFailedCallback,
// 		Addr:                     net.JoinHostPort(w.config.Host, w.config.Port),
// 		IdleTimeout:              w.config.Timeouts.Idle,
// 	}

// 	// Load data
// 	w.LoadData()

// 	if err := server.ListenAndServe(); err != nil {
// 		logrus.WithError(err).Error("Could not start server")
// 	}
// }

// func (w *World) ConnectionFailedCallback(conn net.Conn, err error) {
// 	defer conn.Close()
// 	logrus.WithError(err).Error("Connection failed")
// }

// func (w *World) HandleTimeout(s ssh.Session) {
// 	logrus.Debug("Timeout handler started")
// 	i := 0
// 	for {
// 		i += 1
// 		select {
// 		case <-time.After(time.Second):
// 			continue
// 		case <-s.Context().Done():
// 			logrus.Info("Connection closed")
// 			return
// 		}
// 	}
// }

// type ConnectionState int

// const (
// 	StateBanner ConnectionState = iota
// 	StatePromptUsername
// 	StateConfirmUsername
// 	StatePromptPassword
// 	StateAuthenticate
// 	StateMOTD
// 	StateMenu
// 	StateGameLoop
// 	StateQuit
// )

// func (w *World) Handler(s ssh.Session) {
// 	logrus.WithFields(logrus.Fields{"user": s.User(), "remote_addr": s.RemoteAddr()}).Info("New connection")

// 	// BANNER
// 	// banner:
// 	io.WriteString(s, cfmt.Sprint(`
//     {{     ::::::::  :::    :::     :::     :::::::::   ::::::::  :::       ::: ::::    ::::  :::    ::: :::::::::  }}::#ff8700
//     {{    :+:    :+: :+:    :+:   :+: :+:   :+:    :+: :+:    :+: :+:       :+: +:+:+: :+:+:+ :+:    :+: :+:    :+: }}::#ff5f00
//     {{    +:+        +:+    +:+  +:+   +:+  +:+    +:+ +:+    +:+ +:+       +:+ +:+ +:+:+ +:+ +:+    +:+ +:+    +:+ }}::#ff0000
//     {{    +#++:++#++ +#++:++#++ +#++:++#++: +#+    +:+ +#+    +:+ +#+  +:+  +#+ +#+  +:+  +#+ +#+    +:+ +#+    +:+ }}::#d70000
//     {{           +#+ +#+    +#+ +#+     +#+ +#+    +#+ +#+    +#+ +#+ +#+#+ +#+ +#+       +#+ +#+    +#+ +#+    +#+ }}::#af0000
//     {{    #+#    #+# #+#    #+# #+#     #+# #+#    #+# #+#    #+#  #+#+# #+#+#  #+#       #+# #+#    #+# #+#    #+# }}::#870000
//     {{     ########  ###    ### ###     ### #########   ########    ###   ###   ###       ###  ########  #########  }}::#5f0000

//     {{Enter your username to continue}}::#c0c0c0
//     `))

// 	t := term.NewTerminal(s, "")

// userlogin:
// 	io.WriteString(s, cfmt.Sprint("{{Username: }}::#ffffff|bold"))
// 	username, errReadLine := t.ReadLine()
// 	if errReadLine != nil {
// 		logrus.WithError(errReadLine).Error("Error reading username")
// 		return
// 	}

// 	username = strings.TrimSpace(username)
// 	logrus.WithField("username", username).Info("Received username")

// 	if new := common.IsNewUsername(username); new {
// 		logrus.WithField("new", new).Info("Username is new")
// 		goto usercreation
// 	}

// 	cfmt.Sprint("{{Password: }}::#ffffff|bold")
// 	passwordBytes, err := t.ReadPassword(cfmt.Sprint("{{Password: }}::#ffffff|bold"))
// 	if err != nil {
// 		logrus.WithError(err).Error("Error reading password")
// 		return
// 	}

// 	password := strings.TrimSpace(string(passwordBytes))
// 	logrus.WithField("password", password).Info("Received password")

// 	// Auth user
// 	cfmt.Sprintf("{{Welcome back %s!}}::#ffffff|bold", username)

// 	// USER CREATION
// usercreation:
// 	cfmt.Sprintf("{{Welcome %s! is this the username you wish to use? (y/n)}}::#ffffff|bold", username)
// 	goodUsername, errReadLine := t.ReadLine()
// 	if errReadLine != nil {
// 		logrus.WithError(errReadLine).Error("Error reading username")
// 		return
// 	}

// 	if !strings.EqualFold(goodUsername, "y") {
// 		goto userlogin
// 	} else {

// 	}
// }

// func (w *World) DisplayBanner(s ssh.Session) {

// 	io.WriteString(s, cfmt.Sprint(`
// {{     ::::::::  :::    :::     :::     :::::::::   ::::::::  :::       ::: ::::    ::::  :::    ::: :::::::::  }}::#ff8700
// {{    :+:    :+: :+:    :+:   :+: :+:   :+:    :+: :+:    :+: :+:       :+: +:+:+: :+:+:+ :+:    :+: :+:    :+: }}::#ff5f00
// {{    +:+        +:+    +:+  +:+   +:+  +:+    +:+ +:+    +:+ +:+       +:+ +:+ +:+:+ +:+ +:+    +:+ +:+    +:+ }}::#ff0000
// {{    +#++:++#++ +#++:++#++ +#++:++#++: +#+    +:+ +#+    +:+ +#+  +:+  +#+ +#+  +:+  +#+ +#+    +:+ +#+    +:+ }}::#d70000
// {{           +#+ +#+    +#+ +#+     +#+ +#+    +#+ +#+    +#+ +#+ +#+#+ +#+ +#+       +#+ +#+    +#+ +#+    +#+ }}::#af0000
// {{    #+#    #+# #+#    #+# #+#     #+# #+#    #+# #+#    #+#  #+#+# #+#+#  #+#       #+# #+#    #+# #+#    #+# }}::#870000
// {{     ########  ###    ### ###     ### #########   ########    ###   ###   ###       ###  ########  #########  }}::#5f0000

// {{Enter your username to continue}}::#c0c0c0
// `))
// }

// func (w *World) PromptUsername(s ssh.Session) (string, error) {
// 	t := term.NewTerminal(s, "")

// 	io.WriteString(s, cfmt.Sprint("{{Username: }}::#ffffff|bold"))
// 	username, errReadLine := t.ReadLine()
// 	if errReadLine != nil {
// 		logrus.WithError(errReadLine).Error("Error reading username")
// 		return "", errReadLine
// 	}

// 	username = strings.TrimSpace(username)
// 	logrus.WithField("username", username).Info("Received username")

// 	return username, nil
// }

// func (w *World) ConfirmUsername(s ssh.Session, username string) (string, error) {
// 	t := term.NewTerminal(s, "")

// 	io.WriteString(s, cfmt.Sprintf("{{Welcome %s, is this correct (y/n): }}::#ffffff|bold", username))
// 	usernameCorrect, errReadLine := t.ReadLine()
// 	if errReadLine != nil {
// 		logrus.WithError(errReadLine).Error("Error reading username")
// 		return "", errReadLine
// 	}

// 	usernameCorrect = strings.TrimSpace(usernameCorrect)
// 	if strings.EqualFold(usernameCorrect, "y") {
// 		logrus.WithField("username", username).Info("Username confirmed")
// 		return username, nil
// 	} else {
// 		logrus.WithField("username", username).Info("Username not confirmed")
// 		return "", nil
// 	}

// 	return username, nil
// }

// func PromptPassword(s ssh.Session) (string, error) {
// 	t := term.NewTerminal(s, "")

// 	// Collect password without echoing
// 	cfmt.Sprint("{{Password: }}::#ffffff|bold")
// 	passwordBytes, err := t.ReadPassword(cfmt.Sprint("{{Password: }}::#ffffff|bold"))
// 	if err != nil {
// 		logrus.WithError(err).Error("Error reading password")
// 		return "", err
// 	}

// 	password := strings.TrimSpace(string(passwordBytes))
// 	logrus.WithField("password", password).Info("Received password")

// 	return password, nil
// }

// func (w *World) NewHandler(s ssh.Session) {
// 	logrus.WithFields(logrus.Fields{"user": s.User(), "remote_addr": s.RemoteAddr()}).Info("New connection")
// 	u := common.NewUser(s)

// 	conState := StateBanner

// 	go w.HandleTimeout(s)

// 	for {
// 		switch conState {
// 		case StateBanner:
// 			w.DisplayBanner(s)
// 			conState = StateAuthenticate
// 		case StatePromptUsername:
// 			username, _ := w.PromptUsername(s)
// 			if isNew := common.IsNewUsername(username); isNew {
// 				conState = StateConfirmUsername
// 			}
// 			conState = StatePromptUsername

// 			if strings.EqualFold(username, "new") {
// 				conState = StateAuthenticate
// 			}
// 		case StateAuthenticate:
// 			if err := common.AuthenticateUser(s); err != nil {
// 				logrus.WithError(err).Error("Error authenticating user")
// 			}
// 			conState = StateGameLoop
// 		case StateGameLoop:
// 			if err := u.GameLoop(); err != nil {
// 				logrus.WithError(err).Error("Error in game loop")
// 			}
// 		}
// 		// output := termenv.NewOutput(s)
// 		// style := output.String("Welcome to ShadowMUD").Foreground(output.Color("#FFFFFF"))
// 		// io.WriteString(s, style.Styled())

// 		// output.WriteString("Hello World")

// 		// styleWhite := output.Foreground(output.Color("#ffffff")).Background(output.Color("#0000ff"))

// 		// buf := output.String("Hello World").Foreground(output.Color("#abcdef")).Background(output.Color("#0000ff"))
// 		// io.WriteString(s, buf.String())

// 		// // Supports hex values
// 		// // Will automatically degrade colors on terminals not supporting RGB
// 		// s.Foreground(output.Color("#abcdef"))
// 		// // but also supports ANSI colors (0-255)
// 		// s.Background(output.Color("69"))
// 		// // ...or the color.Color interface
// 		// s.Foreground(output.FromColor(color.RGBA{255, 128, 0, 255}))

// 		// // Combine fore- & background colors
// 		// s.Foreground(output.Color("#ffffff")).Background(output.Color("#0000ff"))

// 		// // Supports the fmt.Stringer interface
// 		// fmt.Println(s)

// 		// io.WriteString(s, "Username: ")
// 		// output := termenv.NewOutput(s)
// 		// st := output.String("Hello World")
// 		// st.Foreground(output.Color("#ffffff")).Background(output.Color("#0000ff"))
// 		// output.WriteString(st)
// 		// output.RestoreScreen()
// 		// output.SaveScreen()
// 		// output.Reset()

// 		// ColorProfile := output.ColorProfile().String()
// 		// FGcolor := output.ForegroundColor()
// 		// FGcolor.
// 		// BGcolor := output.BackgroundColor()
// 		// IsDarkBackground := output.HasDarkBackground()

// 		// c := common.NewCharacter(s)
// 		// Commented out for now
// 		// if ok := c.Authenticate(); !ok {
// 		// color.New(color.FgHiRed).Fprintln(s, "Username or Password is incorrect")
// 		// return
// 		// }

// 		// w.characters[c.ID] = c
// 		// defer func() {
// 		// 	delete(w.characters, c.ID)
// 		// }()

// 		// c.Room = w.GetRoom("the_void:limbo")
// 		// // c.Room = w.rooms["the_void:limbo"]
// 		// c.Room.AddCharacter(c)

// 		// if err := c.GameLoop(); err != nil {
// 		// 	logrus.WithError(err).Error("Error in game loop")
// 		// }
// 	}
// }

// func (w *World) LoadData() {
// 	w.LoadZones()
// 	w.LoadRooms()
// 	w.BuildRooms()
// }

// func (w *World) GetZone(id string) *common.Zone {
// 	logrus.WithFields(logrus.Fields{"zone_id": id}).Debug("Getting zone")

// 	if zone, ok := w.zones[id]; ok {
// 		logrus.WithFields(logrus.Fields{"zone_id": id}).Debug("Zone found")
// 		return zone
// 	}

// 	logrus.WithFields(logrus.Fields{"zone_id": id}).Error("Zone not found")
// 	return nil
// }

// func (w *World) LoadZones() {
// 	logrus.Info("Loading zones")
// 	rootpath := w.config.Data.BaseDir + w.config.Data.ZonesDir

// 	if err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
// 		if info.IsDir() {
// 			return nil
// 		}

// 		if filepath.Ext(path) == ".yaml" {
// 			var spec common.ZoneSpec
// 			if err := utils.LoadStructFromYAML(path, &spec); err != nil {
// 				logrus.WithError(err).WithFields(logrus.Fields{"path": path}).Error("Failed to load zone")
// 				return err
// 			}

// 			// Load the zone
// 			v := common.NewZone(&spec)
// 			w.zones[v.ID] = v
// 			logrus.WithFields(logrus.Fields{"path": path, "id": v.ID}).Debug("Found zone")
// 		}

// 		return nil
// 	}); err != nil {
// 		logrus.WithError(err).Error("file walk error")
// 	}

// 	logrus.Infof("Finished loading %d zones", len(w.zones))
// }

// func (w *World) GetRoom(id string) *common.Room {
// 	logrus.WithFields(logrus.Fields{"room_id": id}).Debug("Getting room")

// 	if room, ok := w.rooms[id]; ok {
// 		logrus.WithFields(logrus.Fields{"room_id": id}).Debug("Room found")
// 		return room
// 	}

// 	logrus.WithFields(logrus.Fields{"room_id": id}).Error("Room not found, loading default room")
// 	return w.GetDefaultRoom()
// }

// func (w *World) GetDefaultRoom() *common.Room {
// 	r := common.NewRoom(&common.RoomSpec{
// 		ID:               "limbo",
// 		ZoneID:           "the_void",
// 		Name:             "Limbo",
// 		ShortDescription: "You are in a void of nothingness.",
// 		Description:      "You are in a void of nothingness. There is nothing here.",
// 		Exits:            common.Exits{},
// 	})
// 	r.ID = "limbo"

// 	r.Zone = common.NewZone(&common.ZoneSpec{
// 		ID: "the_void",
// 	})
// 	r.Zone.Rooms[r.ID] = r

// 	return r
// }
// func (w *World) LoadRooms() {
// 	logrus.Info("Loading rooms")

// 	rootpath := w.config.Data.BaseDir + w.config.Data.RoomsDir

// 	if err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
// 		if info.IsDir() {
// 			return nil
// 		}

// 		if filepath.Ext(path) == ".yaml" {
// 			var spec common.RoomSpec
// 			if err := utils.LoadStructFromYAML(path, &spec); err != nil {
// 				logrus.WithError(err).WithFields(logrus.Fields{"path": path}).Error("Failed to load room")
// 				return err
// 			}

// 			// Load our rooms
// 			v := common.NewRoom(&spec)
// 			logrus.WithFields(logrus.Fields{"zone_id": v.Spec.ZoneID, "path": path, "id": v.ID}).Debug("Found room")
// 			w.rooms[fmt.Sprintf("%s:%s", v.Spec.ZoneID, v.ID)] = v

// 		}

// 		return nil
// 	}); err != nil {
// 		logrus.WithError(err).Error("file walk error")
// 	}

// 	logrus.Infof("Finished loading %d rooms", len(w.rooms))
// }

// func (w *World) BuildRooms() {
// 	logrus.Info("Building rooms")
// 	for _, r := range w.rooms {
// 		logrus.WithFields(logrus.Fields{"zone_id": r.Spec.ZoneID, "room_id": r.ID}).Debug("Building room")
// 		// Add this room to the zone's room list
// 		w.zones[r.Spec.ZoneID].Rooms[r.ID] = r
// 		// Set the room's zone
// 		r.Zone = w.zones[r.Spec.ZoneID]
// 	}
// }
