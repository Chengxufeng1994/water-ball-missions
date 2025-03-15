package card

type Suit int

const (
	_ Suit = iota
	Club
	Diamond
	Heart
	Spade
)

func NewSuit(i int) Suit {
	switch i {
	case 1:
		return Club
	case 2:
		return Diamond
	case 3:
		return Heart
	case 4:
		return Spade
	default:
		return Unknown
	}
}

func (s Suit) String() string {
	switch s {
	case Club:
		return "C"
	case Diamond:
		return "D"
	case Heart:
		return "H"
	case Spade:
		return "S"
	default:
		return ""
	}
}
