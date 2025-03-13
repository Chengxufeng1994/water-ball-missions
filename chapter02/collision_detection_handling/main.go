package main

import (
	"math/rand"
	"time"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/collisionhandler"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/world"
)

const (
	WorldSize      = 30
	InitialSprites = 10
)

func main() {
	rand.Seed(time.Now().UnixNano())

	handler := collisionhandler.NewFireFireCollisionHandler(
		collisionhandler.NewWaterWaterCollisionHandler(
			collisionhandler.NewWaterFireCollisionHandler(
				collisionhandler.NewHeroFireCollisionHandler(
					collisionhandler.NewHeroWaterCollisionHandler(
						nil,
					),
				),
			),
		),
	)
	w := world.NewWorld(WorldSize, InitialSprites, handler)
	w.Prepare()
	w.Run()
}
