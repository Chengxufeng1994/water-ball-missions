package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/state"
)

type Petrochemical struct {
	*BasedSkill
}

var _ domain.Action = (*Petrochemical)(nil)

func NewPetrochemical() *Petrochemical {
	return &Petrochemical{
		BasedSkill: NewBasedSkill("石化", 100, TARGET_TYPE_ALL_ENEMY, 1, 0),
	}
}

func (p *Petrochemical) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(p.manaCost)
	targets[0].RetrieveState(state.NewPetrochemicalState())
}
