package card

import "fmt"

type Card interface {
	CompareTo(Card) bool
}

type ShowdownCard struct {
	Suit Suit
	Rank Rank
}

var _ Card = (*ShowdownCard)(nil)

func NewShowdownCard(suit Suit, rank Rank) *ShowdownCard {
	return &ShowdownCard{Suit: suit, Rank: rank}
}

func (c *ShowdownCard) CompareTo(other Card) bool {
	if c.Rank.EqualTo(other.(*ShowdownCard).Rank) {
		return c.Suit.GreaterThan(other.(*ShowdownCard).Suit)
	}
	return c.Rank.GreaterThan(other.(*ShowdownCard).Rank)
}

func (c *ShowdownCard) String() string {
	return fmt.Sprintf("[%s][%s]", c.Suit.String(), c.Rank.String())
}

type UnoCard struct {
	Color  Color
	Number int
}

var _ Card = (*UnoCard)(nil)

func NewUnoCard(color Color, number int) *UnoCard {
	return &UnoCard{Color: color, Number: number}
}

func (c *UnoCard) CompareTo(other Card) bool {
	return true
}

func (c *UnoCard) String() string {
	return fmt.Sprintf("[%s][%d]", c.Color.String(), c.Number)
}
