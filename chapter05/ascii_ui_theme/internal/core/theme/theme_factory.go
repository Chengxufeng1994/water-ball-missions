package theme

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/numberlist"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text"
)

type ASCIIThemeFactory interface {
	CreateButton(text string, padding *button.Padding, position *core.Position) button.Button
	CreateNumberList(list []string, position *core.Position) numberlist.NumberList
	CreateText(value []string, position *core.Position) text.Text
}
