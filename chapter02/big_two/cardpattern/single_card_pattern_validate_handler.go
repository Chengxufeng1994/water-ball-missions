package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type SingleCardPatternValidateHandler struct {
	next ICardPatternValidateHandler
}

var _ ICardPatternValidateHandler = (*SingleCardPatternValidateHandler)(nil)

func NewSingleCardPatternValidateHandler(next ICardPatternValidateHandler) *SingleCardPatternValidateHandler {
	return &SingleCardPatternValidateHandler{next: next}
}

func (h SingleCardPatternValidateHandler) Validate(cards []card.Card) (bool, ICardPattern) {
	ret := NewSingleCardPattern(cards)
	if ret != nil {
		return true, ret
	}
	if h.next != nil {
		return h.next.Validate(cards)
	}
	return false, nil
}
