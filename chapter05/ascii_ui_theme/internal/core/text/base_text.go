package text

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
)

type BaseText struct {
	lines    []string
	position *core.Position
}

var _ Text = (*BaseText)(nil)

func NewBaseText(value []string, position *core.Position) *BaseText {
	return &BaseText{
		lines:    value,
		position: position,
	}
}

func (b *BaseText) Lines() []string {
	return b.lines
}

func (b *BaseText) Position() *core.Position {
	return b.position
}

func (b *BaseText) Render(canvas [][]string) {
	for i, line := range b.Lines() {
		row := b.position.Y + i
		if row < 0 || row >= len(canvas) {
			continue
		}
		for j, ch := range line {
			col := b.position.X + j
			if col < 0 || col >= len(canvas[row]) {
				break
			}
			canvas[row][col] = string(ch)
		}
	}
}
