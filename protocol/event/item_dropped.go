package event

// ItemDropped is sent by the client when it drops an item on the ground from its inventory.
type ItemDropped struct {
	// Type is the numerical ID/type of the item dropped.
	Type int
	// AuxType is the metadata value of the item dropped.
	AuxType int
}
