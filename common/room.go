package common

import (
	"io"
	"sync"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

const (
	RoomsDataPath = "data/rooms"
	RoomFilename  = RoomsDataPath + "/%s/%s.yaml"
)

type Rooms map[string]*Room

type RoomSpec struct {
	ID               string            `yaml:"id,omitempty"`
	ZoneID           string            `yaml:"zone_id,omitempty"`
	Name             string            `yaml:"name,omitempty"`
	ShortDescription string            `yaml:"short_description,omitempty"`
	Description      string            `yaml:"description,omitempty"`
	Exits            map[string]string `yaml:"exits,omitempty"`
}

func NewRoom(spec *RoomSpec) *Room {
	return &Room{
		ID:         spec.ID,
		Spec:       spec,
		Characters: make(Charcters),
	}
}

type Room struct {
	sync.Mutex
	ID              string                `yaml:"id,omitempty"`
	Spec            *RoomSpec             `yaml:"-"`
	Characters      map[string]*Character `yaml:"-"`
	CharactersCount int                   `yaml:"-"`
	Zone            *Zone                 `yaml:"-"`
}

func (r *Room) AddCharacter(c *Character) {
	logrus.WithFields(logrus.Fields{"room_id": r.ID, "character_id": c.ID, "character_name": c.Name}).Debug("Adding character to room")
	r.Lock()
	r.Characters[c.ID] = c
	r.CharactersCount++
	r.Unlock()

	for k, v := range r.Characters {
		logrus.WithFields(logrus.Fields{"id": k, "name": v.Name}).Debug("Character in room")
		if k == c.ID {
			continue
		}
		// TODO: add direction of entrance
		io.WriteString(v.Session, c.Name+" has entered the room.\n")
		logrus.WithFields(logrus.Fields{"character": c.Name, "room": r.ID}).Info("Character entered room")
	}
}

func (r *Room) RemoveCharacter(c *Character) {
	logrus.WithFields(logrus.Fields{"room_id": r.ID, "character_id": c.ID, "character_name": c.Name}).Debug("Removing character from room")
	// TODO: add direction of exit
	for k, v := range r.Characters {
		logrus.WithFields(logrus.Fields{"id": k, "name": v.Name}).Debug("Character in room")
		if k == c.ID {
			continue
		}
		io.WriteString(v.Session, c.Name+" has left the room.\n")
		logrus.WithFields(logrus.Fields{"character": c.Name, "room": r.ID}).Info("Character left room")
	}
	r.Lock()
	delete(r.Characters, c.ID)
	r.CharactersCount--
	r.Unlock()
}

func (r *Room) DisplayRoom(c *Character) {
	color.New(color.FgBlue, color.Underline).Fprintln(c.Session, r.Spec.Name)
	color.New(color.FgWhite).Fprintln(c.Session, r.Spec.ShortDescription)
	color.New(color.FgGreen).Fprintln(c.Session, "Exits:")
	if len(r.Spec.Exits) == 0 {
		color.New(color.FgGreen).Fprintln(c.Session, "\tNone")
	} else {
		for k, v := range r.Spec.Exits {
			color.New(color.FgGreen).Fprintf(c.Session, "\t%s - %s", k, v)
		}
	}
}

var CoreRooms = []RoomSpec{
	{
		ID:               "the_void",
		Name:             "The Void",
		ZoneID:           "the_void",
		ShortDescription: "You step out into ......",
		Description:      "You don't think that you are not floating in nothing.",
		Exits: map[string]string{
			"north": "limbo",
			"south": "the_chat_room",
		},
	},
	{
		ID:               "limbo",
		Name:             "Limbo",
		ZoneID:           "the_void",
		ShortDescription: "You are floating in a formless void, detached from all sensation of physical matter, surrounded by swirling glowing light, which fades into the relative darkness around you without any trace of edges or shadow.",
		Description:      "There is a \"No Tipping\" notice pinned to the darkness.",
		Exits: map[string]string{
			"south": "the_void",
		},
	},
	{
		ID:               "the_chat_room",
		Name:             "The Chat Room",
		ZoneID:           "the_void",
		ShortDescription: "You are lounging in a quiet cosy parlour, warmed by a gentle magical fire which twinkles happily in a warm fireplace.  There are no doors out.  Clearly the owner of this room needs none.",
		Description:      "You are lounging in a quiet cosy parlour, warmed by a gentle magical fire which twinkles happily in a warm fireplace.  There are no doors out.  Clearly the owner of this room needs none.",
		Exits: map[string]string{
			"north": "the_void",
		},
	},
}
