package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Three, Suit: Spade})
	fmt.Println(Card{Rank: Ace, Suit: Club})
	fmt.Println(Card{Rank: Four, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Queen, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	//Output:
	//Three of Spades
	//Ace of Clubs
	//Four of Hearts
	//Jack of Spades
	//Nine of Diamonds
	//Queen of Spades
	//Joker
}

func TestNew(t *testing.T) {
	deck := New()
	if len(deck) != 13*4 {
		t.Error("Wrong number of cards in new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(DefaultSort)
	if (deck[0] != Card{Rank: Ace, Suit: Spade}) {
		t.Error("First card is supposed to be Ace of Spades")
	}
}

func TestGenericSort(t *testing.T) {
	deck := New(Sort(Less))
	if (deck[0] != Card{Rank: Ace, Suit: Spade}) {
		t.Error("First card is supposed to be Ace of Spades")
	}
}

func TestJokers(t *testing.T) {
	deck := New(Jokers(5))
	count := 0
	for _, card := range deck {
		if card.Suit == Joker {
			count++
		}
	}
	if count != 5 {
		t.Errorf("Got %d Jokers, wanted 5", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == King || card.Rank == Queen
	}
	deck := New(Filter(filter))
	for _, card := range deck {
		if card.Rank == King || card.Rank == Queen {
			t.Errorf("Got a %s, but the filter should have removed it", card.String())
		}
	}
}
