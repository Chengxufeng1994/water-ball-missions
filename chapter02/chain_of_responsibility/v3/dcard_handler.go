package main

type DcardMessageHandler struct {
	*BaseHandler
}

var _ IMessageHandler = (*DcardMessageHandler)(nil)

func NewDcardMessageHandler(next IMessageHandler) *DcardMessageHandler {
	return &DcardMessageHandler{
		BaseHandler: NewBaseHandler(next),
	}
}

func (h *DcardMessageHandler) Match(message Message) bool {
	return message.EventName == "dcard"
}

func (h *DcardMessageHandler) DoHandling(Message) {
	panic("unimplemented")
}
