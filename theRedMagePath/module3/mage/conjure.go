package mage

import (
	"fmt"
	"module3/operations"
)

func Conjure(){
	a:= 2
	b:= 2

		funcs := []func(int, int) int{
		operations.Soma,
		operations.Subtrair,
		operations.Multiplicacao,
		operations.Divisao,
	}

	for _, f := range funcs {
		fmt.Println(f(a, b))
	}
}