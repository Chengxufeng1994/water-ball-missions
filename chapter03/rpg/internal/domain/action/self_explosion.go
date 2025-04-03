package action

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type SelfExplosion struct {
	*BasedSkill
}

var _ domain.Action = (*SelfExplosion)(nil)

func NewSelfExplosion() *SelfExplosion {
	return &SelfExplosion{
		BasedSkill: NewBasedSkill("自爆", 200, TARGET_TYPE_ALL, -1, 150),
	}
}

func (s *SelfExplosion) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.Suicide()
	for _, target := range targets {
		target.OnDamage(s.amount)
		fmt.Printf("%v 對 %v 造成 %d 點傷害。\n", attacker, target, s.amount)
	}
}
