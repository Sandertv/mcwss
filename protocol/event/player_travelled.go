package event

const (
	TravelWalking = iota
	TravelSwimmingInWater
	TravelFalling
	TravelClimbing
	TravelSwimmingInLava
	TravelFlying
	TravelVehicle
	TravelSneaking
	TravelSprinting
	TravelBouncing
)

// PlayerTravelled is sent by the client when it travels to a new destination.
type PlayerTravelled struct {
	measurements Measurements
	// TravelMethodType is the method type indicating the way the player travelled.
	TravelMethodType int
	// HasRelevantBuff specifies if a player has an effect that affects its movement, such as slowness, jump
	// boost, slow falling etc.
	HasRelevantBuff bool
	// MobType no longer seems to be used. It appears always to be 0.
	MobType int
	// IsUnderwater indicates if the movement was underwater.
	IsUnderwater bool `json:",omitempty"`
}

// Measurements returns all measurements associated with the player travelled event.
func (event *PlayerTravelled) Measurements() Measurements {
	return event.measurements
}

// ConsumeMeasurements takes the movement measurements from the Measurements struct.
func (event *PlayerTravelled) ConsumeMeasurements(measurements Measurements) {
	event.measurements = measurements
}
