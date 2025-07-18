package battle

import (
	domeconection "fightForYourLife/domeConection"
	"fightForYourLife/fighters"
	"fightForYourLife/utils"
	"fmt"
	"time"
)

func Battle(a *fighters.Mage, p *fighters.Enemy) {

	for {
		var result bool
		for a.HP > 0 {
			EnemyTurn(a)
			MageTurn(*a)
			_, existe := a.Atack["time stop"]
			if existe {
				fmt.Println("'AZAZEL:' O DRAGÃO ME ESCOLHEU POR UM MOTIVO.")
				vitoria(&result)
				break
			}
		}

		if result {
			fmt.Println("após uma batalha complexa e cheia de dúvidas, Azazel saiu vitorioso")
			break
		} else {
			fmt.Println("Azazel toma o ultimo golpe... infelizmente... Protogen ganhou...")
			break
		}
	}

}
func EnemyTurn(a *fighters.Mage) {
	fmt.Println("Protogen  utiliza um ataque manipulando a matéria para atingir Azazel...")
	time.Sleep(3 * time.Second)
	rolagem := utils.DiceRoll(20)
	fmt.Println("Rolagem de esquiva:", rolagem)
	if rolagem >= 10 {

		fmt.Println("Azazel consegue se esquivar por pouco...")
	} else {
		fmt.Println("A massa de matéria energizada acerta Azazel...")
		a.HP -= 1
		fmt.Println("vida Azazel:", a.HP)
	}
	time.Sleep(6 * time.Second)
}

func MageTurn(azazel fighters.Mage) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	done := make(chan string)
	go domeconection.Domeconection(ch1, ch2, ch3, done, &azazel)
	<-done
}

func vitoria(n *bool) {
	*n = true
}
