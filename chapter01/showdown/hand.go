package main

import "slices"

const CardLimit = 13

type Hand struct {
	limit int
	cards []Card
}

func NewHand() *Hand {
	return &Hand{
		limit: CardLimit,
		cards: make([]Card, 0),
	}
}

// TODO: error handling
func (h *Hand) AddCard(card Card) {
	if len(h.cards) >= h.limit {
		return
	}
	h.cards = append(h.cards, card)
}

func (h *Hand) Size() int {
	return len(h.cards)
}

func (h *Hand) SelectCard(i int) Card {
	if i >= len(h.cards) {
		return Card{}
	}

	old := h.cards
	ret := old[i]
	h.cards = slices.Delete(h.cards, i, i+1)
	return ret
}

func (h *Hand) IsEmpty() bool {
	return len(h.cards) == 0
}
