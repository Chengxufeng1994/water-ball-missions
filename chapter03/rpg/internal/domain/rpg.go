package domain

import (
	"errors"
	"fmt"
)

const TROOP_SIZE = 2

var (
	ErrTooManyTroops = errors.New("too many troops")
	ErrTooLessTroops = errors.New("too less troops")
)

type RPG struct {
	troops            []*Troop
	currentTroop      int
	currentUnit       Unit
	round             int
	oppositeTroopMaps map[int]int
}

func NewRPG(troops ...*Troop) (*RPG, error) {
	if len(troops) < TROOP_SIZE {
		return nil, ErrTooLessTroops
	}

	if len(troops) > TROOP_SIZE {
		return nil, ErrTooManyTroops
	}

	allyTroopIndex := 0
	enemyTroopIndex := 1
	oppositeTroopMaps := map[int]int{
		allyTroopIndex:  enemyTroopIndex,
		enemyTroopIndex: allyTroopIndex,
	}

	return &RPG{
		troops:            troops,
		currentTroop:      0,
		currentUnit:       nil,
		round:             0,
		oppositeTroopMaps: oppositeTroopMaps,
	}, nil
}

func (rpg *RPG) Start() {
	isGameOver := false
	for !isGameOver {
		queue := rpg.BuildTurnQueue()

		for len(queue) > 0 {
			rpg.currentUnit = queue[0]
			queue = queue[1:]

			rpg.currentUnit.TakeTurn(rpg)

			// 檢查新增的 Unit
			newUnits := rpg.troops[rpg.currentTroop].GetNewUnitsAndClear() // 新增的單位
			if len(newUnits) > 0 {
				// 加入到當前的 turnQueue（尾端或其他策略）
				queue = append(queue, newUnits...)
			}

			if isGameOver = !rpg.CheckVictory(); isGameOver {
				break
			}
		}

		rpg.currentTroop = (rpg.currentTroop + 1) % len(rpg.troops)
		rpg.round++
	}

	rpg.AnnounceWinner()
}

func (rpg *RPG) BuildTurnQueue() []Unit {
	return rpg.troops[rpg.currentTroop].GetAliveUnits()
}

func (rpg *RPG) CheckVictory() bool {
	for _, troop := range rpg.troops {
		if !troop.IsHeroAlive() {
			return false
		}
		if troop.IsAnnihilated() {
			return false
		}
	}

	return true
}

func (rpg *RPG) AnnounceWinner() {
	for _, troop := range rpg.troops {
		if troop.IsHeroAlive() {
			if troop.Name == "Ally" {
				fmt.Println("你獲勝了!")
			} else {
				fmt.Println("你失敗了!")
			}
			break
		}
	}
}

func (rpg *RPG) GetCurrentTroop() *Troop {
	return rpg.troops[rpg.currentTroop]
}

func (rpg *RPG) GetOppositeTroop() *Troop {
	return rpg.troops[rpg.oppositeTroopMaps[rpg.currentTroop]]
}

func (rpg *RPG) GetCurrentUnitTroopUnits() []Unit {
	troop := rpg.GetCurrentTroop()
	units := make([]Unit, 0)
	for _, unit := range troop.Units {
		if unit.IsAlive() {
			if unit.GetID() == rpg.currentUnit.GetID() {
				continue
			}
			units = append(units, unit)
		}
	}

	return units
}

func (rpg *RPG) GetCurrentUnitOppositeTroopUnits() []Unit {
	troop := rpg.GetOppositeTroop()
	units := make([]Unit, 0)
	for _, unit := range troop.Units {
		if unit.IsAlive() {
			units = append(units, unit)
		}
	}

	return units
}

func (rpg *RPG) GetAllAliveUnits() []Unit {
	var units []Unit
	for _, troop := range rpg.troops {
		for _, unit := range troop.Units {
			if unit.IsAlive() {
				units = append(units, unit)
			}
		}
	}
	return units
}
