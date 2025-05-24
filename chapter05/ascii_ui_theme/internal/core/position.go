package core

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewPosition(x, y int) *Position {
	return &Position{X: x, Y: y}
}
