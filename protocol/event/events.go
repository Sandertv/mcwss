package event

// Name is the name of an event.
type Name string

const (
	NameBlockPlaced     Name = "BlockPlaced"
	NameBlockBroken     Name = "BlockBroken"
	NamePlayerMessage   Name = "PlayerMessage"
	NameItemUsed        Name = "ItemUsed"
	NamePlayerTravelled Name = "PlayerTravelled"
)

// Events is a map used to find the corresponding event for an event name.
var Events = map[Name]interface{}{
	NameBlockPlaced:     &BlockPlaced{},
	NameBlockBroken:     &BlockBroken{},
	NamePlayerMessage:   &PlayerMessage{},
	NameItemUsed:        &ItemUsed{},
	NamePlayerTravelled: &PlayerTravelled{},
}
