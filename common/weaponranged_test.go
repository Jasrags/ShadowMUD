package common_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"

	"github.com/stretchr/testify/assert"
)

func TestSaveCoreWeaponRanged(t *testing.T) {
	for _, v := range common.CoreWeaponRanged {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.WeaponRangedFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}

func TestToggleFiringMode(t *testing.T) {
	firingModes := []common.WeaponFiringMode{
		common.WeaponFiringModeSemiAutomatic,
		common.WeaponFiringModeBurstFire,
		common.WeaponFiringModeFullAuto,
	}
	dt := []struct {
		Name            string
		CurrentFireMode common.WeaponFiringMode
		FiringModes     []common.WeaponFiringMode
		Expected        string
	}{
		{
			Name:     "No Firing Modes",
			Expected: "[Ares Light Fire 70] No firing modes available",
		},
		{
			Name:            "Semi-Automatic to Burst Fire",
			CurrentFireMode: common.WeaponFiringModeSemiAutomatic,
			FiringModes:     firingModes,
			Expected:        "[Ares Light Fire 70] Firing mode changed to Burst Fire",
		},
		{
			Name:            "Burst Fire to Full Auto",
			CurrentFireMode: common.WeaponFiringModeBurstFire,
			FiringModes:     firingModes,
			Expected:        "[Ares Light Fire 70] Firing mode changed to Full Auto",
		},
		{
			Name:            "Full Auto to Semi-Automatic",
			CurrentFireMode: common.WeaponFiringModeFullAuto,
			FiringModes:     firingModes,
			Expected:        "[Ares Light Fire 70] Firing mode changed to Semi-Automatic",
		},
	}

	for _, tt := range dt {
		t.Run(tt.Name, func(t *testing.T) {
			w := &common.WeaponRanged{}
			w.Spec.Name = "Ares Light Fire 70"
			w.Spec.FiringModes = tt.FiringModes
			w.SelectedFiringMode = tt.CurrentFireMode

			result := w.ToggleFiringMode()
			assert.Equal(t, tt.Expected, result)
		})
	}
}

// func TestLoad(t *testing.T) {
// 	td := []struct {
// 		Name string
// 	}{
// 		{Name: "Test"},
// 	}

// 	for _, tt := range td {
// 		t.Run(tt.Name, func(t *testing.T) {
// 		})
// 	}
// 	w := &common.WeaponRanged{
// 		AmmoRemaining: 10,
// 		Spec: common.WeaponRangedSpec{
// 			AmmoCapacity: 20,
// 		},
// 	}

// 	result := w.Load()
// 	expected := "Ammo remaining: 10, Ammo capacity: 20"
// 	assert.Equal(t, expected, result)
// }

// func TestLoadEmptyAmmo(t *testing.T) {
// 	w := &common.WeaponRanged{
// 		AmmoRemaining: 0,
// 		Spec: common.WeaponRangedSpec{
// 			AmmoCapacity: 20,
// 		},
// 	}

// 	result := w.Load()
// 	expected := "Ammo remaining: 0, Ammo capacity: 20"
// 	assert.Equal(t, expected, result)
// }

// func TestLoadNoSpec(t *testing.T) {
// 	w := &common.WeaponRanged{
// 		AmmoRemaining: 10,
// 	}

// 	result := w.Load()
// 	expected := "Ammo remaining: 10, Ammo capacity: 0"
// 	assert.Equal(t, expected, result)
// }
