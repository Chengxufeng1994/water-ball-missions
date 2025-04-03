package unit

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/decisionstrategy"
)

type Hero struct {
	*BasedUnit
}

var _ domain.Unit = (*Hero)(nil)

func NewHeroUnit(id int, name string, hp, mp, str int, actions ...domain.Action) *Hero {
	return &Hero{
		BasedUnit: NewBasedUnit(
			id, name, hp, mp, str, actions, decisionstrategy.NewPlayerDecisionStrategy(),
		),
	}
}

func (hero *Hero) IsHero() bool {
	return true
}
