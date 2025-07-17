package rituals

import (
	"fmt"
	"sync"
	"time"
)

func Conjure(spell string, rounds int, wg *sync.WaitGroup){
	defer wg.Done()
	for i:= 1; i <= rounds; i++{
		fmt.Println("conjurando", spell, "- round", i)
		time.Sleep(500 * time.Millisecond)
	}
}


func SummonAll(){
	var wg sync.WaitGroup


	spells:= []string{"denyto","slingshot","blackhotton"}

	wg.Add(len(spells))

	for _, spell:= range spells{
		go  Conjure(spell, 3, &wg)
	}

	wg.Wait()
	fmt.Println("Azazel revigora o ambiente com energia")
}