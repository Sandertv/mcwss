package command

import (
	"bytes"
	"encoding/json"
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
	// Details is a slice with details for all targets matching the query. It is an escaped JSON string,
	// unlike a proper JSON array as might be expected.
	Details *QueryResults `json:"details"`
	// StatusCode is the status code of the response. If successful, this is 0.
	StatusCode int `json:"statusCode"`
	// StatusMessage is the same as Details, except a string.
	StatusMessage string `json:"statusMessage"`
}

// QueryResults is a slice with details for all targets matching the query.
type QueryResults queryResults

// queryResults is a slice with details for all targets matching the query.
type queryResults []struct {
	// Dimension is the dimension the entity is currently in.
	Dimension int `json:"dimension"`
	// Position is the current position of the entity.
	Position mctype.Position `json:"position"`
	// UniqueID is the entity unique ID of the entity. For players this is a UUID, for entities this is
	// a negative number.
	UniqueID string `json:"uniqueId"`
	// YRotation is the rotation on the Y axis of the entity. (yaw)
	YRotation float64 `json:"yRot"`
}

// UnmarshalJSON unmarshals a data slice passed and implements the json.Unmarshaler. It is implemented to make
// sure the details string is unmarshaled to an array properly.
func (results *QueryResults) UnmarshalJSON(data []byte) error {
	data = bytes.Replace(data, []byte{'\\', 'n'}, []byte{}, -1)
	data = bytes.Replace(data, []byte{'\\'}, []byte{}, -1)
	data = bytes.Replace(data, []byte{' '}, []byte{}, -1)

	res := queryResults(*results)
	if err := json.Unmarshal(bytes.Trim(data, `"`), &res); err != nil {
		return err
	}

	*results = QueryResults(res)
	return nil
}
