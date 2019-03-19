package event

const (
	DestructionBreak = iota
)

// BlockBroken is sent by the client when it breaks a block.
type BlockBroken struct {
	// AuxType is the metadata value of the block broken.
	AuxType int
	// Block is the block that was broken. This is a string like 'stone'.
	Block string
	// Variant is the variant of the block broken. This is the same as AuxType.
	Variant int
	// Namespace is the namespace to which the block belongs. Usually, this is 'minecraft'.
	Namespace string
	// Type is the numerical ID of the block broken. For stone, this is 1.
	Type int
	// DestructionMethod is the method used to break the block.
	DestructionMethod int
	// ToolItemType is the item type of the item that was used to break the block.
	ToolItemType int
}
