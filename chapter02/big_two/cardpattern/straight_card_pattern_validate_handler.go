package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type StraightCardPatternValidateHandler struct {
	next ICardPatternValidateHandler
}

var _ ICardPatternValidateHandler = (*StraightCardPatternValidateHandler)(nil)

func NewStraightCardPatternValidateHandler(next ICardPatternValidateHandler) *StraightCardPatternValidateHandler {
	return &StraightCardPatternValidateHandler{next: next}
}

func (h StraightCardPatternValidateHandler) Validate(cards []card.Card) (bool, ICardPattern) {
	ret := NewStraightCardPattern(cards)
	if ret != nil {
		return true, ret
	}
	if h.next != nil {
		return h.next.Validate(cards)
	}
	return false, nil
}
