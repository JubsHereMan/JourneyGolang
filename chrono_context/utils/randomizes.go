package utils

import "math/rand"

func DiceRoll(l int) int {
	return rand.Intn(l+1)
}