package protocol

// ErrorResponse is sent by the client when an error occurs during the process of communicating.
type ErrorResponse struct {
	// StatusMessage is a string explaining the error that occurred.
	StatusMessage string `json:"statusMessage"`
	// StatusCode is the type of the error that occurred.
	StatusCode int `json:"statusCode"`
}
