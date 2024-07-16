package util_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jasrags/ShadowMUD/core/util"
)

func TestRollDice(t *testing.T) {
	// Test case 1: Rolling 0 dice should return 0 total and an empty slice
	gotTotal, gotDiceRolled := util.RollDice(0)
	wantTotal := 0
	wantDiceRolled := []int{}

	fmt.Printf("gotTotal: %d, gotDiceRolled: %v\n", gotTotal, gotDiceRolled)
	if gotTotal != wantTotal || !reflect.DeepEqual(gotDiceRolled, wantDiceRolled) {
		t.Errorf("rollDice(0) = (%d, %v), want (%d, %v)", gotTotal, gotDiceRolled, wantTotal, wantDiceRolled)
	}

	// Test case 2: Rolling 1 dice should return a total between 1 and 6 (inclusive) and a slice with a single element
	gotTotal, gotDiceRolled = util.RollDice(1)
	fmt.Printf("gotTotal: %d, gotDiceRolled: %v\n", gotTotal, gotDiceRolled)
	if gotTotal < 1 || gotTotal > 6 || len(gotDiceRolled) != 1 {
		t.Errorf("rollDice(1) = (%d, %v), want a total between 1 and 6 (inclusive) and a slice with a single element", gotTotal, gotDiceRolled)
	}

	// Test case 3: Rolling 5 dice should return a total between 5 and 30 (inclusive) and a slice with 5 elements
	gotTotal, gotDiceRolled = util.RollDice(5)
	fmt.Printf("gotTotal: %d, gotDiceRolled: %v\n", gotTotal, gotDiceRolled)
	if gotTotal < 5 || gotTotal > 30 || len(gotDiceRolled) != 5 {
		t.Errorf("rollDice(5) = (%d, %v), want a total between 5 and 30 (inclusive) and a slice with 5 elements", gotTotal, gotDiceRolled)
	}
}
