package button

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"

type ButtonFactory interface {
	CreateButton(text string, padding *Padding, position *core.Position) Button
}
