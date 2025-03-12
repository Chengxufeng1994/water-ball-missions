package main

import "fmt"

type HelpMessageHandler struct {
	next IMessageHandler
}

var _ IMessageHandler = (*HelpMessageHandler)(nil)

func NewHelpMessageHandler(next IMessageHandler) *HelpMessageHandler {
	return &HelpMessageHandler{
		next: next,
	}
}

func (h *HelpMessageHandler) Handle(message Message) {
	if message.EventName == "help" {
		fmt.Println("help")
	} else if h.next != nil {
		h.next.Handle(message)
	}
}
