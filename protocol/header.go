package protocol

// Header describes the header of a packet. Each packet shares the same header.
type Header struct {
	// RequestID is a UUID specific to the request.
	RequestID string `json:"requestId"`
	// MessagePurpose is the purpose of the request. To subscribe to an event, this is 'subscribe'. To
	// unsubscribe from an event, this is 'unsubscribe'. For an event response, this is 'event'.
	MessagePurpose MessagePurpose `json:"messagePurpose"`
	// Version is the version of the request. Currently 1.
	Version int `json:"version"`
}
