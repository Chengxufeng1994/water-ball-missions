package collisionhandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/sprite"
)

type GenericeWorld interface {
	RemoveSprite(source sprite.ISprite)
	RemoveSprites(sprites ...sprite.ISprite)
	MoveSprite(source sprite.ISprite, target sprite.ISprite)
}

type ICollisionHandler interface {
	Handle(source sprite.ISprite, target sprite.ISprite, world GenericeWorld)
}
