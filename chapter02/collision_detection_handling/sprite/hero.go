package sprite

import "fmt"

type Hero struct {
	Sprite
	HP int
}

var _ ISprite = (*Hero)(nil)

func NewHero(position int) *Hero {
	return &Hero{
		Sprite: NewSprite(position),
		HP:     30,
	}
}

func (h *Hero) Damage(number int) {
	h.HP -= number
}

func (h *Hero) Recover(number int) {
	h.HP += number
}

func (h Hero) Type() SpriteType {
	return SpriteTypeHero
}

func (h Hero) String() string {
	format := "Hero {HP: %d, Position: %d}"
	return fmt.Sprintf(format, h.HP, h.Position())
}
