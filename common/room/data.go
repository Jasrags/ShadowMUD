package room

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
