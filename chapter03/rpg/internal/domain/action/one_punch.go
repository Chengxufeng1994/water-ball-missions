package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/onepunchhandler"
)

type OnePunch struct {
	*BasedSkill
	handler onepunchhandler.OnePunchHandler
}

func NewOnePunch(handler onepunchhandler.OnePunchHandler) *OnePunch {
	return &OnePunch{
		BasedSkill: NewBasedSkill("一拳攻擊", 180, TARGET_TYPE_ALL_ENEMY, 1, 0),
		handler:    handler,
	}
}

func (o *OnePunch) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(o.manaCost)
	target := targets[0]
	o.handler.Handle(attacker, target)
}
