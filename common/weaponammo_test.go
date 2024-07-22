package common_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"
)

func TestSaveCoreWeaponAmunition(t *testing.T) {
	for _, v := range common.CoreWeaponAmunition {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.WeaponAmunitionFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
