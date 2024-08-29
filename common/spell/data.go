package spell

import "github.com/Jasrags/ShadowMUD/common/shared"

var CoreSpells = []Spell{
	{
		ID:          "acid_stream",
		Name:        "Acid Stream",
		Description: "A stream of acid that deals physical damage.",
		// Indirect, Elemental
		Duration:    DurationInstantaneous,
		Drain:       5,
		DamageValue: -3, // F-3
		DamageType:  shared.DamageTypePhysical,
		Type:        TypeDirect,
		Range:       RangeLOS,
		Category:    CategoryCombat,
	},
}
