package screen

// type Screens struct {
// 	log     *logrus.Entry
// 	cfg     *config.Server
// 	session ssh.Session
// 	pty     ssh.Pty
// 	window  <-chan ssh.Window
// 	term    *term.Terminal
// }

// func New(s ssh.Session, cfg *config.Server) *Screens {
// 	pty, ptyWindow, _ := s.Pty()

// 	screens := &Screens{
// 		cfg:     cfg,
// 		session: s,
// 		pty:     pty,
// 		window:  ptyWindow,
// 		term:    term.NewTerminal(s, ""),
// 	}
// 	screens.log = logrus.WithFields(logrus.Fields{"package": "screen"})
// 	return screens
// }
