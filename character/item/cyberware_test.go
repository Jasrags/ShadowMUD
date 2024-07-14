package item

import (
	"testing"
)

func TestModifyReaction(t *testing.T) {
	dt := []struct {
		Name       string
		IsActive   bool
		Rating     int
		WantResult int
	}{
		{
			Name:       "Active Cyberware",
			IsActive:   true,
			Rating:     3,
			WantResult: 3,
		},
		{
			Name:       "Inactive Cyberware",
			IsActive:   false,
			Rating:     5,
			WantResult: 0,
		},
	}

	for _, tt := range dt {
		c := &Cyberware{
			Name:     tt.Name,
			IsActive: tt.IsActive,
			Rating:   tt.Rating,
		}

		got := c.ModifyReaction()

		if got != tt.WantResult {
			t.Errorf("ModifyReaction() = %d, want %d", got, tt.WantResult)
		}
	}
}
