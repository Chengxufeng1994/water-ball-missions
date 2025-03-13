package collisionhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/sprite"
)

type WaterWaterCollisionHandler struct {
	next ICollisionHandler
}

var _ ICollisionHandler = (*WaterWaterCollisionHandler)(nil)

func NewWaterWaterCollisionHandler(next ICollisionHandler) *WaterWaterCollisionHandler {
	return &WaterWaterCollisionHandler{next: next}
}

func (w WaterWaterCollisionHandler) Handle(source sprite.ISprite, target sprite.ISprite, world GenericeWorld) {
	if source.Type() == sprite.SpriteTypeWater && target.Type() == sprite.SpriteTypeWater {
		fmt.Printf("source %v, target %v, move failed\n", source, target)
	} else if w.next != nil {
		w.next.Handle(source, target, world)
	}
}
