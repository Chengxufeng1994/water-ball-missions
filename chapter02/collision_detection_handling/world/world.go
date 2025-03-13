package world

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/collisionhandler"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/collision_detection_handling/sprite"
)

type World struct {
	Length             int
	initialNumOfSprite int
	Sprites            []sprite.ISprite
	CollisionHandler   collisionhandler.ICollisionHandler
}

func NewWorld(length int, initialNumOfSprite int, collisionHandler collisionhandler.ICollisionHandler) World {
	return World{
		Length:             length,
		initialNumOfSprite: initialNumOfSprite,
		Sprites:            make([]sprite.ISprite, length),
		CollisionHandler:   collisionHandler,
	}
}

func (w *World) Prepare() {
	fmt.Println("World prepare...")
	spriteTypes := []sprite.SpriteType{sprite.SpriteTypeHero, sprite.SpriteTypeWater, sprite.SpriteTypeFire}

	for i := 0; i < w.initialNumOfSprite; i++ {
		pos := rand.Intn(w.Length)
		if w.Sprites[pos] == nil {
			typeIndex := rand.Intn(len(spriteTypes))
			var spr sprite.ISprite
			if spriteTypes[typeIndex] == sprite.SpriteTypeHero {
				spr = sprite.NewHero(pos)
			} else if spriteTypes[typeIndex] == sprite.SpriteTypeWater {
				spr = sprite.NewWater(pos)
			} else if spriteTypes[typeIndex] == sprite.SpriteTypeFire {
				spr = sprite.NewFire(pos)
			}
			w.Sprites[pos] = spr
		}
	}
}

func (w *World) PrintSprites() {
	for i := range w.Sprites {
		fmt.Printf("position %d: %v\n", i, w.Sprites[i])
	}
}

func (w *World) Run() {
	fmt.Println("World running...")
	reader := bufio.NewReader(os.Stdin)
	for {
		w.PrintSprites()
		fmt.Print(">> ") // 提示符
		input := get(reader)

		if !shouldContinue(input) {
			os.Exit(0)
		}

		positions := strings.Split(input, " ")
		if len(positions) != 2 {
			fmt.Println("Invalid position, please try again")
			continue
		}

		source, _ := strconv.Atoi(positions[0])
		target, _ := strconv.Atoi(positions[1])
		if source < 0 || source >= w.Length || target < 0 || target >= w.Length {
			fmt.Println("Invalid position, please try again")
			continue
		}

		fmt.Printf("move source(c1) %d to target(c2) %d\n", source, target)
		w.Handle(w.Sprites[source], w.Sprites[target])
	}
}

func (w *World) Handle(source sprite.ISprite, target sprite.ISprite) {
	w.CollisionHandler.Handle(source, target, w)
}

func (w *World) RemoveSprite(s sprite.ISprite) {
	w.Sprites[s.Position()] = nil
}

func (w *World) RemoveSprites(sprites ...sprite.ISprite) {
	for _, s := range sprites {
		w.RemoveSprite(s)
	}
}

func (w *World) MoveSprite(source sprite.ISprite, target sprite.ISprite) {
	sourcePos := source.Position()
	targetPos := target.Position()

	source.Move(targetPos)
	target.Move(sourcePos)

	w.Sprites[sourcePos] = target
	w.Sprites[targetPos] = source
}

func get(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func shouldContinue(text string) bool {
	if strings.EqualFold("exit", text) {
		fmt.Println("Bye!")
		return false
	}
	return true
}
