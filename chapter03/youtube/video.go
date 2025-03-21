package main

type Video struct {
	Title       string
	Description string
	Length      int
	Likes       int
}

func NewVideo(title, description string, length int) Video {
	return Video{
		Title:       title,
		Description: description,
		Length:      length,
		Likes:       0,
	}
}

func (v *Video) Like() {
	// Do something.
	v.Likes++
}
