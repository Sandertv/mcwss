package event

// ScriptInternalError is sent by the client when an internal 'critical' error is encountered during the
// running of a script. Generally, these errors are encountered during the initial join, and they are often
// related to syntax errors or related bugs.
type ScriptInternalError struct {
	// ErrorMessage is the error message. Usually this error message is not very clear. It is often something
	// like 'Something went wrong...'.
	ErrorMessage string
}
