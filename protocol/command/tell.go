package command

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
)

// TellRequest produces the command required to private message a target.
func TellRequest(target mctype.Target, message string) string {
	return fmt.Sprintf("tell %v %v", target, message)
}

// Tell is sent by the server to send a private message from one player to another.
type Tell struct {
	// StatusMessage is a message describing the status of the command execution. Executing /tell when the
	// privacy settings of the XBOX Live account of the client did not allow chatting to other players will
	// fail. It will come up with a status message like 'You cannot chat with other players because of how
	// your Xbox Live account is set up.  This can be changed in your privacy & online safety settings on
	// Xbox.com.'
	StatusMessage string `json:"statusMessage"`
	// StatusCode describes the status of the command in a code. 0 if successful.
	StatusCode int `json:"statusCode"`
}
