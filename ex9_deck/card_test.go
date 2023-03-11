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
