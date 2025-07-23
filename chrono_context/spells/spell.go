package spells

import(
	"chrono_context/utils"
)

func TimeCircle(ch chan string) {
	roll :=utils.DiceRoll(7)
	if roll >=4{
		ch <- "o tempo"
	}
}

func SpaceCircle(ch chan string) {
	roll :=utils.DiceRoll(7)
	if roll >=4{
			ch <- "foi"
		}
}

func ConscienceCircle(ch chan string) {
	roll :=utils.DiceRoll(7)
	if roll >=4{
			ch <- "dominado"
		}
}
