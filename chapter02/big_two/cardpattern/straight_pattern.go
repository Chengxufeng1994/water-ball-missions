package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type StraightCardPattern struct {
	CardPattern
}

var _ ICardPattern = (*StraightCardPattern)(nil)

func NewStraightCardPattern(cards []card.Card) *StraightCardPattern {
	if !isStraight(cards) {
		return nil
	}
	return &StraightCardPattern{
		CardPattern{
			cardPatternType: CardPatternTypeStraight,
			cards:           cards,
		},
	}
}

func (p *StraightCardPattern) FindMaxCard() card.Card {
	var maxCard card.Card
	for _, card := range p.cards {
		if card.Rank > maxCard.Rank {
			maxCard = card
		}
	}
	return maxCard
}

func (p *StraightCardPattern) CompareTo(other ICardPattern) bool {
	if !p.EqualCardType(other) {
		return false
	}
	return p.FindMaxCard().Compare(other.FindMaxCard())
}

// isStraight 判斷是否為順子
func isStraight(cards []card.Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[i].Rank-cards[i-1].Rank != 1 {
			return false
		}
	}
	return true
}
