package main

type IMessageHandler interface {
	Handle(Message)
}
