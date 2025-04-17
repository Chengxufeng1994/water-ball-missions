package community

import "fmt"

var (
	ErrInvalidPayload = fmt.Errorf("invalid payload")
	ErrCommunityEvent = fmt.Errorf("invalid community event")
)
