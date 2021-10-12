package cards

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Color string
	Value string
}

type Deck struct {
	Discarded []Card `default:"[]"`
	Shuffled  bool   `default:"false"`
	Stack     []Card `default:"[]"`
	Hands     []Hand `default:"[]"`
}

type Hand struct {
	Cards []Card `default:"[]"`
	Owner string
}

func Create_Deck() Deck {
	var d Deck
	d.Shuffled = false

	var all_cards []Card
	for _, suit := range []string{"Spades", "Clubs", "Hearts", "Diamonds"} {
		for _, value := range []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} {
			this_card := new(Card)
			this_card.Suit = suit
			this_card.Value = value
			if suit == "Spades" || suit == "Clubs" {
				this_card.Color = "Black"
			} else {
				this_card.Color = "Red"
			}
			all_cards = append(all_cards, *this_card)
		}
	}

	d.Stack = all_cards

	return d
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Stack), func(i, j int) { d.Stack[i], d.Stack[j] = d.Stack[j], d.Stack[i] })
	d.Shuffled = true
}

func (d *Deck) Deal(players int, cards_per_player int) {
	// Assign Player Names
	for x := 1; x <= players; x++ {
		player_hand := new(Hand)
		player_hand.Owner = fmt.Sprintf("Player %d", x)
		d.Hands = append(d.Hands, *player_hand)
	}

	for i := 0; i < cards_per_player; i++ {
		for j := 0; j < players; j++ {
			card := d.Stack[0]
			d.Stack = d.Stack[1:]
			d.Hands[j].Cards = append(d.Hands[j].Cards, card)
		}
	}
}
