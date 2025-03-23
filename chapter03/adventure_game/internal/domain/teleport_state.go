package domain

type TeleportState struct {
	role          Role
	remainedRound int
}

var _ IState = (*TeleportState)(nil)

func NewTeleportState() IState {
	return &TeleportState{
		remainedRound: 1,
	}
}

// SetRole implements IState.
func (t *TeleportState) SetRole(role Role) {
	t.role = role
}

// RetrieveState implements IState.
func (t *TeleportState) RetrieveState() {
	// do nothing
}

// LoseState implements IState.
func (t *TeleportState) LoseState() {
	t.role.RandomMove()
	t.role.RetrieveState(NewNormalState())
	t.role = nil
}

// DeduceRound implements IState.
func (t *TeleportState) DeduceRound() {
	t.remainedRound--
}

// PreRound implements IState.
func (t *TeleportState) PreRound() {
	// do nothing
}

// PostRound implements IState.
func (t *TeleportState) PostRound() {
	if t.remainedRound == 0 {
		t.LoseState()
	}
}

// OnDamage implements IState.
func (t *TeleportState) OnDamage() {
	// do nothing
}

func (t *TeleportState) String() string {
	return "TeleportState, RemainedRound: " + string(t.remainedRound)
}
