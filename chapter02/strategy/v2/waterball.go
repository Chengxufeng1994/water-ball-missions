package main

type Waterball struct{}

var _ AttackType = (*Waterball)(nil)

func (Waterball) Attack(attacker, attacked Hero) {
	amount := float64(attacker.hp) * 0.5
	attacked.Damage(int(amount))
}
