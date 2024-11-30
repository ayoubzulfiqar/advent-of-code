package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadLines() []string {
	result := []string{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	Check(err, "error reading lines")

	return result
}
func TotalWining() int64 {
	lines := ReadLines()

	var g Game
	for _, ln := range lines {
		hand := NewHand(ln)
		g = append(g, hand)
	}

	sort.Sort(g)

	var sum int64
	for i, v := range g {
		sum += int64(i+1) * v.Bid
	}
	return sum
}

type HandType int64

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Type  HandType
	Cards []rune
	Bid   int64
}

// Create hand by parsing line
func NewHand(ln string) *Hand {
	parts := strings.Fields(ln)

	cards := []rune(parts[0])
	bid, err := strconv.ParseInt(parts[1], 10, 64)
	Check(err, "Unable to parse %s to int64", parts[1])

	counts := make(map[rune]int)

	for _, c := range cards {
		counts[c]++
	}

	cardCounts := make([]int, 0, 5)

	for _, v := range counts {
		cardCounts = append(cardCounts, v)
	}

	sort.Ints(cardCounts)

	var hand HandType
	switch {
	case equal(cardCounts, []int{5}):
		hand = FiveOfAKind
	case equal(cardCounts, []int{1, 4}):
		hand = FourOfAKind
	case equal(cardCounts, []int{2, 3}):
		hand = FullHouse
	case equal(cardCounts, []int{1, 1, 3}):
		hand = ThreeOfAKind
	case equal(cardCounts, []int{1, 2, 2}):
		hand = TwoPair
	case equal(cardCounts, []int{1, 1, 1, 2}):
		hand = OnePair
	case equal(cardCounts, []int{1, 1, 1, 1, 1}):
		hand = HighCard
	default:
		panic(fmt.Sprintf("Unexpected card counts: %v", cardCounts))
	}

	return &Hand{hand, cards, bid}
}

func equal(a, b []int) bool {
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

type Game [](*Hand)

func (g Game) Len() int      { return len(g) }
func (g Game) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g Game) Less(i, j int) bool {
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
		'J': 9,
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

func Check(e error, format string, a ...any) {
	if e != nil {
		message := fmt.Sprintf(format, a...)
		panic(fmt.Errorf("%s: %s", message, e))
	}
}
