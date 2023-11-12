package main

import (
	"bufio"
	"fmt"
	"os"
)

var jetPos int
var jets []rune
var numRocks int
var rockPattern [][]Point
var chamber Chamber

func nextJet() (next rune) {
	next = jets[jetPos]
	jetPos = (jetPos + 1) % len(jets)
	return
}

type Chamber struct {
	Rocks       []*Rock
	Grid        map[Point]bool // true is a rock, false is air
	HighestRock int
}

type Rock struct {
	Points          []Point
	LowerLeftAnchor Point
}

type Point struct {
	X, Y int
}

func init() {
	rockPattern = [][]Point{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // 0: -
		{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}, // 1: +
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, // 2: L backwards
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},         // 3: |
		{{0, 0}, {1, 0}, {0, 1}, {1, 1}},         // 4: Square
	}

	chamber = Chamber{
		Rocks:       make([]*Rock, 0),
		Grid:        make(map[Point]bool),
		HighestRock: -1,
	}

	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Scan()
	jets = []rune(fileScanner.Text())
}

func main() {

	var rock *Rock

	for {
		// Rock Falls
		rock = RockFalls(rock)
		// ChamberPrint()

		if numRocks > 2022 {
			break
		}

		// Rock Push by Jet
		rock.PushByJet()
		// ChamberPrint()
	}

	fmt.Println("Part 1 Units Tall:", chamber.HighestRock+1)
}

func NextRock() (rock *Rock) {
	rock = &Rock{
		Points:          rockPattern[(numRocks)%5],
		LowerLeftAnchor: Point{2, 4 + chamber.HighestRock},
	}
	numRocks++
	chamber.Rocks = append(chamber.Rocks, rock)
	for _, v := range rock.Points {
		chamber.Grid[Point{
			X: rock.LowerLeftAnchor.X + v.X,
			Y: rock.LowerLeftAnchor.Y + v.Y,
		}] = true
	}
	return
}

func (rock *Rock) PushByJet() {
	jet := nextJet()
	ChamberRock(rock, false)
	switch jet {
	case '<':
		if ChamberRockFallCheck(rock, Point{-1, 0}) {
			rock.LowerLeftAnchor.X--
		}
	case '>':
		if ChamberRockFallCheck(rock, Point{1, 0}) {
			rock.LowerLeftAnchor.X++
		}
	}
	ChamberRock(rock, true)

}

func ChamberPrint() {
	fmt.Println("")
	for y := chamber.HighestRock + 5; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < 7; x++ {
			if chamber.Grid[Point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|\n")
	}
	fmt.Println("+-------+")
}

func ChamberRock(rock *Rock, val bool) {
	for _, v := range rock.Points {
		chamber.Grid[Point{
			X: rock.LowerLeftAnchor.X + v.X,
			Y: rock.LowerLeftAnchor.Y + v.Y,
		}] = val
	}
}

func ChamberRockFallCheck(rock *Rock, offset Point) (valid bool) {
	valid = true
	for _, v := range rock.Points {
		newX := rock.LowerLeftAnchor.X + v.X + offset.X
		newY := rock.LowerLeftAnchor.Y + v.Y + offset.Y
		if chamber.Grid[Point{
			X: newX,
			Y: newY,
		}] || newY < 0 || newX < 0 || newX > 6 {
			valid = false
			break
		}
	}
	return
}

func ChamberUpdateHighestRock(rock *Rock) {
	for _, v := range rock.Points {
		chamber.HighestRock = max(chamber.HighestRock, v.Y+rock.LowerLeftAnchor.Y)
	}
}

func RockFalls(rock *Rock) *Rock {
	// initial condition
	if rock == nil {
		return NextRock()
	}

	// Clear current pos
	ChamberRock(rock, false)

	if ChamberRockFallCheck(rock, Point{0, -1}) {
		rock.LowerLeftAnchor.Y -= 1
		ChamberRock(rock, true)
	} else {
		ChamberRock(rock, true)
		ChamberUpdateHighestRock(rock)
		rock = NextRock()
	}

	return rock
}
