package event

// ScriptRan is sent by the client when it initially runs a script after having downloaded it from the world
// the client joined.
type ScriptRan struct {
	// ScriptHash is a numeric hash of the script that was ran.
	ScriptHash int
	// ScriptName is the name/path of the script that was ran.
	ScriptName string
	// ScriptRanClientSide indicates if the script that was ran was running client side.
	ScriptRanClientSide bool `json:"ScriptRanClientside"`
	// ScriptSucceeded indicates if the script that was ran ran without any errors occurring. If an error was
	// encountered, this is sent in either a ScriptError event or a ScriptInternalError event.
	ScriptSucceeded bool
}
