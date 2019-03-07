package command

import "fmt"

// LocalPlayerNameRequest produces a command used to get the name of the player connected to a websocket.
func LocalPlayerNameRequest() string {
	return fmt.Sprintf("getlocalplayername")
}

// LocalPlayerName is sent by the server to find out the player name of the player connected to the websocket.
// This command is executed automatically by the server and stored in a field of the Player struct.
//
// This is a client-side command which can be called even on third-party servers that do not implement it.
type LocalPlayerName struct {
	// LocalPlayerName is the name of the player connected.
	LocalPlayerName string `json:"localplayername"`
	// StatusCode is the status code of the command. This is generally 0 for this command.
	StatusCode int `json:"statusCode"`
	// StatusMessage is the same as LocalPlayerName for this command.
	StatusMessage string `json:"statusMessage"`
}
