package core_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"

	"testing"
)

func TestSaveCoreBioware(t *testing.T) {
	for _, v := range core.CoreBioware {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.BiowareFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
