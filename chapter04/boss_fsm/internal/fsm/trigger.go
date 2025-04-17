package fsm

type Trigger interface {
	Match(event Event) bool
}
