package numberlist

import (
	"fmt"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
)

type PrettyNumberList struct {
	*BasicNumberList
}

func NewPrettyNumberList(list []string, position *core.Position) *PrettyNumberList {
	return &PrettyNumberList{
		BasicNumberList: NewBasicNumberList(list, position),
	}
}

func (p *PrettyNumberList) Render(canvas [][]string) {
	newList := make([]string, len(p.list))
	for i, line := range p.list {
		newList[i] = fmt.Sprintf("%s. %s", ToRoman(i+1), line)
	}

	startRow := p.position.Y
	startCol := p.position.X

	for i, line := range newList {
		row := startRow + i
		if row >= len(canvas) {
			break // 超出 canvas 高度
		}

		for j, ch := range line {
			col := startCol + j
			if col >= len(canvas[row]) {
				break // 超出 canvas 寬度
			}
			canvas[row][col] = string(ch)
		}
	}
}

// 支援範圍：1~3999
func ToRoman(num int) string {
	if num <= 0 || num >= 4000 {
		return "Invalid input (must be between 1 and 3999)"
	}

	var romanBuilder strings.Builder

	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			romanBuilder.WriteString(syms[i])
		}
	}
	return romanBuilder.String()
}
