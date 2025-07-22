package main

import (
	"doppleganger/spell"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Azazel se encontra em uma situação onde uma bomba assincrona está ameaçando destruir o distrito 7")
	time.Sleep(5 * time.Second)
	fmt.Println("Azazel conjura 5 clones para desacelerar a explosão. Fazendo algo assincrono tentar ser algo sincrono")
	var wg sync.WaitGroup
	y := 5
	spell.Doppleganger(y,&wg)

}