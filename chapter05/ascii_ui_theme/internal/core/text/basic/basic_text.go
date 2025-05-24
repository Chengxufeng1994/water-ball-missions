package basic

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
)

type BasicText struct {
	*text.BaseText
}

func NewBasicText(value []string, position *core.Position) *BasicText {
	return &BasicText{
		BaseText: text.NewBaseText(value, position),
	}
}
