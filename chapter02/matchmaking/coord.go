package main

type Coord struct {
	X float64
	Y float64
}

func NewCoord(x, y float64) Coord {
	return Coord{
		X: x,
		Y: y,
	}
}
