package numberlist

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"

type NumberListFactory interface {
	CreateNumberList(list []string, position *core.Position) NumberList
}
