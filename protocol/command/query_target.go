package command

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
)

// QueryTargetRequest produces the command used to query information about a target.
func QueryTargetRequest(target mctype.Target) string {
	return fmt.Sprintf("querytarget %v", target)
}

// QueryTarget is sent by the server to find out information about entities in the world, in particular the
// position related information.
type QueryTarget struct {
	// Details is a slice with details for all targets matching the query.
	Details []struct {
		// Dimension is the dimension the entity is currently in.
		Dimension int `json:"dimension"`
		// Position is the current position of the entity.
		Position mctype.Position `json:"position"`
		// UniqueID is the entity unique ID of the entity. For players this is a UUID, for entities this is
		// a negative number.
		UniqueID string `json:"uniqueId"`
		// YRotation is the rotation on the Y axis of the entity. (yaw)
		YRotation float64 `json:"yRot"`
	} `json:"details"`
	// StatusCode is the status code of the response. If successful, this is 0.
	StatusCode int `json:"statusCode"`
	// StatusMessage is the same as Details, except a string.
	StatusMessage string `json:"statusMessage"`
}
