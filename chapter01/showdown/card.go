package main

import "fmt"

type Card struct {
	rank Rank
	suit Suit
}

func (c Card) GreaterThan(other Card) bool {
	if c.rank.EqualTo(other.rank) {
		return c.suit.GreaterThan(other.suit)
	}

	return c.rank.GreaterThan(other.rank)
}

func (c Card) String() string {
	return fmt.Sprintf("[%s][%s]", c.suit.String(), c.rank.String())
}
