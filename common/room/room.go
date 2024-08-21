package room

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Jasrags/ShadowMUD/common/character"

	"github.com/sirupsen/logrus"
)

const (
	RoomsFilepath = "_data/rooms"
)

type (
	Exits map[string]Exit
	Exit  struct {
		Direction string `yaml:"direction"`
		RoomID    string `yaml:"room_id"`
		Hidden    bool   `yaml:"hidden"`
		// Lock      bool   `yaml:"lock"
	}
	Specs map[string]*Spec
	Spec  struct {
		ID               string `yaml:"id"`
		ZoneID           string `yaml:"zone_id"`
		Name             string `yaml:"name"`
		ShortDescription string `yaml:"short_description"`
		Description      string `yaml:"description"`
		Exits            Exits  `yaml:"exits"`
	}
	Rooms map[string]*Room
	Room  struct {
		sync.Mutex `yaml:"-"`
		log        *logrus.Entry `yaml:"-"`

		ID              string               `yaml:"id,omitempty"`
		Spec            *Spec                `yaml:"-"`
		Characters      character.Characters `yaml:"-"`
		CharactersCount int                  `yaml:"-"`
	}
)

func (r Spec) Filepath() string {
	return fmt.Sprintf("%s/%s/%s.yaml",
		RoomsFilepath, strings.ToLower(r.ZoneID), strings.ToLower(r.ID))

}

func (r Spec) Validate() error {
	if r.ID == "" {
		return fmt.Errorf("id is required")
	}

	return nil
}

func NewRoomFromSpec(spec *Spec) *Room {
	return NewRoom(spec)
}

func NewRoom(spec *Spec) *Room {
	r := &Room{
		ID:         spec.ID,
		Spec:       spec,
		Characters: make(character.Characters),
	}
	r.log = logrus.WithFields(logrus.Fields{"package": "common", "type": "room", "room_id": r.ID, "room_name": r.Spec.Name})

	return r
}

func (r *Room) AddCharacter(c *character.Character) {
	logrus.WithFields(logrus.Fields{"room_id": r.ID, "character_id": c.ID, "character_name": c.Name}).Debug("Adding character to room")
	r.Lock()
	r.Characters[c.ID] = c
	r.CharactersCount++
	r.Unlock()

	// for k, v := range r.Characters {
	// 	logrus.WithFields(logrus.Fields{"id": k, "name": v.Name}).Debug("Character in room")
	// 	if k == c.ID {
	// 		continue
	// 	}
	// 	// TODO: add direction of entrance
	// 	io.WriteString(v.Session, c.Name+" has entered the room.\n")
	// 	logrus.WithFields(logrus.Fields{"character": c.Name, "room": r.ID}).Info("Character entered room")
	// }
}

func (r *Room) RemoveCharacter(c *character.Character) {
	logrus.WithFields(logrus.Fields{"room_id": r.ID, "character_id": c.ID, "character_name": c.Name}).Debug("Removing character from room")
	// TODO: add direction of exit
	// for k, v := range r.Characters {
	// 	logrus.WithFields(logrus.Fields{"id": k, "name": v.Name}).Debug("Character in room")
	// 	if k == c.ID {
	// 		continue
	// 	}
	// 	io.WriteString(v.Session, c.Name+" has left the room.\n")
	// 	logrus.WithFields(logrus.Fields{"character": c.Name, "room": r.ID}).Info("Character left room")
	// }
	r.Lock()
	delete(r.Characters, c.ID)
	r.CharactersCount--
	r.Unlock()
}

func (r *Room) DisplayRoom(c *character.Character) {
	// io.WriteString(c.Session, "\n")
	// color.New(color.FgBlue, color.Underline).Fprintln(c.Session, r.Spec.Name)
	// color.New(color.FgWhite).Fprintln(c.Session, r.Spec.ShortDescription)
	// color.New(color.FgGreen).Fprintln(c.Session, "Exits:")
	// if len(r.Spec.Exits) == 0 {
	// 	color.New(color.FgGreen).Fprintln(c.Session, "\tNone")
	// 	// } else {
	// 	// for k, v := range r.Spec.Exits {
	// 	// color.New(color.FgGreen).Fprintf(c.Session, "\t%s - %s", k, v)
	// 	// }
	// }
}

var CoreRooms = []Spec{
	{
		ID:               "the_void",
		Name:             "The Void",
		ZoneID:           "the_void",
		ShortDescription: "You step out into ......",
		Description:      "You don't think that you are not floating in nothing.",
		Exits: Exits{
			"north": {Direction: "north", RoomID: "the_void:limbo"},
			"south": {Direction: "south", RoomID: "the_void:the_chat_room"},
		},
	},
	{
		ID:               "limbo",
		Name:             "Limbo",
		ZoneID:           "the_void",
		ShortDescription: "You are floating in a formless void, detached from all sensation of physical matter, surrounded by swirling glowing light, which fades into the relative darkness around you without any trace of edges or shadow.",
		Description:      "There is a \"No Tipping\" notice pinned to the darkness.",
		Exits: Exits{
			"south": {Direction: "south", RoomID: "the_void:the_void"},
			"down":  {Direction: "down", RoomID: "seattle:empty"},
		},
	},
	{
		ID:               "the_chat_room",
		Name:             "The Chat Room",
		ZoneID:           "the_void",
		ShortDescription: "You are lounging in a quiet cosy parlour, warmed by a gentle magical fire which twinkles happily in a warm fireplace.  There are no doors out.  Clearly the owner of this room needs none.",
		Description:      "You are lounging in a quiet cosy parlour, warmed by a gentle magical fire which twinkles happily in a warm fireplace.  There are no doors out.  Clearly the owner of this room needs none.",
		Exits: Exits{
			"north": {Direction: "north", RoomID: "the_void:the_void"},
		},
	},
}
