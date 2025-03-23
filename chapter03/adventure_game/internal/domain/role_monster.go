package domain

const (
	MONSTER_HP    = 1
	MONSTER_POWER = 50
)

type Monster struct {
	BasedRole
}

var _ Role = (*Monster)(nil)

func NewMonster(game *AdventureGame, x, y int) *Monster {
	return &Monster{
		BasedRole: NewBasedRole(
			game, x, y, SymbolMonster, MONSTER_HP, 50,
		),
	}
}

func (m *Monster) TakeTurn(game *AdventureGame) {
	panic("unimplemented")
}

func (m *Monster) InAttackRange(attacked Role) bool {
	attackRange := 1
	targetX, targetY := attacked.Position()

	// 只允許水平或垂直攻擊
	if m.X == targetX {
		return abs(m.Y-targetY) <= attackRange
	}
	if m.Y == targetY {
		return abs(m.X-targetX) <= attackRange
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
