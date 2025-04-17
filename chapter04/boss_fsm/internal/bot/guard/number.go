package guard

type NumberGuard interface {
	Check(value int) bool
}

type GreaterThanEqualNumberGuard struct {
	value int
}

func NewGreaterThanEqualNumberGuard(value int) NumberGuard {
	return &GreaterThanEqualNumberGuard{value: value}
}

func (g *GreaterThanEqualNumberGuard) Check(value int) bool {
	return value >= g.value
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
