package domain

type Role interface {
	MapObject
	HP() int
	TakeTurn(game *AdventureGame)
	Move(x, y int)
	Attack(attacked Role)
	OnDamage(amount int)
	RetrieveState(State State)
}

type BasedRole struct {
	BasedMapObject
	game  *AdventureGame
	hp    int
	power int
	state State
}

var _ Role = (*BasedRole)(nil)

func NewBasedRole(game *AdventureGame, x, y int, symbol Symbol, hp, power int) BasedRole {
	return BasedRole{
		BasedMapObject: NewBasedMapObject(symbol, x, y),
		game:           game,
		hp:             hp,
		power:          power,
		state:          NewNormalState(),
	}
}

// HP implements Role.
func (b *BasedRole) HP() int {
	return b.hp
}

// TakeTurn implements Role.
func (b *BasedRole) TakeTurn(game *AdventureGame) {
}

// Attack implements Role.
func (b *BasedRole) Attack(attacked Role) {
	attacked.OnDamage(b.power)
}

// Move implements Role.
func (b *BasedRole) Move(x int, y int) {
	b.X = x
	b.Y = y
}

// OnDamage implements Role.
func (b *BasedRole) OnDamage(amount int) {
	b.hp -= amount

	if b.hp <= 0 {
		b.hp = 0
	}

	if b.hp == 0 {
		b.game.RemoveMapObject(b.X, b.Y)
	}
}

// RetrieveState implements Role.
func (b *BasedRole) RetrieveState(State State) {
	b.state = State
}
