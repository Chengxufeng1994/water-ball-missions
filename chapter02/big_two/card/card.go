package card

import "fmt"

type Card struct {
	Rank Rank
	Suit Suit
}

// Compare 比較兩張牌大小
func (c Card) Compare(other Card) bool {
	if c.Rank == other.Rank {
		return int(c.Suit) > int(other.Suit)
	}
	return int(c.Rank) > int(other.Rank)
}

func (c Card) String() string {
	return fmt.Sprintf("%s[%s]", c.Suit.String(), c.Rank.String())
}
