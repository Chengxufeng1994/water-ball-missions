package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type Summon struct {
	*BasedSkill
}

var _ domain.Action = (*Summon)(nil)

func NewSummon() *Summon {
	return &Summon{
		BasedSkill: NewBasedSkill("召喚", 150, TARGET_TYPE_SELF, -1, 0),
	}
}

func (s *Summon) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(s.manaCost)
	attacker.Summon()
}
