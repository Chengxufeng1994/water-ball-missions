package main

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/action"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/unit"
)

func main() {
	// Initialize the troops
	troopAlly := domain.NewTroop(1, "Ally",
		unit.NewHeroUnit("英雄", 300, 500, 100, []domain.Action{action.NewFireBall(), action.NewWaterBall()}...),
	)
	troopEnemy := domain.NewTroop(2, "Enemy",
		unit.NewAIUnit("Slime1", 200, 60, 49, []domain.Action{action.NewFireBall()}...),
		unit.NewAIUnit("Slime2", 200, 200, 50, []domain.Action{action.NewFireBall(), action.NewWaterBall()}...),
	)

	// Initialize the game
	rpg, err := domain.NewRPG([]*domain.Troop{troopAlly, troopEnemy}...)
	if err != nil {
		panic(err)
	}

	// Start the game
	rpg.Start()
}
