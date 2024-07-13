package character

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRollInitiative(t *testing.T) {
	c := &Character{}
	got := c.RollInitiative()
	want := 0

	if got != want {
		t.Errorf("RollInitiative() = %d, want %d", got, want)
	}
}
func TestRollDice(t *testing.T) {
	// Test case 1: Rolling 0 dice should return 0 total and an empty slice
	gotTotal, gotDiceRolled := rollDice(0)
	wantTotal := 0
	wantDiceRolled := []int{}

	fmt.Printf("gotTotal: %d, gotDiceRolled: %v\n", gotTotal, gotDiceRolled)
	if gotTotal != wantTotal || !reflect.DeepEqual(gotDiceRolled, wantDiceRolled) {
		t.Errorf("rollDice(0) = (%d, %v), want (%d, %v)", gotTotal, gotDiceRolled, wantTotal, wantDiceRolled)
	}

	// Test case 2: Rolling 1 dice should return a total between 1 and 6 (inclusive) and a slice with a single element
	gotTotal, gotDiceRolled = rollDice(1)
	fmt.Printf("gotTotal: %d, gotDiceRolled: %v\n", gotTotal, gotDiceRolled)
	if gotTotal < 1 || gotTotal > 6 || len(gotDiceRolled) != 1 {
		t.Errorf("rollDice(1) = (%d, %v), want a total between 1 and 6 (inclusive) and a slice with a single element", gotTotal, gotDiceRolled)
	}

	// Test case 3: Rolling 5 dice should return a total between 5 and 30 (inclusive) and a slice with 5 elements
	gotTotal, gotDiceRolled = rollDice(5)
	fmt.Printf("gotTotal: %d, gotDiceRolled: %v\n", gotTotal, gotDiceRolled)
	if gotTotal < 5 || gotTotal > 30 || len(gotDiceRolled) != 5 {
		t.Errorf("rollDice(5) = (%d, %v), want a total between 5 and 30 (inclusive) and a slice with 5 elements", gotTotal, gotDiceRolled)
	}
}

func TestGetInitiative(t *testing.T) {
	dt := []struct {
		Reaction  int
		Intuition int
		want      int
	}{
		{5, 7, 12},
		{5, 5, 10},
	}

	for _, tt := range dt {
		c := &Character{
			Reaction:  tt.Reaction,
			Intuition: tt.Intuition,
		}

		got := c.GetInitiative()

		if got != tt.want {
			t.Errorf("GetInitiative() = %d, want %d", got, tt.want)
		}
	}
}

func TestGetPhysicalLimit(t *testing.T) {
	dt := []struct {
		Strength int
		Body     int
		Reaction int
		want     int
	}{
		{5, 7, 6, 8},
		{5, 7, 5, 8},
	}

	for _, tt := range dt {
		c := &Character{
			Strength: tt.Strength,
			Body:     tt.Body,
			Reaction: tt.Reaction,
		}

		got := c.GetPhysicalLimit()

		if got != tt.want {
			t.Errorf("GetPhysicalLimit() = %d, want %d", got, tt.want)
		}
	}
}

func TestGetMentalLimit(t *testing.T) {
	dt := []struct {
		Logic     int
		Intuition int
		Willpower int
		want      int
	}{
		{5, 7, 6, 8},
		{5, 7, 5, 8},
	}

	for _, tt := range dt {
		c := &Character{
			Logic:     tt.Logic,
			Intuition: tt.Intuition,
			Willpower: tt.Willpower,
		}

		got := c.GetMentalLimit()

		if got != tt.want {
			t.Errorf("GetMentalLimit() = %d, want %d", got, tt.want)
		}
	}
}

func TestGetSocialLimit(t *testing.T) {
	dt := []struct {
		Charisma  int
		Willpower int
		Essence   float64
		want      int
	}{
		{5, 7, 6, 8},
		{5, 7, 3.4, 7},
	}

	for _, tt := range dt {
		c := &Character{
			Charisma:  tt.Charisma,
			Willpower: tt.Willpower,
			Essence:   tt.Essence,
		}

		got := c.GetSocialLimit()

		if got != tt.want {
			t.Errorf("GetSocialLimit() = %d, want %d", got, tt.want)
		}
	}
}
