# Deck of Cards

Implements deck of cards with functional options

## Default New

```go
cards := deck.New()
```

## New Deck with Shuffle

```go
cards := deck.New(
    deck.Shuffle
)
```

## New Deck with Jokers

```go
cards := deck.New(
    deck.Jokers(2)
)
```

## New Deck with Default Sort

```go
cards := deck.New(
    deck.DefaultSort
)
```

## New Deck with Custom Sort

```go
cards := deck.New(
    deck.Sort(deck.Less)
)
```

## New Deck with Filter

```go
filter := func(card deck.Card) bool {
		return card.Rank == Two || card.Rank == Three
}
cards := deck.New(
    deck.Filter(filter)
)
```

## New Deck with multiple decks combined

```go
//blackjack requires 3 standard decks
cards := deck.New(
    deck.AddDeck(2)
)
```

## New Deck with multiple options

```go
addJackOfSpades := func(c *[]deck.Card){
   cards := *c
   cards = append(cards, Card{Suit: deck.Spade, Rank: deck.Jack})
   *c = cards
}
cards := deck.New(
    deck.Shuffle,
    deck.Jokers(2),
    deck.Sort(deck.Less),
    deck.AddDeck(3),
    addJackOfSpades
)
```
