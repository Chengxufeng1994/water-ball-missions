package factory

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/numberlist"
)

type PrettyNumberListFactory struct{}

func NewPrettyNumberListFactory() *PrettyNumberListFactory {
	return &PrettyNumberListFactory{}
}

func (b *PrettyNumberListFactory) CreateNumberList(
	list []string, position *core.Position,
) numberlist.NumberList {
	return numberlist.NewPrettyNumberList(list, position)
}
