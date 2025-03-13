package sprite

import "fmt"

type Fire struct {
	Sprite
}

var _ ISprite = (*Fire)(nil)

func NewFire(position int) *Fire {
	return &Fire{
		Sprite: NewSprite(position),
	}
}

func (f Fire) Type() SpriteType {
	return SpriteTypeFire
}

func (f Fire) String() string {
	format := "Fire {Position: %d}"
	return fmt.Sprintf(format, f.Position())
}
