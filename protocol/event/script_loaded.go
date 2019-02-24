package event

// ScriptLoaded is sent by the client after it loaded a script that was sent, but before the script was
// actually ran.
type ScriptLoaded struct {
	// ScriptHash is the numerical hash of the script.
	ScriptHash int
	// ScriptName is the name/path of the script that was loaded.
	ScriptName string
}
