package hero

type LevelSheet struct{}

func (ls *LevelSheet) queryLevel(exp int) int {
	return exp/1000 + 1
}
