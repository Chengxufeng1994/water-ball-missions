package domain

type NormalState struct {
	role        Role
	remainRound int
}

var _ IState = (*NormalState)(nil)

func NewNormalState() *NormalState {
	return &NormalState{
		remainRound: 0,
	}
}
func (s *NormalState) SetRole(role Role) {
	s.role = role
}

func (s *NormalState) RetrieveState() {}

func (s *NormalState) LoseState() {}

func (s *NormalState) DeduceRound() {}

func (s *NormalState) PreRound() {}

func (s *NormalState) PostRound() {}

func (s *NormalState) OnDamage() {}

func (s NormalState) String() string {
	return "Normal"
}
