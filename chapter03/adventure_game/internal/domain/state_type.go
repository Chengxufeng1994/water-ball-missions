package domain

// 狀態

// 當角色當前狀態的時效性結束後，如果其沒有特別說明，則角色狀態會回復成正常狀態。

// 名稱	時效性	說明
// 無敵 (Invincible)	2回合	受到攻擊時並不會有任何生命損失
// 中毒 (Poisoned)	3回合	每回合開始時失去15點生命值
// 加速 (Accelerated)	3回合	每回合中可以進行「兩次動作」，若在期間遭受攻擊則立刻恢復至正常狀態
// 恢復 (Healing)	5回合	每回合開始時恢復30點生命值，直到滿血。若滿血則立刻恢復至正常狀態
// 混亂 (Orderless)	3回合	每回合隨機取得以下其中一種效果：1. 只能進行上下移動 2. 只能進行左右移動（角色只能移動，不能選擇做其他操作）
// 蓄力 (Stockpile)	2回合	兩回合後進入爆發狀態，若在期間遭受攻擊則立刻恢復至正常狀態
// 爆發 (Erupting)	3回合	角色的攻擊範圍擴充至「全地圖」，且攻擊行為變成「全場攻擊」：每一次攻擊時都會攻擊到地圖中所有其餘角色，且攻擊力為50。三回合過後取得瞬身狀態。
// 瞬身 (Teleport)	1回合	一回合後角色的位置將被隨機移動至任一空地

type StateType int

const (
	StateTypeNormal StateType = iota
	StateTypeInvincible
	StateTypePoisoned
	StateTypeAccelerated
	StateTypeHealing
	StateTypeOrderless
	StateTypeStockpile
	StateTypeErupting
	StateTypeTeleport
)

func (state StateType) String() string {
	switch state {
	case StateTypeNormal:
		return "NormalState"
	case StateTypeInvincible:
		return "Invincible"
	case StateTypePoisoned:
		return "Poisoned"
	case StateTypeAccelerated:
		return "Accelerated"
	case StateTypeHealing:
		return "Healing"
	case StateTypeOrderless:
		return "Orderless"
	case StateTypeStockpile:
		return "Stockpile"
	case StateTypeErupting:
		return "Erupting"
	case StateTypeTeleport:
		return "Teleport"
	}

	return ""
}
