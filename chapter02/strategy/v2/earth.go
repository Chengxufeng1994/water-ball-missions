package main

type Earth struct{}

func (Earth) Attack(attacker, attacked Hero) {
	for range 10 {
		attacked.Damage(20)
	}
}
