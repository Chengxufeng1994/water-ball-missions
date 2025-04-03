package state

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type CheerUpState struct {
	*BasedState
}

var _ domain.State = (*CheerUpState)(nil)

func NewCheerUpState() *CheerUpState {
	return &CheerUpState{
		BasedState: NewBasedState("受到鼓舞", 3),
	}
}

func (state *CheerUpState) PreTurn() {
	state.remainRound--
}

func (state *CheerUpState) PostTurn() {
	if state.remainRound == 0 {
		state.ExitState()
		state.unit.RetrieveState(NewNormalState())
	}
}

func (state *CheerUpState) BonusStrength() int {
	return 50
}
