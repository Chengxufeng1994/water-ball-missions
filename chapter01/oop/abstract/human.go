package abstract

import (
	"bufio"
	"fmt"
	"os"
)

type Human struct {
	id string
}

var _ Player = (*Human)(nil)

func NewHuman(id string) *Human {
	return &Human{
		id: id,
	}
}

func (h *Human) ID() string {
	return h.id
}

// MakeDecide implements Player.
func (h *Human) MakeDecide() Decision {
	scanner := bufio.NewScanner(os.Stdout)
	msg := "Paper, Scissors, or Stone?"
	fmt.Fprintln(os.Stdout, msg)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	text := scanner.Text()
	return Decision(text)
}
