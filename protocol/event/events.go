package event

// Name is the name of an event.
type Name string

const (
	NameAwardAchievement     Name = "AwardAchievement"
	NameBlockPlaced          Name = "BlockPlaced"
	NameBlockBroken          Name = "BlockBroken"
	NameEndOfDay             Name = "EndOfDay"
	NameGameRulesLoaded      Name = "GameRulesLoaded"
	NameGameRulesUpdated     Name = "GameRulesUpdated"
	NamePlayerMessage        Name = "PlayerMessage"
	NamePlayerTeleported     Name = "PlayerTeleported"
	NamePlayerTravelled      Name = "PlayerTravelled"
	NamePlayerTransform      Name = "PlayerTransform"
	NameItemAcquired         Name = "ItemAcquired"
	NameItemCrafted          Name = "ItemCrafted"
	NameItemDropped          Name = "ItemDropped"
	NameItemEquipped         Name = "ItemEquipped"
	NameItemInteracted       Name = "ItemInteracted"
	NameItemNamed            Name = "ItemNamed"
	NameItemSmelted          Name = "ItemSmelted"
	NameItemUsed             Name = "ItemUsed"
	NameBookEdited           Name = "BookEdited"
	NameSignedBookOpened     Name = "SignedBookOpened"
	NameMobBorn              Name = "MobBorn"
	NameMobInteracted        Name = "MobInteracted"
	NameMobKilled            Name = "MobKilled"
	NameStartWorld           Name = "StartWorld"
	NameWorldLoaded          Name = "WorldLoaded"
	NameWorldGenerated       Name = "WorldGenerated"
	NameScriptLoaded         Name = "ScriptLoaded"
	NameScriptRan            Name = "ScriptRan"
	NameScreenChanged        Name = "ScreenChanged"
	NameSlashCommandExecuted Name = "SlashCommandExecuted"
	NameSignInToXBOXLive     Name = "SignInToXboxLive"
	NameSignOutOfXBOXLive    Name = "SignOutOfXboxLive"
	NameVehicleExited        Name = "VehicleExited"
)

// Events is a map used to find the corresponding event for an event name.
var Events = map[Name]func() interface{}{
	NameAwardAchievement:     func() interface{} { return &AwardAchievement{} },
	NameBlockPlaced:          func() interface{} { return &BlockPlaced{} },
	NameBlockBroken:          func() interface{} { return &BlockBroken{} },
	NameEndOfDay:             func() interface{} { return &EndOfDay{} },
	NameGameRulesLoaded:      func() interface{} { return &GameRulesLoaded{} },
	NameGameRulesUpdated:     func() interface{} { return &GameRulesUpdated{} },
	NamePlayerMessage:        func() interface{} { return &PlayerMessage{} },
	NamePlayerTeleported:     func() interface{} { return &PlayerTeleported{} },
	NamePlayerTravelled:      func() interface{} { return &PlayerTravelled{} },
	NamePlayerTransform:      func() interface{} { return &PlayerTransform{} },
	NameItemAcquired:         func() interface{} { return &ItemAcquired{} },
	NameItemCrafted:          func() interface{} { return &ItemCrafted{} },
	NameItemDropped:          func() interface{} { return &ItemDropped{} },
	NameItemEquipped:         func() interface{} { return &ItemEquipped{} },
	NameItemInteracted:       func() interface{} { return &ItemInteracted{} },
	NameItemNamed:            func() interface{} { return &ItemNamed{} },
	NameItemSmelted:          func() interface{} { return &ItemSmelted{} },
	NameItemUsed:             func() interface{} { return &ItemUsed{} },
	NameBookEdited:           func() interface{} { return &BookEdited{} },
	NameSignedBookOpened:     func() interface{} { return &SignedBookOpened{} },
	NameMobBorn:              func() interface{} { return &MobBorn{} },
	NameMobInteracted:        func() interface{} { return &MobInteracted{} },
	NameMobKilled:            func() interface{} { return &MobKilled{} },
	NameStartWorld:           func() interface{} { return &StartWorld{} },
	NameWorldLoaded:          func() interface{} { return &WorldLoaded{} },
	NameWorldGenerated:       func() interface{} { return &WorldGenerated{} },
	NameScriptLoaded:         func() interface{} { return &ScriptLoaded{} },
	NameScriptRan:            func() interface{} { return &ScriptRan{} },
	NameScreenChanged:        func() interface{} { return &ScreenChanged{} },
	NameSlashCommandExecuted: func() interface{} { return &SlashCommandExecuted{} },
	NameSignInToXBOXLive:     func() interface{} { return &SignInToXBOXLive{} },
	NameSignOutOfXBOXLive:    func() interface{} { return &SignOutOfXBOXLive{} },
	NameVehicleExited:        func() interface{} { return &VehicleExited{} },
}
