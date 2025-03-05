package main

type Hero struct {
	hp         int
	name       string
	attackType string
}

func NewHero(name, attackType string) *Hero {
	return &Hero{
		hp:         500,
		name:       name,
		attackType: attackType,
	}
}

func (h Hero) Attack(attacked Hero) {
	switch h.attackType {
	case "waterball":
		dmg := float64(h.hp) * 0.5
		attacked.Damage(int(dmg))
	case "fireball":
		for range 3 {
			attacked.Damage(50)
		}
	case "earth":
		for range 10 {
			attacked.Damage(20)
		}
	default:
	}
}

func (h *Hero) Damage(number int) {
	h.hp -= number
}

func (h *Hero) IsDead() bool {
	return h.hp <= 0
}
