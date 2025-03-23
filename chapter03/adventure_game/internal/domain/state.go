package domain

type State struct {
	StateType StateType
}

func NewNormalState() State {
	return State{StateType: StateTypeNormal}
}

func NewInvincible() State {
	return State{StateType: StateTypeInvincible}
}

func NewPoisoned() State {
	return State{StateType: StateTypePoisoned}
}

func NewAccelerated() State {
	return State{StateType: StateTypeAccelerated}
}

func NewHealing() State {
	return State{StateType: StateTypeHealing}
}

func NewOrderless() State {
	return State{StateType: StateTypeOrderless}
}

func NewStockpile() State {
	return State{StateType: StateTypeStockpile}
}

func NewErupting() State {
	return State{StateType: StateTypeErupting}
}

func NewTeleport() State {
	return State{StateType: StateTypeTeleport}
}
