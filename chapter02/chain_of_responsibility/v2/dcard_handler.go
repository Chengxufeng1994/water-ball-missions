package main

import "fmt"

type DcardMessageHandler struct {
	next IMessageHandler
}

var _ IMessageHandler = (*DcardMessageHandler)(nil)

func NewDcardMessageHandler(next IMessageHandler) *DcardMessageHandler {
	return &DcardMessageHandler{
		next: next,
	}
}

func (h *DcardMessageHandler) Handle(message Message) {
	if message.EventName == "dcard" {
		fmt.Println("dcard")
	} else if h.next != nil {
		h.next.Handle(message)
	}
}
