package command

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/core"
)

type MoveTankForward struct {
	tank core.TankInterface
}

var _ Command = (*MoveTankForward)(nil)

func NewMoveTankForward(tank core.TankInterface) *MoveTankForward {
	return &MoveTankForward{
		tank: tank,
	}
}

// Name implements Command.
func (m *MoveTankForward) Name() string {
	return "MoveTankForward"
}

// Execute implements Command.
func (m *MoveTankForward) Execute() {
	m.tank.MoveForward()
}

// Undo implements Command.
func (m *MoveTankForward) Undo() {
	m.tank.MoveBackward()
}
