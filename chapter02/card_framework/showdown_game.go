package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/card"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/deck"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/player"
)

type ShowdownGame struct {
	deck deck.IDeck
	turn int
}

var _ IGame = (*ShowdownGame)(nil)

func NewShowdownGame() *ShowdownGame {
	return &ShowdownGame{}
}

func (g *ShowdownGame) NewDeck() deck.IDeck {
	return deck.NewShowdownDeck()
}

func (g *ShowdownGame) InitializeCard() int {
	return 13
}

func (g *ShowdownGame) PreGame(game *Game) {
	// TODO: do nothing
}

func (g *ShowdownGame) TakeTurn(player player.IPlayer, topCard *card.Card, tableCards *[]card.Card, deck deck.IDeck) card.Card {
	card := player.Show(nil)
	*tableCards = append(*tableCards, card)
	return card
}

// func (g *ShowdownGame) PlayGame() {
// 	fmt.Println("========== Start Showdown Game ==========")
// 	type played struct {
// 		player player.IPlayer
// 		card   card.Card
// 	}
// 	showCardMap := make(map[string]played, len(g.players))
// 	for _, player := range g.players {
// 		card := player.Show(nil)
// 		showCardMap[player.Name()] = played{
// 			player: player,
// 			card:   card,
// 		}
// 	}

// 	var maximum card.Card
// 	var turnWinner player.IPlayer
// 	for name, played := range showCardMap {
// 		if maximum == nil || played.card.CompareTo(maximum) {
// 			maximum = played.card
// 			turnWinner = played.player
// 		}

//			fmt.Printf("turn %d, %s show a card %v\n", g.turn, name, played.card)
//		}
//		fmt.Println("turn", g.turn, ", maximum card is", maximum, "by", turnWinner.Name())
//		turnWinner.GainPoint()
//		g.turn++
//	}

func (g *ShowdownGame) PostRound(players []player.IPlayer, tableCards *[]card.Card) {
	var bigger card.Card
	var idx int
	for i, card := range *tableCards {
		if bigger == nil {
			bigger = card
			idx = i
			continue
		}

		if card.CompareTo(bigger) {
			bigger = card
			idx = i
		}
	}
	players[idx].GainPoint()
	*tableCards = make([]card.Card, 0)
}

func (g *ShowdownGame) EndGame(players []player.IPlayer) bool {
	for i := range players {
		if len(players[i].Hand()) != 0 {
			return false
		}
	}
	return true
}

func (g *ShowdownGame) DetermineWinner(players []player.IPlayer) {
	var winner player.IPlayer
	for _, player := range players {
		fmt.Printf("%s point: %d\n", player.Name(), player.ShowPoint())
		if winner == nil || player.ShowPoint() > winner.ShowPoint() {
			winner = player
		}
	}
	fmt.Printf("Winner: %s\n", winner.Name())
}
