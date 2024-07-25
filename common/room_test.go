package common_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"
)

func TestSaveCoreRooms(t *testing.T) {
	for _, v := range common.CoreRooms {
		if err := utils.SaveStructToYAML(fmt.Sprintf("../"+common.RoomFilename, v.ZoneID, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
