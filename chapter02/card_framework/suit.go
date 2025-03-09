package main

type Suit int

const (
	_ Suit = iota
	Club
	Diamond
	Heart
	Spade
)

func (s Suit) GreaterThan(other Suit) bool {
	return s > other
}

func (s Suit) String() string {
	switch s {
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Spade:
		return "Spade"
	}
	return ""
}
