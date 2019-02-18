package event

// Measurable is implemented by events that have information stored in the measurements JSON object, such as
// the PlayerTravelled event.
type Measurable interface {
	// ConsumeMeasurements consumes the measurements passed and stores it in fields in the event.
	ConsumeMeasurements(measurements Measurements)
}

// Measurements contains information about measurements made for this event. Most events use only the top 4
// fields, others such as player travelled event use the others too.
type Measurements struct {
	// Count ...
	Count int `json:"count"`
	// RecordCount ...
	RecordCount int `json:"recordCnt"`
	// SequenceMax ...
	SequenceMax int `json:"seqMax"`
	// SequenceMin ...
	SequenceMin int `json:"seqMin"`

	// MetresTravelled is the amount of metres travelled since the last PlayerTravelled event.
	MetresTravelled float64 `json:"MetersTravelled"`
	// NewBiome is the new biome to which the player moved.
	NewBiome int
	// PositionAverageX is the average position on the X axis of a player during a travelled event.
	PositionAverageX float64 `json:"PosAvgX"`
	// PositionAverageY is the average position on the Y axis of a player during a travelled event.
	PositionAverageY float64 `json:"PosAvgY"`
	// PositionAverageZ is the average position on the Z axis of a player during a travelled event.
	PositionAverageZ float64 `json:"PosAvgZ"`
}
