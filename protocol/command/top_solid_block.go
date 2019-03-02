package command

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
)

// TopSolidBlockRequest produces the command required to request the top solid block at a location.
func TopSolidBlockRequest(x, z, maxY int) string {
	return fmt.Sprintf("gettopsolidblock %v %v %v", x, maxY, z)
}

// TopSolidBlock is sent by the server to find the top solid block under a given position.
type TopSolidBlock struct {
	// Block is the string block type. For a grass block for example, this is 'grass'.
	Block string `json:"blockName"`
	// AuxType is the metadata value of a block. For a grass block this is 0.
	AuxType int `json:"blockData"`
	// Position is the block position of the block. In particular the Y value is important here, as it
	// indicates the height of the top solid block.
	Position mctype.BlockPosition `json:"position"`
	// StatusCode is the status code of the command. This is 0 on success.
	StatusCode int `json:"statusCode"`
}
