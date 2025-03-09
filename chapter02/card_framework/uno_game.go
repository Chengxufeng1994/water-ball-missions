package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/card"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/deck"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/player"
)

type UnoGame struct {
}

var _ IGame = (*UnoGame)(nil)

func NewUnoGame() *UnoGame {
	return &UnoGame{}
}

func (g *UnoGame) NewDeck() deck.IDeck {
	return deck.NewUnoDeck()
}

func (g *UnoGame) InitializeCard() int {
	return 5
}

func (g *UnoGame) PreGame(game *Game) {
	game.topCard = game.deck.Draw()
}

func (g *UnoGame) TakeTurn(player player.IPlayer, topCard *card.Card, tableCards *[]card.Card, deck deck.IDeck) card.Card {
	showCard := player.Show(*topCard)
	if showCard != nil {
		*tableCards = append(*tableCards, *topCard)
		*topCard = showCard
		return showCard
	}

	if deck.IsEmpty() {
		for _, tc := range *tableCards {
			deck.Add(tc)
		}
		deck.Shuffle()
		*tableCards = make([]card.Card, 0)
	}

	player.DrawCardIntoHand(deck.Draw())

	return g.TakeTurn(player, topCard, tableCards, deck)
}

func (g *UnoGame) PostRound(players []player.IPlayer, tableCards *[]card.Card) {
	// do nothing
}

func (g *UnoGame) EndGame(players []player.IPlayer) bool {
	for _, player := range players {
		if len(player.Hand()) == 0 {
			return true
		}
	}
	return false
}

func (g *UnoGame) DetermineWinner(players []player.IPlayer) {
	var winner player.IPlayer
	for _, player := range players {
		if len(player.Hand()) == 0 {
			winner = player
			break
		}
	}
	fmt.Printf("Winner: %s\n", winner.Name())
}
