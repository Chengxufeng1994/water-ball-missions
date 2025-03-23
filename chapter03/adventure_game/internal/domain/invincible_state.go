package domain

import "strconv"

type InvincibleState struct {
	role        Role
	remainRound int
}

var _ IState = (*InvincibleState)(nil)

func NewInvincibleState() *InvincibleState {
	return &InvincibleState{
		remainRound: 2,
	}
}

func (state *InvincibleState) SetRole(role Role) {
	state.role = role
}

func (state *InvincibleState) RetrieveState() {
	state.role.UpdateInvincible(true)
}

func (state *InvincibleState) LoseState() {
	state.role.UpdateInvincible(false)
	state.role.RetrieveState(NewNormalState())
}

func (state *InvincibleState) DeduceRound() {
	state.remainRound--
}

func (state *InvincibleState) PreRound() {
	// do nothing
}

func (state *InvincibleState) PostRound() {
	if state.remainRound == 0 {
		state.LoseState()
	}
}

func (state *InvincibleState) OnDamage() {
	// do nothing
}

func (state *InvincibleState) String() string {
	return "Invincible, RemainRound: " + strconv.Itoa(state.remainRound)
}
