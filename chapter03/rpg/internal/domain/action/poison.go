package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/state"
)

type Poison struct {
	*BasedSkill
}

var _ domain.Action = (*Poison)(nil)

func NewPoison() *Poison {
	return &Poison{
		BasedSkill: NewBasedSkill("下毒", 80, TARGET_TYPE_ALL_ENEMY, 1, 0),
	}
}

func (p *Poison) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(p.manaCost)
	targets[0].RetrieveState(state.NewPoisonedState())
}
