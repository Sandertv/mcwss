package mcwss

// Config is the configuration of the websocket server.
type Config struct {
	// HandlerPattern is the handler pattern on which the websocket server may be found. By default, this
	// value is '/ws'.
	HandlerPattern string
	// Address is the address+port on which the websocket server is running. By default, this is ':8000'.
	Address string
}

// defaultConfig returns a default configuration that can be connected to using the URI 'localhost:8000/ws'.
func defaultConfig() Config {
	return Config{
		HandlerPattern: "/ws",
		Address: ":8000",
	}
}