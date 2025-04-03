package unit

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/decisionstrategy"
)

type AI struct {
	*BasedUnit
}

var _ domain.Unit = (*AI)(nil)

func NewAIUnit(id int, name string, hp, mp, str int, actions ...domain.Action) *AI {
	return &AI{
		BasedUnit: NewBasedUnit(
			id, name, hp, mp, str, actions, decisionstrategy.NewSeedAIDecisionStrategy(),
		),
	}
}
