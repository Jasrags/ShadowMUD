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

	CategoryCombat       Category = "Combat"
	CategoryDetection    Category = "Detection"
	CategoryHealth       Category = "Health"
	CategoryIllusion     Category = "Illusion"
	CategoryManipulation Category = "Manipulation"
	CategoryEnchantment  Category = "Enchantment"
	CategoryRitual       Category = "Ritual"

	DurationInstantaneous Duration = "instantaneous"
	DurationSustained     Duration = "sustained"
	DurationPermanent     Duration = "permanent"

	DescriptorIndirect      Descriptor = "Indirect"
	DescriptorDirect        Descriptor = "Direct"
	DescriptorElemental     Descriptor = "Elemental"
	DescriptorEssence       Descriptor = "Essence"
	DescriptorRealistic     Descriptor = "Realistic"
	DescriptorSingleSense   Descriptor = "Single-Sense"
	DescriptorMultiSense    Descriptor = "Multi-Sense"
	DescriptorObvious       Descriptor = "Obvious"
	DescriptorPhysical      Descriptor = "Physical"
	DescriptorMental        Descriptor = "Mental"
	DescriptorDamaging      Descriptor = "Damaging"
	DescriptorPsychic       Descriptor = "Psychic"
	DescriptorEnvironmental Descriptor = "Environmental"
	DescriptorMaterialLink  Descriptor = "Material Link"
	DescriptorSpell         Descriptor = "Spell"
	DescriptorSpotter       Descriptor = "Spotter"
	DescriptorAnchored      Descriptor = "Anchored"
	DescriptorMinion        Descriptor = "Minion"
	DescriptorOrganicLink   Descriptor = "Organic Link"
	DescriptorGeomancy      Descriptor = "Geomancy"
	DescriptorContractual   Descriptor = "Contractual"
	DescriptorArea          Descriptor = "Area"
	DescriptorBlood         Descriptor = "Blood"
	DescriptorExtendedArea  Descriptor = "Extended Area"
	DescriptorActive        Descriptor = "Active"
	DescriptorPassive       Descriptor = "Passive"
	DescriptorDirectional   Descriptor = "Directional"
)

type (
	Type       string
	DamageType string
	Range      string
	Category   string
	Duration   string
	Descriptor string
	Spells     map[string]*Spell
	Spell      struct {
		ID          string
		Name        string
		Description string
		Duration    Duration
		Drain       int
		DamageValue int
		DamageType  shared.DamageType
		Descriptors []Descriptor
		Type        Type
		Range       Range
		Category    Category
	}
)
