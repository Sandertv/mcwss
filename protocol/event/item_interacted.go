package event

const (
	InteractedUse = iota
	InteractedPlace
)

// ItemInteracted is sent by the client when it interacts with a block using an item.
type ItemInteracted struct {
	// AuxType is the metadata value or variant of the item.
	AuxType int `json:"Aux"`
	// Item is the item string of the item, excluding the namespace. For example 'wooden_door'. If no item is
	// held, this string is empty.
	Item string `json:"Id"`
	// Count is the count of the item stack the player held. If no item was held, the count is 0.
	Count int
	// Method is the method of the interaction. The method is one of the constants above.
	Method int
}
