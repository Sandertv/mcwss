package event

// Name is the name of an event.
type Name string

const (
	NameBlockPlaced      Name = "BlockPlaced"
	NameBlockBroken      Name = "BlockBroken"
	NameEndOfDay         Name = "EndOfDay"
	NamePlayerMessage    Name = "PlayerMessage"
	NamePlayerTravelled  Name = "PlayerTravelled"
	NamePlayerTransform  Name = "PlayerTransform"
	NameItemCrafted      Name = "ItemCrafted"
	NameItemInteracted   Name = "ItemInteracted"
	NameItemUsed         Name = "ItemUsed"
	NameBookEdited       Name = "BookEdited"
	NameSignedBookOpened Name = "SignedBookOpened"
)

// Events is a map used to find the corresponding event for an event name.
var Events = map[Name]interface{}{
	NameBlockPlaced:      &BlockPlaced{},
	NameBlockBroken:      &BlockBroken{},
	NameEndOfDay:         &EndOfDay{},
	NamePlayerMessage:    &PlayerMessage{},
	NamePlayerTravelled:  &PlayerTravelled{},
	NamePlayerTransform:  &PlayerTransform{},
	NameItemCrafted:      &ItemCrafted{},
	NameItemInteracted:   &ItemInteracted{},
	NameItemUsed:         &ItemUsed{},
	NameBookEdited:       &BookEdited{},
	NameSignedBookOpened: &SignedBookOpened{},
}
