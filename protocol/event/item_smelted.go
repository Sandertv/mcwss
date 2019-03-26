package event

// ItemSmelted is sent by the client when it smelts an item and takes the item out of the result slot. It is
// not sent when the item is actually smelted.
type ItemSmelted struct {
	// AuxType is the metadata value of the resulting item.
	AuxType int
	// Type is the numerical ID of the resulting item.
	Type int
	// FuelSourceAuxType is the metadata value of the fuel item.
	FuelSourceAuxType int
	// FuelSourceType is the numerical ID of the fuel item.
	FuelSourceType int
}
