package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Grid map[Point]rune
type LoopPipes map[Point]bool

var (
	grid  Grid
	start Point
	loop  LoopPipes
)

var (
	N = Point{0, -1}
	E = Point{1, 0}
	S = Point{0, 1}
	W = Point{-1, 0}
)

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (p Point) Ins() (a, b Point) {
	o := &a

	switch grid[p.Add(N)] {
	case '|', '7', 'F':
		*o = p.Add(N)
		o = &b
	}

	switch grid[p.Add(E)] {
	case '-', '7', 'J':
		*o = p.Add(E)
		o = &b
	}

	switch grid[p.Add(S)] {
	case '|', 'L', 'J':
		*o = p.Add(S)
		o = &b
	}

	switch grid[p.Add(W)] {
	case '-', 'L', 'F':
		*o = p.Add(W)
		o = &b
	}

	return
}

func (p Point) Outs(prev Point) (a, b Point) {
	fmt.Println(string(grid[p]))
	switch grid[p] {
	case '|':
		return p.Add(N), p.Add(S)
	case '-':
		return p.Add(E), p.Add(W)
	case 'L':
		return p.Add(N), p.Add(E)
	case 'J':
		return p.Add(N), p.Add(W)
	case '7':
		return p.Add(S), p.Add(W)
	case 'F':
		return p.Add(E), p.Add(S)
	}
	panic("not a valid pipe: " + string(grid[p]))
}

func (p Point) Next(prev Point) Point {
	p1, p2 := p.Outs(p)
	if p1 == prev {
		return p2
	}
	return p1
}

func (l LoopPipes) Mark(p Point) {
	l[p] = true
}

func main() {
	inputFile, _ := os.Open("sample2.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	y := 0
	grid = make(Grid)
	loop = make(LoopPipes)
	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
		for x, r := range line {
			grid[Point{x, y}] = r
			if r == 'S' {
				start = Point{x, y}
				loop.Mark(start)
			}
		}
		y++
	}
	fmt.Println("Start:", start)
	for row := 0; row < y; row++ {
		for col := range line {
			fmt.Print(string(grid[Point{col, row}]))
		}
		fmt.Print("\n")
	}
	p1, p2 := start.Ins()
	l1, l2 := start, start
	fmt.Println(p1, p2)

	steps := 1
	for p1 != p2 {
		n1 := p1.Next(l1)
		l1 = p1
		p1 = n1
		n2 := p2.Next(l2)
		l2 = p2
		p2 = n2
		fmt.Println(n1, n2)
		loop.Mark(n1)
		loop.Mark(n2)
		steps++
	}

	//

	fmt.Println("Part 1: ", steps)

	for row := 0; row < y; row++ {
		for col := range line {
			if loop[Point{col, row}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("Part 2: ", line)
}
