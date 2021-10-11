package cards

import "testing"

func TestCreateDeck(t *testing.T) {
	testDeck := Create_Deck()

	// Check not shuffled
	want_shuffled_value := false
	test_shuffled_value := testDeck.Shuffled
	if test_shuffled_value != want_shuffled_value {
		t.Errorf("Create_Deck() Shuffled: %v, want %v", test_shuffled_value, want_shuffled_value)
	}

	// Check Discared is Emtpy
	test_discarded := testDeck.Discarded
	if len(test_discarded) != 0 {
		t.Errorf("Create_Deck() Discarded: %s, want []", test_discarded)
	}

	// Check Hands are Empty
	test_hands := testDeck.Hands
	if len(test_hands) != 0 {
		t.Errorf("Create_Deck() Hands: %s, want []", test_hands)
	}
}

func TestShuffle(t *testing.T) {
	testDeck := Create_Deck()
	testDeck2 := Create_Deck()
	testDeck2.Shuffle()

	same := false
	for index, card := range testDeck.Stack {
		// Make sure cards in same position consecutively to at least verify some difference
		if testDeck2.Stack[index] == card && testDeck2.Stack[index+1] == testDeck.Stack[index+1] {
			same = true
			break
		}
	}
	if same != false {
		t.Errorf("Shuffle() Decks are the same.")
	}
}

func TestDeal(t *testing.T) {
	testDeck := Create_Deck()
	testDeck.Deal(3, 4)

	want_players := 3
	if len(testDeck.Hands) != want_players {
		t.Errorf("Deal() Player Count: got %d, want %d", len(testDeck.Hands), want_players)
	}

	want_cards_in_hand := 4
	for _, player := range testDeck.Hands {
		if len(player.Cards) != want_cards_in_hand {
			t.Errorf("Deal() Card in Hand Count: got %d, want %d", len(testDeck.Hands), want_players)
		}
	}

}
