package deck

import "fmt"

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
