package common_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"

	"testing"
)

func TestSaveCoreArmor(t *testing.T) {
	for _, v := range common.CoreArmor {
		filename := fmt.Sprintf("../%s/%s.yaml", common.ArmorFilepath, v.ID)
		if err := utils.SaveStructToYAML(filename, &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
