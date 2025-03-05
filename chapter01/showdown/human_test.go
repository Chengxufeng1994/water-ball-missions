package main

import (
	"fmt"
	"testing"
)

func TestHuman(t *testing.T) {
	h1 := NewHuman("human one")
	h2 := NewHuman("human two")
	h1.DrawCardIntoHand(Card{
		rank: Ace,
		suit: Spade,
	})
	h2.DrawCardIntoHand(Card{
		rank: Ace,
		suit: Diamond,
	})

	h1.ExchangeHand(h2, 0)
	fmt.Println(h1.hand)
	fmt.Println(h2.hand)
}
