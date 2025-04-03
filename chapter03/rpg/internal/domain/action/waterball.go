package action

import "github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"

type WaterBall struct {
	name              string
	manaCost          int
	requiredOfTargets int
	damage            int
}

var _ domain.Action = (*WaterBall)(nil)

func NewWaterBall() *WaterBall {
	return &WaterBall{
		name:              "水球",
		manaCost:          50,
		requiredOfTargets: 1,
		damage:            120,
	}
}

func (w *WaterBall) RequiredOfTargets() int {
	return w.requiredOfTargets
}

func (w *WaterBall) MagicPointCost() int {
	return w.manaCost
}

func (w *WaterBall) Damage() int {
	return w.damage
}

func (w *WaterBall) Execute(target domain.Unit) {
	target.OnDamage(w.damage)
}

func (w *WaterBall) String() string {
	return w.name
}
