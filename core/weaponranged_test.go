package core_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"
)

func TestSaveCoreWeaponRanged(t *testing.T) {
	for _, v := range core.CoreWeaponRanged {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.WeaponRangedFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
