package protocol

import (
	"github.com/google/uuid"
	"github.com/sandertv/mcwss/protocol/event"
)

// EventRequest is sent by the server to request the client to start sending events of a particular type to
// the server.
type EventRequest struct {
	// EventName is the name of the event, for example 'BlockPlaced'.
	EventName event.Name `json:"eventName"`
}

// NewEventRequest returns an event request for a particular event name and action.
func NewEventRequest(eventName event.Name, purpose MessagePurpose) Packet {
	return Packet{
		Header: Header{
			RequestID:      uuid.New().String(),
			MessagePurpose: purpose,
			Version:        1,
		},
		Body: EventRequest{
			EventName: eventName,
		},
	}
}
