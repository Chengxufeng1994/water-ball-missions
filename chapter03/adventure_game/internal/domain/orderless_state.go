package domain

import (
	"math/rand"
	"strconv"
)

type OrderlessState struct {
	role        Role
	remainRound int
}

var _ IState = (*OrderlessState)(nil)

func NewOrderlessState() *OrderlessState {
	return &OrderlessState{
		remainRound: 3,
	}
}

func (state *OrderlessState) SetRole(role Role) {
	state.role = role
}

// RetrieveState implements IState.
func (state *OrderlessState) RetrieveState() {
	// do nothing
}

// LoseState implements IState.
func (state *OrderlessState) LoseState() {
	state.role.UpdateDirections(Directions)
	state.role.RetrieveState(NewNormalState())
}

// DeduceRound implements IState.
func (state *OrderlessState) DeduceRound() {
	state.remainRound--
}

// PreRound implements IState.
func (state *OrderlessState) PreRound() {
	dirs := []map[string][2]int{VerticalDirections, HorizontalDirections}
	state.role.UpdateDirections(dirs[rand.Intn(len(dirs))])
}

// PostRound implements IState.
func (state *OrderlessState) PostRound() {
	if state.remainRound == 0 {
		state.LoseState()
	}
}

// OnDamage implements IState.
func (state *OrderlessState) OnDamage() {
	// do nothing
}

func (state *OrderlessState) String() string {
	return "Orderless, RemainRound: " + strconv.Itoa(state.remainRound)
}
