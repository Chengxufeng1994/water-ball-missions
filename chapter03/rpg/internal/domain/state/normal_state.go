package state

type NormalState struct {
	*BasedState
}

func NewNormalState() *NormalState {
	return &NormalState{
		BasedState: NewBasedState("正常", 0),
	}
}

func (state *NormalState) PreTurn() {
	// do nothing
}

func (state *NormalState) PostTurn() {
	// do nothing
}
