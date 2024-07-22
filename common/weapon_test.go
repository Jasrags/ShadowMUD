package common_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"
)

func TestSaveCoreWeapon(t *testing.T) {
	for _, v := range common.CoreWeapons {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.WeaponFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
