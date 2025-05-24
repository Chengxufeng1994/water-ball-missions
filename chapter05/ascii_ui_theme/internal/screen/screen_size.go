package screen

type ScreenSize struct {
	Width  int
	Height int
}

func NewScreenSize(width, height int) *ScreenSize {
	return &ScreenSize{
		Width:  width,
		Height: height,
	}
}
