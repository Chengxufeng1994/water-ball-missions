package cardpattern

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"

type FullHouseCardPattern struct {
	CardPattern
}

var _ ICardPattern = (*FullHouseCardPattern)(nil)

func NewFullHouseCardPattern(cards []card.Card) *StraightCardPattern {
	rankCount := make(map[int]int)
	for _, card := range cards {
		rankCount[int(card.Rank)]++
	}
	if !isFullHouse(rankCount) {
		return nil
	}
	return &StraightCardPattern{
		CardPattern{
			cardPatternType: CardPatternTypeFullHouse,
			cards:           cards,
		},
	}
}

func (p *FullHouseCardPattern) FindMaxCard() card.Card {
	m := make(map[int]int)
	for _, card := range p.cards {
		m[int(card.Rank)]++
	}
	var maxCard card.Card
	for _, card := range p.cards {
		if m[int(card.Rank)] == 3 && card.Suit > maxCard.Suit {
			maxCard = card
		}
	}
	return maxCard
}

func (p *FullHouseCardPattern) CompareTo(other ICardPattern) bool {
	if !p.EqualCardType(other) {
		return false
	}
	return p.FindMaxCard().Compare(other.FindMaxCard())
}

// isFullHouse 判斷是否為葫蘆
func isFullHouse(rankCount map[int]int) bool {
	three := false
	pair := false
	for _, count := range rankCount {
		if count == 3 {
			three = true
		} else if count == 2 {
			pair = true
		}
	}
	return three && pair
}
