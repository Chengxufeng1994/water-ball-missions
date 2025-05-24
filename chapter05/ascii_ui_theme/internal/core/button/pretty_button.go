package button

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button/borderstyle"
)

type PrettyButton struct {
	*BaseButton
}

func NewPrettyButton(text string, padding *Padding, position *core.Position) *PrettyButton {
	return &PrettyButton{
		BaseButton: NewBaseButton(text, padding, borderstyle.NewPrettyBorderStyle(), position),
	}
}
