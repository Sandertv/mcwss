package event

// MobBorn is sent by the client when a mob is born by having 2 mobs breed.
type MobBorn struct {
	// BabyColour is the colour of the baby. This is 0 for all animals that do not have a colour. For sheep
	// and horses, this colour is set.
	BabyColour int `json:"BabyColor"`
	// BabyType is the numerical ID of the entity born. For a chicken for example, this is 10.
	BabyType int
	// BabyVariant is the variant of the baby born. This variant is 0 for most mobs.
	BabyVariant int
}
