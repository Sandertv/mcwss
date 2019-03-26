package event

// ItemNamed is sent by the client when it names an item using an anvil.
type ItemNamed struct {
	// AuxType is the metadata value of the item named.
	AuxType int
	// Type is the numerical ID of the item named.
	Type int
}
