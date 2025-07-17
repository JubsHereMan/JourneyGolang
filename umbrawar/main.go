package main

import (
	"fmt"
	"umbrawar/enemy"
	"umbrawar/mage"
	"umbrawar/utils"
)

func main() {
	azazel := mage.Mage{
		Name:     "Azazel",
		Material: "negative",
		Mana:     100,
		Atacks: map[string]int{
			"Denyto":      30,
			"blackHotton": 50,
			"slingshot":   40,
		},
		HP: 100,
	}

	orc := enemy.Enemy{Name: "Orc of Grief", HP: 100}

	battle(&azazel, &orc)

}

func battle(m *mage.Mage, e *enemy.Enemy) {
	count := 0

	for m.HP > 0 && e.HP > 0 {
		play := utils.RollDice(20)
		damage := utils.RollDice(12)
		if count % 2 == 0 {
			fmt.Println(m.Name, "makes an atack, he roll the dices and the result is", play)
			if play >= 10 {
				fmt.Println("Azazel hits the orc and now will roll the demage")
				e.TakeDamage(damage)
				fmt.Println("the orc takes the demage  and now have", e.HP)
			} else {
				fmt.Println("azazel misses")
			}
		} else {
			fmt.Println(e.Name, "makes an atack, he roll the dices and the result is", play)
			if play >= 10 {
				fmt.Println(e.Name, "hits the orc and now will roll the demage")
				m.TakeDamage(damage)
				fmt.Println("Azazel takes the demage  and now have", m.HP)
			} else {
				fmt.Println(e.Name, "misses")
			}
		}

		count++
	}
}
