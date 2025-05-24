package main

import (
	btnfactory "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button/factory"
	nlfactory "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/numberlist/factory"
	basictext "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text/basic"
	prettytext "github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/text/pretty"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/theme"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/screen"
)

func main() {
	basicButtonFactory := btnfactory.NewBasicButtonFactory()
	basicNumberListFactory := nlfactory.NewBasicNumberListFactory()
	basicTextFactory := basictext.NewBasicTextFactory()
	basicAsciiTheme := theme.NewAsciiTheme(
		basicButtonFactory, basicNumberListFactory, basicTextFactory,
	)
	screen := screen.NewScreen(screen.NewScreenSize(22, 13), basicAsciiTheme)
	screen.Render()

	prettyButtonFactory := btnfactory.NewPrettyButtonFactory()
	prettyNumberListFactory := nlfactory.NewPrettyNumberListFactory()
	prettyTextFactory := prettytext.NewPrettyTextFactory()
	prettyAsciiTheme := theme.NewAsciiTheme(
		prettyButtonFactory, prettyNumberListFactory, prettyTextFactory,
	)
	screen.SetAsciiTheme(prettyAsciiTheme)
	screen.Render()
}
