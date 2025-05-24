package numberlist

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
)

type BaseNumberList struct {
	list     []string
	position *core.Position
}

var _ NumberList = (*BaseNumberList)(nil)

func NewBaseNumberList(list []string, position *core.Position) *BaseNumberList {
	return &BaseNumberList{
		list:     list,
		position: position,
	}
}

func (b *BaseNumberList) Render(canvas [][]string) {
	newList := make([]string, len(b.list))
	for i, line := range b.list {
		newList[i] = fmt.Sprintf("* %s", line)
	}

	startRow := b.position.Y
	startCol := b.position.X

	for i, line := range newList {
		row := startRow + i
		if row >= len(canvas) {
			break // 超出 canvas 高度
		}

		for j, ch := range line {
			col := startCol + j
			if col >= len(canvas[row]) {
				break // 超出 canvas 寬度
			}
			canvas[row][col] = string(ch)
		}
	}
}
