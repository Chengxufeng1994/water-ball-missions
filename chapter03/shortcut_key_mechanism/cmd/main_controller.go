package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/command"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/model"
)

type MainController struct {
	keyboard    *model.Keyboard
	registry    *command.CommandRegistry
	undoHistory []command.Command
	redoHistory []command.Command
}

func NewMainController(keyboard *model.Keyboard) MainController {
	return MainController{
		keyboard:    keyboard,
		registry:    command.Registry,
		undoHistory: make([]command.Command, 0),
		redoHistory: make([]command.Command, 0),
	}
}

func (mc *MainController) Press(key model.Key) {
	cmd := mc.keyboard.GetKey(key)
	if cmd == nil {
		fmt.Println("尚未綁定:", key)
		return
	}
	cmd.Execute()
	mc.undoHistory = append(mc.undoHistory, cmd)
	mc.redoHistory = make([]command.Command, 0)
}

func (mc *MainController) Undo() {
	if len(mc.undoHistory) > 0 {
		cmd := mc.undoHistory[len(mc.undoHistory)-1]
		cmd.Undo()
		mc.redoHistory = append(mc.redoHistory, cmd)
		mc.undoHistory = mc.undoHistory[:len(mc.undoHistory)-1]
	}
}

func (mc *MainController) Redo() {
	if len(mc.redoHistory) > 0 {
		cmd := mc.redoHistory[len(mc.redoHistory)-1]
		cmd.Execute()
		mc.undoHistory = append(mc.undoHistory, cmd)
		mc.redoHistory = mc.redoHistory[:len(mc.redoHistory)-1]
	}
}

func (mc *MainController) BindKey() {
	fmt.Printf("設置巨集指令 (y/n): ")

	var isMacro bool
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "y" {
		isMacro = true
	}

	fmt.Print("Key: ")
	reader = bufio.NewReader(os.Stdin)
	keyInput, _ := reader.ReadString('\n')
	key := model.Key(strings.ToUpper(strings.TrimSpace(keyInput)))

	if isMacro {
		mc._bindMacro(key)
	} else {
		mc._bindKey(key)
	}

}

func (mc *MainController) _bindKey(key model.Key) {
	fmt.Printf("要將哪一道指令設置到快捷鍵 %s 上:\n", key)
	commandNames := mc.registry.GetCommands()
	for i, cmdName := range commandNames {
		fmt.Printf("(%d) %s\n", i, cmdName)
	}

	reader := bufio.NewReader(os.Stdin)
	cmdInput, _ := reader.ReadString('\n')
	cmdIndex, _ := strconv.Atoi(strings.TrimSpace(cmdInput))

	if cmdIndex >= 0 && cmdIndex < len(commandNames) {
		cmd := mc.registry.GetCommand(commandNames[cmdIndex])
		if cmd != nil {
			mc.keyboard.BindKey(key, cmd)
		}
	}
}

func (mc *MainController) _bindMacro(key model.Key) {
	fmt.Printf("要將哪些指令設置成快捷鍵 %s 的巨集（輸入多個數字，以空白隔開: \n", key)
	commandNames := mc.registry.GetCommands()
	for i, cmdName := range commandNames {
		fmt.Printf("(%d) %s\n", i, cmdName)
	}

	reader := bufio.NewReader(os.Stdin)
	cmdInput, _ := reader.ReadString('\n')
	cmdIndexStrs := strings.Split(strings.TrimSpace(cmdInput), " ")
	cmdIndexes := make([]int, 0)
	for _, cmdIndexStr := range cmdIndexStrs {
		cmdIndex, _ := strconv.Atoi(cmdIndexStr)
		cmdIndexes = append(cmdIndexes, cmdIndex)
	}

	commands := make([]command.Command, 0)
	for _, cmdIndex := range cmdIndexes {
		if cmdIndex >= 0 && cmdIndex < len(commandNames) {
			cmd := mc.registry.GetCommand(commandNames[cmdIndex])
			if cmd != nil {
				commands = append(commands, cmd)
			}
		}
	}

	mc.keyboard.BindKey(key, command.NewMacroCommand(commands))
}
