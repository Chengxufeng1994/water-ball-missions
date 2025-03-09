package player

import (
	"fmt"
	"slices"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/card"
)

type BasedPlayer struct {
	ID       int
	name     string
	point    int
	hand     []card.Card
	strategy SelectCardStrategy
}

var _ IPlayer = (*BasedPlayer)(nil)

func NewBasedPlayer(id int, strategy SelectCardStrategy) *BasedPlayer {
	return &BasedPlayer{
		ID:       id,
		hand:     make([]card.Card, 0),
		point:    0,
		strategy: strategy,
	}
}

func (p *BasedPlayer) Name() string {
	return p.name
}

func (p *BasedPlayer) NamingHimself() {
	p.name = fmt.Sprintf("player_%d", p.ID)
}

func (p *BasedPlayer) DrawCardIntoHand(card card.Card) {
	p.hand = append(p.hand, card)
}

func (p *BasedPlayer) Show(compare card.Card) card.Card {
	if len(p.hand) == 0 {
		return nil
	}

	i := p.strategy.Select(compare, p.hand)
	if i < 0 {
		return nil
	}
	ret := p.hand[i]
	p.hand = slices.Delete(p.hand, i, i+1)
	return ret
}

func (p *BasedPlayer) Hand() []card.Card {
	return p.hand
}

func (p *BasedPlayer) GainPoint() {
	p.point++
}

func (p *BasedPlayer) ShowPoint() int {
	return p.point
}

type Human struct {
	*BasedPlayer
}

var _ IPlayer = (*Human)(nil)

func NewHuman(id int, strategy SelectCardStrategy) *Human {
	return &Human{
		BasedPlayer: NewBasedPlayer(id, strategy),
	}
}
