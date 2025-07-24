package main

import (
	"fmt"
	"module4/internal"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go internal.ControlParadox(&wg)
	go internal.ControlSpace(&wg)
	go internal.ControlTime(&wg)
	wg.Wait()
	fmt.Println("a fenda temporal foi fechada")
}