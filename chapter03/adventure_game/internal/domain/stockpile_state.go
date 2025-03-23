package domain

import "strconv"

type StockpileState struct {
	role        Role
	remainRound int
}

var _ IState = (*StockpileState)(nil)

func NewStockpileState() IState {
	return &StockpileState{
		remainRound: 2,
	}
}

func (state *StockpileState) SetRole(role Role) {
	state.role = role
}

// RetrieveState implements IState.
func (state *StockpileState) RetrieveState() {
	// do nothing
}

// LoseState implements IState.
func (state *StockpileState) LoseState() {
	state.role.RetrieveState(NewNormalState())
}

// DeduceRound implements IState.
func (state *StockpileState) DeduceRound() {
	state.remainRound--
}

// PreRound implements IState.
func (state *StockpileState) PreRound() {
	// do nothing
}

// PostRound implements IState.
func (state *StockpileState) PostRound() {
	if state.remainRound == 0 {
		state.role.RetrieveState(NewEruptingState())
	}
}

// OnDamage implements IState.
func (state *StockpileState) OnDamage() {
	state.LoseState()
}

func (state *StockpileState) String() string {
	return "Stockpile, RemainRound: " + strconv.Itoa(state.remainRound)
}
