package command

// CloseChatRequest produces the command required to close the chat window of a player.
func CloseChatRequest() string {
	return "closechat"
}

// CloseChat is sent by the server to force the client to close its chat window.
//
// This is a client-side command which can be called even on third-party servers that do not implement it.
type CloseChat struct {
	// StatusCode is the status code of the command. This is 0 if successful.
	StatusCode int `json:"statusCode"`
	// StatusMessage is the status message of the command. If successful, this is 'Chat closed'.
	StatusMessage string `json:"statusMessage"`
}
