package event

// EndOfDay is sent by the client when a day progresses naturally. The event is called at sunrise if the time
// was not modified using a command. (daytime = 23000)
type EndOfDay struct {
	// EndOfDay has no fields.
}
