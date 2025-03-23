package domain

type IState interface {
	SetRole(role Role)

	RetrieveState()
	LoseState()

	DeduceRound()
	PreRound()
	PostRound()

	OnDamage()
}
