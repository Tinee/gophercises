package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Value int
	Color Color
}

type Color int

type Deck []Card

// Different Colors
const (
	Heart Color = iota
	Spade
	Square
	Clubs
)

func main() {
	c := NewCards()
	c.Shuffle()
	fmt.Println(c)
}

func NewCards() Deck {
	var ret []Card
	for i := 0; i < 4; i++ {
		temp := Card{
			Color: Color(i),
		}
		for j := 1; j < 14; j++ {
			temp.Value = j
			ret = append(ret, temp)
		}
	}

	return ret
}

func (c Deck) Shuffle() {
	s := rand.NewSource(time.Now().Unix())
	rand.New(s).Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
}
