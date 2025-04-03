package state

type PoisonedState struct {
	*BasedState
}

func NewPoisonedState() *PoisonedState {
	return &PoisonedState{
		BasedState: NewBasedState("中毒", 3),
	}
}

func (state *PoisonedState) PreTurn() {
	state.remainRound--
	state.unit.OnPoisoned(30)
}

func (state *PoisonedState) PostRound() {
	if state.remainRound == 0 {
		state.unit.RetrieveState(NewNormalState())
	}
}
