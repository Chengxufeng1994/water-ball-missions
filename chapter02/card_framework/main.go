package main

func main() {
	players := []*Player{NewPlayer(1), NewPlayer(2), NewPlayer(3), NewPlayer(4)}
	game := NewShowdownGame(players)
	game.Initialize()
	game.StartGame()
}
