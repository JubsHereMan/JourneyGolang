
package utils

import (
	"math/rand"
)

func RollDice(sides int) int {
	return rand.Intn(sides) + 1
}
