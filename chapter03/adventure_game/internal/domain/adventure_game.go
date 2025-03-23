package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

type AdventureGame struct {
	Width             int
	Height            int
	NumberOfMonsters  int
	NumberOfTreasures int
	NumberOfObstacles int
	Round             int
	GameMap           [][]MapObject
	Character         *Character
	Monsters          []*Monster
	Treasures         []*Treasure
	Obstacles         []*Obstacle
}

func NewAdventureGame(width, height, numberOfMonsters, numberOfTreasures, numberOfObstacles int) AdventureGame {
	gameMap := make([][]MapObject, height)
	for i := range height {
		gameMap[i] = make([]MapObject, width)
	}

	return AdventureGame{
		Width:             width,
		Height:            height,
		NumberOfMonsters:  numberOfMonsters,
		NumberOfTreasures: numberOfTreasures,
		NumberOfObstacles: numberOfObstacles,
		GameMap:           gameMap,
		Character:         nil,
		Monsters:          make([]*Monster, 0),
		Treasures:         make([]*Treasure, 0),
		Obstacles:         make([]*Obstacle, 0),
	}
}

func (game *AdventureGame) GenerateMap() {
	// TODO: 假設主角初始位置都在 (0, 0), 隨機面向方向
	character := NewCharacter(game, 0, 0, GetRandomDirection())
	game.GameMap[0][0] = character
	game.Character = character

	monsterCount := 0
	for monsterCount < game.NumberOfMonsters {
		monster := NewMonster(game, 0, 0)
		game.PlaceMapObjectRandomly(monster)
		game.Monsters = append(game.Monsters, monster)
		monsterCount++
	}

	treasureCount := 0
	for treasureCount < game.NumberOfTreasures {
		treasure := GenerateRandomTreasure(0, 0)
		if treasure == nil {
			continue
		}
		game.PlaceMapObjectRandomly(treasure)
		game.Treasures = append(game.Treasures, treasure)
		treasureCount++
	}

	obstacleCount := 0
	for obstacleCount < game.NumberOfObstacles {
		obstacle := NewObstacle(0, 0)
		game.PlaceMapObjectRandomly(obstacle)
		game.Obstacles = append(game.Obstacles, obstacle)
		obstacleCount++
	}
}

func (game *AdventureGame) PlaceMapObjectRandomly(mapObject MapObject) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		x := r.Intn(game.Width)
		y := r.Intn(game.Height)
		if game.GameMap[x][y] == nil {
			mapObject.SetPosition(x, y)
			game.GameMap[x][y] = mapObject
			break
		}
	}
}

// 檢查新位置是否可移動
func (game *AdventureGame) IsMoveAllowed(x, y int) error {
	// 檢查是否超出邊界
	if x < 0 || x >= game.Width || y < 0 || y >= game.Height {
		return errors.New("out of bounds")
	}

	// 檢查障礙物與怪物
	mapObject := game.GameMap[x][y]
	if mapObject != nil {

		if mapObject.GetSymbol() == SymbolObstacle {
			return errors.New("cannot move to obstacle")
		}

		if mapObject.GetSymbol() == SymbolMonster {
			return errors.New("cannot move to monster")
		}

		if mapObject.GetSymbol() == DirectionUp ||
			mapObject.GetSymbol() == DirectionDown ||
			mapObject.GetSymbol() == DirectionLeft ||
			mapObject.GetSymbol() == DirectionRight {
			return errors.New("cannot move to character")
		}
	}

	return nil
}

func (game *AdventureGame) Move(role Role, srcX, srcY, dstX, dstY int) error {
	if err := game.IsMoveAllowed(dstX, dstY); err != nil {
		return err
	}

	for i, treasure := range game.Treasures {
		// 檢查是否有寶物
		treasureX, treasureY := treasure.Position()
		if dstX == treasureX && dstY == treasureY {
			treasure.ApplyTreasure(game.Character)
			game.Treasures = slices.Delete(game.Treasures, i, i+1)
			break
		}
	}

	game.UpdateMapObjectPosition(srcX, srcY, dstX, dstY)

	return nil
}

