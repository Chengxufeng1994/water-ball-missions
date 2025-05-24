package button

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button/borderstyle"
)

type BasicButton struct {
	*BaseButton
}

var _ Button = (*BasicButton)(nil)

func NewBasicButton(text string, padding *Padding, position *core.Position) *BasicButton {
	return &BasicButton{
		BaseButton: NewBaseButton(text, padding, borderstyle.NewBasicBorderStyle(), position),
	}
}
