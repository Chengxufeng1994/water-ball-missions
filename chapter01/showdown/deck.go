package main

import "math/rand"

type Deck struct {
	cards []Card
}

func NewPokerDeck() *Deck {
	return &Deck{
		cards: []Card{
			{rank: Ace, suit: Spade},
			{rank: Two, suit: Spade},
			{rank: Three, suit: Spade},
			{rank: Four, suit: Spade},
			{rank: Five, suit: Spade},
			{rank: Six, suit: Spade},
			{rank: Seven, suit: Spade},
			{rank: Eight, suit: Spade},
			{rank: Nine, suit: Spade},
			{rank: Ten, suit: Spade},
			{rank: Jack, suit: Spade},
			{rank: Queen, suit: Spade},
			{rank: King, suit: Spade},
			{rank: Ace, suit: Heart},
			{rank: Two, suit: Heart},
			{rank: Three, suit: Heart},
			{rank: Four, suit: Heart},
			{rank: Five, suit: Heart},
			{rank: Six, suit: Heart},
			{rank: Seven, suit: Heart},
			{rank: Eight, suit: Heart},
			{rank: Nine, suit: Heart},
			{rank: Ten, suit: Heart},
			{rank: Jack, suit: Heart},
			{rank: Queen, suit: Heart},
			{rank: King, suit: Heart},
			{rank: Ace, suit: Diamond},
			{rank: Two, suit: Diamond},
			{rank: Three, suit: Diamond},
			{rank: Four, suit: Diamond},
			{rank: Five, suit: Diamond},
			{rank: Six, suit: Diamond},
			{rank: Seven, suit: Diamond},
			{rank: Eight, suit: Diamond},
			{rank: Nine, suit: Diamond},
			{rank: Ten, suit: Diamond},
			{rank: Jack, suit: Diamond},
			{rank: Queen, suit: Diamond},
			{rank: King, suit: Diamond},
			{rank: Ace, suit: Club},
			{rank: Two, suit: Club},
			{rank: Three, suit: Club},
			{rank: Four, suit: Club},
			{rank: Five, suit: Club},
			{rank: Six, suit: Club},
			{rank: Seven, suit: Club},
			{rank: Eight, suit: Club},
			{rank: Nine, suit: Club},
			{rank: Ten, suit: Club},
			{rank: Jack, suit: Club},
			{rank: Queen, suit: Club},
			{rank: King, suit: Club},
		},
	}
}

func (d *Deck) Shuffle() {
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) Draw() Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}
