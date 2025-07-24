package utils

import "math/rand"

func RollDice(l int) int {
	return rand.Intn(l)+1
}