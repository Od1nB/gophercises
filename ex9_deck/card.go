//go:generate stringer -type=Suit, Color, Rank
package deck

type Suit uint8

const (
	Spades Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Color uint8

const (
	Black Color = iota
	Red
)

type Rank uint

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Color
	Rank
}

func (c Card) String() string {
	return ""
}
