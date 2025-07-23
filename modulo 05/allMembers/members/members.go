package members

type Character struct {
	Name   string
	Class  string
	Weapon string
	Level  int
}

type Party struct {
	Members []Character
}

var Members = Party{
	Members: []Character{
		{Name: "Azazel", Class: "Mage", Weapon: "the negative", Level: 65},
		{Name: "Kaya", Class: "Sniper", Weapon: "Fal", Level: 35},
		{Name: "Marcos", Class: "Velocist", Weapon: "Katana", Level: 65},
	},
}
