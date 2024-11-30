package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func JokerWining() int64 {
	lines := ReadLines()

	var g JokerGame
	for _, ln := range lines {
		hand := NewJokerHand(ln)
		g = append(g, hand)
	}

	sort.Sort(g)

	var sum int64
	for i, v := range g {
		sum += int64(i+1) * v.Bid
	}
	return sum
}

type JokerHandType int

const (
	JokerHighCard JokerHandType = iota
	JokerOnePair
	JokerTwoPair
	JokerThreeOfAKind
	JokerFullHouse
	JokerFourOfAKind
	JokerFiveOfKind
)

type JokerHand struct {
	Type  JokerHandType
	Cards []rune
	Bid   int64
}

// Create hand by parsing line
func NewJokerHand(ln string) *JokerHand {
	parts := strings.Fields(ln)

	cards := []rune(parts[0])
	bid, err := strconv.ParseInt(parts[1], 10, 64)
	Check(err, "Unable to parse %s to int", parts[1])

	hand := getBestHand(cards)

	return &JokerHand{hand, cards, bid}
}

func getBestHand(cards []rune) JokerHandType {
	counts := make(map[rune]int)

	for _, c := range cards {
		counts[c]++
	}

	if counts['J'] != 0 {
		maxValue := 0
		var maxCount rune
		for c, v := range counts {
			if c != 'J' && v > maxValue {
				maxCount = c
				maxValue = v
			}
		}

		counts[maxCount] += counts['J']
		counts['J'] = 0
	}

	cardCounts := make([]int, 0, 5)

	for _, v := range counts {
		if v != 0 {
			cardCounts = append(cardCounts, v)
		}
	}

	sort.Ints(cardCounts)

	var hand JokerHandType
	switch {
	case jokerEqual(cardCounts, []int{5}):
		hand = JokerFiveOfKind
	case jokerEqual(cardCounts, []int{1, 4}):
		hand = JokerFourOfAKind
	case jokerEqual(cardCounts, []int{2, 3}):
		hand = JokerFullHouse
	case jokerEqual(cardCounts, []int{1, 1, 3}):
		hand = JokerThreeOfAKind
	case jokerEqual(cardCounts, []int{1, 2, 2}):
		hand = JokerTwoPair
	case jokerEqual(cardCounts, []int{1, 1, 1, 2}):
		hand = JokerOnePair
	case jokerEqual(cardCounts, []int{1, 1, 1, 1, 1}):
		hand = JokerHighCard
	default:
		panic(fmt.Sprintf("Unexpected card counts: %v", cardCounts))
	}

	return hand
}

func jokerEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type JokerGame [](*JokerHand)

func (g JokerGame) Len() int      { return len(g) }
func (g JokerGame) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g JokerGame) Less(i, j int) bool {
	a := *g[i]
	b := *g[j]

	if a.Type != b.Type {
		return a.Type < b.Type
	}

	cardValues := map[rune]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': -1,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	for k := 0; k < 5; k++ {
		if cardValues[a.Cards[k]] != cardValues[b.Cards[k]] {
			return cardValues[a.Cards[k]] < cardValues[b.Cards[k]]
		}
	}

	return false
}
