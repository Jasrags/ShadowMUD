package core_test

import (
	"testing"

	"github.com/Jasrags/ShadowMUD/core"
)

func TestSaveCoreContacts(t *testing.T) {
	if err := core.SaveCoreContacts(core.ContactsDataPath); err != nil {
		t.Errorf("Error saving core contacts: %s", err)
	}
}
