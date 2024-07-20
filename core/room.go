package core

const (
	RoomsDataPath = "data/rooms"
	RoomFilename  = RoomsDataPath + "/%s/%s.yaml"
)

type RoomSpec struct {
	ID               string            `yaml:"id,omitempty"`
	Name             string            `yaml:"name,omitempty"`
	ShortDescription string            `yaml:"short_description,omitempty"`
	Description      string            `yaml:"description,omitempty"`
	Exits            map[string]string `yaml:"exits,omitempty"`
	Zone             string            `yaml:"zone,omitempty"`
}

type Room struct {
	ID   string    `yaml:"id,omitempty"`
	Spec *RoomSpec `yaml:"-"`
}

var CoreRooms = []RoomSpec{
	{
		ID:               "the_void",
		Name:             "The Void",
		Zone:             "the_void",
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
		Zone:             "the_void",
		ShortDescription: "You are floating in a formless void, detached from all sensation of physical matter, surrounded by swirling glowing light, which fades into the relative darkness around you without any trace of edges or shadow.",
		Description:      "There is a \"No Tipping\" notice pinned to the darkness.",
		Exits: map[string]string{
			"south": "the_void",
		},
	},
	{
		ID:               "the_chat_room",
		Name:             "The Chat Room",
		Zone:             "the_void",
		ShortDescription: "You are lounging in a quiet cosy parlour, warmed by a gentle magical fire which twinkles happily in a warm fireplace.  There are no doors out.  Clearly the owner of this room needs none.",
		Description:      "You are lounging in a quiet cosy parlour, warmed by a gentle magical fire which twinkles happily in a warm fireplace.  There are no doors out.  Clearly the owner of this room needs none.",
		Exits: map[string]string{
			"north": "the_void",
		},
	},
}
