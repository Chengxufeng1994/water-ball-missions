package theme

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/numberlist"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text/basic"
)

type BasicAsciiThemeFactory struct{}

var _ ASCIIThemeFactory = (*BasicAsciiThemeFactory)(nil)

func NewBasicAsciiThemeFactory() *BasicAsciiThemeFactory {
	return &BasicAsciiThemeFactory{}
}

func (b *BasicAsciiThemeFactory) CreateButton(
	text string, padding *button.Padding, position *core.Position,
) button.Button {
	return button.NewBasicButton(text, padding, position)
}

func (b *BasicAsciiThemeFactory) CreateNumberList(
	list []string, position *core.Position,
) numberlist.NumberList {
	return numberlist.NewBasicNumberList(list, position)
}

func (b *BasicAsciiThemeFactory) CreateText(
	value []string, position *core.Position,
) text.Text {
	return basic.NewBasicText(text.NewBaseText(value, position))
}
