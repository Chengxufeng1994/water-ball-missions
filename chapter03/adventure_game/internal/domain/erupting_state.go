package domain

import "strconv"

type EruptingState struct {
	role          Role
	remainedRound int
}

var _ IState = (*EruptingState)(nil)

func NewEruptingState() IState {
	return &EruptingState{
		remainedRound: 3,
	}
}

func (state *EruptingState) SetRole(role Role) {
	state.role = role
}

func (state *EruptingState) RetrieveState() {
	state.role.UpdateAttackBehavior(AttackBehaviorFullMap)
}

func (state *EruptingState) LoseState() {
	state.role.UpdateAttackBehavior(AttackBehaviorOnLine)
	state.role.RetrieveState(NewTeleportState())
}

func (state *EruptingState) DeduceRound() {
	state.remainedRound--
}

func (state *EruptingState) PreRound() {
	// do nothing
}

func (state *EruptingState) PostRound() {
	if state.remainedRound == 0 {
		state.LoseState()
	}
}

func (state *EruptingState) OnDamage() {
}

func (state *EruptingState) String() string {
	return "Erupting, RemainRound: " + strconv.Itoa(state.remainedRound)
}
