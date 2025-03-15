package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type FullHouseCardPatternValidateHandler struct {
	next ICardPatternValidateHandler
}

func NewFullHouseCardPatternValidateHandler(next ICardPatternValidateHandler) *FullHouseCardPatternValidateHandler {
	return &FullHouseCardPatternValidateHandler{next: next}
}

func (h FullHouseCardPatternValidateHandler) Validate(cards []card.Card) (bool, ICardPattern) {
	ret := NewFullHouseCardPattern(cards)
	if ret != nil {
		return true, ret
	}
	if h.next != nil {
		return h.next.Validate(cards)
	}
	return false, nil
}
