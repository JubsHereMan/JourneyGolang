package utils

import "math/rand"


func rollDice(l int)int{
	return rand.Intn(l)+1
}