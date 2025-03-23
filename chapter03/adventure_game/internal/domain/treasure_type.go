package domain

type TreasureType int

const (
	TreasureTypeSuperStar TreasureType = iota + 1
	TreasureTypePoison
	TreasureTypeAccelerationPotion
	TreasureTypeHealingPotion
	TreasureTypeDevilFruit
	TreasureTypeKingsRock
	TreasureTypeDokodemoDoor
)

func (t TreasureType) String() string {
	switch t {
	case TreasureTypeSuperStar:
		return "SuperStar"
	case TreasureTypePoison:
		return "Poison"
	case TreasureTypeAccelerationPotion:
		return "AccelerationPotion"
	case TreasureTypeHealingPotion:
		return "HealingPotion"
	case TreasureTypeDevilFruit:
		return "DevilFruit"
	case TreasureTypeKingsRock:
		return "KingsRock"
	case TreasureTypeDokodemoDoor:
		return "DokodemoDoor"
	default:
		return ""
	}
}
