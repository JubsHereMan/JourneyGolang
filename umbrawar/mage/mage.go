package mage

import "fmt"

type Mage struct {
	Name     string
	Material string
	Mana     int
	Atacks   map[string]int
	HP       int
}

func (m *Mage) UseSpell(spell string) {
	custo, ok := m.Atacks[spell]
	if !ok {
		fmt.Println(m.Name, "does not have the spell:", spell)
		return
	}

	if m.Mana < custo {
		fmt.Println(m.Name, "tries to cast", spell, "but lacks mana!")
		return
	}

	m.Mana -= custo
	fmt.Println(m.Name, "casts", spell, "— costing", custo, "mana! Remaining mana:", m.Mana)
}

func (m *Mage) TakeDamage(dano int) {
	m.HP -= dano
}
