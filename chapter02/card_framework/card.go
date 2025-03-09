package main

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
	return true
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
