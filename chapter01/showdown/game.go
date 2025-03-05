package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"slices"
)

type Game struct {
	players         []Player
	deck            *Deck
	exchangeRecords []ExchangeRecord
}

func NewGame(players []Player, deck *Deck) *Game {
	return &Game{
		players:         players,
		deck:            deck,
		exchangeRecords: make([]ExchangeRecord, 0),
	}
}

func (g *Game) Start() {
	g.deck.Shuffle()
	log.Println("Deck shuffle")

	log.Println("All player naming start")
	for _, player := range g.players {
		// TODO: can duplication naming?
		player.Naming()
	}
	log.Println("All player naming end")

	ready := make(map[string]bool, len(g.players))
	var i int
	for i < 13 {
		for _, player := range g.players {
			if ready[player.Name()] {
				continue
			}
			player.DrawCardIntoHand(g.deck.Draw())
			if player.CheckHandSize() == 13 {
				ready[player.Name()] = true
			}
		}
		i++
	}
	log.Println("All player draw 13 cards")
	for _, player := range g.players {
		fmt.Println(player.Hand())
	}

	turn := 1
	for turn <= 13 {
		log.Printf("========== turn %d ==========", turn)

		// Take Turn
		showCards := make(map[Player]Card, len(g.players))
		for _, player := range g.players {
			if player.MakeDecisionForExchange() {
				exchangeRecord := player.ExchangeHand(g.players[0], turn)
				if !reflect.DeepEqual(exchangeRecord, ExchangeRecord{}) {
					g.exchangeRecords = append(g.exchangeRecords, exchangeRecord)
				}
			}

			g.CheckExchangeReset(turn)

			card := player.SelectCardForShow()
			showCards[player] = card
		}

		// Show Turn
		var turnWinner Player
		var prev Card
		for player, card := range showCards {
			log.Printf("%s show: %s", player.Name(), card.String())
			if reflect.DeepEqual(prev, Card{}) {
				prev = card
				turnWinner = player
				continue
			}
			if !prev.GreaterThan(card) {
				prev = card
				turnWinner = player
			}
		}

		log.Printf("This turn %d Winner: %s\n", turn, turnWinner.Name())
		turnWinner.GainPoint()
		log.Printf("=============================")
		turn++
	}

	g.announceWinner()
}

func (g *Game) announceWinner() {
	// Show Point
	max := math.MinInt32
	var winner Player
	for _, player := range g.players {
		log.Printf("%s point: %d", player.Name(), player.CheckPoint())
		if player.CheckPoint() > max {
			max = player.CheckPoint()
			winner = player
		}
	}
	// Show Winner
	log.Printf("Winner: %s", winner.Name())
}

func (g *Game) CheckExchangeReset(currentTurn int) {
	//
	for i := 0; i < len(g.exchangeRecords); i++ {
		elem := g.exchangeRecords[i]
		if currentTurn-elem.Round >= 3 {
			elem.FromPlayer.ExchangeHandReset(elem.ToPlayer)
			g.exchangeRecords = slices.Delete(g.exchangeRecords, i, i+1)
			i-- // adjust index after removal
		}
	}
}
