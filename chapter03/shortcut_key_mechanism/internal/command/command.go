package command

type Command interface {
	Name() string
	Execute()
	Undo()
}

type CommandRegistry struct {
	Commands map[string]Command
	Order    []string
}

var Registry *CommandRegistry

func init() {
	if Registry == nil {
		Registry = &CommandRegistry{
			Commands: make(map[string]Command),
			Order:    make([]string, 0),
		}
	}
}

func (r *CommandRegistry) Register(cmd Command) {
	if _, ok := r.Commands[cmd.Name()]; ok {
		return
	}
	r.Commands[cmd.Name()] = cmd
	r.Order = append(r.Order, cmd.Name())
}

func (registry *CommandRegistry) GetCommand(commandName string) Command {
	return registry.Commands[commandName]
}

func (registry *CommandRegistry) GetCommands() []string {
	return registry.Order
}
