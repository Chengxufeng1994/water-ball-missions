package shared

type State interface {
	EntryState()
	ExitState()
	GenerateMessage(channelType, authorId string) string
	SetContext(ctx Context)
}
