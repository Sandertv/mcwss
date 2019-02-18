package protocol

import (
	"encoding/json"
	"github.com/sandertv/mcwss/protocol/event"
)

// EventResponse is sent by the client. It holds information about a particular event listened on by the
// sever.
type EventResponse struct {
	// EventName is the name of the event, for example 'BlockPlaced'.
	EventName event.Name `json:"eventName"`
	// Measurements ...
	Measurements event.Measurements `json:"measurements"`
	// Properties is a collection of properties (un)specific to an event. A part of these properties are
	// shared among all events, others are specific to this event.
	Properties json.RawMessage `json:"properties"`
}
