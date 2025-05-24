package borderstyle

type DefaultBorderStyle struct{}

var _ BorderStyle = (*DefaultBorderStyle)(nil)

func NewDefaultBorderStyle() *DefaultBorderStyle { return &DefaultBorderStyle{} }

func (d *DefaultBorderStyle) TopLeftCorner() string     { return "" }
func (d *DefaultBorderStyle) TopRightCorner() string    { return "" }
func (d *DefaultBorderStyle) BottomLeftCorner() string  { return "" }
func (d *DefaultBorderStyle) BottomRightCorner() string { return "" }
func (d *DefaultBorderStyle) HorizontalEdge() string    { return "" }
func (d *DefaultBorderStyle) VerticalEdge() string      { return "" }
