package common_test

import (
	"fmt"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"

	"testing"
)

func TestSaveCoreSkillGroups(t *testing.T) {
	for _, v := range common.CoreSkillGroups {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.SkillGroupFilename, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
