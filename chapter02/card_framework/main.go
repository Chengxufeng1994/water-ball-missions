package main

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/player"

func main() {
	players := []player.IPlayer{
		player.NewHuman(1, &player.ManualSelectCardStrategy{}),
		player.NewHuman(2, &player.ManualSelectCardStrategy{}),
		player.NewHuman(3, &player.ManualSelectCardStrategy{}),
		player.NewHuman(4, &player.ManualSelectCardStrategy{}),
	}
	game := NewGame(players, NewShowdownGame())
	// players := []player.IPlayer{
	// 	player.NewHuman(1, &player.UnoSelectCardColorPriorityStrategy{}),
	// 	player.NewHuman(2, &player.UnoSelectCardColorPriorityStrategy{}),
	// 	player.NewHuman(3, &player.UnoSelectCardColorPriorityStrategy{}),
	// 	player.NewHuman(4, &player.UnoSelectCardColorPriorityStrategy{}),
	// }
	// game := NewGame(players, NewUnoGame())
	game.Start()
}
