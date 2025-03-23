package domain

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidMove = errors.New("invalid move")

const FULL_HP = 300

type Character struct {
	*BasedRole
}

var _ Role = (*Character)(nil)

func NewCharacter(game *AdventureGame, x, y int, symbol Symbol) *Character {
	character := &Character{
		BasedRole: NewBasedRole(
			game, x, y, symbol, FULL_HP, MONSTER_HP,
		),
	}

	character.turnBehavior = character
	return character
}

func (c *Character) RoundAction() {
	for {
		fmt.Println("請選擇行動: ↑(W), ↓(S), ←(A), →(D), F(攻擊)")
		for k := range c.directions {
			fmt.Printf("%s: %v\n", k, c.directions[k])
		}

		var input string
		fmt.Scanln(&input)
		action := strings.ToLower(input)
		fmt.Println("你選擇了:", action)

		switch action {
		case "w", "s", "a", "d":
			if err := c.moveAction(action); err != nil {
				fmt.Println(err)
				continue
			}
			return
		case "f":
			if err := c.attackAction(); err != nil {
				fmt.Println(err)
				continue
			}
			return
		default:
			fmt.Println("無效的行動")
			continue
		}
	}
}

func (c *Character) moveAction(action string) error {
	direction := c.directions[action]

	origX, origY := c.X, c.Y
	newX, newY := origX+direction[0], origY+direction[1]

	if err := c.game.Move(c, origX, origY, newX, newY); err != nil {
		return err
	}
	c.Move(newX, newY)
	c.SetSymbol(SymbolMaps[action])

	return nil
}

func (c *Character) attackAction() error {
	if c.attack == AttackBehaviorOnLine {
		monsters, err := c.game.FindMonstersBasedOnDirection(c.X, c.Y, c.Symbol)
		if err != nil {
			return err
		}

		fmt.Println("Found", len(monsters), "monsters to attack")

		for _, monster := range monsters {
			c.Attack(monster)
		}
	} else if c.attack == AttackBehaviorFullMap {

		monsters, err := c.game.FindRolesExcludePosition(c.X, c.Y)
		if err != nil {
		}

		fmt.Println("Found", len(monsters), "monsters to attack")

		for _, monster := range monsters {
			c.Attack(monster)
		}
	}

	return nil
}

func (c Character) String() string {
	x, y := c.Position()
	format := "Character HP: %d, State: %s, Direction: %s, Current Position: %d, %d"
	return fmt.Sprintf(format, c.hp, c.state, c.Symbol, x, y)
}
