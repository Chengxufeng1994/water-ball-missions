package shared

type Context interface {
	GetPrefix() string
	GetState() State
	GetFSM() FiniteStateMachine
	GetValue(key string) any
	SetValue(key string, value any)
}
