package borderstyle

type BorderStyle interface {
	TopLeftCorner() string
	TopRightCorner() string
	BottomLeftCorner() string
	BottomRightCorner() string
	HorizontalEdge() string
	VerticalEdge() string
}
