package main

type AttackType interface {
	Attack(attacker, attacked Hero)
}
