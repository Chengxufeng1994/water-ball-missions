package main

import "math/rand"

type Deck struct {
	cards []Card
}

func (b *Deck) Shuffle() {
	for i := range b.cards {
		j := rand.Intn(i + 1)
		b.cards[i], b.cards[j] = b.cards[j], b.cards[i]
	}
}

func (b *Deck) Draw() Card {
	card := b.cards[0]
	b.cards = b.cards[1:]
	return card
}

func NewUnoDeck() *Deck {
	return &Deck{
		cards: []Card{
			NewUnoCard(ColorBlue, 1),
			NewUnoCard(ColorBlue, 2),
			NewUnoCard(ColorBlue, 3),
			NewUnoCard(ColorBlue, 4),
			NewUnoCard(ColorBlue, 5),
			NewUnoCard(ColorBlue, 6),
			NewUnoCard(ColorBlue, 7),
			NewUnoCard(ColorBlue, 8),
			NewUnoCard(ColorBlue, 9),
			NewUnoCard(ColorBlue, 10),
			NewUnoCard(ColorRed, 1),
			NewUnoCard(ColorRed, 2),
			NewUnoCard(ColorRed, 3),
			NewUnoCard(ColorRed, 4),
			NewUnoCard(ColorRed, 5),
			NewUnoCard(ColorRed, 6),
			NewUnoCard(ColorRed, 7),
			NewUnoCard(ColorRed, 8),
			NewUnoCard(ColorRed, 9),
			NewUnoCard(ColorRed, 10),
			NewUnoCard(ColorYellow, 1),
			NewUnoCard(ColorYellow, 2),
			NewUnoCard(ColorYellow, 3),
			NewUnoCard(ColorYellow, 4),
			NewUnoCard(ColorYellow, 5),
			NewUnoCard(ColorYellow, 6),
			NewUnoCard(ColorYellow, 7),
			NewUnoCard(ColorYellow, 8),
			NewUnoCard(ColorYellow, 9),
			NewUnoCard(ColorYellow, 10),
			NewUnoCard(ColorGreen, 1),
			NewUnoCard(ColorGreen, 2),
			NewUnoCard(ColorGreen, 3),
			NewUnoCard(ColorGreen, 4),
			NewUnoCard(ColorGreen, 5),
			NewUnoCard(ColorGreen, 6),
			NewUnoCard(ColorGreen, 7),
			NewUnoCard(ColorGreen, 8),
			NewUnoCard(ColorGreen, 9),
			NewUnoCard(ColorGreen, 10),
		},
	}
}

func NewShowdownDeck() *Deck {
	return &Deck{
		cards: []Card{
			NewShowdownCard(Club, Ace),
			NewShowdownCard(Club, Two),
			NewShowdownCard(Club, Three),
			NewShowdownCard(Club, Four),
			NewShowdownCard(Club, Five),
			NewShowdownCard(Club, Six),
			NewShowdownCard(Club, Seven),
			NewShowdownCard(Club, Eight),
			NewShowdownCard(Club, Nine),
			NewShowdownCard(Club, Ten),
			NewShowdownCard(Club, Jack),
			NewShowdownCard(Club, Queen),
			NewShowdownCard(Club, King),
			NewShowdownCard(Diamond, Ace),
			NewShowdownCard(Diamond, Two),
			NewShowdownCard(Diamond, Three),
			NewShowdownCard(Diamond, Four),
			NewShowdownCard(Diamond, Five),
			NewShowdownCard(Diamond, Six),
			NewShowdownCard(Diamond, Seven),
			NewShowdownCard(Diamond, Eight),
			NewShowdownCard(Diamond, Nine),
			NewShowdownCard(Diamond, Ten),
			NewShowdownCard(Diamond, Jack),
			NewShowdownCard(Diamond, Queen),
			NewShowdownCard(Diamond, King),
			NewShowdownCard(Heart, Ace),
			NewShowdownCard(Heart, Two),
			NewShowdownCard(Heart, Three),
			NewShowdownCard(Heart, Four),
			NewShowdownCard(Heart, Five),
			NewShowdownCard(Heart, Six),
			NewShowdownCard(Heart, Seven),
			NewShowdownCard(Heart, Eight),
			NewShowdownCard(Heart, Nine),
			NewShowdownCard(Heart, Ten),
			NewShowdownCard(Heart, Jack),
			NewShowdownCard(Heart, Queen),
			NewShowdownCard(Heart, King),
			NewShowdownCard(Spade, Ace),
			NewShowdownCard(Spade, Two),
			NewShowdownCard(Spade, Three),
			NewShowdownCard(Spade, Four),
			NewShowdownCard(Spade, Five),
			NewShowdownCard(Spade, Six),
			NewShowdownCard(Spade, Seven),
			NewShowdownCard(Spade, Eight),
			NewShowdownCard(Spade, Nine),
			NewShowdownCard(Spade, Ten),
			NewShowdownCard(Spade, Jack),
			NewShowdownCard(Spade, Queen),
			NewShowdownCard(Spade, King),
		},
	}
}
