package event

// PlayerTransform is sent by the client when it is 'transformed', meaning it was moved from one place to
// another. This appears not to be sent during movement, (see PlayerTravelled for that) but rather when the
// player is, for example, teleported.
type PlayerTransform struct {
	// PlayerYRot is the Y rotation of the player after the transformation.
	PlayerYRotation float64 `json:"PlayerYRot"`
	// PlayerID is the UUID of the player.
	PlayerID string `json:"PlayerId"`
	// Dimension is the dimension of the player.
	Dimension int `json:"Dimension"`
	// PositionX is the position on the X axis of the player after the transformation.
	PositionX float64 `json:"PosX"`
	// PositionY is the position on the Y axis of the player after the transformation.
	PositionY float64 `json:"PosY"`
	// PositionZ is the position on the Z axis of the player after the transformation.
	PositionZ float64 `json:"PosZ"`
}
