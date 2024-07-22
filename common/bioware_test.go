package common_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"

	"testing"
)

func TestSaveCoreBioware(t *testing.T) {
	for _, v := range common.CoreBioware {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.BiowareFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
