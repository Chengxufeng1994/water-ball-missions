package borderstyle

type BasicBorderStyle struct{}

var _ BorderStyle = (*BasicBorderStyle)(nil)

func NewBasicBorderStyle() *BasicBorderStyle { return &BasicBorderStyle{} }

func (b *BasicBorderStyle) TopLeftCorner() string     { return "+" }
func (b *BasicBorderStyle) TopRightCorner() string    { return "+" }
func (b *BasicBorderStyle) BottomLeftCorner() string  { return "+" }
func (b *BasicBorderStyle) BottomRightCorner() string { return "+" }
func (b *BasicBorderStyle) HorizontalEdge() string    { return "-" }
func (b *BasicBorderStyle) VerticalEdge() string      { return "|" }
