package collisionhandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/sprite"
)

type HeroFireCollisionHandler struct {
	next ICollisionHandler
}

var _ ICollisionHandler = (*HeroFireCollisionHandler)(nil)

func NewHeroFireCollisionHandler(next ICollisionHandler) *HeroFireCollisionHandler {
	return &HeroFireCollisionHandler{
		next: next,
	}
}

func (w HeroFireCollisionHandler) Handle(source sprite.ISprite, target sprite.ISprite, world GenericeWorld) {
	if source.Type() == sprite.SpriteTypeHero && target.Type() == sprite.SpriteTypeFire {
		source.(*sprite.Hero).Damage(10)
		world.MoveSprite(source, target)
		world.RemoveSprite(target)
	} else if w.next != nil {
		w.next.Handle(source, target, world)
	}
}
