package domain

type Cusrer interface {
	GetID() int
	OnCursedDead(amount int)
	AddCurser(curser Cusrer)
	RemoveCurser(curser Cusrer)
}
