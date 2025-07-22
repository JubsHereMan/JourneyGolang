package spell

import (
	"doppleganger/mage"
	"fmt"
	"sync"
	"doppleganger/utils"
)


var mu sync.Mutex

func Doppleganger(c int,wg *sync.WaitGroup){
	Azazel := mage.Mage{Name: "Azazel",Material: "Negative",Mana: 100}
	
	clones := []string{"A1","A2","A3","A4","A5"}
for c > 0{
	for _, clone := range clones {
	wg.Add(1)
	cloneCopy := clone
	go castDenyto(&c, &Azazel, &cloneCopy, wg)
	}
	wg.Wait()

	mu.Lock()
	if Azazel.Mana <= 0{
		fmt.Println("a bomba assincrona detonou...")
		mu.Unlock()
		break
	}
	if c ==0{
		fmt.Println("Azazel conseguiu desarmar a bomba e salvar a cidade")
		fmt.Println("restou",Azazel.Mana,"de mana para Azazel")
		mu.Unlock()
		break
	}

}

	
}



func castDenyto(c *int, a *mage.Mage, d *string, wg *sync.WaitGroup){
	defer wg.Done()
	mu.Lock()
	a.Mana -=10
	fmt.Println("Azazel está tentando desarmar a bomba para um estado onde ela não tenha entrado em estado de explosão")
	 role :=utils.RollDice(20)
	 fmt.Println("rolagem de Azazel",role)
	 if role >=10{
		*c = *c-1
		fmt.Println("Azazel está conseguindo") 
	 }
	mu.Unlock()
}