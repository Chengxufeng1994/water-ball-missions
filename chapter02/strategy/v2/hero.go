package main

type Hero struct {
	hp         int
	name       string
	attackType AttackType
}

func NewHero(name string, attackType AttackType) *Hero {
	return &Hero{
		hp:         500,
		name:       name,
		attackType: attackType,
	}
}

func (h Hero) Attack(attacked Hero) {
	h.attackType.Attack(h, attacked)
}

func (h *Hero) Damage(number int) {
	h.hp -= number
}

func (h *Hero) IsDead() bool {
	return h.hp <= 0
}
