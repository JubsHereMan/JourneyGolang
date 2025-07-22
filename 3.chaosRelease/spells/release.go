package spells

import (
	"fmt"
	"time"
)


func MaterialPassage(ch2 chan string) {
	count := 1000000000000
	for i := 0; i < count; i++ {
		delay := time.Duration((i + 1) * 100) * time.Millisecond
		<-time.After(delay)
		fmt.Println("material passage:", delay)
		ch2 <- "SAVED"
	}
}

func TimePassage(ch3 chan string) {
	count := 10
	for i := 0; i < count; i++ {
		delay := time.Duration((i + 1) * 100) * time.Millisecond
		<-time.After(delay)
		fmt.Println("time passage:", delay)
		ch3 <- "SAVED"
	}
}


func incrementar (n *int){
	*n= *n+1
}