package domain

import "strconv"

type AccelerationPotionState struct {
	role        Role
	remainRound int
}

var _ IState = (*AccelerationPotionState)(nil)

func NewAccelerationPotionState() *AccelerationPotionState {
	return &AccelerationPotionState{
		remainRound: 3,
	}
}

func (state *AccelerationPotionState) SetRole(role Role) {
	state.role = role
}

func (state *AccelerationPotionState) RetrieveState() {
	state.role.UpdateNumOfAction(2)
}

func (state *AccelerationPotionState) LoseState() {
	state.role.UpdateNumOfAction(1)
	state.role.RetrieveState(NewNormalState())
}

func (state *AccelerationPotionState) DeduceRound() {
	state.remainRound--
}

func (state *AccelerationPotionState) PreRound() {
	// do nothing
}

func (state *AccelerationPotionState) PostRound() {
	if state.remainRound == 0 {
		state.LoseState()
	}
}

func (state *AccelerationPotionState) OnDamage() {
	state.LoseState()
}

func (state *AccelerationPotionState) String() string {
	return "AccelerationPotion, RemainRound: " + strconv.Itoa(state.remainRound)
}
