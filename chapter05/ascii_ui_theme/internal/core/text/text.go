package text

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"

type Text interface {
	core.Element
	Lines() []string
	Position() *core.Position
}
