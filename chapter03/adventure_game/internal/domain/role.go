package domain

import (
	"math/rand"
)

var (
	VerticalDirections = map[string][2]int{
		"w": {0, -1},
		"s": {0, 1},
	}
	HorizontalDirections = map[string][2]int{
		"a": {-1, 0},
		"d": {1, 0},
	}
	Directions = map[string][2]int{
		"w": {0, -1},
		"s": {0, 1},
		"a": {-1, 0},
		"d": {1, 0},
	}
)

type RoundBehavior interface {
	RoundAction()
}

type AttackBehavior string

const (
	AttackBehaviorOnLine  AttackBehavior = "onLine"
	AttackBehaviorFullMap AttackBehavior = "fullMap"
)

type Role interface {
	MapObject
	HP() int
	IsFullHP() bool
	LoseHP(amount int)
	RecoverHP(amount int)
	UpdateDirections(directions map[string][2]int)
	Move(x, y int)
	RandomMove()
	Attack(attacked Role)
	UpdateAttackBehavior(behavior AttackBehavior)
	OnDamage(amount int)
	RetrieveState(State IState)
	UpdateInvincible(invincible bool)
	UpdateNumOfAction(numOfAction int)
	TakeTurn()
	RoundAction()
}

type BasedRole struct {
	BasedMapObject
	game         *AdventureGame
	turnBehavior RoundBehavior
	hp           int
	fullHP       int
	power        int
	numOfAction  int
	remainRound  int
	directions   map[string][2]int
	attack       AttackBehavior
	attackRange  int
	invincible   bool
	state        IState
}

var _ Role = (*BasedRole)(nil)

func NewBasedRole(game *AdventureGame, posX, posY int, roleSymbol Symbol, initialHP, attackPower int) *BasedRole {
	role := &BasedRole{
		BasedMapObject: NewBasedMapObject(roleSymbol, posX, posY),
		game:           game,
		hp:             initialHP,
		fullHP:         initialHP,
		power:          attackPower,
		numOfAction:    1,
		directions:     Directions,
		attack:         AttackBehaviorOnLine,
		attackRange:    1,
	}
	role.RetrieveState(NewNormalState())

	return role
}

// HP implements Role.
func (b *BasedRole) HP() int {
	return b.hp
}

func (b *BasedRole) LoseHP(amount int) {
	b.hp -= amount
}

func (b *BasedRole) RecoverHP(amount int) {
	if b.hp+amount > b.fullHP {
		b.hp = b.fullHP
		return
	}
	b.hp += amount
}

func (b *BasedRole) IsFullHP() bool {
	return b.hp == b.fullHP
}

func (b *BasedRole) UpdateDirections(directions map[string][2]int) {
	b.directions = directions
}

func (b *BasedRole) UpdateNumOfAction(numOfAction int) {
	b.numOfAction = numOfAction
}

func (b *BasedRole) TakeTurn() {
	b.state.PreRound()

	b.state.DeduceRound()

	for range b.numOfAction {
		if b.turnBehavior != nil {
			b.turnBehavior.RoundAction()
		} else {
			b.RoundAction()
		}
	}

	b.state.PostRound()
}

func (b *BasedRole) RoundAction() {
}

// Attack implements Role.
func (b *BasedRole) Attack(attacked Role) {
	attacked.OnDamage(b.power)
}

func (b *BasedRole) UpdateAttackBehavior(ab AttackBehavior) {
	switch ab {
	case AttackBehaviorOnLine:
		b.attack = AttackBehaviorOnLine
		b.power = MONSTER_HP
		b.attackRange = 1
	case AttackBehaviorFullMap:
		b.attack = AttackBehaviorFullMap
		b.power = 50
		b.attackRange = max(b.game.Height, b.game.Width)
	}
}

// Move implements Role.
func (b *BasedRole) Move(x int, y int) {
	b.X = x
	b.Y = y
}

// generate function the role can random move to empty position
func (b *BasedRole) RandomMove() {
	positions := b.game.FindEmptyPosition()
	position := positions[rand.Intn(len(positions))]
	srcX, srcY := b.X, b.Y
	dstX, dstY := srcX+position[0], srcY+position[1]
	b.game.Move(b, srcX, srcY, dstX, dstY)
	b.Move(position[0], position[1])
}

// OnDamage implements Role.
func (b *BasedRole) OnDamage(amount int) {
	b.state.OnDamage()

	if b.invincible {
		amount = 0
	}

	b.hp -= amount

	if b.hp <= 0 {
		b.hp = 0
	}

	if b.hp == 0 {
		b.game.RemoveMapObject(b.X, b.Y)
	}
}

func (b *BasedRole) UpdateInvincible(invincible bool) {
	b.invincible = invincible
}

func (b *BasedRole) RetrieveState(newState IState) {
	newState.SetRole(b)
	b.state = newState
	b.state.RetrieveState()
}
