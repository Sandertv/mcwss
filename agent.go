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
	agent.player.Exec(command.AgentTurnRequest(mctype.Right), nil)
}

// TurnLeft turns the agent 90 degrees to the left.
func (agent *Agent) TurnLeft() {
	agent.player.Exec(command.AgentTurnRequest(mctype.Left), nil)
}

// Attack makes the agent attack up to one block in a given direction.
func (agent *Agent) Attack(direction mctype.Direction) {
	agent.player.Exec(command.AgentAttackRequest(direction), nil)
}

// UseHeldItem makes the agent place its held item, provided it is a block, in a given direction.
func (agent *Agent) UseHeldItem(direction mctype.Direction) {
	agent.player.Exec(command.AgentPlaceRequest(direction), nil)
}

// DestroyBlock makes the agent destroy a block in a direction. The agent can break any block, regardless of
// whether the block is breakable in survival mode or not. Bedrock, water etc. can be broken.
func (agent *Agent) DestroyBlock(direction mctype.Direction) {
	agent.player.Exec(command.AgentDestroyRequest(direction), nil)
}

// TillBlock makes the agent till a block in a given direction using a hoe. Tilling does not appear to work
// in non-edu mode, and seems to often crash the game instead.
func (agent *Agent) TillBlock(direction mctype.Direction) {
	agent.player.Exec(command.AgentTillRequest(direction), nil)
}
