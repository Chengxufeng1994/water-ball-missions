package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"
)

type Deck struct {
	cards []card.Card
}

func NewPokerDeck() *Deck {
	suits := []card.Suit{card.Club, card.Diamond, card.Heart, card.Spade}
	ranks := []card.Rank{card.Three, card.Four, card.Five, card.Six, card.Seven, card.Eight, card.Nine, card.Ten, card.Jack, card.Queen, card.King, card.Ace, card.Two}

	cards := make([]card.Card, len(suits)*len(ranks))
	for i, suit := range suits {
		for j, rank := range ranks {
			cards[i*len(ranks)+j] = card.Card{Rank: rank, Suit: suit}
		}
	}

	return &Deck{cards: cards}
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Draw() card.Card {
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
}

func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d Deck) String() string {
	str := make([]string, len(d.cards))
	for i := range d.cards {
		str[i] = fmt.Sprintf("%v", d.cards[i])
	}
	return strings.Join(str, " ")
}
