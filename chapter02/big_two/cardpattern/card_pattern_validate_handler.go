package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type ICardPatternValidateHandler interface {
	Validate(cards []card.Card) (bool, ICardPattern)
}
