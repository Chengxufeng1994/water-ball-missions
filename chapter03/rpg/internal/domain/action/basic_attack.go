package action

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type BasicAttack struct {
	*BasedSkill
}

var _ domain.Action = (*BasicAttack)(nil)

func NewBasicAttack() *BasicAttack {
	return &BasicAttack{
		BasedSkill: NewBasedSkill("基本攻擊", 0, TARGET_TYPE_ALL_ENEMY, 1, 0),
	}
}

func (b *BasicAttack) Description(attacker domain.Unit, targets []domain.Unit) string {
	return fmt.Sprintf("%v 攻擊 %v。", attacker, targets[0])
}

func (b *BasicAttack) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(b.manaCost)
	targets[0].OnDamage(attacker.GetStrength())
	fmt.Printf("%v 對 %v 造成 %d 點傷害。\n", attacker, targets[0], attacker.GetStrength())
}
