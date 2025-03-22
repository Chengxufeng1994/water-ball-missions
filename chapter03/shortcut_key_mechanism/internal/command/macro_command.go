package command

type MacroCommand struct {
	Commands []Command
}

var _ Command = (*MacroCommand)(nil)

func NewMacroCommand(commands []Command) *MacroCommand {
	return &MacroCommand{
		Commands: commands,
	}
}

func (m *MacroCommand) Name() string {
	return "MacroCommand"
}

func (m *MacroCommand) Execute() {
	for i := 0; i < len(m.Commands); i++ {
		m.Commands[i].Execute()
	}
}

func (m *MacroCommand) Undo() {
	for i := len(m.Commands) - 1; i >= 0; i-- {
		m.Commands[i].Undo()
	}
}
