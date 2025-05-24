package pretty

import (
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
)

type PrettyText struct {
	*text.BaseTextDecorator
}

func NewPrettyText(inner text.Text) *PrettyText {
	return &PrettyText{
		BaseTextDecorator: text.NewBaseTextDecorator(inner),
	}
}

func (pt *PrettyText) Lines() []string {
	newLines := make([]string, len(pt.Inner.Lines()))
	for i, line := range pt.Inner.Lines() {
		newLines[i] = strings.ToUpper(line)
	}
	return newLines
}

func (pt *PrettyText) Render(canvas [][]string) {
	for i, line := range pt.Lines() {
		row := pt.Position().Y + i
		if row < 0 || row >= len(canvas) {
			continue
		}
		for j, ch := range line {
			col := pt.Position().X + j
			if col < 0 || col >= len(canvas[row]) {
				break
			}
			canvas[row][col] = string(ch)
		}
	}
}
