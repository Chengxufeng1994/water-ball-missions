package action

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type BasedSkill struct {
	name               string
	manaCost           int
	requiredTargetType int
	requiredOfTargets  int
	amount             int
}

var _ domain.Action = (*BasedSkill)(nil)

func NewBasedSkill(name string, manaCost int, requiredTargetType int, requiredOfTargets int, damage int) *BasedSkill {
	return &BasedSkill{
		name:               name,
		manaCost:           manaCost,
		requiredTargetType: requiredTargetType,
		requiredOfTargets:  requiredOfTargets,
		amount:             damage,
	}
}

func (b *BasedSkill) Amount() int {
	return b.amount
}

func (b *BasedSkill) MagicPointCost() int {
	return b.manaCost
}

func (b *BasedSkill) RequiredTargetType() int {
	return b.requiredTargetType
}

func (b *BasedSkill) RequiredOfTargets() int {
	return b.requiredOfTargets
}

func (b *BasedSkill) Description(attacker domain.Unit, attacked []domain.Unit) string {
	if b.requiredOfTargets == TARGET_TYPE_SELF {
		return fmt.Sprintf("%v 使用了 %s。", attacker, b.name)
	}

	targets := ""
	for i, target := range attacked {
		targets += fmt.Sprintf("%s", target)
		if i < len(attacked)-1 {
			targets += ", "
		}
	}

	return fmt.Sprintf("%v 對 %v 使用了 %s。", attacker, targets, b.name)
}

func (b *BasedSkill) Execute(rpg *domain.RPG, attacker domain.Unit, targets []domain.Unit) {
	attacker.LoseMagicPoint(b.manaCost)
	for _, target := range targets {
		target.OnDamage(b.amount)
		fmt.Printf("%v 對 %v 造成 %d 點傷害。\n", attacker, target, b.amount)
	}
}

func (b *BasedSkill) String() string {
	return b.name
}
