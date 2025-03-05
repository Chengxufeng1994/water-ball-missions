package main

type Player interface {
	ID() string
	Name() string
	Naming()
	Hand() *Hand
	SetHand(*Hand)
	SetExchangeable(bool)
	SetExchangeTurn(int)
	CheckHandSize() int
	DrawCardIntoHand(Card)
	SelectCardForShow() Card
	GainPoint()
	CheckPoint() int
	IsExchangeable() bool
	MakeDecisionForExchange() bool
	ExchangeHand(opponent Player, turn int) ExchangeRecord
	ExchangeHandReset(opponent Player)
}
