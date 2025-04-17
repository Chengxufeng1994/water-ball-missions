package content

type Speak struct {
	*BasedContent
}

var _ CommunityContent = (*Speak)(nil)

func NewSpeak(authorID string, content string) *Speak {
	return &Speak{
		BasedContent: NewBasedContent(authorID, content, nil),
	}
}
