package state

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type BasedState struct {
	name        string
	remainRound int
	unit        domain.Unit
}

var _ domain.State = (*BasedState)(nil)

func NewBasedState(name string, remainRound int) *BasedState {
	return &BasedState{
		name:        name,
		remainRound: remainRound,
	}
}

func (state *BasedState) EntryState() {
	// do nothing
}

func (state *BasedState) ExitState() {
	// do nothing
}

func (state *BasedState) BonusStrength() int {
	return 0
}

func (state *BasedState) PreTurn() {
	state.remainRound--
}

func (state *BasedState) PostTurn() {
	// do nothing
}

func (state *BasedState) SetUnit(unit domain.Unit) {
	state.unit = unit
}

func (state *BasedState) Equal(other domain.State) bool {
	return state.String() == fmt.Sprintf("%v", other)
}

func (state *BasedState) String() string {
	return state.name
}
