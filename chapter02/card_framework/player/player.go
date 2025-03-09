package player

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/card"
)

type IPlayer interface {
	Name() string
	NamingHimself()
	DrawCardIntoHand(card card.Card)
	Hand() []card.Card
	Show(compare card.Card) card.Card
	GainPoint()
	ShowPoint() int
}
