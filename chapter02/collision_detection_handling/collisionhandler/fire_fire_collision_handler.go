package collisionhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/sprite"
)

type FireFireCollisionHandler struct {
	next ICollisionHandler
}

var _ ICollisionHandler = (*FireFireCollisionHandler)(nil)

func NewFireFireCollisionHandler(next ICollisionHandler) *FireFireCollisionHandler {
	return &FireFireCollisionHandler{next: next}
}

func (h FireFireCollisionHandler) Handle(source sprite.ISprite, target sprite.ISprite, world GenericeWorld) {
	if source.Type() == sprite.SpriteTypeFire && target.Type() == sprite.SpriteTypeFire {
		fmt.Printf("source %v, target %v, move failed\n", source, target)
	} else if h.next != nil {
		h.next.Handle(source, target, world)
	}
}
