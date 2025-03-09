package card

type Color int

const (
	ColorBlue Color = iota
	ColorRed
	ColorYellow
	ColorGreen
)

func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "blue"
	case ColorRed:
		return "red"
	case ColorYellow:
		return "yellow"
	case ColorGreen:
		return "green"
	}

	return ""
}
