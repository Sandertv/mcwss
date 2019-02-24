package event

// ItemCrafted is sent by the client when it crafts an item either in its inventory or in a crafting table.
type ItemCrafted struct {
	// Type is the numerical type of the crafted item. For wooden planks for example, this is 5.
	Type int
	// AuxType is the metadata or variant of the item. For oak wooden planks for example, this is 0.
	AuxType int
	// UsedCraftingTable indicates if the player was using a crafting table to craft this item. If the
	// inventory itself was used, this is false.
	UsedCraftingTable bool
	// CraftingSessionID is an ID identifying this crafting session.
	CraftingSessionID int

	// RecipeBookShown indicates if a player had its recipe book shown. If so, the fields below apply.
	RecipeBookShown bool
	// NumberOfTabsChanged is the number of times a player changed the inventory tabs while having the
	// recipe book open. There are 5 tabs in total, (Construction, Equipment, Items, Nature and Craftable) but
	// the NumberOfTabsChanged may exceed this number.
	NumberOfTabsChanged int
	// EndingTabID is the ID of the last tab the player had opened in the recipe book. The tabs Construction,
	// Equipment, Items, Nature and Craftable have the IDs 1, 2, 3, 4 and 5 respectively.
	EndingTabID int
	// StartingTabID is the ID of the first tab the player had opened when it opened the recipe book in its
	// crafting inventory. The tabs Construction, Equipment, Items, Nature and Craftable have the IDs 1, 2, 3,
	// 4 and 5 respectively.
	StartingTabID int
	// HasCraftableFilterOn indicates if the Craftable filter was enabled (it is enabled by default). When
	// disabled, the recipe book shows items that cannot be crafted with the collected items with a red
	// background.
	HasCraftableFilterOn bool
	// CraftedAutomatically indicates if the item was directly crafted using the recipe book. This is done by
	// shift clicking an item in the recipe book, causing the products to be put directly into the inventory.
	CraftedAutomatically bool
	// UsedSearchBar indicates if the search bar was used to find the recipe of the item crafted.
	UsedSearchBar bool
}
