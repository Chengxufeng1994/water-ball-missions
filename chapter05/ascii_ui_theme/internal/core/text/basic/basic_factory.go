package basic

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
)

type BasicTextFactory struct{}

var _ text.TextFactory = (*BasicTextFactory)(nil)

func NewBasicTextFactory() *BasicTextFactory {
	return &BasicTextFactory{}
}

func (b *BasicTextFactory) CreateText(value []string, position *core.Position) text.Text {
	baseText := text.NewBaseText(value, position)
	return NewBasicText(baseText)
}
