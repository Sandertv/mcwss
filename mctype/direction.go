package mctype

// Direction is a readable direction used in agent controlling commands.
type Direction string

const (
	Forward Direction = "forward"
	Back    Direction = "back"
	Right   Direction = "right"
	Left    Direction = "left"
	Up      Direction = "up"
	Down    Direction = "down"
)
