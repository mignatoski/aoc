package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X, Y int
}

type Number struct {
	Value, Length int
	Chars         string
	Start         Point
}

type Sign struct {
	Position Point
	Type     rune
	Numbers  []int
}

func NewNumberRef() *Number {
	return &Number{}
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	lineNum := 0
	numbers := make([]*Number, 0)
	signPositions := make(map[Point]rune)
	gearPositions := make(map[Point]*Sign)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		var currentNumber *Number

		for i, r := range []rune(line) {
			p := Point{i, lineNum}
			if r >= '0' && r <= '9' {
				if currentNumber == nil {
					currentNumber = NewNumberRef()
					currentNumber.Chars = string(r)
					currentNumber.Start = p
					currentNumber.Length = 1
					numbers = append(numbers, currentNumber)
				} else {
					currentNumber.Chars += string(r)
					currentNumber.Length++
				}
			} else if r == '.' {
				currentNumber = nil
			} else {
				currentNumber = nil
				signPositions[p] = r

				if r == '*' {
					gearPositions[p] = &Sign{Numbers: make([]int, 0)}
				}
			}

		}
		lineNum++
	}

	totalPartNumbers := 0
	for _, n := range numbers {
		fmt.Println(n)
		for x := n.Start.X - 1; x <= n.Start.X+n.Length; x++ {
			for y := n.Start.Y - 1; y <= n.Start.Y+1; y++ {
				p := Point{x, y}
				integer, _ := strconv.ParseInt(n.Chars, 0, 0)
				if signPositions[p] != 0 {
					fmt.Println("yay")
					totalPartNumbers += int(integer)
				}

				if gearPositions[p] != nil {
					gearPositions[p].Numbers = append(gearPositions[p].Numbers, int(integer))
				}
			}
		}
	}

	totalGearRatio := 0
	for _, g := range gearPositions {
		if len(g.Numbers) == 2 {
			totalGearRatio += g.Numbers[0] * g.Numbers[1]
		}
	}

	fmt.Println("Part 1: ", totalPartNumbers)

	fmt.Println("Part 2: ", totalGearRatio)
}
