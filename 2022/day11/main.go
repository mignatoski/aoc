package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Round   int
	Monkeys []Monkey
}

type Monkey struct {
	Number                                  int
	Items                                   []int
	Operation, Test                         string
	TrueTarget, FalseTarget, ItemsInspected int
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
			monkey.Items = make([]int, len(items))
			for i, v := range items {
				tmpInt, _ := strconv.ParseInt(v, 0, 0)
				monkey.Items[i] = int(tmpInt)
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

	for round := 1; round <= 20; round++ {
		game.Round = round

		for i := range game.Monkeys {
			m := &game.Monkeys[i]
			for j, v := range m.Items {
				m.Items[j] = ApplyOperation(v, m.Operation) / 3
				m.ItemsInspected++
			}
		}
	}
	fmt.Printf("%+v\n", game)
}

func ApplyOperation(old int, op string) int {
	return old * 5
}
