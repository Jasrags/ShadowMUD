package core_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/core"
	"github.com/Jasrags/ShadowMUD/core/util"
)

func TestSaveCoreRooms(t *testing.T) {
	for _, v := range core.CoreRooms {
		if err := util.SaveStructToYAML(fmt.Sprintf("../"+core.RoomFilename, v.Zone, v.ID), &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
