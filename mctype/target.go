package mctype

// Target is a target indicator in a command. The target may resolve to multiple entities rather than one.
type Target string

const (
	AllPlayers    Target = "@a"
	AllEntities   Target = "@e"
	NearestPlayer Target = "@p"
	RandomPlayer  Target = "@r"
	Self          Target = "@s"
)
