package event

// Properties is a list of properties sent in an event response that are shared among every event.
type Properties struct {
	AccountType                    int
	ActiveSessionID                string
	AppSessionID                   string
	Biome                          int
	Build                          string
	BuildPlatform                  int    `json:"BuildPlat"`
	Achievements                   bool   `json:"Cheevos"`
	ClientID                       string `json:"ClientId"`
	CurrentInput                   int
	CurrentNumPlayers              int    `json:",omitempty"`
	DeviceSessionID                string `json:"DeviceSessionId"`
	Dimension                      int    `json:"Dim"`
	GlobalMultiplayerCorrelationID string `json:"GlobalMultiplayerCorrelationId,omitempty"`
	Mode                           int
	MultiplayerCorrelationID       string `json:"MultiplayerCorrelationId"`
	NetworkType                    int
	Platform                       string `json:"Plat"`
	PlayerGameMode                 int
	SchemaCommitHash               string
	Sequence                       int    `json:"Seq,omitempty"`
	ServerID                       string `json:"ServerId"`
	Treatments                     string
	// UserID is a UUID that is only set when a player is logged into XBOX Live.
	UserID         string `json:"UserId,omitempty"`
	WorldFeature   int
	WorldSessionID string `json:"WorldSessionId,omitempty"`
	IsTrial        int    `json:"isTrial"`
	EditionType    string `json:"editionType"`
	Locale         string `json:"locale"`
	VRMode         bool   `json:"vrMode"`
	// BuildTypeID is a key that is only set when the Minecraft build is a development build. It is not
	// present for normal clients.
	BuildTypeID int `json:",omitempty"`
}
