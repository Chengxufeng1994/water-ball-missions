package command

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/core"
)

type ConnectTelecomCommand struct {
	telecom core.TelecomInterface
}

var _ Command = (*ConnectTelecomCommand)(nil)

func NewConnectTelecomCommand(telecom core.TelecomInterface) *ConnectTelecomCommand {
	return &ConnectTelecomCommand{telecom: telecom}
}

func (t *ConnectTelecomCommand) Name() string {
	return "ConnectTelecom"
}

func (t *ConnectTelecomCommand) Execute() {
	t.telecom.Connect()
}

func (t *ConnectTelecomCommand) Undo() {
	t.telecom.Disconnect()
}
