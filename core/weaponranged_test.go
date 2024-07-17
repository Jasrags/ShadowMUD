package core_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"

	"github.com/stretchr/testify/assert"
)

func TestSaveCoreWeaponRanged(t *testing.T) {
	for _, v := range core.CoreWeaponRanged {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.WeaponRangedFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
func TestToggleFiringMode(t *testing.T) {
	dt := []struct {
		Name            string
		CurrentFireMode core.WeaponFiringMode
		FiringModes     []core.WeaponFiringMode
		Expected        string
	}{
		{
			Name:     "No Firing Modes",
			Expected: "[Ares Light Fire 70] No firing modes available",
		},
		{
			Name:            "Semi-Automatic to Burst Fire",
			CurrentFireMode: core.WeaponFiringModeSemiAutomatic,
			FiringModes: []core.WeaponFiringMode{
				core.WeaponFiringModeSemiAutomatic,
				core.WeaponFiringModeBurstFire,
				core.WeaponFiringModeFullAuto,
			},
			Expected: "[Ares Light Fire 70] Firing mode changed to Burst Fire",
		},
		{
			Name:            "Burst Fire to Full Auto",
			CurrentFireMode: core.WeaponFiringModeBurstFire,
			FiringModes: []core.WeaponFiringMode{
				core.WeaponFiringModeSemiAutomatic,
				core.WeaponFiringModeBurstFire,
				core.WeaponFiringModeFullAuto,
			},
			Expected: "[Ares Light Fire 70] Firing mode changed to Full Auto",
		},
		{
			Name:            "Full Auto to Semi-Automatic",
			CurrentFireMode: core.WeaponFiringModeFullAuto,
			FiringModes: []core.WeaponFiringMode{
				core.WeaponFiringModeSemiAutomatic,
				core.WeaponFiringModeBurstFire,
				core.WeaponFiringModeFullAuto,
			},
			Expected: "[Ares Light Fire 70] Firing mode changed to Semi-Automatic",
		},
	}

	for _, tt := range dt {
		t.Run(tt.Name, func(t *testing.T) {

			w := &core.WeaponRanged{
				Name:               "Ares Light Fire 70",
				FiringModes:        tt.FiringModes,
				SelectedFiringMode: tt.CurrentFireMode,
			}

			result := w.ToggleFiringMode()
			assert.Equal(t, tt.Expected, result)
		})
	}
}
