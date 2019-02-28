package event

import "github.com/sandertv/mcwss/mctype"

// ItemEquipped is sent by the client when it equips an item. This includes armour, elytras, pumpkins and
// other items that may be worn.
type ItemEquipped struct {
	// ItemEnchantmentCount is the number of enchants the item equipped has. Even though this event only has 3
	// enchantment types, an item may have more than 3 enchantments on it. That means that the ItemEnchantment
	// Count may exceed 3.
	ItemEnchantmentCount int `json:"ItemEnchantCount"`
	// ItemEnchantmentTypeA is the enchantment type/ID of the first enchantment on the item.
	ItemEnchantmentTypeA int `json:"ItemEnchantTypeA"`
	// ItemEnchantmentLevelA is the level of the first enchantment on the item.
	ItemEnchantmentLevelA int `json:"ItemEnchantLevelA"`
	// ItemEnchantmentTypeB is the enchantment type/ID of the second enchantment on the item.
	ItemEnchantmentTypeB int `json:"ItemEnchantTypeB"`
	// ItemEnchantmentLevelB is the level of the second enchantment on the item.
	ItemEnchantmentLevelB int `json:"ItemEnchantLevelB"`
	// ItemEnchantmentTypeB is the enchantment type/ID of the second enchantment on the item.
	ItemEnchantmentTypeC int `json:"ItemEnchantTypeC"`
	// ItemEnchantmentLevelC is the level of the third enchantment on the item.
	ItemEnchantmentLevelC int `json:"ItemEnchantLevelC"`

	// Slot is the slot the item was equipped in. This is a number ranging from 2-5, with 2 being the helmet
	// and 5 being the boots.
	Slot mctype.ArmourSlot
	// Type is the numerical ID of the item equipped.
	Type int
	// AuxType is the aux value of the item equipped. For most items equipped, this will be the durability or
	// damage of the item.
	AuxType int
}
