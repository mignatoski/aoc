package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	Round   int
	Monkeys []Monkey
}

type Monkey struct {
	Number                  int
	Items                   []uint64
	Operation, Test         string
	TrueTarget, FalseTarget int
	ItemsInspected          uint64
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	lineCount := 0

	game := Game{}
	game.Monkeys = make([]Monkey, 0)
	var monkey *Monkey

	for fileScanner.Scan() {
		line = fileScanner.Text()
		_, value, _ := strings.Cut(line, ": ")

		switch lineCount % 7 {
		case 0:
			monkey = &Monkey{Number: lineCount / 7}
		case 1:
			items := strings.Split(value, ", ")
			monkey.Items = make([]uint64, len(items))
			for i, v := range items {
				tmpInt, _ := strconv.ParseInt(v, 0, 0)
				monkey.Items[i] = uint64(tmpInt)
			}
		case 2:
			_, monkey.Operation, _ = strings.Cut(value, " = ")
		case 3:
			_, monkey.Test, _ = strings.Cut(value, "divisible by ")
		case 4:
			tmpInt, _ := strconv.ParseInt(value[len(value)-1:], 0, 0)
			monkey.TrueTarget = int(tmpInt)
		case 5:
			tmpInt, _ := strconv.ParseInt(value[len(value)-1:], 0, 0)
			monkey.FalseTarget = int(tmpInt)
		case 6:
			game.Monkeys = append(game.Monkeys, *monkey)
		}

		lineCount++
	}

	fmt.Printf("%+v\n\n", game)
	for round := 1; round <= 10000; round++ {
		game.Round = round

		for i := range game.Monkeys {
			m := &game.Monkeys[i]
			itemsTrue := make([]uint64, 0)
			itemsFalse := make([]uint64, 0)

			for _, v := range m.Items {
				item := v
				item = ApplyOperation(item, m.Number) // /* part 1  */ / 3

				item = item % (5 * 11 * 2 * 13 * 7 * 3 * 17 * 19)

				if ApplyTest(item, m.Number) {
					itemsTrue = append(itemsTrue, item)
				} else {
					itemsFalse = append(itemsFalse, item)
				}

				m.ItemsInspected++
			}
			m.Throw(itemsTrue, &game.Monkeys[m.TrueTarget])
			m.Throw(itemsFalse, &game.Monkeys[m.FalseTarget])
			m.Items = make([]uint64, 0)
		}
		fmt.Printf("%+v\n\n", game)
	}

	sort.Slice(game.Monkeys, func(i, j int) bool {
		return game.Monkeys[i].ItemsInspected < game.Monkeys[j].ItemsInspected
	})

	fmt.Println(game.Monkeys[6].ItemsInspected, game.Monkeys[7].ItemsInspected)
	fmt.Println(game.Monkeys[6].ItemsInspected * game.Monkeys[7].ItemsInspected)
}

func (g Game) String() string {
	var s string
	var totalInspected uint64
	s = fmt.Sprintln("Round :", g.Round)
	for _, v := range g.Monkeys {
		totalInspected += v.ItemsInspected
		s += fmt.Sprintln(v.Number, ": ", v.Items, " - ", v.ItemsInspected)
	}
	s += fmt.Sprintln("Inspected: ", totalInspected)
	return s
}

func (m *Monkey) Throw(arr []uint64, target *Monkey) {
	target.Items = append(target.Items, arr...)
}

func ApplyTest(worry uint64, number int) bool {
	/* Cycles
	0 -> 2 -> 5 -> 4 -> 7
	*/
	var result bool

	switch number {
	case 0:
		result = worry%5 == 0
	case 1:
		result = worry%11 == 0
	case 2:
		result = worry%2 == 0
	case 3:
		result = worry%13 == 0
	case 4:
		result = worry%7 == 0
	case 5:
		result = worry%3 == 0
	case 6:
		result = worry%17 == 0
	case 7:
		result = worry%19 == 0
	}

	return result
}

func ApplyOperation(old uint64, number int) uint64 {
	var result uint64

	switch number {
	case 0:
		result = old * 3
	case 1:
		result = old + 8
	case 2:
		result = old + 2
	case 3:
		result = old + 4
	case 4:
		result = old * 19
	case 5:
		result = old + 5
	case 6:
		result = old * old
	case 7:
		result = old + 1
	}

	return result
}
