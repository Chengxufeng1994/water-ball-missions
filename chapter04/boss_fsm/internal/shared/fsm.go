package shared

type FiniteStateMachine interface {
	GetCurrentState() State
	ProcessEvent(event *Event, ctx Context)
}
