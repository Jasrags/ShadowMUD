package core_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"
)

func TestSaveCoreWeaponMelee(t *testing.T) {
	for _, v := range core.CoreWeaponMelee {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.WeaponMeleeFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
