package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type FireBall struct {
	*BasedSkill
}

var _ domain.Action = (*FireBall)(nil)

func NewFireBall() *FireBall {
	return &FireBall{
		BasedSkill: NewBasedSkill("火球", 50, TARGET_TYPE_ALL_ENEMY, -1, 50),
	}
}
