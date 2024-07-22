package common_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"

	"testing"
)

func TestSaveCoreWeaponModifications(t *testing.T) {
	for _, v := range common.CoreWeaponModifications {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.WeaponModificationFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
