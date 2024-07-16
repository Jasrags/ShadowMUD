package core_test

import (
	"shadowrunmud/core"

	"testing"
)

func TestReloadData(t *testing.T) {
	am := core.ArmorMod{
		// Name:         "Radiation Shielding",
		Rating:   2,
		Capacity: 6, // Rating
		// Cost:         3000, // Rating x 200
		// Availability: 12,   // Rating x 2
		// RuleSource:   "SR5:R&G",
		// FileVersion:  "0.0.1",
	}
	want := 2

	got := am.GetCapacity(func(am *core.ArmorMod) int {
		return am.Rating
	})

	if got != want {
		t.Errorf("Test() = %d, want %d", got, want)
	}

}
