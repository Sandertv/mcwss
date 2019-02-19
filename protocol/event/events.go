package event

// Name is the name of an event.
type Name string

const (
	NameBlockPlaced     Name = "BlockPlaced"
	NameBlockBroken     Name = "BlockBroken"
	NamePlayerMessage   Name = "PlayerMessage"
	NamePlayerTravelled Name = "PlayerTravelled"
	NamePlayerTransform Name = "PlayerTransform"
	NameItemUsed        Name = "ItemUsed"
	NameBookEdited      Name = "BookEdited"
)

// Events is a map used to find the corresponding event for an event name.
var Events = map[Name]interface{}{
	NameBlockPlaced:     &BlockPlaced{},
	NameBlockBroken:     &BlockBroken{},
	NamePlayerMessage:   &PlayerMessage{},
	NamePlayerTravelled: &PlayerTravelled{},
	NamePlayerTransform: &PlayerTransform{},
	NameItemUsed:        &ItemUsed{},
	NameBookEdited:      &BookEdited{},
}
