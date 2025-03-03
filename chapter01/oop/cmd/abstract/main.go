package main

import "github.com/Chengxufeng1994/water-ball-missions/chapter01/oop/abstract"

func main() {
	player1 := abstract.NewHuman("player1")
	player2 := abstract.NewAI("player2")
	game := abstract.NewGame(player1, player2)
	game.Start()
}
