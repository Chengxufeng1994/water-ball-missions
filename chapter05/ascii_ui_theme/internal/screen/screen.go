package screen

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/button"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/theme"
)

type Screen struct {
	ScreenSize   *ScreenSize
	CurrentTheme *theme.AsciiTheme
}

func NewScreen(
	screenSize *ScreenSize,
	theme *theme.AsciiTheme,
) *Screen {
	return &Screen{
		ScreenSize:   screenSize,
		CurrentTheme: theme,
	}
}

func (ui *Screen) GetScreenSize() *ScreenSize {
	return ui.ScreenSize
}

func (ui *Screen) SetScreenSize(screenSize *ScreenSize) {
	ui.ScreenSize = screenSize
}

func (ui *Screen) SetAsciiTheme(theme *theme.AsciiTheme) {
	ui.CurrentTheme = theme
}

func (s Screen) Render() {
	canvas := make([][]string, s.ScreenSize.Height)
	for i := range canvas {
		canvas[i] = make([]string, s.ScreenSize.Width)
		for j := range canvas[i] {
			canvas[i][j] = " " // 每格初始化為空白
		}
	}

	// 加邊框（點點）
	for y := 0; y < s.ScreenSize.Height; y++ {
		for x := 0; x < s.ScreenSize.Width; x++ {
			if y == 0 || y == s.ScreenSize.Height-1 || x == 0 || x == s.ScreenSize.Width-1 {
				canvas[y][x] = "."
			}
		}
	}

	// 渲染元件（這邊假設每個元件已改為接受 [][]string 並考慮寬度）
	s.CurrentTheme.ButtonFactory.CreateButton(
		"Hi, I miss u", button.NewPadding(0, 0, 1, 1), core.NewPosition(3, 1)).
		Render(canvas)

	s.CurrentTheme.ButtonFactory.CreateButton(
		"No", button.NewPadding(0, 0, 1, 1), core.NewPosition(3, 6)).
		Render(canvas)

	s.CurrentTheme.ButtonFactory.CreateButton(
		"Yes", button.NewPadding(0, 0, 1, 1), core.NewPosition(12, 6)).
		Render(canvas)
	s.CurrentTheme.TextFactory.CreateText(
		[]string{"Do u love me ?", "Please tell..."}, core.NewPosition(4, 4)).
		Render(canvas)
	s.CurrentTheme.NumberListFactory.CreateNumberList(
		[]string{"Let's Travel", "Back to home", "Have dinner"}, core.NewPosition(3, 9)).
		Render(canvas)

	// 輸出畫面
	for _, row := range canvas {
		line := ""
		for _, cell := range row {
			line += cell
		}
		fmt.Println(line)
	}
}
