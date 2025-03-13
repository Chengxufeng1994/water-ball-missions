package sprite

type ISprite interface {
	Position() int
	Move(newPosition int)
	Type() SpriteType
}

type Sprite struct {
	position int
}

var _ ISprite = (*Sprite)(nil)

func NewSprite(position int) Sprite {
	return Sprite{position: position}
}

func (s *Sprite) Move(newPosition int) {
	s.position = newPosition
}

func (s *Sprite) Position() int {
	return s.position
}

func (s *Sprite) Type() SpriteType {
	panic("unimplemented")
}
