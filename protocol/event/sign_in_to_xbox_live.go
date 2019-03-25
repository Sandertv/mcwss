package event

// SignInToXBOXLive is sent by the client when it signs into XBOX Live by clicking the button in the Minecraft
// app.
type SignInToXBOXLive struct {
	// SignInUI ...
	SignInUI bool
	// Stage ...
	Stage int
	// Timestamp is the timestamp at which the user was logged into XBOX Live.
	Timestamp float64
}
