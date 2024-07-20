package core_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"
)

func TestSaveCoreWeapon(t *testing.T) {
	for _, v := range core.CoreWeapons {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.WeaponFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
