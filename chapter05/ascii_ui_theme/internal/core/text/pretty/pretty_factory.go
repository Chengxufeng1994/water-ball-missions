package pretty

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
)

type PrettyTextFactory struct{}

func NewPrettyTextFactory() *PrettyTextFactory {
	return &PrettyTextFactory{}
}

func (p *PrettyTextFactory) CreateText(value []string, position *core.Position) text.Text {
	baseText := text.NewBaseText(value, position)
	return NewPrettyText(baseText)
}
