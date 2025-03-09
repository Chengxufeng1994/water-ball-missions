package main

import "fmt"

type Game struct {
	players         []*Player
	deck            *Deck
	cardDrawingSize int
}

func NewBasedGame(players []*Player, deck *Deck, cardDrawingSize int) *Game {
	return &Game{
		players:         players,
		deck:            deck,
		cardDrawingSize: cardDrawingSize,
	}
}

func (b *Game) Initialize() {
	fmt.Println("========== Player Naming Phase ==========")
	b.PlayerNamingTurn()
	fmt.Println("========== Deck Shuffle Phase ==========")
	b.DeckShuffle()
	fmt.Println("========== Card Drawing Phase ==========")
	b.CardDrawingPhase(b.cardDrawingSize)
}

func (b *Game) PlayerNamingTurn() {
	for _, player := range b.players {
		player.NamingHimself()
	}
}

func (b *Game) DeckShuffle() {
	b.deck.Shuffle()
}

func (b *Game) CardDrawingPhase(size int) {
	for range size {
		for _, player := range b.players {
			// fmt.Printf("turn %d, %s draw a card\n", i+1, player.Name)
			player.DrawCardIntoHand(b.deck.Draw())
		}
	}
}

func (b *Game) StartGame() {}

type ShowdownGame struct {
	*Game
}

func NewShowdownGame(players []*Player) *ShowdownGame {
	game := NewBasedGame(players, NewShowdownDeck(), 13)
	return &ShowdownGame{
		Game: game,
	}
}

func (g *ShowdownGame) StartGame() {
	fmt.Println("========== Start Showdown Game ==========")
	turn := 0
	for turn < 13 {
		showCardMap := make(map[string]Card, len(g.Game.players))
		for _, player := range g.Game.players {
			showCardMap[player.Name] = player.Show()
		}

		for name, card := range showCardMap {
			fmt.Printf("turn %d, %s show a card %v\n", turn, name, card)
		}
		turn++
	}
}

type UnoGame struct {
	*Game
}

func NewUnoGame(players []*Player) *UnoGame {
	game := NewBasedGame(players, NewUnoDeck(), 5)
	return &UnoGame{Game: game}
}
