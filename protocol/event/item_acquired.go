package event

const (
	_ = iota
	AcquisitionPickedUp
	AcquisitionCrafted
	AcquisitionTakenFromContainer
	_
	_
	AcquisitionForged
	AcquisitionSmelted
	AcquisitionBrewed
	AcquisitionFilledBottle
	AcquistionTraded
	AcquisitionFished
)

// ItemAcquired is sent by the client when it acquires an item. This is most notably done when a player picks
// an item up from the ground, but also when it crafts a new item, gets one from a chest etc.
type ItemAcquired struct {
	// Type is the numerical ID/type of the item acquired.
	Type int
	// AuxType is the metadata value of the item acquired.
	AuxType int
	// AcquisitionMethod was the method used to acquire the item. This is one of the methods in the constants
	// above.
	AcquisitionMethodID int
}
