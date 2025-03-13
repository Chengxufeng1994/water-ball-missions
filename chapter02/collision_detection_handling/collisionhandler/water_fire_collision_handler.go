package collisionhandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/sprite"
)

type WaterFireCollisionHandler struct {
	next ICollisionHandler
}

var _ ICollisionHandler = (*WaterFireCollisionHandler)(nil)

func NewWaterFireCollisionHandler(next ICollisionHandler) *WaterFireCollisionHandler {
	return &WaterFireCollisionHandler{next: next}
}

func (w WaterFireCollisionHandler) Handle(source sprite.ISprite, target sprite.ISprite, world GenericeWorld) {
	if source.Type() == sprite.SpriteTypeWater && target.Type() == sprite.SpriteTypeFire {
		world.RemoveSprites(source, target)
	} else if w.next != nil {
		w.next.Handle(source, target, world)
	}
}
