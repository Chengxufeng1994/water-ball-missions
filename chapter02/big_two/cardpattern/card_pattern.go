package cardpattern

import (
	"fmt"
	"sort"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"
)

type ICardPattern interface {
	Type() CardPatternType
	EqualCardType(other ICardPattern) bool
	CompareTo(other ICardPattern) bool
	FindMaxCard() card.Card
	ListCard() []card.Card
	String() string
}

type CardPattern struct {
	cardPatternType CardPatternType
	cards           []card.Card
}

var _ ICardPattern = (*CardPattern)(nil)

func NewCardPattern(cards []card.Card) ICardPattern {
	n := len(cards)
	rankCount := make(map[int]int)
	for _, card := range cards {
		rankCount[int(card.Rank)]++
	}
	switch n {
	case 1:
		return NewSingleCardPattern(cards)
	case 2:
		return NewPairCardPattern(cards)
	case 5:
		sort.Slice(cards, func(i, j int) bool {
			if cards[i].Rank == cards[j].Rank {
				return cards[i].Suit < cards[j].Suit
			}
			return cards[i].Rank < cards[j].Rank
		})
		ret := NewStraightCardPattern(cards)
		if ret != nil {
			return ret
		}
		ret = NewFullHouseCardPattern(cards)
		if ret != nil {
			return ret
		}
	}
	return nil
}

func (cp *CardPattern) ListCard() []card.Card {
	return cp.cards
}

func (cp *CardPattern) Type() CardPatternType {
	return cp.cardPatternType
}

func (cp *CardPattern) EqualCardType(other ICardPattern) bool {
	return cp.Type() == other.Type()
}

func (cp *CardPattern) FindMaxCard() card.Card {
	panic("unimplemented")
}

func (cp *CardPattern) CompareTo(other ICardPattern) bool {
	panic("unimplemented")
}

func (cp *CardPattern) String() string {
	return fmt.Sprintf("%s %v", cp.Type(), cp.cards)
}
