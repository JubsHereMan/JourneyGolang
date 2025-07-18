package main

import (
	"chaosRelease/spells"
	"fmt"
)

func main(){
	AzazelIsFree:= false
	ch1:= make(chan string)
	ch2:= make(chan string)
	ch3:= make(chan string)

	go spells.TimePassage(ch1)
	go spells.RunicPassage(ch2)
	go spells.MaterialPassage(ch3)
	
	
	select {
	case msg1:= <-ch1:
		fmt.Println("temporal",msg1)
		free(&AzazelIsFree)
	case msg2:= <- ch2:
		fmt.Println("runica",msg2)
	case msg3:= <- ch3:
		fmt.Println("materia;",msg3)
	}
	
	
}

func free(b *bool){
	*b= true
}