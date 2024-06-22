//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	Ace Rank = iota + 1
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
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func(*[]Card)) []Card {
	cs := make([]Card, 0, 52)

	cs = append(cs, makeSuit(Spade)...)
	cs = append(cs, makeSuit(Diamond)...)
	cs = append(cs, makeSuit(Club)...)
	cs = append(cs, makeSuit(Heart)...)

	for _, opt := range opts {
		opt(&cs)
	}

	return cs
}

func DefaultSort(c *[]Card) {
	cards := *c
	sort.Slice(cards, Less(cards))
	*c = cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func(*[]Card) {
	return func(c *[]Card) {
		cards := *c
		sort.Slice(cards, less(cards))
		*c = cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(King) + int(c.Rank)
}

func Shuffle(c *[]Card) {
	cards := *c
	ret := make([]Card, len(cards))
	perm := rand.Perm(len(cards))

	for i, j := range perm {
		ret[i] = cards[j]
	}
	*c = ret
}

func Jokers(i int) func(*[]Card) {
	return func(c *[]Card) {
		cards := *c
		for i > 0 {
			cards = append(cards, Card{Suit: Joker})
			i--
		}
		*c = cards
	}
}

func Filter(f func(Card) bool) func(*[]Card) {
	return func(c *[]Card) {
		cards := *c
		var ret []Card
		for _, card := range cards {
			if !f(card) {
				ret = append(ret, card)
			}
		}
		*c = ret
	}
}

func AddDeck(i int) func(*[]Card) {
	return func(c *[]Card) {
		cards := *c
		var ret []Card
		for i > 0 {
			ret = append(ret, cards...)
			i--
		}
		*c = ret
	}
}

func makeSuit(s Suit) []Card {
	sp := make([]Card, 0, 13)

	for i := 1; i <= 13; i++ {
		sp = append(sp, Card{Rank: Rank(i), Suit: s})
	}

	return sp
}
