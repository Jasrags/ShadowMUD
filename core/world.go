package core

import (
	"io"
	"net"
	"sync"
	"time"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/config"
	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
)

const (
	ConfigFilepath = "_data/config/server.yaml"
)

type World struct {
	config     config.Server
	characters sync.Map
}

func NewWorld() *World {
	var serverConfig config.Server
	utils.LoadStructFromYAML(ConfigFilepath, &serverConfig)

	logrus.WithField("serverConfig", serverConfig).Info("Loaded server configuration")

	logrusLevel, err := logrus.ParseLevel(serverConfig.LogLevel)
	if err != nil {
		logrus.WithError(err).Warn("Could not parse log level, defaulting to INFO")
		logrusLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logrusLevel)

	logrus.WithField("log_level", logrusLevel).Info("Logger level set")

	w := &World{
		config: serverConfig,
	}

	return w
}

func (w *World) Start() {
	logrus.WithFields(logrus.Fields{"host": w.config.Host, "port": w.config.Port}).Info("Starting server")
	ssh.Handle(w.Handler)

	server := &ssh.Server{
		BannerHandler: w.BannerHandler,
		// PasswordHandler: w.PasswordHandler,
		Handler:                  w.Handler,
		ConnectionFailedCallback: w.ConnectionFailedCallback,
		Addr:                     net.JoinHostPort(w.config.Host, w.config.Port),
		IdleTimeout:              time.Second * 10,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Error("Could not start server")
	}

	// srv, err := wish.NewServer(
	// 	wish.WithAddress(net.JoinHostPort(w.config.Host, w.config.Port)),
	// 	wish.WithHostKeyPath(".ssh/id_ed25519"),
	// 	ssh.AllocatePty(),

	// 	wish.WithMiddleware(
	// 		func(next ssh.Handler) ssh.Handler {
	// 			return func(s ssh.Session) {

	// 				pty, _, _ := s.Pty()
	// 				renderer := bubbletea.MakeRenderer(s)
	// 				// textStyle := renderer.NewStyle().Foreground(lipgloss.Color("10"))

	// 				color.New(color.FgBlue).Fprintln(s, "blue color!")

	// 				bg := "light"
	// 				if renderer.HasDarkBackground() {
	// 					bg = "dark"
	// 				}

	// wish.Printf(s, lipgloss.JoinVertical(
	// 	lipgloss.Top,
	// 	textStyle.Render("Hello, world!\r\n"),
	// 	fmt.Sprintf("Term: %s\r\n", pty.Term),
	// 	fmt.Sprintf("PTY: %s\r\n", pty.Slave.Name()),
	// 	fmt.Sprintf("FD: %d\r\n", pty.Slave.Fd()),
	// 	fmt.Sprintf("Background: %v\r\n", bg),
	// ))

	// wish.Printf(s, color.Green("Hello, world!\r\n").String())
	// wish.Printf(s, textStyle.Render("Hello, world!\r\n"))
	// wish.Printf(s, "Hello, world!\r\n")
	// wish.Printf(s, "Term: %s\r\n", pty.Term)
	// wish.Printf(s, "PTY: %s\r\n", pty.Slave.Name())
	// wish.Printf(s, "FD: %d\r\n", pty.Slave.Fd())
	// wish.Printf(s, "Background: %v\r\n", bg)

	// wish.Printf(s, textStyle.Render("Hello, world!\r\n"))

	// next(s)
	// }
	// 		},

	// 		activeterm.Middleware(),
	// 		logging.Middleware(),
	// 	),
	// )
	// if err != nil {
	// 	log.Error("Could not start server", "error", err)
	// }

	// done := make(chan os.Signal, 1)
	// signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	log.Info("Starting SSH server", "host", w.config.Host, "port", w.config.Port)
	// 	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
	// 		log.Error("Could not start server", "error", err)
	// 		done <- nil
	// 	}
	// }()

	// <-done

	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer func() { cancel() }()
	// log.Info("Stopping SSH server")
	// if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
	// 	log.Error("Could not stop server", "error", err)
	// }

	// listener, errListen := net.Listen("tcp", net.JoinHostPort(w.config.Host, w.config.Port))
	// if errListen != nil {
	// 	logrus.WithError(errListen).Panic("Could not start server")
	// }
	// defer listener.Close()

	// for {
	// 	conn, errAccept := listener.Accept()
	// 	if errAccept != nil {
	// 		logrus.WithError(errAccept).Error("Error accepting connection")
	// 	}

	// 	// Handle the connection in a new goroutine
	// 	go w.handleConnection(conn)
	// }
}

