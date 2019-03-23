package event

// StartWorld is sent by the client when it starts a world by clicking it in the main menu.
type StartWorld struct {
	// Host is the host IP address of the world. For local worlds, this field is not set and is empty. For
	// servers, this is the IP address of the server.
	Host string `json:",omitempty"`
}
