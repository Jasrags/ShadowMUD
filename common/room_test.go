package common_test

import (
	"fmt"
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/Jasrags/ShadowMUD/utils"
)

// TODO: add zone folder before loading rooms
func TestSaveCoreRooms(t *testing.T) {
	for _, v := range common.CoreRooms {
		filename := fmt.Sprintf("../%s/%s.yaml", common.RoomsFilepath, v.ID)
		if err := utils.SaveStructToYAML(filename, &v); err != nil {
			t.Errorf("Error saving %s: %s", v.ID, err)
		}
	}
}
