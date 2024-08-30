package spell

import "github.com/Jasrags/ShadowMUD/common/shared"

const (
	TypeDirect   Type = "direct"
	TypeIndirect Type = "indirect"

	RangeLOSArea Range = "line-of-sight-area"
	RangeLOS     Range = "line-of-sight"
	RangeTouch   Range = "touch"
	RangeArea    Range = "area"
	RangeSpecial Range = "special"

	DamageTypePhysical DamageType = "physical"
	DamageTypeMana     DamageType = "mana"

	CategoryCombat       Category = "combat"
	CategoryDetection    Category = "detection"
	CategoryHealth       Category = "health"
	CategoryIllusion     Category = "illusion"
	CategoryManipulation Category = "manipulation"
	CategoryEnchantment  Category = "enchantment"
	CategoryRitual       Category = "ritual"

	DurationInstantaneous Duration = "instantaneous"
	DurationSustained     Duration = "sustained"
	DurationPermanent     Duration = "permanent"
)

type (
	Type       string
	DamageType string
	Range      string
	Category   string
	Duration   string
	Spells     map[string]*Spell
	Spell      struct {
		ID          string
		Name        string
		Description string
		Duration    Duration
		Drain       int
		DamageValue int
		DamageType  shared.DamageType
		Type        Type
		Range       Range
		Category    Category
	}
)
