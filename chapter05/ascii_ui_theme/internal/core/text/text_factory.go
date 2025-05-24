package text

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"

type TextFactory interface {
	CreateText(value []string, position *core.Position) Text
}
