package utils

import "math/rand"


func DiceRoll(s int)int{
	return rand.Intn(s)
}