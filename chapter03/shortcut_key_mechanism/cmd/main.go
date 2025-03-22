package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/command"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/model"
)

func main() {
	registry := command.Registry

	tank := model.NewTank()
	telecom := model.NewTelecom()
	keyboard := model.NewKeyboard()
	moveTankForwardCmd := command.NewMoveTankForward(tank)
	moveTankBackwardCmd := command.NewMoveTankBackward(tank)
	connectTelecomCmd := command.NewConnectTelecomCommand(telecom)
	disconnectTelecomCmd := command.NewDisconnectTelecomCommand(telecom)
	resetMainControlKeyboardCmd := command.NewResetMainControlKeyboard(keyboard)

	registry.Register(moveTankForwardCmd)
	registry.Register(moveTankBackwardCmd)
	registry.Register(connectTelecomCmd)
	registry.Register(disconnectTelecomCmd)
	registry.Register(resetMainControlKeyboardCmd)

	mainController := NewMainController(keyboard)

	for {
		fmt.Printf("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵: ")

		input := bufio.NewReader(os.Stdin)
		text, _ := input.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text {
		case "1":
			mainController.BindKey()
		case "2":
			mainController.Undo()
		case "3":
			mainController.Redo()
		case "q":
			os.Exit(0)
		default:
			key := model.Key(strings.ToUpper(text))
			mainController.Press(key)
		}
	}
}
