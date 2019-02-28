package mctype

// ArmourSlot is an inventory slot in which an armour item may be put.
type ArmourSlot int

const (
	SlotHelmet ArmourSlot = iota + 2
	SlotChestplate
	SlotLeggings
	SlotBoots
)
