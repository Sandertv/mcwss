package protocol

import "github.com/google/uuid"

// CommandRequest is used to send a raw command line to the client. The client itself will do the
// processing of its arguments.
type CommandRequest struct {
	// CommandLine is the full raw command line to be executed. The command line must not include the leading
	// slash. For example 'tell player this was sent using ws'.
	CommandLine string `json:"commandLine"`
	// Version is the version of the packet. Currently this is 1.
	Version int `json:"version"`
}

// NewCommandRequest returns a packet for a raw command to be executed.
func NewCommandRequest(commandLine string) Packet {
	return Packet{
		Header: Header{
			RequestID:      uuid.New().String(),
			MessagePurpose: Command,
			Version:        1,
		},
		Body: CommandRequest{
			Version:     1,
			CommandLine: commandLine,
		},
	}
}
