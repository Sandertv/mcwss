package mctype

// Position is the position of an entity in a 3D world.
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// BlockPosition is the position of a block in a 3D world.
type BlockPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}
