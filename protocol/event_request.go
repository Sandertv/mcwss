package protocol

import "github.com/sandertv/mcwss/protocol/event"

// EventRequest is sent by the server to request the client to start sending events of a particular type to
// the server.
type EventRequest struct {
	// EventName is the name of the event, for example 'BlockPlaced'.
	EventName event.Name `json:"eventName"`
}
