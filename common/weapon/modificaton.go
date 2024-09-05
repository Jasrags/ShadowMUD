package weapon

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	WeaponModificationsFilepath = "_data/items/weapons/modifications"

	MountPointUnderBarrel MountPoint = "Under-Barrel"
	MountPointBarrel      MountPoint = "Barrel"
	MountPointStock       MountPoint = "Stock"
	MountPointTop         MountPoint = "Top"
	MountPointSide        MountPoint = "Side"
	MountPointInternal    MountPoint = "Internal"
)

type (
	MountPoint       string
	ModificationSpec struct {
		ID          string       `yaml:"id,omitempty"`
		Name        string       `yaml:"name,omitempty"`
		Description string       `yaml:"description,omitempty"`
		MountPoints []MountPoint `yaml:"mount_points"`
		// ArmorRating  int          `yaml:"armor_rating,omitempty"`
		Cost         int                 `yaml:"cost,omitempty"`
		Capacity     int                 `yaml:"capacity,omitempty"`
		Availability int                 `yaml:"availability,omitempty"`
		Legality     shared.LegalityType `yaml:"legality,omitempty"`
		ItemTags     []shared.ItemTag    `yaml:"tags"`
		// Modifiers    []Modifier          `yaml:"modifiers"`
		RuleSource shared.RuleSource `yaml:"rule_source,omitempty"`
	}
)
