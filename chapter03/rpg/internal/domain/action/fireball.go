package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type FireBall struct {
	name              string
	manaCost          int
	requiredOfTargets int
	damage            int
}

var _ domain.Action = (*FireBall)(nil)

func NewFireBall() *FireBall {
	return &FireBall{
		name:              "火球",
		manaCost:          50,
		requiredOfTargets: ALL,
		damage:            50,
	}
}

// RequiredOfTargets implements domain.Action.
func (f *FireBall) RequiredOfTargets() int {
	return f.requiredOfTargets
}

// MagicPointCost implements domain.Action.
func (f *FireBall) MagicPointCost() int {
	return f.manaCost
}

// Damage implements domain.Action.
func (f *FireBall) Damage() int {
	return f.damage
}

// Execute implements domain.Action.
func (f *FireBall) Execute(target domain.Unit) {
	target.OnDamage(f.damage)
}

func (f *FireBall) String() string {
	return f.name
}
