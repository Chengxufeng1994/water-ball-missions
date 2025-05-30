package main

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/theme"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/screen"
)

func main() {
	basicAsciiThemeFactory := theme.NewBasicAsciiThemeFactory()
	screen := screen.NewScreen(screen.NewScreenSize(22, 13), basicAsciiThemeFactory)
	screen.Render()

	prettyAsciiThemeFactory := theme.NewPrettyAsciiThemeFactory()
	screen.SetAsciiThemeFactory(prettyAsciiThemeFactory)
	screen.Render()
}
