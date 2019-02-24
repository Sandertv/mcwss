package mcwss

import (
	"fmt"
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

// escapeMessage escapes characters in a string so that it can safely be sent in a command to the player.
func escapeMessage(message string) string {
	message = strings.Replace(message, `\`, `\\`, -1)
	message = strings.Replace(message, `'`, `\'`, -1)
	return strings.Replace(message, `"`, `\"`, -1)
}
