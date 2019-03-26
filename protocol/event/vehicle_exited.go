package event

// VehicleExited is sent by the client when it exits a vehicle such as a horse or a minecart.
type VehicleExited struct {
	// FurthestAxisMetersTravelled is the total amount of blocks that was travelled during the time that the
	// client was on the vehicle.
	FurthestAxisMetersTravelled int
	// MobType is the numerical ID of the mob that was ridden. For minecarts as an example, this is 84.
	MobType int
	// TravelMethodID is the ID of the travel method. This always appears to be 0.
	TravelMethodID int
	// TripDuration is the duration the trip took. This is measured in seconds, rather than minutes, as the
	// field actually implies.
	TripDuration int `json:"TripDurationMinutes"`
}
