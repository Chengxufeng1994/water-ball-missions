package sprite

import "fmt"

type Water struct {
	Sprite
}

var _ ISprite = (*Water)(nil)

func NewWater(position int) *Water {
	return &Water{
		Sprite: NewSprite(position),
	}
}

func (w Water) Type() SpriteType {
	return SpriteTypeWater
}

func (w Water) String() string {
	format := "Water {Position: %d}"
	return fmt.Sprintf(format, w.Position())
}
