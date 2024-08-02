package common_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"

	"testing"
)

func TestSaveCoreWeaponModifications(t *testing.T) {
	for _, v := range common.CoreWeaponModifications {
		filename := fmt.Sprintf("../%s/%s.yaml", common.WeaponModificationsFilepath, v.ID)
		if err := utils.SaveStructToYAML(filename, &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
