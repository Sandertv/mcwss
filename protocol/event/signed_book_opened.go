package event

// SignedBookOpened is sent by the client when it opens a book that was signed.
type SignedBookOpened struct {
	// IsAuthor reports if the player that opened the book was the author of the book.
	IsAuthor bool
	// NetworkType no longer seems to be used. It is always 0.
	NetworkType int
}
