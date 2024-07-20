package core_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"
)

func TestSaveCoreZones(t *testing.T) {
	for _, v := range core.CoreZones {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.ZoneFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
