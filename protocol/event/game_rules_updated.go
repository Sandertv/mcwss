package event

// GameRulesUpdated is sent by the client when a game rule was updated. It is sent for each specific game rule
// change.
type GameRulesUpdated struct {
	// GameRuleName is the name of the game rule that was updated.
	GameRuleName string `json:"UpdatedOptionName"`
	// NewValue is the new value of the game rule that was updated. This may be either a bool, a float64 or
	// an int.
	NewValue interface{} `json:"UpdatedOptionNewValue"`
	// OldValue is the value before the game rule was updated. This may be either a bool, a float64 or an int,
	// but is guaranteed to be of the same type as NewValue.
	OldValue interface{} `json:"UpdatedOptionOldValue"`
}
