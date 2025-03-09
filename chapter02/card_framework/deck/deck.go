package deck

import (
	"math/rand"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/card_framework/card"
)

type IDeck interface {
	Shuffle()
	Draw() card.Card
	Add(card.Card)
	IsEmpty() bool
}

type BasedDeck struct {
	cards []card.Card
}

var _ IDeck = (*BasedDeck)(nil)

func NewBasedDeck(cards []card.Card) *BasedDeck {
	return &BasedDeck{cards: cards}
}

func (b *BasedDeck) Shuffle() {
	// fmt.Printf("before shuffle %v\n", b.cards)
	rand.Shuffle(len(b.cards), func(i, j int) {
		b.cards[i], b.cards[j] = b.cards[j], b.cards[i]
	})
	// fmt.Printf("after shuffle %v\n", b.cards)
}

func (b *BasedDeck) Draw() card.Card {
	card := b.cards[0]
	b.cards = b.cards[1:]
	return card
}

func (b *BasedDeck) Add(card card.Card) {
	b.cards = append(b.cards, card)
}

func (b *BasedDeck) IsEmpty() bool {
	return len(b.cards) == 0
}

func NewUnoDeck() *BasedDeck {
	colors := []card.Color{
		card.ColorBlue,
		card.ColorRed,
		card.ColorYellow,
		card.ColorGreen,
	}
	cards := make([]card.Card, 0)
	for _, color := range colors {
		for num := 1; num <= 10; num++ {
			cards = append(cards, card.NewUnoCard(color, num))
		}
	}
	return &BasedDeck{cards: cards}
}

func NewShowdownDeck() *BasedDeck {
	return &BasedDeck{
		cards: []card.Card{
			card.NewShowdownCard(card.Club, card.Ace),
			card.NewShowdownCard(card.Club, card.Two),
			card.NewShowdownCard(card.Club, card.Three),
			card.NewShowdownCard(card.Club, card.Four),
			card.NewShowdownCard(card.Club, card.Five),
			card.NewShowdownCard(card.Club, card.Six),
			card.NewShowdownCard(card.Club, card.Seven),
			card.NewShowdownCard(card.Club, card.Eight),
			card.NewShowdownCard(card.Club, card.Nine),
			card.NewShowdownCard(card.Club, card.Ten),
			card.NewShowdownCard(card.Club, card.Jack),
			card.NewShowdownCard(card.Club, card.Queen),
			card.NewShowdownCard(card.Club, card.King),
			card.NewShowdownCard(card.Diamond, card.Ace),
			card.NewShowdownCard(card.Diamond, card.Two),
			card.NewShowdownCard(card.Diamond, card.Three),
			card.NewShowdownCard(card.Diamond, card.Four),
			card.NewShowdownCard(card.Diamond, card.Five),
			card.NewShowdownCard(card.Diamond, card.Six),
			card.NewShowdownCard(card.Diamond, card.Seven),
			card.NewShowdownCard(card.Diamond, card.Eight),
			card.NewShowdownCard(card.Diamond, card.Nine),
			card.NewShowdownCard(card.Diamond, card.Ten),
			card.NewShowdownCard(card.Diamond, card.Jack),
			card.NewShowdownCard(card.Diamond, card.Queen),
			card.NewShowdownCard(card.Diamond, card.King),
			card.NewShowdownCard(card.Heart, card.Ace),
			card.NewShowdownCard(card.Heart, card.Two),
			card.NewShowdownCard(card.Heart, card.Three),
			card.NewShowdownCard(card.Heart, card.Four),
			card.NewShowdownCard(card.Heart, card.Five),
			card.NewShowdownCard(card.Heart, card.Six),
			card.NewShowdownCard(card.Heart, card.Seven),
			card.NewShowdownCard(card.Heart, card.Eight),
			card.NewShowdownCard(card.Heart, card.Nine),
			card.NewShowdownCard(card.Heart, card.Ten),
			card.NewShowdownCard(card.Heart, card.Jack),
			card.NewShowdownCard(card.Heart, card.Queen),
			card.NewShowdownCard(card.Heart, card.King),
			card.NewShowdownCard(card.Spade, card.Ace),
			card.NewShowdownCard(card.Spade, card.Two),
			card.NewShowdownCard(card.Spade, card.Three),
			card.NewShowdownCard(card.Spade, card.Four),
			card.NewShowdownCard(card.Spade, card.Five),
			card.NewShowdownCard(card.Spade, card.Six),
			card.NewShowdownCard(card.Spade, card.Seven),
			card.NewShowdownCard(card.Spade, card.Eight),
			card.NewShowdownCard(card.Spade, card.Nine),
			card.NewShowdownCard(card.Spade, card.Ten),
			card.NewShowdownCard(card.Spade, card.Jack),
			card.NewShowdownCard(card.Spade, card.Queen),
			card.NewShowdownCard(card.Spade, card.King),
		},
	}
}
