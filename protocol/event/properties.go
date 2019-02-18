package event

// Properties is a list of properties sent in an event response that are shared among every event.
type Properties struct {
	AccountType int
	ActiveSessionID string
	AppSessionID string
	Biome int
	Build string
	BuildPlatform int `json:"BuildPlat"`
	Achievements bool `json:"Cheevos"`
	ClientID string `json:"ClientId"`
	CurrentInput int
	CurrentNumPlayers int
	DeviceSessionID string `json:"DeviceSessionId"`
	Dimension int `json:"Dim"`
	GlobalMultiplayerCorrelationID string `json:"GlobalMultiplayerCorrelationId"`
	Mode int
	MultiplayerCorrelationID string `json:"MultiplayerCorrelationId"`
	Namespace string
	NetworkType int
	PlacementMethod int
	Platform string `json:"Plat"`
	PlayerGameMode int
	SchemaCommitHash string
	ServerID string `json:"ServerId"`
	ToolItemType int
	Treatments string
	UserID string `json:"UserId"`
	WorldFeature int
	WorldSessionID string `json:"WorldSessionId"`
	EditionType string `json:"editionType"`
	Locale string `json:"locale"`
	VRMode bool `json:"vrMode"`
}
