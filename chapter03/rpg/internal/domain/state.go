package domain

type State interface {
	EntryState()
	ExitState()

	PreTurn()
	PostTurn()

	BonusStrength() int
	SetUnit(unit Unit)
	Equal(state State) bool
}
