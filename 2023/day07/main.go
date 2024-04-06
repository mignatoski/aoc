package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func (h *Hand) Rank() int {
	if h.rank > 0 {
		return h.rank
	}
	r := 0

	m := make(map[Card]int, 0)
	cnt1, cnt2 := 1, 1
	for _, c := range h.Cards {
		m[c] = m[c] + 1
		if m[c] > cnt1 {
			cnt1 = m[c]
		} else if m[c] > cnt2 {
			cnt2 = m[c]
		}
		if cnt1 < cnt2 {
			cnt1, cnt2 = cnt2, cnt1
		}
	}

	fmt.Println(m, cnt1, cnt2)

	switch {
	case cnt1 == 5:
		r = 7
	case cnt1 == 4:
		r = 6
	case cnt1 == 3 && cnt2 == 2:
		r = 5
	case cnt1 == 3:
		r = 4
	case cnt1 == 2 && cnt2 == 2:
		r = 3
	case cnt1 == 2:
		r = 2
	default:
		r = 1
	}

	h.rank = r
	return h.rank
}

func cmpHand(a, b *Hand) int {
	r1 := a.Rank()
	r2 := b.Rank()

	switch {
	case r1 < r2:
		return -1
	case r1 > r2:
		return 1
	default:
		for i := range 5 {
			if a.Cards[i].Value() < b.Cards[i].Value() {
				return -1
			} else if a.Cards[i].Value() > b.Cards[i].Value() {
				return 1
			}
		}
		return 0
	}
}

var (
	hands      []*Hand
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
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	hands = make([]*Hand, 0)
	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
		var (
			c string
			b int
		)

		fmt.Sscanf(line, "%v %d", &c, &b)

		ca := Cards(c)
		// cant sort due to second rule :( sort.Sort(sort.Reverse(ca))
		hands = append(hands, &Hand{Cards: ca, Bid: b})

	}

	slices.SortFunc(hands, cmpHand)

	part1 := 0
	for i, h := range hands {
		fmt.Println(h)
		part1 += (i + 1) * h.Bid
	}
	fmt.Println("Part 1: ", part1)

	fmt.Println("Part 2: ", line)
}
