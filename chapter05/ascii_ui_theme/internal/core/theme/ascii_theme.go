package theme

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/numberlist"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
)

type Theme interface{}

type AsciiTheme struct {
	ButtonFactory     button.ButtonFactory
	NumberListFactory numberlist.NumberListFactory
	TextFactory       text.TextFactory
}

func NewAsciiTheme(
	buttonFactory button.ButtonFactory,
	numberListFactory numberlist.NumberListFactory,
	textFactory text.TextFactory,
) *AsciiTheme {
	return &AsciiTheme{
		ButtonFactory:     buttonFactory,
		NumberListFactory: numberListFactory,
		TextFactory:       textFactory,
	}
}
