package common_test

import (
	"testing"

	"github.com/Jasrags/ShadowMUD/common"
)

func TestSaveCoreContacts(t *testing.T) {
	if err := common.SaveCoreContacts(common.ContactsDataPath); err != nil {
		t.Errorf("Error saving core contacts: %s", err)
	}
}
