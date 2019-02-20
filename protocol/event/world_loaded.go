package event

// WorldLoaded is sent by the client when its world is loaded. (or, more specifically, when it receives the
// StartGamePacket) This is sent for both multiplayer and singleplayer worlds.
type WorldLoaded struct {
	// ContainsBehaviourPacks indicates if the world that is loaded has any global behaviour packs/addons
	// enabled for it.
	ContainsBehaviourPacks bool `json:"ContainsAddons"`
	// ContainsTexturePacks indicates if the world that is loaded has any texture packs enabled for it.
	ContainsTexturePacks bool `json:"ContainsTextures"`
	// IsFromWorldTemplate indicates if the world that is loaded is from a world template, or in other words,
	// comes from a template bought from the marketplace.
	IsFromWorldTemplate bool
	// RequiredHostTexturesToJoin indicates if the texture pack of the world was required to be downloaded to
	// join the world.
	RequiredHostTexturesToJoin bool
	// SaveID is the save ID of the world. This is a base64 string.
	SaveID string `json:"SaveId"`
	// UsingExperimentalGameplay indicates if the world had experimental gameplay enabled in it.
	UsingExperimentalGameplay bool
	// WorldSeed is the numerical seed of the world.
	WorldSeed int
}
