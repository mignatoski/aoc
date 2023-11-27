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

type Grid map[Point]rune

var grid Grid

var size, start, pos Point

var (
	UP    Point = Point{0, -1}
	RIGHT Point = Point{1, 0}
	LEFT  Point = Point{-1, 0}
	DOWN  Point = Point{0, 1}
)

type Direction struct {
	D     Point
	Value int
	L, R  *Direction
}

type SideName int
type EdgeName int

const (
	eRight EdgeName = iota
	eDown
	eLeft
	eUp
)

const (
	sTop SideName = iota
	sBottom
	sLeft
	sRight
	sFront
	sBack
)

type Side struct {
	TopLeft Point
	Edges   map[EdgeName]Edge
}

type Edge struct {
	S         SideName
	E         EdgeName
	IsFlipped bool
}

var dir *Direction

var dR, dD, dL, dU Direction

var cube map[SideName]Side

func init() {
	dR = Direction{D: RIGHT, Value: 0}
	dD = Direction{D: DOWN, Value: 1}
	dL = Direction{D: LEFT, Value: 2}
	dU = Direction{D: UP, Value: 3}

	dR.L, dR.R = &dU, &dD
	dD.L, dD.R = &dR, &dL
	dL.L, dL.R = &dD, &dU
	dU.L, dU.R = &dL, &dR

	dir = &dR

	grid = make(Grid)

	cube = make(map[SideName]Side, 6)
	cube[sTop] = Side{TopLeft: Point{50, 0}, Edges: make(map[EdgeName]Edge)}
	cube[sBottom] = Side{TopLeft: Point{50, 100}, Edges: make(map[EdgeName]Edge)}
	cube[sLeft] = Side{TopLeft: Point{0, 100}, Edges: make(map[EdgeName]Edge)}
	cube[sRight] = Side{TopLeft: Point{100, 0}, Edges: make(map[EdgeName]Edge)}
	cube[sFront] = Side{TopLeft: Point{50, 50}, Edges: make(map[EdgeName]Edge)}
	cube[sBack] = Side{TopLeft: Point{0, 150}, Edges: make(map[EdgeName]Edge)}

	cube[sTop].Edges[eUp] = Edge{sBack, eRight, false}

	// think about rotation and direction after folding

}

var sideLength int

func main() {
	fileName := "sample.txt"
	inputFile, _ := os.Open(fileName)
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	switch fileName {
	case "sample.txt":
		sideLength = 4
	default:
		sideLength = 50
	}

	var line string
	// map
	var p Point
	for fileScanner.Scan() {
		line = fileScanner.Text()
		if line == "" {
			break
		}

		for i, r := range []rune(line) {
			if start.X == 0 && r == '.' {
				start.X = i
			}

			switch r {
			case ' ':
			// ignore
			case '.', '#':
				p.X = i
				p.Y = size.Y
				grid[p] = r
			}
		}

		size.X = max(size.X, len(line))
		size.Y++
	}
	fmt.Println(start, dir)
	pos = start

	// actions
	fileScanner.Scan()
	line = fileScanner.Text()

	// part 1
	var buf string
	for i, r := range []rune(line) {
		switch {
		case r >= '0' && r <= '9':
			// number
			buf += string(r)
		case r == 'L' || r == 'R':
			// turn
			fmt.Println("Move: ", buf)
			m, _ := strconv.ParseInt(buf, 0, 0)
			Move(int(m))
			fmt.Println("Turn: ", string(r))
			Turn(r)
			buf = ""
		default:
			panic("uh oh")
		}

		if i == len(line)-1 {
			//end
			fmt.Println("Move: ", buf)
			m, _ := strconv.ParseInt(buf, 0, 0)
			Move(int(m))
		}
	}

	fmt.Println("Part 1: ", Password())

	// part 2
	buf = ""
	pos = start
	dir = &dR
	fmt.Println(line)
	for i, r := range []rune(line) {
		switch {
		case r >= '0' && r <= '9':
			// number
			buf += string(r)
		case r == 'L' || r == 'R':
			// turn
			fmt.Println("Move: ", buf)
			m, _ := strconv.ParseInt(buf, 0, 0)
			Move2(int(m))
			fmt.Println("Turn: ", string(r))
			Turn(r)
			buf = ""
		default:
			panic("uh oh")
		}

		if i == len(line)-1 {
			//end
			fmt.Println("Move: ", buf)
			m, _ := strconv.ParseInt(buf, 0, 0)
			Move2(int(m))
		}
	}

	fmt.Println("Part 2: ", Password())

}

func Password() int {
	return (1000 * (pos.Y + 1)) + (4 * (pos.X + 1)) + (dir.Value)
}

func Move2(steps int) {
	for i := 0; i < steps; i++ {
		n, valid := Check2()
		if valid {
			pos = n
		} else {
			break
		}
	}
}

func Wrap() Point {
	var w Point
	var s Side
	var ox, oy int

	for _, v := range cube {
		tl := v.TopLeft
		if pos.X >= tl.X && pos.X < tl.X+sideLength &&
			pos.Y >= tl.Y && pos.Y < tl.Y+sideLength {
			ox = pos.X - tl.X
			oy = pos.Y - tl.Y

			s = v
			break
		}
	}

	e := EdgeName(dir.Value)
	ns := s.Edges[e].S
	w = cube[ns].TopLeft
	w.X += ox
	w.Y += oy

	return w
}

func Check2() (next Point, valid bool) {
	next.X = pos.X + dir.D.X
	next.Y = pos.Y + dir.D.Y
	for {
		fmt.Println(pos, next, grid[next], dir.D)
		switch grid[next] {
		case 0:
			w := Wrap()
			next.X = w.X
			next.Y = w.Y
		case '.':
			valid = true
			return
		case '#':
			valid = false
			return
		default:
			panic("???")
		}
	}
}

func Move(steps int) {
	for i := 0; i < steps; i++ {
		n, valid := Check()
		if valid {
			pos = n
		} else {
			break
		}
	}
}

func Check() (next Point, valid bool) {
	next.X = pos.X + dir.D.X
	next.Y = pos.Y + dir.D.Y
	for {
		fmt.Println(pos, next, grid[next], dir.D)
		switch grid[next] {
		case 0:
			next.X += dir.D.X
			next.Y += dir.D.Y
			if next.X < 0 || next.Y < 0 ||
				next.X >= size.X || next.Y >= size.Y {
				switch dir.D {
				case RIGHT:
					next.X = 0
				case DOWN:
					next.Y = 0
				case LEFT:
					next.X = size.X - 1
				case UP:
					next.Y = size.Y - 1
				default:
					panic("awwp")
				}
			}
		case '.':
			valid = true
			return
		case '#':
			valid = false
			return
		default:
			panic("???")
		}
	}
}

func Turn(d rune) {
	switch d {
	case 'L':
		dir = dir.L
	case 'R':
		dir = dir.R
	}
	fmt.Println("Facing: ", dir.Value)
}
