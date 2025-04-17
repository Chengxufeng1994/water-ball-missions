package fsm

type Trigger interface {
	Match() bool
}
