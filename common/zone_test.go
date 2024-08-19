package common_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"
)

func TestSaveCoreZones(t *testing.T) {
	for _, v := range common.CoreZones {
		filename := fmt.Sprintf("../%s/%s.yaml", common.ZonesFilepath, v.ID)
		if err := utils.SaveStructToYAML(filename, v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
