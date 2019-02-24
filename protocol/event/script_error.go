package event

// ScriptError is sent by the client when one of its scripts runs into a 'non-critical' error. It does not
// stop execution of the script.
type ScriptError struct {
	// ErrorMessage is an error message describing accurately what the error was in the script, and in what
	// line this error can be found.
	ErrorMessage string
	// ScriptName is supposed to be the name/path of the script, but this appears to be bugged in version 1.9.
	// The value contained is rather part of the error message above.
	ScriptName string
}
