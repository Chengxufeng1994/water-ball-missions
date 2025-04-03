package state

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type PetrochemicalState struct {
	*BasedState
}

var _ domain.State = (*PetrochemicalState)(nil)

func NewPetrochemicalState() *PetrochemicalState {
	return &PetrochemicalState{
		BasedState: NewBasedState("石化", 3),
	}
}

func (state *PetrochemicalState) EntryState() {
	state.unit.SetActionable(false)
}

func (state *PetrochemicalState) ExitState() {
	state.unit.SetActionable(true)
}

func (state *PetrochemicalState) PostTurn() {
	if state.remainRound == 0 {
		state.ExitState()
		state.unit.RetrieveState(NewNormalState())
	}
}
