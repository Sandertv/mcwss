package event

const (
	_ = iota
	_
	InteractTame
	_
	InteractWithGolem
	InteractShave
	InteractMilk
	InteractTrade
	InteractFeed
	InteractIgnite
	InteractDye
	InteractNamed
	InteractLeash
	InteractUnleash
	_
	_
	InteractToggleSitting
)

// MobInteracted is sent when a player interacts with a mob in a way that activates something. Simply right
// clicking a mob seems not to send the event.
type MobInteracted struct {
	// InteractionType is the method that was used to interact with an entity.
	InteractionType int
	// MobType is the numerical type of an entity.
	MobType int
	// MobColor is the colour of the mob after a dye was used on it. This is present with sheep and dogs that
	// have their collar dyed.
	MobColor int
	// MobVariant is the variant of the mob. This is present with for example horses, that have different
	// markings and colours.
	MobVariant int
}
