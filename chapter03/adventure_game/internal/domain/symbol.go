package domain

import (
	"errors"
	"math/rand"
)

type Symbol string

const (
	DirectionUp    Symbol = "↑"
	DirectionRight        = "→"
	DirectionDown         = "↓"
	DirectionLeft         = "←"
	SymbolEmpty           = "."
	SymbolObstacle        = "□"
	SymbolTreasure        = "X"
	SymbolMonster         = "M"
)

var SymbolMaps = map[string]Symbol{
	"DirectionUp":    DirectionUp,
	"DirectionDown":  DirectionDown,
	"DirectionLeft":  DirectionLeft,
	"DirectionRight": DirectionRight,
	"w":              DirectionUp,
	"s":              DirectionDown,
	"a":              DirectionLeft,
	"d":              DirectionRight,
	"Obstacle":       SymbolObstacle,
	"Treasure":       SymbolTreasure,
	"Monster":        SymbolMonster,
}

var ErrInvalidSymbol = errors.New("invalid symbol")

func NewSymbol(symbol string) (Symbol, error) {
	if _, ok := SymbolMaps[symbol]; !ok {
		return "", ErrInvalidSymbol
	}

	return Symbol(symbol), nil
}

var SymbolDirections = []Symbol{DirectionUp, DirectionRight, DirectionDown, DirectionLeft}

func GetRandomDirection() Symbol {
	return SymbolDirections[rand.Intn(len(SymbolDirections))]
}

var DirectionMoves = map[Symbol][2]int{
	DirectionUp:    {0, -1},
	DirectionRight: {1, 0},
	DirectionDown:  {0, 1},
	DirectionLeft:  {-1, 0},
}
