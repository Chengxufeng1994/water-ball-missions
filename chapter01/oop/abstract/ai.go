package abstract

import "math/rand"

type AI struct {
	id string
}

var _ Player = (*AI)(nil)

func NewAI(id string) *AI {
	return &AI{
		id: id,
	}
}

func (a AI) ID() string {
	return a.id
}

func (a AI) MakeDecide() Decision {
	list := []Decision{Paper, Scissors, Stone}

	return list[rand.Intn(len(list))]
}
