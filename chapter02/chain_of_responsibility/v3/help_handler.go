package main

import "fmt"

type HelpMessageHandler struct {
	*BaseHandler
}

var _ IMessageHandler = (*HelpMessageHandler)(nil)

func NewHelpMessageHandler(next IMessageHandler) *HelpMessageHandler {
	return &HelpMessageHandler{
		BaseHandler: NewBaseHandler(next),
	}
}

func (h *HelpMessageHandler) Handle(message Message) {
	if h.Match(message) {
		h.DoHandling(message)
	} else if h.next != nil {
		h.next.Handle(message)
	}
}

func (h *HelpMessageHandler) Match(message Message) bool {
	return message.EventName == "help"
}

func (h *HelpMessageHandler) DoHandling(message Message) {
	fmt.Println("help")
}
