package player

import (
	"math/rand"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/card"
)

type SelectCardStrategy interface {
	Select(card card.Card, hand []card.Card) int
}

type RandomSelectCardStrategy struct{}

var _ SelectCardStrategy = RandomSelectCardStrategy{}

func (s RandomSelectCardStrategy) Select(card card.Card, hand []card.Card) int {
	i := rand.Intn(len(hand))
	return i
}

type ManualSelectCardStrategy struct {
}

var _ SelectCardStrategy = ManualSelectCardStrategy{}

func (s ManualSelectCardStrategy) Select(card card.Card, hand []card.Card) int {
	var i int
	// fmt.Scan(&i)
	i = rand.Intn(len(hand))
	return i
}

type UnoSelectCardColorPriorityStrategy struct{}

var _ SelectCardStrategy = (*UnoSelectCardColorPriorityStrategy)(nil)

func (strategy *UnoSelectCardColorPriorityStrategy) Select(top card.Card, hand []card.Card) int {
	c, _ := top.(*card.UnoCard)
	for i, item := range hand {
		hc, _ := item.(*card.UnoCard)
		if c.Color == hc.Color {
			return i
		}
	}

	for i, item := range hand {
		hc := item.(*card.UnoCard)
		if c.Number == hc.Number {
			return i
		}
	}

	return -1
}
