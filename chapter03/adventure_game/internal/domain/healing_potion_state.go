package domain

import "strconv"

type HealingPotionState struct {
	role        Role
	remainRound int
}

var _ IState = (*HealingPotionState)(nil)

func NewHealingPotionState() *HealingPotionState {
	return &HealingPotionState{
		remainRound: 5,
	}
}

func (state *HealingPotionState) SetRole(role Role) {
	state.role = role
}

func (state *HealingPotionState) RetrieveState() {
	// do nothing
}

func (state *HealingPotionState) LoseState() {
	state.role.RetrieveState(NewNormalState())
}

func (state *HealingPotionState) DeduceRound() {
	state.remainRound--
}

func (state *HealingPotionState) PreRound() {
	state.role.RecoverHP(30)
}

func (state *HealingPotionState) PostRound() {
	if state.remainRound == 0 || state.role.IsFullHP() {
		state.LoseState()
	}
}

func (state *HealingPotionState) OnDamage() {
	// do nothing
}

func (state *HealingPotionState) String() string {
	return "HealingPotion, RemainRound: " + strconv.Itoa(state.remainRound)
}
