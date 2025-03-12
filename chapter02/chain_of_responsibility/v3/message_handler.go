package main

type IMessageHandler interface {
	Handle(Message)
	Match(Message) bool
	DoHandling(Message)
}
