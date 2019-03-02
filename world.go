package mcwss

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
	"github.com/sandertv/mcwss/protocol/command"
	"strings"
)

// World represents a world a player has joined. This might be either a singleplayer world, or a multiplayer
// server. Multiworld servers are treated as if they had only one world.
type World struct {
	player *Player
}

// NewWorld returns a new world representing the world the player passed is in.
func NewWorld(player *Player) *World {
	return &World{player: player}
}

// Broadcast broadcasts a message with optional parameters (operates using fmt.Sprintf) to all players in the
// world. The message is prefixed with '[External]'.
func (world *World) Broadcast(message string, parameters ...interface{}) {
	message = fmt.Sprintf(message, parameters...)
	world.player.Exec(command.SayRequest(message), nil)
}

// SetBlock sets a block at a given position. It uses the data value, provided it is a value of 0-15. If not,
// the function panics.
func (world *World) SetBlock(position mctype.BlockPosition, block string, dataValue byte) {
	if dataValue > 15 {
		panic("block data value " + string(dataValue) + " exceeds the max value of 15")
	}
	world.player.Exec(command.SetBlockRequest(position, block, dataValue, "replace"), nil)
}

// DestroyBlock destroys a block at the position passed. It will show the block breaking particles and sounds
// that would normally be sent when destroying a block.
func (world *World) DestroyBlock(position mctype.BlockPosition) {
	world.player.Exec(command.SetBlockRequest(position, "air", 0, "destroy"), nil)
}

// SpawnParticle spawns a particle with a given name into the world at the position passed.
func (world *World) SpawnParticle(particle string, position mctype.Position) {
	world.player.Exec(command.ParticleRequest(particle, position), nil)
}

// escapeMessage escapes characters in a string so that it can safely be sent in a command to the player.
func escapeMessage(message string) string {
	message = strings.Replace(message, `\`, `\\`, -1)
	message = strings.Replace(message, `'`, `\'`, -1)
	return strings.Replace(message, `"`, `\"`, -1)
}
