package factory

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button"
)

type BasicButtonFactory struct{}

var _ button.ButtonFactory = (*BasicButtonFactory)(nil)

func NewBasicButtonFactory() *BasicButtonFactory {
	return &BasicButtonFactory{}
}

func (b *BasicButtonFactory) CreateButton(
	text string, padding *button.Padding, position *core.Position,
) button.Button {
	return button.NewBasicButton(text, padding, position)
}
