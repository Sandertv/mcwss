package event

// ScriptBroadcastEvent is sent by the client when one of its running scripts broadcasts an event, for example
// a display chat message event or a command execution event.
type ScriptBroadcastEvent struct {
	// ScriptBroadcastEventSucceeded indicates if the broadcast event succeeded. This is false if the script
	// event name below did not exist.
	ScriptBroadcastEventSucceeded bool
	// ScriptEventName returns the name of the event that was broadcasted, for example
	// 'minecraft:execute_command'.
	ScriptEventName string
	// ScriptEventRegistration ???
	ScriptEventRegistration int
}
