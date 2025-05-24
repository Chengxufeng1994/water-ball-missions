package borderstyle

type PrettyBorderStyle struct{}

var _ BorderStyle = (*PrettyBorderStyle)(nil)

func NewPrettyBorderStyle() *PrettyBorderStyle { return &PrettyBorderStyle{} }

func (p *PrettyBorderStyle) TopLeftCorner() string     { return "-" }
func (p *PrettyBorderStyle) BottomLeftCorner() string  { return "-" }
func (p *PrettyBorderStyle) TopRightCorner() string    { return "-" }
func (p *PrettyBorderStyle) BottomRightCorner() string { return "-" }
func (b *PrettyBorderStyle) HorizontalEdge() string    { return "-" }
func (b *PrettyBorderStyle) VerticalEdge() string      { return "|" }
