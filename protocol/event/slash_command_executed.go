package event

// SlashCommandExecuted is called by the command when it executes a command that existed.
type SlashCommandExecuted struct {
	// CommandName is the name of the command executed, for example 'setblock'. This command is guaranteed to
	// be registered. Attempting to execute commands that do not exist will not call this event.
	CommandName string
	// ErrorCount is the amount of errors the command encountered. For normal commands, this will always be
	// either 0 or 1. For the so called 'functions', the possible amount of errors encountered depends on the
	// amount of commands executed in the function.
	ErrorCount int
	// ErrorList is a list of the command error codes encountered, joined by '\n'. For example
	// 'commands.generic.num.tooBig'
	ErrorList string `json:",omitempty"`
	// SuccessCount is the amount of success returns the command encountered. For normal commands, this will
	// always be either 0 or 1. For the so called 'functions', the possible amount of success returns depends
	// on the amount of commands executed in the function.
	SuccessCount int
}
