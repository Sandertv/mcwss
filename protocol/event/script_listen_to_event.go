package event

// ScriptListenToEvent is sent by the client when a script currently active starts listening for a specific
// event. It is possible to use this event for Script -> websocket server communication.
type ScriptListenToEvent struct {
	// ScriptEventName is the name of the scripting event listened on, for example 'minecraft:entity_death'.
	// Note that this script event name does not have to be an event that is actually called by Minecraft.
	// Script -> Websocket server communication is possible using this.
	ScriptEventName string
}
