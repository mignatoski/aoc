package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var jetPos int
var jets []rune
var numRocks int
var rockPattern [][]Point
var chamber Chamber
var numFalls int
var cycle bool
var states map[State]int

func nextJet() (next rune) {
	next = jets[jetPos]
	jetPos = (jetPos + 1) % len(jets)
	return
}

type Chamber struct {
	Rocks                        []*Rock
	Grid                         map[Point]bool // true is a rock, false is air
	HighestRock, PrevHighestRock int
}

type Rock struct {
	Points          []Point
	LowerLeftAnchor Point
}

type State struct {
	rockPattern, jetPos, numFalls, x, delta int
	fill                                    string
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
	states = make(map[State]int)
}

func main() {

	var rock *Rock

	for !cycle {
		// Rock Falls
		rock = RockFalls(rock)
		// ChamberPrint()

		// if numRocks > 2022 {
		// if numRocks > 1_000_000_000_000 {
		// Move up floor by checking complete lines?j

		// Rock Push by Jet
		rock.PushByJet()
		// ChamberPrint()
	}

	// fmt.Println("Part 1 Units Tall:", chamber.HighestRock+1)
	fmt.Println(numRocks)
}

func CheckForCycle(rock *Rock) bool {
	s := State{
		rockPattern: numRocks % 5,
		jetPos:      jetPos,
		numFalls:    numFalls,
		x:           rock.LowerLeftAnchor.X,
		fill:        ChamberFill(),
		delta:       chamber.HighestRock - chamber.PrevHighestRock,
	}
	if states[s] > 0 {
		fmt.Println(states[s])
		var pre, oneCycle, cycleLen, loop, post, rem int
		stateArray := make([]State, numRocks)
		for k, v := range states {
			stateArray[v] = k
			if v < states[s] {
				pre += k.delta
			} else if v >= states[s] {
				cycleLen++
				oneCycle += k.delta
			}
		}

		rocksLeft := 2022 - states[s] + 1
		loop = oneCycle * (rocksLeft / (numRocks - states[s]))
		rem = rocksLeft % (numRocks - states[s])
		for i := states[s]; i < states[s]+rem; i++ {
			post += stateArray[i].delta
		}
		fmt.Println("Part 1: ", pre+loop+post, pre, loop, post, oneCycle, rocksLeft, numRocks, states[s])

		rocksLeft = 1_000_000_000_000 - states[s] + 1
		loop = oneCycle * (rocksLeft / (cycleLen))
		rem = rocksLeft % (cycleLen)
		post = 0
		for i := states[s]; i < states[s]+rem; i++ {
			post += stateArray[i].delta
		}
		fmt.Println("Part 2: ", pre+loop+post, pre, loop, post, oneCycle, rocksLeft, numRocks, states[s], rocksLeft/cycleLen)
		return true
	} else {
		states[s] = numRocks
	}
	return false
}

func NextRock() (rock *Rock) {
	rock = &Rock{
		Points:          rockPattern[(numRocks)%5],
		LowerLeftAnchor: Point{2, 4 + chamber.HighestRock},
	}
	numRocks++
	numFalls = 0
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

func ChamberFill() string {

	sp := Point{0, chamber.HighestRock + 1}
	var str strings.Builder
	visited := make(map[Point]bool)
	check := make([]Point, 1)
	check[0] = sp
	visited[sp] = true
	var c, n, s, e, w Point
	i := 0
	for len(check) > 0 {
		c = check[0]
		n = Point{c.X, c.Y + 1}
		s = Point{c.X, c.Y - 1}
		e = Point{c.X + 1, c.Y}
		w = Point{c.X - 1, c.Y}
		if w.X >= 0 && !visited[w] {
			visited[w] = true
			if chamber.Grid[w] {
				str.WriteString(fmt.Sprintf("(w,%d)", i))
			} else {
				check = append(check, w)
			}
		}
		if e.X <= 6 && !visited[e] {
			visited[e] = true
			if chamber.Grid[e] {
				str.WriteString(fmt.Sprintf("(e,%d)", i))
			} else {
				check = append(check, e)
			}
		}
		if s.Y >= 0 && !visited[s] {
			visited[s] = true
			if chamber.Grid[s] {
				str.WriteString(fmt.Sprintf("(s,%d)", i))
			} else {
				check = append(check, s)
			}
		}
		if n.Y <= chamber.HighestRock+1 && !visited[n] {
			visited[n] = true
			if chamber.Grid[n] {
				str.WriteString(fmt.Sprintf("(n,%d)", i))
			} else {
				check = append(check, n)
			}
		}
		check = check[1:]
		i++
	}

	return str.String()
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
		numFalls++
		ChamberRock(rock, true)
	} else {
		ChamberRock(rock, true)
		chamber.PrevHighestRock = chamber.HighestRock
		ChamberUpdateHighestRock(rock)
		cycle = CheckForCycle(rock)
		rock = NextRock()
	}

	return rock
}
