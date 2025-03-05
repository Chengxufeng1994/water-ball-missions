package main

import (
	"fmt"
	"math/rand"
)

type Human struct {
	id             string
	name           string
	point          int
	exchangeable   bool
	exchangeTurn   int
	exchangePlayer Player
	hand           *Hand
}

var _ Player = (*Human)(nil)

func NewHuman(id string) *Human {
	return &Human{
		id:           id,
		hand:         NewHand(),
		exchangeable: true,
	}
}

func (h *Human) ID() string {
	return h.id
}

func (h *Human) Name() string {
	return h.name
}

func (h *Human) Naming() {
	h.name = "Human_" + h.id
}

func (h *Human) CheckHandSize() int {
	return h.hand.Size()
}
func (h *Human) DrawCardIntoHand(card Card) {
	h.hand.AddCard(card)
}

func (h *Human) SelectCardForShow() Card {
	if h.hand.IsEmpty() {
		return Card{}
	}

	idx := rand.Intn(h.hand.Size())
	return h.hand.SelectCard(idx)
}

func (h *Human) GainPoint() {
	h.point++
}

func (h *Human) CheckPoint() int {
	return h.point
}

func (h *Human) IsExchangeable() bool {
	return h.exchangeable
}

func (h *Human) MakeDecisionForExchange() bool {
	if !h.exchangeable {
		return false
	}

	// TOD: implement detail
	return true
}
func (h *Human) ExchangeHand(opponent Player, currentTurn int) ExchangeRecord {
	if !h.exchangeable || !opponent.IsExchangeable() {
		fmt.Printf("%s cannot exchange with %s\n", h.Name(), opponent.Name())
		return ExchangeRecord{}
	}

	if h.ID() == opponent.ID() {
		fmt.Printf("player %s cannot exchange with itself\n", h.Name())
		return ExchangeRecord{}
	}

	h.SetExchangeable(false)
	h.SetExchangeTurn(currentTurn)
	opponent.SetExchangeable(false)
	opponent.SetExchangeTurn(currentTurn)

	h.exchangePlayer = opponent
	tmp := opponent.Hand()
	opponent.SetHand(h.Hand())
	h.SetHand(tmp)
	fmt.Printf("player %s exchange with %s\n", h.Name(), opponent.Name())

	return ExchangeRecord{
		Round:      currentTurn,
		FromPlayer: h,
		ToPlayer:   opponent,
	}
}

func (h *Human) ExchangeHandReset(opponent Player) {
	h.exchangePlayer = nil
	tmp := opponent.Hand()
	opponent.SetHand(h.Hand())
	h.SetHand(tmp)
}

func (h *Human) Hand() *Hand {
	return h.hand
}

func (h *Human) SetHand(hand *Hand) {
	h.hand = hand
}

func (h *Human) SetExchangeable(b bool) {
	h.exchangeable = b
}

func (h *Human) SetExchangeTurn(turn int) {
	h.exchangeTurn = turn
}
