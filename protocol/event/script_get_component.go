package event

// ScriptGetComponent is sent by the client when one of its running scripts attempts to get a component of an
// entity.
type ScriptGetComponent struct {
	// ScriptGetComponentSucceeded indicates if the component was returned successfully. This is false if the
	// component name below was not a known component.
	ScriptGetComponentSucceeded bool
	// ScriptComponentName is the name of the component that was attempted to be returned. This name may or
	// may not actually have a valid component associated with it.
	ScriptComponentName string
	// ScriptComponentRegistration ???
	ScriptComponentRegistration int
}
