package event

const (
	ItemWorn = iota
	ItemEaten
	_
	ItemDrunk
	ItemThrown
	ItemReleased
	ItemBlockEntityCreated // This is sent when a music disk is put in a music box, when an item frame is placed, when a brewing stand is placed, but also when a cake is placed...
	ItemBottleFilled
	ItemBucketFilled
	ItemBucketEmptied
	ItemUsedOnBlock
)

// ItemUsed is sent when a player uses an item.
type ItemUsed struct {
	// Type is the numerical type of the item used.
	Type int
	// AuxType is the metadata value (variant) of the item used.
	AuxType int
	// ItemUseMethod is the method of the item usage.
	ItemUseMethod int
}
