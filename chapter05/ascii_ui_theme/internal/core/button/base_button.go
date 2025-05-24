package button

import (
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button/borderstyle"
)

type BaseButton struct {
	Text        string
	Padding     *Padding
	BorderStyle borderstyle.BorderStyle
	Position    *core.Position
}

var _ Button = (*BaseButton)(nil)

func NewBaseButton(
	text string,
	padding *Padding,
	borderStyle borderstyle.BorderStyle,
	position *core.Position,
) *BaseButton {
	return &BaseButton{
		Text:        text,
		Padding:     padding,
		BorderStyle: borderStyle,
		Position:    position,
	}
}

func (b *BaseButton) Render(canvas [][]string) {
	lines := []string{}
	contentWidth := len(b.Text) + b.Padding.Left + b.Padding.Right

	topBorder := b.TopBorder(contentWidth)
	bottomBorder := b.BottomBorder(contentWidth)
	contentLine := b.ContentLine(contentWidth)
	emptyLine := b.EmptyLine(contentWidth)

	lines = append(lines, topBorder)
	for range b.Padding.Top {
		lines = append(lines, emptyLine)
	}

	lines = append(lines, contentLine)

	for range b.Padding.Bottom {
		lines = append(lines, emptyLine)
	}
	lines = append(lines, bottomBorder)

	// ✅ 渲染每一行到 canvas
	for i, line := range lines {
		y := b.Position.Y + i
		if y < 0 || y >= len(canvas) {
			continue // 超出畫面就略過
		}
		for j, r := range line {
			x := b.Position.X + j
			if x < 0 || x >= len(canvas[0]) {
				continue
			}

			char := string(r)
			canvas[y][x] = char
		}
	}
}

// func (b *BaseButton) Render(canvas [][]rune) {
// 	lines := []string{}
// 	contentWidth := len(b.Text) + b.Padding.Left + b.Padding.Right

// 	topBorder := b.TopBorder(contentWidth)
// 	bottomBorder := b.BottomBorder(contentWidth)
// 	contentLine := b.ContentLine(contentWidth)
// 	emptyLine := b.EmptyLine(contentWidth)

// 	lines = append(lines, topBorder)
// 	for range b.Padding.Top {
// 		lines = append(lines, emptyLine)
// 	}

// 	lines = append(lines, contentLine)
// 	for range b.Padding.Bottom {
// 		lines = append(lines, emptyLine)
// 	}

// 	lines = append(lines, bottomBorder)

// fmt.Println(strings.Join(lines, "\n"))
// }

// --- Hook Methods ---
func (b *BaseButton) TopBorder(width int) string {
	return b.BorderStyle.TopLeftCorner() +
		strings.Repeat(b.BorderStyle.HorizontalEdge(), width) +
		b.BorderStyle.TopRightCorner()
}

func (b *BaseButton) BottomBorder(width int) string {
	return b.BorderStyle.BottomLeftCorner() +
		strings.Repeat(b.BorderStyle.HorizontalEdge(), width) +
		b.BorderStyle.BottomRightCorner()
}

func (b *BaseButton) ContentLine(width int) string {
	return b.BorderStyle.VerticalEdge() +
		strings.Repeat(" ", b.Padding.Left) + b.Text + strings.Repeat(" ", b.Padding.Right) +
		b.BorderStyle.VerticalEdge()
}

func (b *BaseButton) EmptyLine(width int) string {
	return b.BorderStyle.VerticalEdge() + strings.Repeat(" ", width) + b.BorderStyle.VerticalEdge()
}
