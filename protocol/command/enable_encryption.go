package command

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// EnableEncryptionRequest produces the command required to enable encryption using a public key byte slice
// and salt passed. The salt passed must be exactly 16 bytes long.
func EnableEncryptionRequest(publicKey []byte, salt []byte) string {
	if len(salt) != 16 {
		panic("invalid salt given: expected salt with length 16, but got salt with length " + strconv.Itoa(len(salt)))
	}
	return fmt.Sprintf(`enableencryption "%v" "%v"`, base64.RawStdEncoding.EncodeToString(publicKey), base64.RawStdEncoding.EncodeToString(salt))
}

// EnableEncryption is sent by the server to enable websocket encryption. The server sends its public key and
// salt in the command itself, and the client submits its own public key in a response.
//
// This is a client-side command which can be called even on third-party servers that do not implement it.
type EnableEncryption struct {
	// PublicKey is the public key of the client. It is base64 encoded, so it must first be decoded before it
	// can be converted to a key.
	PublicKey string `json:"publicKey"`
	// StatusCode is the status code of the command. This is 0 if successful.
	StatusCode int `json:"statusCode"`
	// StatusMessage is the status message of the command. If
	StatusMessage string `json:"statusMessage"`
}
