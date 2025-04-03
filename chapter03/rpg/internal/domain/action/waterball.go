package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type WaterBall struct {
	*BasedSkill
}

var _ domain.Action = (*WaterBall)(nil)

func NewWaterBall() *WaterBall {
	return &WaterBall{
		BasedSkill: NewBasedSkill("水球", 50, TARGET_TYPE_ALL_ENEMY, 1, 120),
	}
}
