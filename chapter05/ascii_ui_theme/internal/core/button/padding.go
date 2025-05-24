package button

type Padding struct {
	Top    int `json:"top"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
	Right  int `json:"right"`
}

func NewPadding(top, bottom, left, right int) *Padding {
	return &Padding{
		Top:    top,
		Bottom: bottom,
		Left:   left,
		Right:  right,
	}
}
