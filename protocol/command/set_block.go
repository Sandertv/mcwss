package command

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
)

// SetBlockRequest produces the command required to place a block in the world of a player. placementMethod
// can be one of either 'replace', 'destroy' and 'keep'. The effects of each of these is explained on the
// Minecraft Wiki page here:
func SetBlockRequest(position mctype.BlockPosition, block string, tileData byte, placementMethod string) string {
	return fmt.Sprintf("setblock %v %v %v %v %v %v", position.X, position.Y, position.Z, block, tileData, placementMethod)
}

// SetBlock is sent by the server to set a block at a specific position using a particular placement method.
type SetBlock struct {
	// Position is the position that the block ended up being placed at.
	Position mctype.Position `json:"position"`
	// StatusMessage is the status message of the command. This is 'Block placed' if the command was
	// successful.
	StatusMessage string `json:"statusMessage"`
	// StatusCode is the status code of the command. If successful, this is 0.
	StatusCode int `json:"statusCode"`
}
