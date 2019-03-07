package command

// EduClientInfoRequest produces a command that is used to get education edition related information about
// a connected player.
func EduClientInfoRequest() string {
	return "geteduclientinfo"
}

// EduClientInfo is sent by the server to receive education information about the player connected to the
// websocket server. This command is available even if the player is not using education edition.
//
// This is a client-side command which can be called even on third-party servers that do not implement it.
type EduClientInfo struct {
	// IsEdu indicates if the player's game was of education edition.
	IsEdu bool `json:"isEdu"`
	// CompanionProtocolVersion is the 'protocol version' of the learn-to-code agent. This is currently 4.
	CompanionProtocolVersion int `json:"companionProtocolVersion"`
	// IsHost indicates if the player was the host of the game that the player is currently in.
	IsHost bool `json:"isHost"`
	// Permission is the permission level of the player.
	Permission int `json:"permission"`
	// PlayerSessionUUID is the UUID of the session of a player.
	PlayerSessionUUID string `json:"playersessionuuid"`
	// ClientUUID is the UUID of the client.
	ClientUUID string `json:"clientuuid"`
	// StatusCode is the status code of the command. This is 0 when successful.
	StatusCode int `json:"statusCode"`
}
