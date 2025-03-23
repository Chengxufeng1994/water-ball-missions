package domain

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidMove = errors.New("invalid move")

const (
	FULL_HP         = 300
	CHARACTER_POWER = 999
)

type Character struct {
	BasedRole
}

var _ Role = (*Character)(nil)

func NewCharacter(game *AdventureGame, x, y int, symbol Symbol) *Character {
	return &Character{
		BasedRole: NewBasedRole(
			game, x, y, symbol, FULL_HP, CHARACTER_POWER,
		),
	}
}

func (c *Character) TakeTurn(game *AdventureGame) {
	for {
		fmt.Println("請選擇行動: ↑(W), ↓(S), ←(A), →(D), F(攻擊)")
		var input string
		fmt.Scanln(&input)
		action := strings.ToLower(input)
		fmt.Println("你選擇了:", action)

		switch action {
		case "w", "s", "a", "d":
			if err := c.moveAction(game, action); err != nil {
				fmt.Println(err)
				continue
			}
			return
		case "f":
			if err := c.attackAction(game); err != nil {
				fmt.Println(err)
				continue
			}
			return
		}
	}
}

func (c *Character) moveAction(game *AdventureGame, action string) error {
	directions := map[string][2]int{
		"w": {0, -1},
		"s": {0, 1},
		"a": {-1, 0},
		"d": {1, 0},
	}
	direction := directions[action]

	origX, origY := c.X, c.Y
	newX, newY := origX+direction[0], origY+direction[1]

	game.Move(c, origX, origY, newX, newY)
	c.Move(newX, newY)
	c.SetSymbol(SymbolMaps[action])

	return nil
}

func (c *Character) attackAction(game *AdventureGame) error {
	monsters, err := game.FindMonstersBasedOnDirection(c.X, c.Y, c.Symbol)
	if err != nil {
		return err
	}

	fmt.Println("Found", len(monsters), "monsters to attack")

	for _, monster := range monsters {
		c.Attack(monster)
	}

	return nil
}

func (c Character) String() string {
	x, y := c.Position()
	format := "Character HP: %d, State: %s, Direction: %s, Current Position: %d, %d"
	return fmt.Sprintf(format, c.hp, c.state.StateType, c.Symbol, x, y)
}
