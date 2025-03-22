package command

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/core"
)

type ResetMainControlKeyboard struct {
	keyboard core.KeyBoardInterface
}

var _ Command = (*ResetMainControlKeyboard)(nil)

func NewResetMainControlKeyboard(keyboard core.KeyBoardInterface) *ResetMainControlKeyboard {
	return &ResetMainControlKeyboard{keyboard: keyboard}
}

func (r *ResetMainControlKeyboard) Name() string {
	return "ResetMainControlKeyboard"
}

// Execute implements Command.
func (r *ResetMainControlKeyboard) Execute() {
	r.keyboard.Reset()
}

// Undo implements Command.
func (r *ResetMainControlKeyboard) Undo() {
	// do nothing
}
