package event

// GameRulesLoaded gets sent by the client when it has its game rules loaded. It contains a list of all
// game rules and their values.
type GameRulesLoaded struct {
	CommandBlockOutput    bool `json:"commandblockoutput"`
	CommandBlocksEnabled  bool `json:"commandblocksenabled"`
	DoDaylightCycle       bool `json:"dodaylightcycle"`
	DoEntityDrops         bool `json:"doentitydrops"`
	DoFireTick            bool `json:"dofiretick"`
	DoImmediateRespawn    bool `json:"doimmediaterespawn"`
	DoInsomnia            bool `json:"doinsomnia"`
	DoMobLoot             bool `json:"domobloot"`
	DoMobSpawning         bool `json:"domobspawning"`
	DoTileDrops           bool `json:"dotiledrops"`
	DoWeatherCycle        bool `json:"doweathercycle"`
	DrowningDamage        bool `json:"drowningdamage"`
	ExperimentalGamePlay  bool `json:"experimentalgameplay"`
	FallDamage            bool `json:"falldamage"`
	FireDamage            bool `json:"firedamage"`
	FunctionCommandLimit  int  `json:"functioncommandlimit"`
	KeepInventory         bool `json:"keepinventory"`
	MaxCommandChainLength int  `json:"maxcommandchainlength"`
	MobGriefing           bool `json:"mobgriefing"`
	NaturalRegeneration   bool `json:"naturalregeneration"`
	PVP                   bool `json:"pvp"`
	RandomTickSpeed       int  `json:"randomtickspeed"`
	SendCommandFeedback   bool `json:"sendcommandfeedback"`
	ShowCoordinates       bool `json:"showcoordinates"`
	ShowDeathMessages     bool `json:"showdeathmessages"`
	TNTExplodes           bool `json:"tntexplodes"`
}
