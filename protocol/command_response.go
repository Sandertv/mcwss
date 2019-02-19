package protocol

import "encoding/json"

// CommandResponse is a merely a JSON raw message, because its value depends on what command was executed.
type CommandResponse json.RawMessage
