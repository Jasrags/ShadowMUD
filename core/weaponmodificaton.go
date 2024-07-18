package core

const (
	WeaponModificationsDataPath  = "data/items/weapons/modifications"
	WeaponModificationFilename   = WeaponModificationsDataPath + "/%s.yaml"
	WeaponModificationMinVersion = "0.0.1"
)

type WeaponModification struct{}

var CoreWeaponModifications = []WeaponModification{}
