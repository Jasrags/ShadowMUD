package common_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"
)

func TestSaveCoreWeaponMelee(t *testing.T) {
	for _, v := range common.CoreWeaponMelee {
		filename := fmt.Sprintf("../%s/%s.yaml", common.MeleeWeaponsFilepath, v.ID)
		if err := utils.SaveStructToYAML(filename, &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
