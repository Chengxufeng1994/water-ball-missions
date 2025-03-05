package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	player1 := NewHuman("player1")
	player2 := NewHuman("player2")
	player3 := NewHuman("player3")
	player4 := NewHuman("player4")
	deck := NewPokerDeck()

	game := NewGame(
		[]Player{player1, player2, player3, player4},
		deck,
	)

	game.Start()
}
