package domain

import "math/rand"

type Treasure struct {
	BasedMapObject
	TreasureType TreasureType
	State        State
}

var _ MapObject = (*Treasure)(nil)

func NewTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
	}
}

func (t *Treasure) ApplyTreasure(mapObject MapObject) {
	if character, ok := mapObject.(*Character); ok {
		character.RetrieveState(t.State)
	}

	if _, ok := mapObject.(*Monster); ok {
	}
}

func NewSuperStarTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
		TreasureType:   TreasureTypeSuperStar,
		State:          NewInvincible(),
	}
}

func NewPoisonTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
		TreasureType:   TreasureTypePoison,
		State:          NewPoisoned(),
	}
}

func NewAccelerationPotionTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
		TreasureType:   TreasureTypeAccelerationPotion,
		State:          NewAccelerated(),
	}
}

func NewHealingPotionTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
		TreasureType:   TreasureTypeHealingPotion,
		State:          NewHealing(),
	}
}

func NewDevilFruitTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
		TreasureType:   TreasureTypeDevilFruit,
		State:          NewOrderless(),
	}
}

func NewKingsRockTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
		TreasureType:   TreasureTypeKingsRock,
		State:          NewStockpile(),
	}
}

func NewDokodemoDoorTreasure(x, y int) *Treasure {
	return &Treasure{
		BasedMapObject: NewBasedMapObject(SymbolTreasure, x, y),
		TreasureType:   TreasureTypeDokodemoDoor,
		State:          NewTeleport(),
	}
}

// 依照機率隨機產生寶物
func GenerateRandomTreasure(x, y int) *Treasure {
	r := rand.Float64()

	switch {
	case r < 0.1:
		return NewSuperStarTreasure(x, y)
	case r < 0.35:
		return NewPoisonTreasure(x, y)
	case r < 0.55:
		return NewAccelerationPotionTreasure(x, y)
	case r < 0.7:
		return NewHealingPotionTreasure(x, y)
	case r < 0.8:
		return NewDevilFruitTreasure(x, y)
	case r < 0.9:
		return NewKingsRockTreasure(x, y)
	default:
		return NewDokodemoDoorTreasure(x, y)
	}
}
