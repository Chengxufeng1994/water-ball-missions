package main

type Fireball struct{}

func (Fireball) Attack(attacker, attacked Hero) {
	for range 3 {
		attacked.Damage(50)
	}
}
