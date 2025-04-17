package guard

type NumberGuard interface {
	Check(value int) bool
}

type GreaterThanNumberGuard struct {
	value int
}

func NewGreaterThanNumberGuard(value int) NumberGuard {
	return &GreaterThanNumberGuard{value: value}
}

func (g *GreaterThanNumberGuard) Check(value int) bool {
	return value > g.value
}

type LessThanNumberGuard struct {
	value int
}

func NewLessThanNumberGuard(value int) NumberGuard {
	return &LessThanNumberGuard{value: value}
}

func (g *LessThanNumberGuard) Check(value int) bool {
	return value < g.value
}
