package event

const (
	TeleportationCommand = iota
	TeleportationEnderPearl
	TeleportationChorusFruit
)

// PlayerTeleported is sent by the client when the player teleports. This is sent for example when using the
// /tp command, but also when, for example, eating a chorus fruit.
type PlayerTeleported struct {
	// MetersTravelled is the total distance that was travelled with the teleportation.
	MetersTravelled float64
	// TeleportationCause is the cause of the teleportation. This is one of the constant values above.
	TeleportationCause int
	// TeleporationItem is the item used to teleport. If teleported using a command, this item value is 0. For
	// an ender pearl, this is 87.
	TeleportationItem int
}
