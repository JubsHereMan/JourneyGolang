package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Mage struct {
	Name    string
	Materia string
}

func main() {
	mage := Mage{Name: "Azazel"}
	fmt.Println("Ancião:Qual o seu nome mago?...")
	time.Sleep(3 * time.Second)
	fmt.Println("O mago olha para o ancião com determinação e diz: \n  ?:eu me chamo", mage.Name)
	time.Sleep(2 * time.Second)
	fmt.Println("Ancião:meu jovem siga-me para o templo para que os grandes elementos te escolham...")
	time.Sleep(2 * time.Second)
	fmt.Println("ambos começam a caminhar...")
	time.Sleep(2 * time.Second)
	fmt.Println("Azazel: como vou saber qual elemento me escolheu?")
	time.Sleep(2 * time.Second)
	fmt.Println("Ancião: Você sentira a conexão...")
	time.Sleep(2 * time.Second)
	fmt.Println("Azazel chega ao templo, e assim que coloca os pés dentro do salão uma energia começa a falar com seu amago...")
	roll:= rollDice(4)
		switch roll{
			case 1:
				fmt.Println("O ELEMENTO QUE ESCOLHEU AZAZEL FOI O FOGO")
				mage.Materia ="fogo"
			case 2: 
				fmt.Println("O ELEMENTO QUE ESCOLHEU AZAZEL FOI A ÁGUA")
				mage.Materia ="água"
			case 3: 
				fmt.Println("O ELEMENTO QUE ESCOLHEU AZAZEL FOI O AR")
				mage.Materia ="ar"
			case 4: 
				fmt.Println("O ELEMENTO QUE ESCOLHEU AZAZEL FOI A TERRA")
				mage.Materia ="terra"
		}
		castSpell(&mage)
	
}


func rollDice(l int)int{
	return rand.Intn(l)+1
}


func castSpell(m *Mage){
	switch m.Materia{
		case "fogo":
			FireConection(m)
		case "água":
			WaterConection(m)
		case "ar":
			AirConection(m)
		case "terra":
			EarthConection(m)
	}
}

func FireConection(m *Mage){
	fmt.Println("Fogo cresce nos olhos de" , m.Name)
}


func WaterConection(m *Mage){
	fmt.Println("Água flui pelo espirito de", m.Name)
}


func AirConection(m *Mage){
	fmt.Println("Ar rodeia",m.Name)
}

func EarthConection(m *Mage){
	fmt.Println("Terra treme sobre os pés de",m.Name)
}


