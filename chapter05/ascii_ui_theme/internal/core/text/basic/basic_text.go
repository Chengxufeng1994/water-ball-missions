package basic

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
)

type BasicText struct {
	*text.BaseTextDecorator
}

func NewBasicText(inner text.Text) *BasicText {
	return &BasicText{
		BaseTextDecorator: text.NewBaseTextDecorator(inner),
	}
}
