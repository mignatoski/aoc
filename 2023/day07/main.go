package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

type Card rune
type Cards []Card

// Len implements sort.Interface.
func (c Cards) Len() int {
	return len(c)
}

// Less implements sort.Interface.
func (c Cards) Less(i int, j int) bool {
	return c[i].Value() < c[j].Value()
}

// Swap implements sort.Interface.
func (c Cards) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cards) String() string {
	return string(c)
}

type Hand struct {
	Cards Cards
	Bid   int
	rank  int
}

func (c Card) Value() int {
	return cardValues[rune(c)]
}

func (h Hand) Rank() int {
	if h.rank > 0 {
		return h.rank
	}

	r := 0

	h.rank = r
	return h.rank
}

func cmpHand(a, b Hand) int {
	r1 := a.Rank()
	r2 := b.Rank()

	switch {
	case r1 < r2:
		return 1
	case r1 > r2:
		return -1
	default:
		for i := range 5 {
			if a.Cards[i].Value() < b.Cards[i].Value() {
				return 1
			} else if a.Cards[i].Value() > b.Cards[i].Value() {
				return -1
			}
		}
		return 0
	}
}

var (
	hands      []Hand
	cardValues map[rune]int
)

func init() {
	cardValues = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 11,
		'J': 12,
		'Q': 13,
		'K': 14,
		'A': 15,
	}
}

func main() {
	inputFile, _ := os.Open("sample.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	hands = make([]Hand, 0)
	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
		var (
			c string
			b int
		)

		fmt.Sscanf(line, "%v %d", &c, &b)

		ca := Cards(c)
		sort.Sort(sort.Reverse(ca))
		hands = append(hands, Hand{Cards: ca, Bid: b})

	}

	slices.SortFunc(hands, cmpHand)

	for _, h := range hands {
		fmt.Println(h)
	}
	fmt.Println("Part 1: ", line)

	fmt.Println("Part 2: ", line)
}
