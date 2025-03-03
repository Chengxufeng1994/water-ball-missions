package abstract

import "testing"

func TestGame(t *testing.T) {
	player1 := NewHuman("player1")
	player2 := NewAI("player2")
	game := NewGame(player1, player2)
	game.Start()
}
