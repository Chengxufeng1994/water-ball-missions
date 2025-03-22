package command

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/core"
)

type DisconnectTelecomCommand struct {
	telecom core.TelecomInterface
}

var _ Command = (*DisconnectTelecomCommand)(nil)

func NewDisconnectTelecomCommand(telecom core.TelecomInterface) *DisconnectTelecomCommand {
	return &DisconnectTelecomCommand{telecom: telecom}
}

func (t *DisconnectTelecomCommand) Name() string {
	return "DisconnectTelecom"
}

func (t *DisconnectTelecomCommand) Execute() {
	t.telecom.Disconnect()
}

func (t *DisconnectTelecomCommand) Undo() {
	t.telecom.Connect()
}
