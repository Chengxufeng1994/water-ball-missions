package domain

import (
	"fmt"
	"math/rand"
)

const (
	MONSTER_HP    = 1
	MONSTER_POWER = 50
)

type Monster struct {
	*BasedRole
}

var _ Role = (*Monster)(nil)

func NewMonster(game *AdventureGame, x, y int) *Monster {
	m := &Monster{
		BasedRole: NewBasedRole(
			game, x, y, SymbolMonster, MONSTER_HP, 50,
		),
	}

	m.turnBehavior = m

	return m
}

func (m *Monster) RoundAction() {
	if m.InAttackRange(m.game.Character) {
		m.attackAction(m.game.Character)
		return
	}

	m.moveAction()
}

func (m *Monster) moveAction() {
	directions := make([][2]int, 0)
	for _, direction := range m.directions {
		directions = append(directions, direction)
	}
	for {
		direction := directions[rand.Intn(len(directions))]
		newX, newY := m.X+direction[0], m.Y+direction[1]

		if err := m.game.Move(m, m.X, m.Y, newX, newY); err != nil {
			continue
		}
		m.Move(newX, newY)
		fmt.Printf("Monster from %d, %d moved to %d, %d\n", m.X, m.Y, newX, newY)
		break
	}
}

func (m *Monster) attackAction(attacked Role) {
	m.Attack(attacked)
	fmt.Println("Monster attack character")
}

func (m *Monster) InAttackRange(attacked Role) bool {
	targetX, targetY := attacked.Position()

	// 只允許水平或垂直攻擊
	if m.X == targetX {
		return abs(m.Y-targetY) <= m.attackRange
	}
	if m.Y == targetY {
		return abs(m.X-targetX) <= m.attackRange
	}

	return false
}

func (m *Monster) String() string {
	return "Monster"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
