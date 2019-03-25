package event

// WorldGenerated is sent by the client when it generates a new singleplayer world.
type WorldGenerated struct {
	// EducationFeaturesEnabled indicates if the newly generated world has education features enabled. Once
	// enabled, it cannot be disabled again.
	EducationFeaturesEnabled bool
	// WorldSeed is the seed of the world that was generated.
	WorldSeed int
	// IsFromWorldTemplate indicates if the world was generated from a world template, or in other words, was
	// bought from the marketplace.
	IsFromWorldTemplate bool `json:"fromTemplate"`
	// SaveID is a base64 encoded string that identifies the save folder of the world.
	SaveID string `json:"SaveId,omitempty"`
}
