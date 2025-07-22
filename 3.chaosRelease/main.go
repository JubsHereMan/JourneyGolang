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
	go spells.MaterialPassage(ch3)
	
	doneCount:=0

	for doneCount < 3{
	select {
	case msg2:= <- ch2:
		fmt.Println("runica",msg2)
			free(&AzazelIsFree)
		doneCount++
	case msg3:= <- ch3:
		fmt.Println("materia;",msg3)
			free(&AzazelIsFree)
		doneCount++
	}
	
}
}

func free(b *bool){
	*b= true
}