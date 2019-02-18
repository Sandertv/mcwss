package protocol

import (
	"github.com/google/uuid"
	"github.com/sandertv/mcwss/protocol/event"
)

// Packet is the main struct that describes every packet sent between client and server to communicate.
// Each packet shares the same header, but a different body.
type Packet struct {
	Header Header `json:"header"`
	Body interface{} `json:"body"`
}

// NewEventRequest returns an event request for a particular event name and action.
func NewEventRequest(eventName event.Name, purpose MessagePurpose) Packet {
	return Packet{
		Header: Header{
			RequestID: uuid.New().String(),
			MessagePurpose: purpose,
			Version: 1,
		},
		Body: EventRequest{
			EventName: eventName,
		},
	}
}
