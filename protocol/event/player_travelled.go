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
	TravelSprintings
	TravelBouncing
)

// PlayerTravelled is sent by the client when it travels to a new destination.
type PlayerTravelled struct {
	// MetresTravelled is the distance a player has travelled since the last player travelled event.
	MetresTravelled float64
	// AverageX is the average X position during the travel.
	AverageX float64
	// AverageY is the average Y position during the travel.
	AverageY float64
	// AverageZ is the average Z position during the travel.
	AverageZ float64
	// NewBiome is the new biome of the player. This is often the same as the current biome.
	NewBiome int

	// TravelMethodType is the method type indicating the way the player travelled.
	TravelMethodType int
	// HasRelevantBuff specifies if a player has an effect that affects its movement, such as slowness, jump
	// boost, slow falling etc.
	HasRelevantBuff bool
	// MobType and NetworkType are both filled out with 0. They no longer seem to be valid, nor do they seem
	// to apply to the event.
	MobType     int
	NetworkType int
}

// ConsumeMeasurements takes the movement measurements from the Measurements struct.
func (event *PlayerTravelled) ConsumeMeasurements(measurements Measurements) {
	event.MetresTravelled = measurements.MetresTravelled
	event.AverageX = measurements.PositionAverageX
	event.AverageY = measurements.PositionAverageY
	event.AverageZ = measurements.PositionAverageZ
	event.NewBiome = measurements.NewBiome
}
