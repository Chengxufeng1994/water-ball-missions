package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	ID    int
	Name  string
	Cards []Card
}

func NewPlayer(id int) *Player {
	return &Player{
		ID:    id,
		Cards: make([]Card, 0),
	}
}

func (p *Player) NamingHimself() {
	p.Name = fmt.Sprintf("player_%d", p.ID)
}

func (p *Player) DrawCardIntoHand(card Card) {
	p.Cards = append(p.Cards, card)
}

func (p *Player) Show() Card {
	i := rand.Intn(len(p.Cards))
	ret := p.Cards[i]
	p.Cards = append(p.Cards[:i], p.Cards[i+1:]...)
	return ret
}
