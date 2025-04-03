package main

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/action"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/onepunchhandler"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/unit"
)

func main() {
	// Initialize the troops
	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 500, 500, 100, []domain.Action{}...),
	// 	unit.NewAIUnit(2, "WaterTA", 200, 200, 70, []domain.Action{}...),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(3, "Slime1", 200, 90, 50, []domain.Action{}...),
	// 	unit.NewAIUnit(4, "Slime1", 200, 90, 50, []domain.Action{}...),
	// 	unit.NewAIUnit(5, "Slime1", 200, 9000, 50, []domain.Action{}...),
	// )

	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 300, 500, 100, []domain.Action{action.NewFireBall(), action.NewWaterBall()}...),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(2, "Slime1", 200, 60, 49, []domain.Action{action.NewFireBall()}...),
	// 	unit.NewAIUnit(3, "Slime2", 200, 200, 50, []domain.Action{action.NewFireBall(), action.NewWaterBall()}...),
	// )

	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 500, 500, 40),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(2, "Slime1", 100, 100, 30, []domain.Action{action.NewSelfHealing()}...),
	// )
	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 400, 9999, 30, []domain.Action{action.NewPetrochemical()}...),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(2, "攻擊力超強的BOSS", 270, 9999, 399, []domain.Action{action.NewPetrochemical()}...),
	// )
	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 1000, 500, 0, []domain.Action{action.NewPoison()}...),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(2, "Slime1", 120, 90, 50, []domain.Action{}...),
	// 	unit.NewAIUnit(3, "Slime2", 120, 90, 50, []domain.Action{}...),
	// 	unit.NewAIUnit(4, "Slime3", 120, 9000, 50, []domain.Action{}...),
	// )
	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 999999, 10000, 30, []domain.Action{action.NewSelfExplosion()}...),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(2, "Slime1", 1000, 0, 99, []domain.Action{}...),
	// )

	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 500, 10000, 30, []domain.Action{action.NewCheerUp()}...),
	// 	unit.NewAIUnit(2, "Servant1", 1000, 0, 0, []domain.Action{}...),
	// 	unit.NewAIUnit(3, "Servant2", 1000, 0, 0, []domain.Action{}...),
	// 	unit.NewAIUnit(4, "Servant3", 1000, 0, 0, []domain.Action{}...),
	// 	unit.NewAIUnit(5, "Servant4", 1000, 0, 0, []domain.Action{}...),
	// 	unit.NewAIUnit(6, "Servant5", 1000, 0, 0, []domain.Action{}...),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(7, "Slime1", 500, 0, 0, []domain.Action{}...),
	// )

	// troopAlly := domain.NewTroop(1, "Ally",
	// 	unit.NewHeroUnit(1, "英雄", 300, 10000, 100, []domain.Action{action.NewCurse()}...),
	// 	unit.NewAIUnit(2, "Ally", 600, 100, 30, []domain.Action{action.NewCurse(), action.NewCurse()}...),
	// )
	// troopEnemy := domain.NewTroop(2, "Enemy",
	// 	unit.NewAIUnit(3, "Slime1", 200, 999, 50, []domain.Action{}...),
	// 	unit.NewAIUnit(4, "Slime2", 200, 999, 100, []domain.Action{}...),
	// )

	onePunchHandler := onepunchhandler.NewHighHpOnePunchHandler(
		onepunchhandler.NewPoisonPetrochemicalOnePunchHandler(
			onepunchhandler.NewCheerUpOnePunchHandler(
				onepunchhandler.NewNormalOnePunchHandler(nil),
			)),
	)

	troopAlly := domain.NewTroop(1, "Ally",
		unit.NewHeroUnit(1, "英雄", 1000, 10000, 100, []domain.Action{action.NewOnePunch(onePunchHandler), action.NewPoison(), action.NewPetrochemical(), action.NewCheerUp()}...),
	)
	troopEnemy := domain.NewTroop(2, "Enemy",
		unit.NewAIUnit(2, "Slime1", 601, 0, 0, []domain.Action{}...),
		unit.NewAIUnit(3, "Slime2", 241, 0, 0, []domain.Action{}...),
		unit.NewAIUnit(4, "Slime3", 101, 999, 0, []domain.Action{action.NewOnePunch(onePunchHandler), action.NewOnePunch(onePunchHandler), action.NewCheerUp()}...),
	)

	// Initialize the game
	rpg, err := domain.NewRPG([]*domain.Troop{troopAlly, troopEnemy}...)
	if err != nil {
		panic(err)
	}

	// Start the game
	rpg.Start()
}
