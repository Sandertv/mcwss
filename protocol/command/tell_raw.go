package command

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
)

// TellRawRequest produces a tell raw command string. The function does not implement logic such as
// translations and formatting.
func TellRawRequest(target mctype.Target, lines ...string) string {
	cmd := fmt.Sprintf(`tellraw %v {"rawtext":[`, target)
	for i, text := range lines {
		cmd += `{"text":"` + text + `"}`
		if i != len(lines)-1 {
			cmd += `,`
		}
	}
	return cmd + `]}`
}

// TellRaw is sent by the server to send a raw message to a player, possibly with additional formatting
// options.
type TellRaw struct {
	// Recipients is a slice of recipients of the raw message sent. This is useful in the case of unspecific
	// target selectors used.
	Recipients []string `json:"recipient"`
	// StatusCode is the status code of the command. It is 0 on success.
	StatusCode int `json:"statusCode"`
}
