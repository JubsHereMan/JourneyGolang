package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)	
	fmt.Println("Azazel começa seu treinamento com sua matéria")
	time.Sleep(3 * time.Second)	
	fmt.Println("Azazel respira e se prepara para desacelerar o tempo")
	time.Sleep(3 * time.Second)
	roll:= RollDice(20)
	fmt.Println(roll)
	if roll >10 {
		fmt.Println("O tempo começa a desacelerar e Azazel se sente cada vez mais poderoso")
	}else{
		fmt.Println("O tempo não muda")
	}
	
}


func RollDice(l int)int{
	return rand.Intn(l+1)
}