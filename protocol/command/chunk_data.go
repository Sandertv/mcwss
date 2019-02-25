package command

import (
	"encoding/base64"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

// ChunkDataRequest produces the command used to get a string of chunk data of a certain chunk.
// The command is only available on education edition games. This may be enabled using minecraft://?edu=1.
func ChunkDataRequest(dimension string, chunkX, chunkZ int, maxY int) string {
	return fmt.Sprintf("getchunkdata %v %v %v %v", dimension, chunkX, chunkZ, maxY)
}

// ChunkData is sent by the server to receive a string of chunk data of the highest blocks in a chunk,
// containing an RGB-24 colour of the blocks and their heights.
type ChunkData struct {
	// Data is a string of data, with the data being joined with commas. The string has some methods to make
	// consume less space. The Data string is ordered in XZ order. Multiple blocks after each other result in
	// a multiplier being added to the previous. The multiplier needs one added to get the total amount.
	// "AAAAAQ*254" translates to 255 times AAAAAQ in a row.
	// If the same block is found twice in the chunk data at different positions, the second position gets a
	// pointer to the data of the first position. This results in values like these: "jMnVPw,c6SuPw,11,12,11".
	// When having all offsets to their correct data, each string must be base64 decoded. Note that these
	// strings do not end with '=='. This needs to be added if needed.
	// Each of the values that were base64 decoded exist out of 4 bytes. The first 3 bytes are the blue, green
	// and red colours respectively. The last byte is the height of the block.
	Data string `json:"data"`
	// StatusCode is the status code of the command. This is 0 on success.
	StatusCode int `json:"statusCode"`
	// StatusMessage is a message indicating if the command was successful.
	StatusMessage string `json:"statusMessage"`
}

// ParseChunkData parses a chunk data string from a ChunkData struct. The values returned are RGB values
// ordered in XZ order, and a byte slice of the heights, ordered in the same order.
// If not successful, an error is returned.
func ParseChunkData(data string) (colourMap []color.RGBA, heightMap []byte, err error) {
	// 16x16 blocks on a chunk layer.
	values := make([]string, 0, 256)

	// Chunk data is a string joined by commas...
	fragments := strings.Split(data, ",")
	for _, fragment := range fragments {
		parts := strings.Split(fragment, "*")
		multiplier := 1
		if len(parts) == 2 {
			var err error
			multiplier, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			// The multiplier given needs another added.
			multiplier++
		}
		value := parts[0]
		if pointer, err := strconv.Atoi(value); err == nil {
			// The value was a number, AKA pointing to an earlier value found.
			value = values[pointer]
		}

		for i := 0; i < multiplier; i++ {
			values = append(values, value)
		}
	}
	// Create RGB and height slices for each block.
	rgba := make([]color.RGBA, 0, 256)
	heights := make([]byte, 0, 256)

	for _, str := range values {
		slice, err := base64.RawStdEncoding.DecodeString(str)
		if err != nil {
			return nil, nil, err
		}
		// Little endian ordered RGB.
		rgba = append(rgba, color.RGBA{B: slice[0], G: slice[1], R: slice[2]})
		heights = append(heights, slice[3])
	}
	return rgba, heights, nil
}
