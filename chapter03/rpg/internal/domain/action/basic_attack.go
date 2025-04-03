package action

import "github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"

type BasicAttack struct {
	name   string
	damage int
}

var _ domain.Action = (*BasicAttack)(nil)

func NewBasicAttack(damage int) *BasicAttack {
	return &BasicAttack{name: "基本攻擊", damage: damage}
}

func (b *BasicAttack) RequiredOfTargets() int {
	return 1
}

func (b *BasicAttack) MagicPointCost() int {
	return 0
}

func (b *BasicAttack) Damage() int {
	return b.damage
}

func (b *BasicAttack) Execute(target domain.Unit) {
	target.OnDamage(b.damage)
}

func (b *BasicAttack) String() string {
	return b.name
}
