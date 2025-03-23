package domain

type MapObject interface {
	GetSymbol() Symbol
	SetSymbol(Symbol)
	Position() (int, int)
	SetPosition(x, y int)
}

type BasedMapObject struct {
	Symbol Symbol
	X      int
	Y      int
}

func NewBasedMapObject(symbol Symbol, x, y int) BasedMapObject {
	return BasedMapObject{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (b BasedMapObject) GetSymbol() Symbol {
	return b.Symbol
}

func (b *BasedMapObject) SetSymbol(symbol Symbol) {
	b.Symbol = symbol
}

func (b BasedMapObject) Position() (int, int) {
	return b.X, b.Y
}

func (b *BasedMapObject) SetPosition(x, y int) {
	b.X, b.Y = x, y
}
