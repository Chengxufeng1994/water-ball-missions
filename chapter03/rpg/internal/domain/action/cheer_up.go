package action

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/state"
)

type CheerUp struct {
	*BasedSkill
}

var _ domain.Action = (*CheerUp)(nil)

func NewCheerUp() *CheerUp {
	return &CheerUp{
		BasedSkill: NewBasedSkill("鼓舞", 100, TARGET_TYPE_ALL_ALLY, 3, 0),
	}
}

func (c *CheerUp) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(c.manaCost)
	for _, target := range targets {
		target.RetrieveState(state.NewCheerUpState())
	}
	fmt.Printf("%v 對 %v 使用了 %s。\n", attacker, targets, c.name)
}
