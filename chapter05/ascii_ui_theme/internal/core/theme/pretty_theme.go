package theme

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/numberlist"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text/pretty"
)

type PrettyAsciiThemeFactory struct{}

var _ ASCIIThemeFactory = (*PrettyAsciiThemeFactory)(nil)

func NewPrettyAsciiThemeFactory() *PrettyAsciiThemeFactory {
	return &PrettyAsciiThemeFactory{}
}

func (p *PrettyAsciiThemeFactory) CreateButton(
	text string, padding *button.Padding, position *core.Position,
) button.Button {
	return button.NewPrettyButton(text, padding, position)
}

func (p *PrettyAsciiThemeFactory) CreateNumberList(
	list []string, position *core.Position,
) numberlist.NumberList {
	return numberlist.NewPrettyNumberList(list, position)
}

func (p *PrettyAsciiThemeFactory) CreateText(
	value []string, position *core.Position,
) text.Text {
	return pretty.NewPrettyText(text.NewBaseText(value, position))
}
