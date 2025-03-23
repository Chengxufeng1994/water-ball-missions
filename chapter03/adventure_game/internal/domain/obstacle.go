package domain

type Obstacle struct {
	BasedMapObject
}

var _ MapObject = (*Obstacle)(nil)

func NewObstacle(x, y int) *Obstacle {
	return &Obstacle{
		BasedMapObject: NewBasedMapObject(SymbolObstacle, x, y),
	}
}
