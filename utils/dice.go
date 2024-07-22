package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func RollDice(numDice int) (int, []int) {
	rand.Seed(uint64(time.Now().UnixNano()))

	diceRolled := make([]int, numDice)
	total := 0
	for i := 0; i < numDice; i++ {
		roll := rand.Intn(6) + 1
		total += roll
		diceRolled[i] = roll
	}

	return total, diceRolled
}
