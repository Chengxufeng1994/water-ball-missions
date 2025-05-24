package text

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"

type BaseTextDecorator struct {
	Inner Text
}

var _ Text = (*BaseTextDecorator)(nil)

func NewBaseTextDecorator(Inner Text) *BaseTextDecorator {
	return &BaseTextDecorator{Inner: Inner}
}

func (b *BaseTextDecorator) Lines() []string {
	return b.Inner.Lines()
}

func (b *BaseTextDecorator) Position() *core.Position {
	return b.Inner.Position()
}

func (b *BaseTextDecorator) Render(canvas [][]string) {
	b.Inner.Render(canvas)
}
