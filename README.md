# Deck of Cards

Implements deck of cards with functional options

## Default New

```
cards := deck.New()
```

## New Deck with Shuffle

```
cards := deck.New(
    deck.Shuffle
)
```

## New Deck with Jokers

```
cards := deck.New(
    deck.Jokers(2)
)
```

## New Deck with Default Sort

```
cards := deck.New(
    deck.DefaultSort
)
```

## New Deck with Custom Sort

```
cards := deck.New(
    deck.Sort(deck.Less)
)
```

## New Deck with Filter

```
filter := func(card deck.Card) bool {
		return card.Rank == Two || card.Rank == Three
}
cards := deck.New(
    deck.Filter(filter)
)
```

## New Deck with multiple decks combined

```
//blackjack requires 3 standard decks
cards := deck.New(
    deck.AddDeck(3)
)
```

## New Deck with multiple options

```
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
