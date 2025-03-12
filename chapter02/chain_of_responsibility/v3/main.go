package main

func main() {
	handler := NewHelpMessageHandler(NewDcardMessageHandler(nil))
	handler.Handle(Message{EventName: "help"})
}
