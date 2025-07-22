package main

import(
	"fightForYourLife/fighters"
	"fightForYourLife/battle"
	
)

func main(){
Azazel:= fighters.Mage{	Name: "Azazel", HP: 5, Atack: map[string]int{"deynto":4, "blackhotton":4, "slingshot":5,},}
Protogen := fighters.Enemy{Name: "Protogen",Defense: 6,HP: 5,}

battle.Battle(&Azazel,&Protogen)
}