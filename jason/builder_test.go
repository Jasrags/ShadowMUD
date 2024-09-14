package jason_test

import (
	"testing"

	"github.com/Jasrags/ShadowMUD/jason"
	assert "github.com/stretchr/testify/assert"
)

func TestBuilder_SetName(t *testing.T) {
	cfg := &jason.DefaultConfig

	tests := []struct {
		name         string
		inputName    string
		expectedErr  string
		expectedName string
	}{
		{
			name:         "Valid name",
			inputName:    "ValidName",
			expectedErr:  "",
			expectedName: "ValidName",
		},
		{
			name:         "Name too short",
			inputName:    "ab",
			expectedErr:  "Name 'ab' must be between 3 and 16 characters",
			expectedName: "",
		},
		{
			name:         "Name too long",
			inputName:    "verylongnamethatistoolog",
			expectedErr:  "Name 'verylongnamethatistoolog' must be between 3 and 16 characters",
			expectedName: "",
		},
		{
			name:         "Name with invalid characters",
			inputName:    "invalidName_",
			expectedErr:  "Name 'invalidName_' must contain only letters, numbers, and underscores",
			expectedName: "",
		},
		{
			name:         "Banned name",
			inputName:    "admin",
			expectedErr:  "Name 'admin' is not allowed",
			expectedName: "",
		},
		{
			name:         "Name already taken",
			inputName:    "Jane",
			expectedErr:  "Name 'Jane' is not allowed",
			expectedName: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := jason.NewBuilder(cfg)
			err := b.SetName(test.inputName)

			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.expectedName, b.Name)
		})
	}
}
