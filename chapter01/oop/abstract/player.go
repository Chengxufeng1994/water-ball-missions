package abstract

type Player interface {
	ID() string
	MakeDecide() Decision
}
