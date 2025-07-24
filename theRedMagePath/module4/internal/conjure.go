package internal

import (
	"fmt"
	"module4/utils"
	"sync"
	"time"
)
func ControlTime(wg *sync.WaitGroup){
	defer wg.Done()
	timeToComplete := utils.RollDice(10)
	fmt.Println("Azazel cuidara da brecha temporal em", timeToComplete,"segundos")
	time.Sleep(time.Duration(timeToComplete)*(time.Second))
}


func ControlSpace(wg *sync.WaitGroup){
	defer wg.Done()
	timeToComplete := utils.RollDice(10)
	fmt.Println("Azazel cuidara da brecha espacial em", timeToComplete,"segundos")
	time.Sleep(time.Duration(timeToComplete)*(time.Second))
}

func ControlParadox(wg *sync.WaitGroup){
	defer wg.Done()
	timeToComplete := utils.RollDice(10)
	fmt.Println("Azazel cuidara da brecha paradoxal em", timeToComplete,"segundos")
	time.Sleep(time.Duration(timeToComplete)*(time.Second))
}