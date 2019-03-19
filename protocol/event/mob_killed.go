package event

const (
	KillMeleeAttack  = 2
	KillShot         = 3
	KillExplosion    = 11
	KillSplashPotion = 14
)

// MobKilled is sent by the client when it kills a living entity.
type MobKilled struct {
	// ArmourFeetAuxType is the aux value/metadata of the boots item worn.
	ArmourFeetAuxType int `json:"ArmorFeetAuxType"`
	// ArmourFeetID is the numerical ID of the boots item worn.
	ArmourFeetID int `json:"ArmorFeetId"`
	// ArmourHeadAuxType is the aux value/metadata of the helmet item worn.
	ArmourHeadAuxType int `json:"ArmorHeadAuxType"`
	// ArmourHeadID is the numerical ID of the helmet item worn.
	ArmourHeadID int `json:"ArmorHeadId"`
	// ArmourLegsAuxType is the aux value/metadata of the leggings item worn.
	AmourLegsAuxType int `json:"ArmorLegsAuxType"`
	// ArmourLegsID is the numerical ID of the leggings item worn.
	ArmourLegsID int `json:"ArmorLegsId"`
	// ArmourTorsoAuxType is the aux value/metadata of the torso item worn.
	ArmourTorsoAuxType int `json:"ArmorTorsoAuxType"`
	// ArmourTorsoID is the numerical ID of the torso item worn.
	ArmourTorsoID int `json:"ArmorTorsoId"`
	// IsMonster indicates if the living entity killed was a monster. Note that this is not true for neutral
	// entities such as wolves that turn angry.
	IsMonster bool
	// KillMethodType is the method used to kill the entity. This is one of the methods above, but there are
	// more methods that are currently unknown.
	KillMethodType int
	// MobType is the numerical monster ID.
	MobType int
	// PlayerIsHiddenFrom indicates if the player was hidden from the entity killed.
	PlayerIsHiddenFrom bool
	// WeaponType is the numerical item ID of the item used to kill the entity. If no item was held, this type
	// is 0.
	WeaponType int
}
