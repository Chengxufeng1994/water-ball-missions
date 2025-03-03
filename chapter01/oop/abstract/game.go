package abstract

import "fmt"

type Game struct {
	player1 Player
	player2 Player
	rule    map[Decision]Decision
}

func NewGame(player1, player2 Player) *Game {
	rule := map[Decision]Decision{
		Paper:    Stone,
		Scissors: Paper,
		Stone:    Scissors,
	}
	return &Game{
		player1: player1,
		player2: player2,
		rule:    rule,
	}
}

func (g *Game) Start() {
	player1Decide := g.player1.MakeDecide()
	player2Decide := g.player2.MakeDecide()

	fmt.Println(player1Decide, player2Decide)

	if player1Decide == player2Decide {
		fmt.Println("deal!")
		return
	}
	if g.rule[player1Decide] == player2Decide {
		fmt.Printf("%s win\n", g.player1.ID())
	} else {
		fmt.Printf("%s win\n", g.player2.ID())
	}
}
