package common_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"

	"testing"
)

func TestSaveCoreCyberwareModifications(t *testing.T) {
	for _, v := range common.CoreCyberware {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.CyberwareModificationFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
