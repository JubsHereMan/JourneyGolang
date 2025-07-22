package summon

import (
	"fmt"
	"summonRyujin/mage"
	"sync"
	"time"
)

func SummonRyujin(){
	var wg sync.WaitGroup
	wg.Add(3)
	Azazel := mage.Mage{Name: "Azazel",Material: "Negative",Mana: 100}
	go wrath(&Azazel,&wg)
	go determination(&Azazel,&wg)
	go purpose(&Azazel,&wg)
	wg.Wait()
	fmt.Println("Watatsumi is among us")
}

func wrath(a *mage.Mage,wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Azazel's Wrath is going to be unleashed")
	a.Mana -=10
	
}


func determination(a *mage.Mage,wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(5* time.Second)
	fmt.Println("Azazel's determination can be felt...")
	a.Mana -=10
}

func purpose(a *mage.Mage, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(10* time.Second)
	fmt.Println("the eyes of Azazel start to glow in yellow")
	a.Mana -=30
}

