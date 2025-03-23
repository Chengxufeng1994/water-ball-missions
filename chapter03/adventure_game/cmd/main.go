package main

import "github.com/Chengxufeng1994/water-ball-missions/chapter03/adventure_game/internal/domain"

const WIDTH = 10
const HEIGHT = 10
const MONSTER_COUNT = 3
const TREASURE_COUNT = 3
const OBSTACLE_COUNT = 5

func main() {
	adventureGame := domain.NewAdventureGame(
		WIDTH,
		HEIGHT,
		MONSTER_COUNT,
		TREASURE_COUNT,
		OBSTACLE_COUNT,
	)
	adventureGame.Start()
}
