package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type SingleCardPattern struct {
	CardPattern
}

var _ ICardPattern = (*SingleCardPattern)(nil)

func NewSingleCardPattern(cards []card.Card) *SingleCardPattern {
	if len(cards) != 1 {
		return nil
	}
	return &SingleCardPattern{
		CardPattern{
			cardPatternType: CardPatternTypeSingle,
			cards:           cards,
		},
	}
}

func (p *SingleCardPattern) FindMaxCard() card.Card {
	return p.cards[0]
}

// 大小比較規則：先比數字再比花色。
func (p *SingleCardPattern) CompareTo(other ICardPattern) bool {
	if !p.EqualCardType(other) {
		return false
	}
	return p.FindMaxCard().Compare(other.FindMaxCard())
}
