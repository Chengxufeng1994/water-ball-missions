package model

import "github.com/Chengxufeng1994/water-ball-missions/chapter03/shortcut_key_mechanism/internal/command"

type Keyboard struct {
	Keys map[Key]command.Command
}

func NewKeyboard() *Keyboard {
	keys := make(map[Key]command.Command)
	for key, _ := range KeyMaps {
		keys[key] = nil
	}

	return &Keyboard{
		Keys: keys,
	}
}

func (kb *Keyboard) BindKey(key Key, cmd command.Command) {
	kb.Keys[key] = cmd
}

func (kb *Keyboard) GetKey(key Key) command.Command {
	return kb.Keys[key]
}

func (kb *Keyboard) Reset() {
	keys := make(map[Key]command.Command)
	for key, _ := range KeyMaps {
		keys[key] = nil
	}
	kb.Keys = keys
}
