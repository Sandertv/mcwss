package protocol

// MessagePurpose is the purpose a JSON message was sent.
type MessagePurpose string

const (
	Subscribe   MessagePurpose = "subscribe"
	Unsubscribe MessagePurpose = "unsubscribe"
	Event       MessagePurpose = "event"
	Error       MessagePurpose = "error"
	Command     MessagePurpose = "commandRequest"
	Response    MessagePurpose = "commandResponse"
)

// Packets is a map to translate message purposes to their corresponding packets.
var Packets = map[MessagePurpose]interface{}{
	Subscribe:   &EventRequest{},
	Unsubscribe: &EventRequest{},
	Event:       &EventResponse{},
	Error:       &ErrorResponse{},
	Response:    &CommandResponse{},
	Command:     &CommandRequest{},
}
