package card

type Rank int

const (
	_ Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	Two

	Unknown = 999
)

func NewRank(i int) Rank {
	switch i {
	case 1:
		return Three
	case 2:
		return Four
	case 3:
		return Five
	case 4:
		return Six
	case 5:
		return Seven
	case 6:
		return Eight
	case 7:
		return Nine
	case 8:
		return Ten
	case 9:
		return Jack
	case 10:
		return Queen
	case 11:
		return King
	case 12:
		return Ace
	case 13:
		return Two
	default:
		return Unknown
	}
}

func (r Rank) String() string {
	switch r {
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	case Two:
		return "2"
	default:
		return ""
	}
}
