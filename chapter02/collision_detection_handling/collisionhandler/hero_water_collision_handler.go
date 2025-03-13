package collisionhandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/sprite"
)

type HeroWaterCollisionHandler struct {
	next ICollisionHandler
}

var _ ICollisionHandler = (*HeroWaterCollisionHandler)(nil)

func NewHeroWaterCollisionHandler(next ICollisionHandler) *HeroWaterCollisionHandler {
	return &HeroWaterCollisionHandler{next: next}
}

func (w HeroWaterCollisionHandler) Handle(source sprite.ISprite, target sprite.ISprite, world GenericeWorld) {
	if source.Type() == sprite.SpriteTypeHero && target.Type() == sprite.SpriteTypeWater {
		source.(*sprite.Hero).Recover(10)
		world.MoveSprite(source, target)
		world.RemoveSprite(target)
	} else if w.next != nil {
		w.next.Handle(source, target, world)
	}
}
