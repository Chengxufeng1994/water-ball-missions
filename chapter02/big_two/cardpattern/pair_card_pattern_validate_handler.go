package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type PairCardPatternValidateHandler struct {
	next ICardPatternValidateHandler
}

var _ ICardPatternValidateHandler = (*PairCardPatternValidateHandler)(nil)

func NewPairCardPatternValidateHandler(next ICardPatternValidateHandler) *PairCardPatternValidateHandler {
	return &PairCardPatternValidateHandler{next: next}
}

func (h PairCardPatternValidateHandler) Validate(cards []card.Card) (bool, ICardPattern) {
	ret := NewPairCardPattern(cards)
	if ret != nil {
		return true, ret
	}

	if h.next != nil {
		return h.next.Validate(cards)
	}

	return false, nil
}
