package core_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"

	"testing"
)

func TestSaveCoreWeaponModifications(t *testing.T) {
	for _, v := range core.CoreWeaponModifications {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.WeaponModificationFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
