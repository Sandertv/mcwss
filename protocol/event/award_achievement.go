package event

const (
	AchievementRenewableEnergy = 40
)

// AwardAchievement is called when a player is awarded an achievement in game. Note that achievements are only
// awarded if cheats are disabled in the world.
type AwardAchievement struct {
	// AchievementID is the numerical ID of the achievement.
	AchievementID int
}
