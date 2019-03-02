package command

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
)

// ParticleRequest produces the command required to spawn a particle using a particle name and position.
func ParticleRequest(particle string, position mctype.Position) string {
	return fmt.Sprintf("particle %v %v %v %v", particle, position.X, position.Y, position.Z)
}

// Particle is sent by the server to spawn a particle in the world. It supports both default particles and
// particles defined using behaviour packs.
type Particle struct {
	// StatusCode is the status code of the command. This status code is always non-0, which appears to be
	// a bug in the command.
	StatusCode int `json:"statusCode"`
}
