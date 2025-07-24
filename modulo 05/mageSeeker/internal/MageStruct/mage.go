package magestruct

type Mage struct{
	Name string `json:"nome"`
	Material string `json:"material"`
	Level int `json:"level"`
}

var Mages = []Mage{
	    {"Azazel", "Negative", 7},
    	{"Kiriko", "Pure", 9},
    	{"Kaya", "Air", 4},
}