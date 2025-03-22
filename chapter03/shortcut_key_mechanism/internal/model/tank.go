package model

import "fmt"

type Tank struct{}

func NewTank() *Tank {
	return &Tank{}
}

func (t *Tank) MoveForward() {
	fmt.Println("The tank has moved forward.")
}

func (t *Tank) MoveBackward() {
	fmt.Println("The tank has moved backward.")
}
