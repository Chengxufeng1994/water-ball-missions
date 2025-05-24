package factory

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/ascii_ui_theme/internal/core/numberlist"
)

type BasicNumberListFactory struct{}

func NewBasicNumberListFactory() *BasicNumberListFactory {
	return &BasicNumberListFactory{}
}

func (b *BasicNumberListFactory) CreateNumberList(
	list []string, position *core.Position,
) numberlist.NumberList {
	return numberlist.NewBasicNumberList(list, position)
}
