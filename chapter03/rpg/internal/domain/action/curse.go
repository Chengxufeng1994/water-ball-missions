package action

import "github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"

type Curse struct {
	*BasedSkill
}

var _ domain.Action = (*Curse)(nil)

func NewCurse() *Curse {
	return &Curse{
		BasedSkill: NewBasedSkill("詛咒", 100, TARGET_TYPE_ALL_ENEMY, 1, 0),
	}
}

func (c *Curse) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(c.manaCost)
	for _, target := range targets {
		target.OnCurse(attacker)
	}
}
