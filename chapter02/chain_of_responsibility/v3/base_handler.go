package main

type BaseHandler struct {
	next IMessageHandler
}

var _ IMessageHandler = (*BaseHandler)(nil)

func NewBaseHandler(next IMessageHandler) *BaseHandler {
	return &BaseHandler{
		next: next,
	}
}

// Handle implements IMessageHandler.
func (base *BaseHandler) Handle(message Message) {
	if base.Match(message) {
		base.DoHandling(message)
	} else if base.next != nil {
		base.next.Handle(message)
	}
}

// Match implements IMessageHandler.
func (b *BaseHandler) Match(Message) bool {
	return false
}

// DoHandling implements IMessageHandler.
func (b *BaseHandler) DoHandling(Message) {
	panic("unimplemented")
}
