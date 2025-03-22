package model

import "fmt"

type Telecom struct {
}

func NewTelecom() *Telecom {
	return &Telecom{}
}

func (t *Telecom) Connect() {
	fmt.Println("The telecom has connected.")
}

func (t *Telecom) Disconnect() {
	fmt.Println("The telecom has disconnected.")
}
