package event

// Name is the name of an event.
type Name string

const (
	NameBlockPlaced          Name = "BlockPlaced"
	NameBlockBroken          Name = "BlockBroken"
	NameEndOfDay             Name = "EndOfDay"
	NamePlayerMessage        Name = "PlayerMessage"
	NamePlayerTravelled      Name = "PlayerTravelled"
	NamePlayerTransform      Name = "PlayerTransform"
	NameItemCrafted          Name = "ItemCrafted"
	NameItemInteracted       Name = "ItemInteracted"
	NameItemUsed             Name = "ItemUsed"
	NameBookEdited           Name = "BookEdited"
	NameSignedBookOpened     Name = "SignedBookOpened"
	NameMobInteracted        Name = "MobInteracted"
	NameStartWorld           Name = "StartWorld"
	NameWorldLoaded          Name = "WorldLoaded"
	NameWorldGenerated       Name = "WorldGenerated"
	NameScriptBroadcastEvent Name = "ScriptBroadcastEvent"
	NameScriptError          Name = "ScriptError"
	NameScriptGetComponent   Name = "ScriptGetComponent"
	NameScriptInternalError  Name = "ScriptInternalError"
	NameScriptListenToEvent  Name = "ScriptListenToEvent"
	NameScriptLoaded         Name = "ScriptLoaded"
	NameScriptRan            Name = "ScriptRan"
	NameScreenChanged        Name = "ScreenChanged"
)

// Events is a map used to find the corresponding event for an event name.
var Events = map[Name]interface{}{
	NameBlockPlaced:          &BlockPlaced{},
	NameBlockBroken:          &BlockBroken{},
	NameEndOfDay:             &EndOfDay{},
	NamePlayerMessage:        &PlayerMessage{},
	NamePlayerTravelled:      &PlayerTravelled{},
	NamePlayerTransform:      &PlayerTransform{},
	NameItemCrafted:          &ItemCrafted{},
	NameItemInteracted:       &ItemInteracted{},
	NameItemUsed:             &ItemUsed{},
	NameBookEdited:           &BookEdited{},
	NameSignedBookOpened:     &SignedBookOpened{},
	NameMobInteracted:        &MobInteracted{},
	NameStartWorld:           &StartWorld{},
	NameWorldLoaded:          &WorldLoaded{},
	NameWorldGenerated:       &WorldGenerated{},
	NameScriptBroadcastEvent: &ScriptBroadcastEvent{},
	NameScriptError:          &ScriptError{},
	NameScriptGetComponent:   &ScriptGetComponent{},
	NameScriptInternalError:  &ScriptInternalError{},
	NameScriptListenToEvent:  &ScriptListenToEvent{},
	NameScriptLoaded:         &ScriptLoaded{},
	NameScriptRan:            &ScriptRan{},
	NameScreenChanged:        &ScreenChanged{},
}
