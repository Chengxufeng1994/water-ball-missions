package cardpattern

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"
)

type PairCardPattern struct {
	CardPattern
}

var _ ICardPattern = (*PairCardPattern)(nil)

func NewPairCardPattern(cards []card.Card) *PairCardPattern {
	if len(cards) != 2 {
		return nil
	}
	if cards[0].Rank != cards[1].Rank {
		return nil
	}
	return &PairCardPattern{
		CardPattern{
			cardPatternType: CardPatternTypePair,
			cards:           cards,
		},
	}
}

func (p *PairCardPattern) FindMaxCard() card.Card {
	if p.cards[0].Suit > p.cards[1].Suit {
		return p.cards[0]
	}

	return p.cards[1]
}

// 大小比較規則：將兩張牌中較大的牌作為比較基準；例如：A-A > 7-7
func (p *PairCardPattern) CompareTo(other ICardPattern) bool {
	if !p.EqualCardType(other) {
		return false
	}

	return p.FindMaxCard().Compare(other.FindMaxCard())
}
