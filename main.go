package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"shadowrunmud/character/metatype"
	"shadowrunmud/character/quality"
	"shadowrunmud/character/skill"
	"shadowrunmud/util"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/google/uuid"
	"github.com/muesli/termenv"
	"github.com/sirupsen/logrus"
)

const (
	host = "localhost"
	port = "23234"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	var wg sync.WaitGroup
	wg.Add(1)
	go metatype.LoadMetatypes(&wg)
	wg.Add(1)
	go skill.LoadLanguageSkills(&wg)
	wg.Add(1)
	go skill.LoadActiveSkills(&wg)
	wg.Add(1)
	go skill.LoadKnowledgeSkills(&wg)
	wg.Add(1)
	go quality.LoadQualities(&wg)
	wg.Wait()

	// for k, m := range metatype.Metatypes {
	// 	fmt.Println(k, m)
	// }
	// for k, m := range item.WeaponsMelee {
	// 	fmt.Println(k, m)
	// }

	// app := newApp()
	// app.Start()

	// id := "3ea978c1-e567-4241-920c-fb863f7c83c4"

	// if _, err := os.Stat(fmt.Sprintf("data/user-%s.yaml", id)); errors.Is(err, os.ErrNotExist) {
	// 	logrus.WithError(err).Fatal("Could not find character")

	// 	id = uuid.New().String()

	// 	c := character.Character{
	// 		ID:        id,
	// 		Name:      "John Doe",
	// 		Metatype:  metatype.Metatypes["Elf"],
	// 		Body:      7,
	// 		Agility:   6,
	// 		Reaction:  5, // (7)
	// 		Strength:  5,
	// 		Willpower: 3,
	// 		Logic:     2,
	// 		Intuition: 3,
	// 		Charisma:  2,
	// 		Essence:   0.88,
	// 		Edge:      1,
	// 	}

	// reloadWeapons()
	// i := 1
	// for k, m := range item.WeaponsMelee {
	// 	m.ID = i
	// 	r := strings.NewReplacer(" ", "_", "-", "_")
	// 	id := r.Replace((strings.ToLower(k)))
	// 	if err := util.SaveStructToYAML(fmt.Sprintf("data/items/weapons/melee/%d-%s.yaml", i, id), &m); err != nil {
	// 		logrus.WithError(err).Fatal("Could not save file")
	// 	}
	// 	i++
	// }

	// var c2 character.Character
	// if err := loadStructFromYAML(fmt.Sprintf("data/user-%s.yaml", id), &c2); err != nil {
	// 	logrus.WithError(err).Fatal("Could not load character")
	// }

	// fmt.Printf("ID: %s\n", c2.ID)
	// fmt.Printf("Name: %s\n", c2.Name)
	// fmt.Printf("Body: %d\n", c2.Body)
	// fmt.Printf("Agility: %d\n", c2.Agility)
	// fmt.Printf("Reaction: %d\n", c2.Reaction)
	// fmt.Printf("Strength: %d\n", c2.Strength)
	// fmt.Printf("Willpower: %d\n", c2.Willpower)
	// fmt.Printf("Logic: %d\n", c2.Logic)
	// fmt.Printf("Intuition: %d\n", c2.Intuition)
	// fmt.Printf("Charisma: %d\n", c2.Charisma)
	// fmt.Printf("Essence: %f\n", c2.Essence)
	// fmt.Printf("Edge: %d\n", c2.Edge)

	// for k, m := range metatype.Metatypes {
	// fmt.Println(k, m)
	// if err := saveStructToYAML(fmt.Sprintf("data/metatypes/%s.yaml", k), &m); err != nil {
	// 	logrus.WithError(err).Fatal("Could not save character")
	// }
	// }

}

func reloadData() {
	for k, m := range skill.LanguageSkills {
		r := strings.NewReplacer(" ", "_", "-", "_")
		id := r.Replace((strings.ToLower(k)))
		if err := util.SaveStructToYAML(fmt.Sprintf("data/skills/languages/%v.yaml", id), &m); err != nil {
			logrus.WithError(err).Fatal("Could not save file")
		}
	}
}

// app contains a wish server and the list of running programs.
type app struct {
	*ssh.Server
	progs sync.Map
}

// send dispatches a message to all running programs.
func (a *app) send(msg tea.Msg) {
	a.progs.Range(func(key, value interface{}) bool {
		go value.(*tea.Program).Send(msg)

		return true
	})
}

func (a *app) sendToId(id string, msg tea.Msg) {
	prog, ok := a.progs.Load(id)
	if !ok {
		return
	}

	go prog.(*tea.Program).Send(msg)
}

func newApp() *app {
	a := new(app)

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.MiddlewareWithProgramHandler(a.ProgramHandler, termenv.ANSI256),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	a.Server = s

	return a
}

func (a *app) Start() {
	var err error
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = a.ListenAndServe(); err != nil {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := a.Shutdown(ctx); err != nil {
		log.Error("Could not stop server", "error", err)
	}
}

func (a *app) ProgramHandler(s ssh.Session) *tea.Program {
	model := initialChatModel(s)
	model.app = a
	model.id = uuid.New().String()
	log.Info("New program", "id", model.id)

	p := tea.NewProgram(model, bubbletea.MakeOptions(s)...)
	a.progs.Store(model.id, p)

	a.sendToId(model.id, chatMsg{id: model.id, text: "Welcome to the chat room!"})

	return p
}

type (
	errMsg  error
	chatMsg struct {
		id   string
		text string
	}
)
