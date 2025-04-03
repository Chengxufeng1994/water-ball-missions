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
		round:             0,
		oppositeTroopMaps: oppositeTroopMaps,
	}, nil
}

func (rpg *RPG) Start() {
	for rpg.CheckVictory() {
		rpg.NextTurn()
		rpg.round++
	}

	rpg.AnnounceWinner()
}

func (rpg *RPG) NextTurn() {
	for _, unit := range rpg.troops[rpg.currentTroop].GetAliveUnits() {
		unit.TakeTurn(rpg)
	}

	rpg.currentTroop = (rpg.currentTroop + 1) % len(rpg.troops)
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
				fmt.Println("敵人獲勝了!")
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
