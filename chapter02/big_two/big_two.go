package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/cardpattern"
)

type BigTwo struct {
	deck         *Deck
	players      []*Player
	currentIndex int
	Round        int
	PassCount    int
	TopPlay      cardpattern.ICardPattern
	TopPlayer    *Player
	Winner       *Player
}

func NewBigTwo() *BigTwo {
	validateHandler := cardpattern.NewSingleCardPatternValidateHandler(
		cardpattern.NewPairCardPatternValidateHandler(
			cardpattern.NewStraightCardPatternValidateHandler(
				cardpattern.NewFullHouseCardPatternValidateHandler(nil),
			),
		),
	)

	players := make([]*Player, 0)

	for range 4 {
		players = append(players, NewPlayer(validateHandler))
	}

	return &BigTwo{
		players: players,
		deck:    NewPokerDeck(),
	}
}

func (bigTwo *BigTwo) Start() {
	fmt.Println("BigTwo Start...")
	// 洗牌階段
	bigTwo.deck.Shuffle()

	// 命名階段
	for _, player := range bigTwo.players {
		player.NamingHimself()
	}

	// 發牌階段
	for !bigTwo.deck.IsEmpty() {
		for _, player := range bigTwo.players {
			player.DrawCardIntoHand(bigTwo.deck.Draw())
		}
	}

	// 整理手牌階段
	for _, player := range bigTwo.players {
		player.SortHand()
	}

	// 找到梅花三的玩家
	bigTwo.currentIndex = bigTwo.FindFirstPlayer()

	// 開始進行遊戲
	bigTwo.Round = 1
LOOP:
	for {
		fmt.Printf("新的回合開始了\n")

		for {
			currentPlayer := bigTwo.players[bigTwo.currentIndex]
			cardPattern := currentPlayer.Play(bigTwo)
			if cardPattern != nil {
				bigTwo.TopPlay = cardPattern
				bigTwo.TopPlayer = currentPlayer
			}

			if currentPlayer.HandSize() == 0 {
				bigTwo.Winner = currentPlayer
				break LOOP
			}

			if bigTwo.PassCount == 3 {
				// 每一回合結束之後，檯面上的牌會被清空，意即，清空頂牌。
				bigTwo.TopPlay = nil
				// 上一回合的頂牌玩家將會成為新的回合的頂牌玩家，並且新的回合由頂牌玩家開始打牌，並且不能喊 PASS。舉例來說，如果上一回合結束時的頂牌玩家為玩家i
				// ，則在新的一回合時時，由玩家i開始打牌
				for i, player := range bigTwo.players {
					if player == bigTwo.TopPlayer {
						bigTwo.currentIndex = i
						break
					}
				}
				bigTwo.PassCount = 0
				break
			}

			bigTwo.currentIndex = (bigTwo.currentIndex + 1) % 4
		}

		bigTwo.Round++
	}

	fmt.Printf("遊戲結束，遊戲的勝利者為 %s\n", bigTwo.Winner.Name)
}

func (bigTwo *BigTwo) FindFirstPlayer() int {
	for i, player := range bigTwo.players {
		if player.HaveClubThree() {
			return i
		}
	}
	return -1
}
