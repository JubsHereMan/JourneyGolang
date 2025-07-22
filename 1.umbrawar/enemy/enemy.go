package enemy

type Enemy struct {
	Name string
	HP   int
}

func (e *Enemy) TakeDamage(dano int) {
	e.HP -= dano
}
