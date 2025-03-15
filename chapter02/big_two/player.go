package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/card"
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/big_two/cardpattern"
)

type Player struct {
	Name            string
	Hand            []card.Card
	ValidateHandler cardpattern.ICardPatternValidateHandler
}

func NewPlayer(
	validateHandler cardpattern.ICardPatternValidateHandler,
) *Player {
	return &Player{
		Name:            "",
		Hand:            make([]card.Card, 0),
		ValidateHandler: validateHandler,
	}
}

func (p *Player) NamingHimself() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	name := strings.TrimSpace(text)
	p.Name = name

}

func (p *Player) DrawCardIntoHand(card card.Card) {
	p.Hand = append(p.Hand, card)
}

func (p *Player) SortHand() {
	sort.Slice(p.Hand, func(i, j int) bool {
		if p.Hand[i].Rank == p.Hand[j].Rank {
			return p.Hand[i].Suit < p.Hand[j].Suit
		}
		return p.Hand[i].Rank < p.Hand[j].Rank
	})
}

func (p *Player) ListHand() {
	p.SortHand()

	for i := range p.Hand {
		fmt.Printf("%-6d", i)
	}
	fmt.Println()

	for _, card := range p.Hand {
		fmt.Printf("%-6s", card)
	}
	fmt.Println()
}

func (p *Player) Play(game *BigTwo) cardpattern.ICardPattern {
	fmt.Printf("輪到 %s 了\n", p.Name)
	p.ListHand()

	var cardPattern cardpattern.ICardPattern
	isValid := false
	for !isValid {
		reader := bufio.NewReader(os.Stdin)
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		// 喊 PASS 放棄出牌機會，不打任何牌。
		if action == "-1" {
			if game.TopPlay == nil {
				fmt.Println("你不能在新的回合中喊 PASS")
				continue
			} else {
				fmt.Printf("玩家 %s 喊 PASS\n", p.Name)
				game.PassCount++
				return nil
			}
		}

		// 打出一副比頂牌還大的合法牌型。此新打出來的牌型便成為新的頂牌。
		list := strings.Split(action, " ")
		if len(list) == 0 {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			continue
		}

		indexes := make([]card.Card, 0, len(list))
		for _, item := range list {
			i, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println("此牌型不合法，請再嘗試一次。")
				continue
			}

			if i < 0 || i >= len(p.Hand) {
				fmt.Println("此牌型不合法，請再嘗試一次。")
				continue
			}

			indexes = append(indexes, p.Hand[i])
		}

		// 判斷打出的牌型是否合法
		isValid, cardPattern = p.ValidateHandler.Validate(indexes)
		if isValid == false || cardPattern == nil {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			continue
		}

		// 合法的牌型應該是比頂牌還大的牌型。
		isValid = p.CheckRules(game, cardPattern)
		if !isValid {
			fmt.Println("此牌型不合法，請再嘗試一次。")
			continue
		}
	}

	// 玩家 <玩家的名字> 打出了 <牌型名稱> <花色>[<數字>] <花色>[<數字>] <花色>[<數字>] ...
	fmt.Printf("玩家 %s 出的牌型是%s\n", p.Name, cardPattern.String())
	p.RemoveCards(cardPattern)
	game.PassCount = 0
	return cardPattern
}

func (p *Player) CheckRules(bigTwo *BigTwo, cardPattern cardpattern.ICardPattern) bool {

	if bigTwo.Round == 1 && bigTwo.TopPlay == nil {
		for _, c := range cardPattern.ListCard() {
			if c.Rank == card.Three && c.Suit == card.Club {
				return true
			}
		}
		return false
	}

	if bigTwo.TopPlay == nil {
		return true
	}

	return cardPattern.CompareTo(bigTwo.TopPlay)
}

func (p *Player) RemoveCards(cardPattern cardpattern.ICardPattern) {
	for _, card := range cardPattern.ListCard() {
		for i, c := range p.Hand {
			if c == card {
				p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
				break
			}
		}
	}
}

func (p Player) HandSize() int {
	return len(p.Hand)
}

func (p Player) HaveClubThree() bool {
	for _, c := range p.Hand {
		if c.Rank == card.Three && c.Suit == card.Club {
			return true
		}
	}

	return false
}
