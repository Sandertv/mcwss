package mcwss

import (
	"github.com/sandertv/mcwss/mctype"
	"github.com/sandertv/mcwss/protocol/command"
)

// Agent is the agent of a player. The agent is an entity that can be created and controlled using websockets.
type Agent struct {
	player *Player
	pos    mctype.Position
}

// NewAgent returns a new agent for the player passed.
func NewAgent(player *Player) *Agent {
	return &Agent{player: player}
}

// Position requests the position of the agent and passes it into the function passed when a response is
// received.
func (agent *Agent) Position(f func(position mctype.Position)) {
	agent.player.Exec(command.AgentPositionRequest(), func(response *command.AgentPosition) {
		f(response.Position)
	})
}

// Rotation requests the rotation of the agent and passes it into the function passed when a response is
// received.
func (agent *Agent) Rotation(f func(yRotation float64)) {
	agent.player.Exec(command.AgentPositionRequest(), func(response *command.AgentPosition) {
		f(response.YRotation)
	})
}

// Move moves the agent in a given direction for a given amount of metres.
func (agent *Agent) Move(direction mctype.Direction, metres int) {
	for i := 0; i < metres; i++ {
		agent.player.Exec(command.AgentMoveRequest(direction), nil)
	}
}

// TurnRight turns the agent 90 degrees to the right.
func (agent *Agent) TurnRight() {
	agent.player.Exec("agent turn right", nil)
}

// TurnLeft turns the agent 90 degrees to the left.
func (agent *Agent) TurnLeft() {
	agent.player.Exec("agent turn left", nil)
}
