package command

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/core"
)

type MoveTankBackward struct {
	tank core.TankInterface
}

var _ Command = (*MoveTankBackward)(nil)

func NewMoveTankBackward(tank core.TankInterface) *MoveTankBackward {
	return &MoveTankBackward{
		tank: tank,
	}
}

// Name implements Command.
func (m *MoveTankBackward) Name() string {
	return "MoveTankBackward"
}

// Execute implements Command.
func (m *MoveTankBackward) Execute() {
	m.tank.MoveBackward()
}

// Undo implements Command.
func (m *MoveTankBackward) Undo() {
	m.tank.MoveForward()
}
