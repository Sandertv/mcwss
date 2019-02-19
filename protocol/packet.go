package protocol

// Packet is the main struct that describes every packet sent between client and server to communicate.
// Each packet shares the same header, but a different body.
type Packet struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}
