package action

import "github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"

type SelfHealing struct {
	*BasedSkill
}

var _ domain.Action = (*SelfHealing)(nil)

func NewSelfHealing() *SelfHealing {
	return &SelfHealing{
		BasedSkill: NewBasedSkill("自我治療", 50, TARGET_TYPE_SELF, -1, 150),
	}
}

func (s *SelfHealing) Execute(rpg *domain.RPG, attacker domain.Unit, target []domain.Unit) {
	attacker.LoseMagicPoint(s.manaCost)
	attacker.OnHeal(s.amount)
}
