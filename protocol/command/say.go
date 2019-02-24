package command

import "fmt"

// SayRequest produces the command used to broadcast a message to all players in a world.
func SayRequest(message string) string {
	return fmt.Sprintf("say %v", message)
}

// Say is sent by the server to broadcast a message to all players in a world.
type Say struct {
	// Message is the message that was broadcasted using the command.
	Message string `json:"message"`
	// StatusCode returns the status code of the command. 0 if successful.
	StatusCode int `json:"statusCode"`
}