// 更新 MapObject 位置
func (game *AdventureGame) UpdateMapObjectPosition(origX, origY, newX, newY int) {
	game.GameMap[newX][newY] = game.GameMap[origX][origY]
	game.GameMap[origX][origY] = nil
}

// 移除 MapObject
func (game *AdventureGame) RemoveMapObject(x, y int) {
	defer func() {
		game.GameMap[x][y] = nil
	}()

	for i := range game.Monsters {
		monster := game.Monsters[i]
		monsterX, monsterY := monster.Position()
		if monsterX == x && monsterY == y {
			game.Monsters = slices.Delete(game.Monsters, i, i+1)
			return
		}
	}
}

func (game *AdventureGame) FindMonsterByPosition(x, y int) *Monster {
	for _, monster := range game.Monsters {
		if monster.X == x && monster.Y == y {
			return monster
		}
	}

	return nil
}

func (game *AdventureGame) FindMonstersBasedOnDirection(x, y int, directionSymbol Symbol) ([]*Monster, error) {
	monsters := make([]*Monster, 0)
	directionOffsets := map[Symbol][2]int{
		DirectionUp:    {0, -1},
		DirectionDown:  {0, 1},
		DirectionLeft:  {-1, 0},
		DirectionRight: {1, 0},
	}
	offset, exists := directionOffsets[directionSymbol]
	if !exists {
		return nil, errors.New("invalid direction symbol")
	}

	for {

		x, y = x+offset[0], y+offset[1]
		if x < 0 || x >= game.Width || y < 0 || y >= game.Height {
			break
		}

		mapObject := game.GameMap[x][y]
		if mapObject != nil {
			if mapObject.GetSymbol() == SymbolObstacle {
				break
			}
		}

		monster := game.FindMonsterByPosition(x, y)
		if monster != nil {
			monsters = append(monsters, monster)
		}

	}

	return monsters, nil
}

func (game *AdventureGame) FindRolesExcludePosition(excludeX, excludeY int) ([]Role, error) {
	var roles []Role
	for _, monster := range game.Monsters {
		if monster.X == excludeX && monster.Y == excludeY {
			continue
		}
		roles = append(roles, monster)
	}

	if game.Character.X != excludeX && game.Character.Y != excludeY {
		roles = append(roles, game.Character)
	}

	return roles, nil
}

func (game *AdventureGame) FindEmptyPosition() [][]int {
	availablePositions := make([][]int, 0)
	for x := 0; x < game.Width; x++ {
		for y := 0; y < game.Height; y++ {
			if game.GameMap[x][y] == nil {
				availablePositions = append(availablePositions, []int{x, y})
			}
		}
	}

	return availablePositions
}

func (game *AdventureGame) Prepare() {
	fmt.Println("準備中...")

	fmt.Println("地圖生成中...")
	game.GenerateMap()

	fmt.Println("準備完畢...")
}

func (game *AdventureGame) IsGameOver() bool {
	if len(game.Monsters) == 0 {
		fmt.Println("怪物全數消滅, 遊戲結束...")
		return true
	}

	if game.Character.HP() <= 0 {
		fmt.Println("角色死亡, 遊戲結束...")
		return true
	}

	return false
}

func (game *AdventureGame) Render() {
	const mapChar = SymbolEmpty

	grid := make([][]string, game.Height)
	for i := range grid {
		grid[i] = make([]string, game.Width)
		for j := range grid[i] {
			grid[i][j] = mapChar
		}
	}

	for _, row := range game.GameMap {
		for _, obj := range row {
			if obj != nil {
				x, y := obj.Position()
				grid[y][x] = string(obj.GetSymbol())
			}
		}
	}

	for _, row := range grid {
		fmt.Println(strings.Join(row, " "))
	}

	fmt.Println(game.Character)
}

func (game *AdventureGame) Start() {
	game.Prepare()
	fmt.Println("遊戲開始...")
	for !game.IsGameOver() {
		game.Render()

		game.Character.TakeTurn()

		for i := range game.Monsters {
			monster := game.Monsters[i]
			monster.TakeTurn()
		}
	}
}
