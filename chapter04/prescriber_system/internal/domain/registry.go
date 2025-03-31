package domain

type PrescribeHandlerRegistry interface {
	Register(name string, prescribeHandler PrescribeHandler)
	Find(name string) PrescribeHandler
}
