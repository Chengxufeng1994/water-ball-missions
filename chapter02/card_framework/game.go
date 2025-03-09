package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/card"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/deck"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/player"
)

type IGame interface {
	NewDeck() deck.IDeck
	InitializeCard() int

	PreGame(*Game)
	PostRound(players []player.IPlayer, tableCards *[]card.Card)
	TakeTurn(player player.IPlayer, topCard *card.Card, tableCards *[]card.Card, deck deck.IDeck) card.Card
	EndGame(players []player.IPlayer) bool

	DetermineWinner(players []player.IPlayer)
}

type Game struct {
	players    []player.IPlayer
	winner     player.IPlayer
	deck       deck.IDeck
	game       IGame
	topCard    card.Card
	tableCards []card.Card
}

func NewGame(players []player.IPlayer, game IGame) *Game {
	return &Game{
		players:    players,
		game:       game,
		topCard:    nil,
		tableCards: make([]card.Card, 0),
	}
}

func (g *Game) Start() {
	fmt.Println("========== Game start ==========")
	for _, player := range g.players {
		player.NamingHimself()
	}

	g.deck = g.game.NewDeck()
	g.deck.Shuffle()

	for range g.game.InitializeCard() {
		for _, player := range g.players {
			player.DrawCardIntoHand(g.deck.Draw())
		}
	}

	g.game.PreGame(g)
	gameover := false
	round := 1
	for !gameover {
		fmt.Printf("==================== round %d ====================\n", round)
		for _, player := range g.players {
			showCard := g.game.TakeTurn(player, &g.topCard, &g.tableCards, g.deck)
			gameover = g.game.EndGame(g.players)
			fmt.Println(player.Name(), "show card:", showCard)
			fmt.Println("top card:", g.topCard)
			fmt.Println("table cards:", g.tableCards)
		}

		g.game.PostRound(g.players, &g.tableCards)
		round++
	}

	g.game.DetermineWinner(g.players)
}
