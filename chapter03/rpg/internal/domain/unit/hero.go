package unit

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/decisionstrategy"
)

type Hero struct {
	*BasedUnit
}

var _ domain.Unit = (*Hero)(nil)

func NewHeroUnit(name string, hp, mp, str int, actions ...domain.Action) *Hero {
	return &Hero{
		BasedUnit: NewBasedUnit(
			name, hp, mp, str, actions, decisionstrategy.NewPlayerDecisionStrategy(),
		),
	}
}
