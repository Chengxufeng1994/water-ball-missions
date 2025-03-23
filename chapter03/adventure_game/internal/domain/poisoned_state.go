package domain

import "strconv"

type PoisonedState struct {
	role        Role
	remainRound int
}

var _ IState = (*PoisonedState)(nil)

func NewPoisonedState() *PoisonedState {
	return &PoisonedState{
		remainRound: 3,
	}
}

func (state *PoisonedState) SetRole(role Role) {
	state.role = role
}

func (state *PoisonedState) RetrieveState() {
	// do nothing
}

func (state *PoisonedState) LoseState() {
	state.role.RetrieveState(NewNormalState())
}

func (state *PoisonedState) DeduceRound() {
	state.remainRound--
}

func (state *PoisonedState) PreRound() {
	state.role.LoseHP(15)
}

func (state *PoisonedState) PostRound() {
	if state.remainRound == 0 {
		state.LoseState()
	}
}

func (state *PoisonedState) OnDamage() {
	// do nothing
}

func (state *PoisonedState) String() string {
	return "Poisoned, RemainRound: " + strconv.Itoa(state.remainRound)
}
