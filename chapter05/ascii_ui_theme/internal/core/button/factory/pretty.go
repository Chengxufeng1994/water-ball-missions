package factory

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button"
)

type PrettyButtonFactory struct{}

var _ button.ButtonFactory = (*PrettyButtonFactory)(nil)

func NewPrettyButtonFactory() *PrettyButtonFactory {
	return &PrettyButtonFactory{}
}

func (p *PrettyButtonFactory) CreateButton(
	text string, padding *button.Padding, position *core.Position,
) button.Button {
	return button.NewPrettyButton(text, padding, position)
}
