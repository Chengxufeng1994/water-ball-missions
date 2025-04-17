package shared

type EventListener interface {
	OnEvent(event *Event)
}

type EventHandler interface {
	HandleEvent(event *Event, ctx Context)
}
