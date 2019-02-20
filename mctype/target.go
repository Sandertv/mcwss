package mctype

// Target is a target indicator in a command. The target may resolve to multiple entities rather than one.
type Target string

const (
	TargetAllPlayers    Target = "@a"
	TargetAllEntities   Target = "@e"
	TargetNearestPlayer Target = "@p"
	TargetRandomPlayer  Target = "@r"
	TargetSelf          Target = "@s"
)