func (w *World) ConnectionFailedCallback(conn net.Conn, err error) {
	defer conn.Close()
	logrus.WithError(err).Error("Connection failed")
}

func (w *World) BannerHandler(ctx ssh.Context) string {
	logrus.Debug("Sending banner")
	return "Welcome to my SSH server, friend!\r\n"
}

func (w *World) HandleTimeout(s ssh.Session) {
	logrus.Debug("Timeout handler started")
	i := 0
	for {
		i += 1
		select {
		case <-time.After(time.Second):
			continue
		case <-s.Context().Done():
			logrus.Info("Connection closed: timeout")
			return
		}
	}
}

func (w *World) Handler(s ssh.Session) {
	logrus.WithFields(logrus.Fields{"user": s.User(), "remote_addr": s.RemoteAddr()}).Info("New connection")

	go w.HandleTimeout(s)

	c := common.NewCharacter(s)
	if ok := c.Authenticate(); !ok {
		io.WriteString(s, "Authentication failed\r\n")
		return
	}

	c.Load()

	w.characters.Store(c.ID, c)
	defer func() {
		w.characters.Delete(c.ID)
	}()

	if err := c.GameLoop(); err != nil {
		logrus.WithError(err).Error("Error in game loop")
	}

	// t := term.NewTerminal(s, "")

	// io.WriteString(s, "Hello world\r\n")
	// io.WriteString(s, fmt.Sprintf("Term: %s\r\n", pty.Term))
	// io.WriteString(s, fmt.Sprintf("Height: %d\r\n", pty.Window.Height))
	// io.WriteString(s, fmt.Sprintf("Width: %d\r\n", pty.Window.Width))

	// io.WriteString(s, "Username: ")
	// username, errReadLine := t.ReadLine()
	// if errReadLine != nil {
	// 	logrus.WithError(errReadLine).Error("Error reading username")
	// 	return
	// }
	// username = strings.TrimSpace(username)
	// logrus.WithField("username", username).Info("Received username")

	// passwordBytes, err := t.ReadPassword("Password: ")
	// if err != nil {
	// 	log.Println("Error reading password:", err)
	// 	return
	// }
	// password := strings.TrimSpace(string(passwordBytes))
	// logrus.WithField("password", password).Info("Received password")

	// // Validate credentials
	// if pass, ok := users[username]; ok && strings.EqualFold(pass, password) {
	// 	logrus.Info("Authentication successful")
	// }

	// for {
	// 	io.WriteString(s, ">")
	// 	line, err := t.ReadLine()
	// 	if err != nil {
	// 		break
	// 	}
	// 	logrus.WithField("line", line).Info("Received line")
	// }
}

// func (w *World) authenticate(reader *bufio.Reader, writer *bufio.Writer) (bool, string) {
// 	logrus.Debug("Authenticating user")

// 	// Prompt for username
// 	writer.WriteString("Username: ")
// 	writer.Flush()
// 	username, errReadUsername := reader.ReadString('\n')
// 	if errReadUsername != nil {
// 		logrus.WithError(errReadUsername).Error("Error reading username")
// 		return false, ""
// 	}
// 	username = strings.TrimSpace(username)
// 	logrus.WithField("username", username).Debug("Received username")

// 	w.DisableEcho(writer)

// 	// Prompt for password
// 	writer.WriteString("Password: ")
// 	writer.Flush()
// 	password, errReadPassword := reader.ReadString('\n')
// 	if errReadPassword != nil {
// 		logrus.WithError(errReadPassword).Error("Error reading password")
// 		return false, ""
// 	}
// 	password = strings.TrimPrefix(strings.TrimSpace(password), "\xff\xfd\x01")
// 	logrus.WithField("password", password).Debug("Received password")

// 	w.EnableEcho(writer)

// 	// Validate credentials
// 	if pass, ok := users[username]; ok && strings.EqualFold(pass, password) {
// 		return true, username
// 	}

// 	return false, ""
// }
