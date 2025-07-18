package domeconection

import (
	"fightForYourLife/fighters"
	"fightForYourLife/utils"
	"fmt"
	"time"
)

func SelfKnowledge(ch chan string) {
	fmt.Println("Em meio ao caos, panico e medo a Luz não se apaga... Azazel procura dentro de sí forças para vencer a batalha")
	time.Sleep(3 * time.Second)

	for {
		rolagem := utils.DiceRoll(20)
		fmt.Println("resultado do dado para auto conhecimento: ", rolagem)

		if rolagem >= 15 {
			fmt.Println("Azazel se lembra dos ensinamentos da arvore... existe vida em lugares que não pensamos")
			ch <- "ok"
			break
		} else {
			fmt.Println("Azazel se distrai por não estar 100'%' focado ")
			time.Sleep(3 * time.Second)
		}
	}
}

func Exalation(chIn, chOut chan string) {
	result := <-chIn
	if result == "ok" {
		time.Sleep(3 * time.Second)
		fmt.Println("Azazel está se sintonizando consigo mesmo... Protogen sente algo diferente nele...")

		for {
			rolagem := utils.DiceRoll(20)
			fmt.Println("resultado do dado para exalação : ", rolagem)

			if rolagem >= 15 {
				fmt.Println("Azazel começa a pensar sobre como o tempo funciona... pequenas distorções começam a acontecer... Protogem começa a sentir medo...")
				chOut <- "ok2"
				break
			} else {
				fmt.Println("Azazel se distrai por não estar 100'%' focado ")
				time.Sleep(3 * time.Second)
			}
		}
	}
}

func Material(chFinal chan string, done chan string) {
	result := <-chFinal
	if result == "ok2" {
		time.Sleep(3 * time.Second)
		fmt.Println("Azazel começa a parar o tempo... Protogem tenta finalizar com tudo")

		for {
			rolagem := utils.DiceRoll(20)
			fmt.Println("resultado do dado para materialização : ", rolagem)

			if rolagem >= 15 {
				fmt.Println("Azazel para o tempo.")
				done <- "acabou" 
				break
			} else {
				fmt.Println("Azazel se distrai por não estar 100'%' focado ")
				time.Sleep(3 * time.Second)
			}
		}
	}
}

func Domeconection(ch1, ch2, ch3, done chan string, a *fighters.Mage) {
	go SelfKnowledge(ch1)

	result1 := <-ch1
	if result1 == "ok" {
		go Exalation(ch1, ch2)
		ch1 <- result1

		go Material(ch3, done)

		result2 := <-ch2
		if result2 == "ok2" {
			ch3 <- result2
		}

		for {
			msg := <-done
			if msg == "acabou" {
				fmt.Println("Azazel adquiriu TIME STOP")
				a.Atack["time stop"] = 7
				break
			}
		}
	}

	fmt.Println("Processo de conexão com o Domo finalizado.")
}
